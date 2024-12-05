#!/bin/sh

[ $# -gt 0 ] || { echo "First arg (0 - 2000000000) is required."; exit 1; }

[ "$1" -eq "$1" ] 2>/dev/null || { echo "Must be a positive number not exceeding 2 billion."; exit 1; }

echo "The results are elapsed time in milliseconds"
echo "============================================"

echo "luajit:"
luajit hello.lua $@ && \
luajit hello.lua $@

echo "c:"
./c_hello $@ && \
./c_hello $@

echo "cpp:"
./cpp_hello $@ && \
./cpp_hello $@

echo "zig:"
./zig-out/bin/zig_hello $@ && \
./zig-out/bin/zig_hello $@

echo "rust:"
./rust_hello $@ && \
./rust_hello $@

echo "java20:"
D:/sdk/jdks/openjdk-20.0.1/bin/java -cp . jhello.Hello $@ && \
D:/sdk/jdks/openjdk-20.0.1/bin/java -cp . jhello.Hello $@

echo "go:"
./go_hello $@ && \
./go_hello $@

echo "dart:"
dart.bat hello.dart $@ && \
dart.bat hello.dart $@