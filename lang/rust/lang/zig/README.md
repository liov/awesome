# .cargo/config.toml
[target.x86_64-unknown-linux-musl]
linker = "zig"
rustflags = ["-C", "linker-flavor=gcc", "-C", "link-arg=cc", "-C", "link-arg=-target", "-C", "link-arg=x86_64-linux-musl"]