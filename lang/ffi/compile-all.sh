#!/bin/sh

tup upd && \
zig build && \
go build -o go_hello-cgo.exe hello-cgo.go && \
go build -o go_hello.exe hello.go && \
rustc -C opt-level=2 -C link-args="-Lnewplus -lnewplus -Wl,-rpath,\$ORIGIN/newplus"  -o rust_hello.exe hello.rs