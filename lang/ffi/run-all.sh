#!/bin/sh

[ $# -gt 0 ] || { echo "First arg (0 - 2000000000) is required."; exit 1; }

[ "$1" -eq "$1" ] 2>/dev/null || { echo "Must be a positive number not exceeding 2 billion."; exit 1; }

echo "The results are elapsed time in milliseconds"
echo "============================================"

echo "\nluajit:"
luajit hello.lua $@ && \
luajit hello.lua $@

echo "\nc:"
./c_hello $@ && \
./c_hello $@

echo "\ncpp:"
./cpp_hello $@ && \
./cpp_hello $@

echo "\nzig:"
./zig-out/zig_hello $@ && \
./zig-out/zig_hello $@

echo "\nrust:"
./rust_hello $@ && \
./rust_hello $@

echo "\njava23:"
D:/sdk/jdks/openjdk-23.0.1/binjava -cp . jhello.Hello $@ && \
D:/sdk/jdks/openjdk-23.0.1/binjava -cp . jhello.Hello $@

echo "\ngo:"
./go_hello $@ && \
./go_hello $@

echo "\ndart:"
dart hello.dart $@ && \
dart hello.dart $@