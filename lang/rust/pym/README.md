# hello-world (PyO3 示例)

这是一个最小可运行的 PyO3 示例，包含：

- Python 扩展模块：`hello_world._lib`
- Python 命令行入口：`sum-cli`
- Rust 可执行文件：`python-hello`

## 安装

```bash
python3 -m venv .venv
.venv/bin/python -m pip install -U pip setuptools setuptools-rust cffi
.venv/bin/python -m pip install -e .
```

## 使用

### 1) CLI

```bash
.venv/bin/sum-cli 1 2
```

输出示例：

```text
1 + 2 = 3
```

### 2) Python 导入

```bash
.venv/bin/python -c "import hello_world; print(hello_world.sum_as_string(2, 3))"
```

输出示例：

```text
5
```
