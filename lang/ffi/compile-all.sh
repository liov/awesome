#!/bin/sh

OPT_FLAG="-O2"
CFLAGS="-fPIC ${OPT_FLAG} -Wall -Wextra -Wno-unused-parameter"
CXXFLAGS="-std=c++11 ${CFLAGS} -fno-rtti"

mkdir out
gcc -o out/plus.o ${CFLAGS} -c ./newplus/plus.c
gcc -o out/libnewplus.dll -fPIC -shared ${OPT_FLAG} -Wl,--whole-archive out/plus.o -Wl,--no-whole-archive
ar -rcs out/_.a out/plus.o
#gendef out/libnewplus.dll
#mv libnewplus.def out/
#dlltool -d out/libnewplus.def -l out/libnewplus.dll.a
DEPS=out/libnewplus.dll

LINK_FLAGS="-Lout -lnewplus -Wl,-rpath,\$ORIGIN"
gcc ${CFLAGS} ${LINK_FLAGS} -o out/c_hello hello.c
g++ ${CXXFLAGS} ${LINK_FLAGS} -o out/cpp_hello hello.cpp


zig build
go build -o out/go_hello-cgo.exe hello-cgo.go
go build -o out/go_hello.exe hello.go
rustc -C opt-level=2 -C link-args="-Lout -lnewplus -Wl,-rpath,\$ORIGIN"  -o out/rust_hello.exe hello.rs

NODE_VERSION=22.3.0

NODE_GYP_FLAGS="-DNODE_GYP_MODULE_NAME=newplus -DUSING_UV_SHARED=1 -DUSING_V8_SHARED=1 -DV8_DEPRECATION_WARNINGS=1 -D_LARGEFILE_SOURCE -D_FILE_OFFSET_BITS=64 -DBUILDING_NODE_EXTENSION -I$HOME/.node-gyp/${NODE_VERSION}/include/node -I$HOME/.node-gyp/${NODE_VERSION}/src -I$HOME/.node-gyp/${NODE_VERSION}/deps/uv/include -I$HOME/.node-gyp/${NODE_VERSION}/deps/v8/include"

NODE_CXX_FLAGS="${CXXFLAGS} -fno-exceptions"