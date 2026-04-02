use std::{env, result};
use std::cell::RefCell;
use tokio::fs;
use shared_memory::*;
use napi_derive::napi;
use napi::bindgen_prelude::*;
use std::ffi::CString;

// sharedMemory 全局句柄，避免对进程重复创建（线程本地）
thread_local! {
    static SHMEM_GLOBAL: RefCell<Option<Shmem>> = const { RefCell::new(None) };
}

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

fn to_napi_err(prefix: &str, err: impl std::fmt::Display) -> napi::Error {
    napi::Error::new(Status::GenericFailure, format!("{prefix}: {err}"))
}

fn with_mem<T>(f: impl FnOnce(&mut Shmem) -> Result<T>) -> Result<T> {
    SHMEM_GLOBAL.with(|cell| {
        if cell.borrow().is_none() {
            let shmem = create_open_mem().map_err(|e| to_napi_err("failed to init shared memory", e))?;
            *cell.borrow_mut() = Some(shmem);
        }
        let mut guard = cell.borrow_mut();
        let shmem = guard
            .as_mut()
            .ok_or_else(|| napi::Error::new(Status::GenericFailure, "shared memory unavailable"))?;
        f(shmem)
    })
}

/**
 * 设置SharedMemory
*/
fn set_cache(set_cache: String) -> Result<()> {
    with_mem(|shmem| unsafe {
        let shared_state = unsafe { shmem.as_slice_mut() };
        let set_string = CString::new(set_cache.as_str())
            .map_err(|e| to_napi_err("failed to encode shared memory value", e))?;
        let bytes = set_string.to_bytes_with_nul();
        if bytes.len() > shared_state.len() {
            return Err(napi::Error::new(Status::InvalidArg, "shared memory value too large"));
        }
        shared_state.fill(0);
        shared_state[..bytes.len()].copy_from_slice(bytes);
        Ok(())
    })
}

/**
 * 读取SharedMemory
*/
fn get_cache() -> Result<String> {
    with_mem(|shmem| unsafe {
        let shared_state = unsafe { shmem.as_slice_mut() };
        let end = shared_state.iter().position(|b| *b == 0).unwrap_or(shared_state.len());
        Ok(String::from_utf8_lossy(&shared_state[..end]).into_owned())
    })
}

/**
 * 暴露给js端get的方法
*/
#[napi]
fn get() -> Result<String> {
    get_cache()
}

/**
 * 暴露给js端的set方法
*/
#[napi]
fn set(val:String)-> Result<()> {
    set_cache(val)
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