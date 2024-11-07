use std::{env, result};
use napi::tokio::fs;
use neon::prelude::*;
extern crate shared_memory;
use neon::prelude::*;
use shared_memory::*;
use napi::bindgen_prelude::*;
use napi_derive::napi;
use std::ffi::{CStr, CString};



static GLOBAL_LOCK_ID: usize = 0;

/**
 * sharedMemory全局句柄，避免对进程重复创建
*/
static mut SHMEM_GLOBAL: Option<Shmem> = None;

/**
 * 创建sharedMemory
*/
fn create_open_mem() -> result::Result<Shmem, ShmemError> {
    let shmem = match ShmemConf::new().size(4096).flink("shared_mem.link").create() {
        Ok(v) => v,
        Err(ShmemError::LinkExists) => ShmemConf::new().flink("shared_mem.link").open()?,
        Err(e) => return Err(e),
    };
    Ok(shmem)
}

/**
 * 设置SharedMemory
*/
fn set_cache(set_cache: String) -> result::Result<String, ShmemError> {
    {
        let mut  shared_state =  unsafe { SHMEM_GLOBAL.as_mut().unwrap().as_slice_mut()};
        let set_string: CString = CString::new(set_cache.as_str()).unwrap();
        shared_state[0..set_string.to_bytes_with_nul().len()]
            .copy_from_slice(set_string.to_bytes_with_nul());
    }
    Ok("".to_owned())
}

/**
 * 读取SharedMemory
*/
fn get_cache() -> result::Result<String, ShmemError> {
    let   result =
    {
        let shared_state = unsafe { SHMEM_GLOBAL.as_mut().unwrap().as_slice_mut()};
        let shmem_str: &CStr = unsafe { CStr::from_bytes_with_nul(shared_state).unwrap() };
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
    let value = cx.argument::<JsString>(0)?.value(&mut cx);
    match set_cache(value) {
        Ok(v) => Ok(cx.string(v)),
        Err(e) => Ok(cx.string("error")),
    }
}


#[neon::main]
fn main(mut cx: ModuleContext) -> NeonResult<()> {
      unsafe {
        SHMEM_GLOBAL = match create_open_mem() {
          Ok(v) => Some(v),
          _ => None,
        };
      }
    cx.export_function("get", get)?;
    cx.export_function("set", set)?;
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

#[napi]
async fn read_file_async(path: String) -> Result<Buffer> {
   match fs::read(path).await {
      Ok(content) => Ok(content.into()),
      Err(e) => Err(napi::Error::new(
        Status::GenericFailure,
        format!("failed to read file, {}", e),
      )),
    }
}