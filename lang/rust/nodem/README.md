# nodem (N-API Rust 示例)

最小可编译的 napi-rs 示例，导出以下接口：

- `fibonacci(n: u32) -> u32`
- `get() -> string`
- `set(val: string) -> void`
- `get_cwd(callback)`
- `read_file_async(path: string) -> Promise<Buffer>`

## 构建检查

```bash
cargo check -p nodem
```

## 说明

- 已启用 `build.rs` 中 `napi_build::setup()`，便于后续接入 `@napi-rs/cli` 打包 `.node` 文件。
- `get/set` 使用线程本地共享内存句柄，避免 Rust 2024 下 `static mut` 的未定义行为风险。
