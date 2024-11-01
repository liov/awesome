
use neon::prelude::*;
extern crate shared_memory;
use neon::prelude::*;
use neon::register_module;
use shared_memory::*;
use napi::bindgen_prelude::*;
use napi_derive::napi;
use std::ffi::{CStr, CString};

/**
 * 定义缓存区块
*/
#[derive(SharedMemCast)]
struct ShmemStructCache {
    num_slaves: u32,
    message: [u8; 256],
}

static GLOBAL_LOCK_ID: usize = 0;

/**
 * sharedMemory全局句柄，避免对进程重复创建
*/
static mut SHMEM_GLOBAL: Option<shared_memory::SharedMem> = None;

/**
 * 创建sharedMemory
*/
fn create_open_mem() -> Result<shared_memory::SharedMem, SharedMemError> {
    let shmem = match SharedMem::create_linked("shared_mem.link", LockType::Mutex, 4096) {
        Ok(v) => v,
        Err(SharedMemError::LinkExists) => SharedMem::open_linked("shared_mem.link")?,
        Err(e) => return Err(e),
    };

    if shmem.num_locks() != 1 {
        return Err(SharedMemError::InvalidHeader);
    }
    Ok(shmem)
}

/**
 * 设置SharedMemory
*/
fn set_cache(set_cache: String) -> Result<String, SharedMemError> {
    {
        let mut  shared_state =  unsafe { SHMEM_GLOBAL.as_mut().unwrap().wlock::<ShmemStructCache>(GLOBAL_LOCK_ID)?};
        let set_string: CString = CString::new(set_cache.as_str()).unwrap();
        shared_state.message[0..set_string.to_bytes_with_nul().len()]
            .copy_from_slice(set_string.to_bytes_with_nul());
    }
    Ok("".to_owned())
}

/**
 * 读取SharedMemory
*/
fn get_cache() -> Result<String, SharedMemError> {
    let   result =
    {
        let shmem = unsafe { SHMEM_GLOBAL.as_mut().unwrap()};
        let shared_state = shmem.rlock::<ShmemStructCache>(GLOBAL_LOCK_ID)?;
        let shmem_str: &CStr = unsafe { CStr::from_ptr(shared_state.message.as_ptr() as *mut i8) };
         shmem_str.to_str().unwrap().into()
    };

    Ok(result)
}

/**
 * 暴露给js端get的方法
*/
fn get(mut cx: FunctionContext) -> JsResult<JsString> {
    match get_cache() {
        Ok(v) => Ok(cx.string(v)),
        Err(_) => Ok(cx.string("error")),
    }
}

/**
 * 暴露给js端的set方法
*/
fn set(mut cx: FunctionContext) -> JsResult<JsString> {
    let value = cx.argument::<JsString>(0)?.value();
    match set_cache(value) {
        Ok(v) => Ok(cx.string(v)),
        Err(e) => Ok(cx.string("error")),
    }
}


fn hello(mut cx: FunctionContext) -> JsResult<JsString> {
    Ok(cx.string("hello node"))
}

#[neon::main]
fn main(mut cx: ModuleContext) -> NeonResult<()> {
      unsafe {
        SHMEM_GLOBAL = match create_open_mem() {
          Ok(v) => Some(v),
          _ => None,
        };
      }
      set_cache("".to_owned());
    cx.export_function("get", get)?;
    cx.export_function("set", set)?;
    cx.export_function("hello", hello)?;
    Ok(())
}

/// module registration is done by the runtime, no need to explicitly do it now.
#[napi]
fn fibonacci(n: u32) -> u32 {
  match n {
    1 | 2 => 1,
    _ => fibonacci(n - 1) + fibonacci(n - 2),
  }
}

/// use `Fn`, `FnMut` or `FnOnce` traits to defined JavaScript callbacks
/// the return type of callbacks can only be `Result`.
#[napi]
fn get_cwd<T: Fn(String) -> Result<()>>(callback: T) {
  callback(env::current_dir().unwrap().to_string_lossy().to_string()).unwrap();
}

/// or, define the callback signature in where clause
#[napi]
fn test_callback<T>(callback: T)
where T: Fn(String) -> Result<()>
{}

/// async fn, require `async` feature enabled.
/// [dependencies]
/// napi = {version="2", features=["async"]}
#[napi]
async fn read_file_async(path: String) -> Result<Buffer> {
  tokio::fs::read(path)
    .map(|r| match r {
      Ok(content) => Ok(content.into()),
      Err(e) => Err(Error::new(
        Status::GenericFailure,
        format!("failed to read file, {}", e),
      )),
    })
    .await
}