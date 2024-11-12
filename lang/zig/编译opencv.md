安装nasm https://www.nasm.us/pub/nasm/releasebuilds/2.16.03/win64/
安装了也会提示SIMD extensions disabled: could not find NASM compiler.  Performance will suffer.
在 CMake 配置中指定 NASM 路径：
安装eigen https://gitlab.com/libeigen/eigen/-/releases/3.4.0
pacman -S mingw-w64-clang-x86_64-openblas
下载opencv源码
修改CMakeLists.txt
set(CMAKE_C_COMPILER "zig cc")
set(CMAKE_CXX_COMPILER "zig c++")  

```cmd
set cvdir=D:\sdk\opencv

cmake %cvdir%\opencv-4.10.0  -B %cvdir%\zigbuild -DCMAKE_C_COMPILER="zig-cc.cmd" -DCMAKE_CXX_COMPILER="zig-c++.cmd" -DCMAKE_ASM_NASM_COMPILER=D:\sdk\msys64\clang64\bin\nasm.exe -DCMAKE_INCLUDE_PATH="D:\sdk\msys64\clang64\include" -DCMAKE_LIBRARY_PATH="D:\sdk\msys64\clang64\lib" -DCMAKE_SIZEOF_VOID_P=8 -DCMAKE_BUILD_TYPE=RELEASE -DENABLE_CXX11=ON -DBUILD_SHARED_LIBS=ON -DOPENCV_EXTRA_MODULES_PATH=%cvdir%\opencv_contrib-4.10.0\modules -DBUILD_DOCS=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=NO -DBUILD_opencv_python=NO -DBUILD_opencv_python2=NO -DBUILD_opencv_python3=NO -DWITH_JASPER=OFF -DWITH_QT=OFF -DWITH_GTK=OFF -DWITH_FFMPEG=OFF -DWITH_TIFF=OFF -DWITH_WEBP=OFF -DWITH_PNG=OFF -DWITH_1394=OFF -DWITH_OPENJPEG=OFF -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_ITT=OFF -Wno-dev
ninja -j%NUMBER_OF_PROCESSORS%
```

```bash
cvdir=/d/sdk/opencv
ASM="zig cc -target x86_64-windows-gnu"
CC="zig cc -target x86_64-windows-gnu"
CXX="zig c++ -target x86_64-windows-gnu" 
AR="zig ar"
RANLIB="zig ranlib"
cmake $cvdir/opencv-4.10.0 -B $cvdir/zigbuild \
-DCMAKE_C_COMPILER="zig-cc.cmd" -DCMAKE_CXX_COMPILER="zig-c++.cmd" -DCMAKE_C_FLAGS="-target x86_64-windows-gnu" -DCMAKE_CXX_FLAGS="-target x86_64-windows-gnu" \
-DCMAKE_INCLUDE_PATH="D:\sdk\msys64\clang64\include" -DCMAKE_LIBRARY_PATH="D:\sdk\msys64\clang64\lib" \
-DCMAKE_ASM_NASM_COMPILER="D:\sdk\msys64\clang64\bin\nasm.exe" \
-DCMAKE_SYSROOT="D:\sdk\msys64\clang64" -DPKG_CONFIG_LIBDIR="D:\sdk\msys64\clang64\lib\pkgconfig" \
-DCMAKE_SIZEOF_VOID_P=8 -DCMAKE_BUILD_TYPE=RELEASE -DENABLE_CXX11=ON -DBUILD_SHARED_LIBS=ON \
-DOPENCV_EXTRA_MODULES_PATH=$cvdir/opencv_contrib-4.10.0/modules \
-DBUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -DBUILD_opencv_java=NO -DBUILD_opencv_python=NO -DBUILD_opencv_python2=NO -DBUILD_opencv_python3=NO \
-DWITH_JASPER=OFF -DWITH_QT=OFF -DWITH_GTK=OFF -DWITH_FFMPEG=OFF -DWITH_TIFF=OFF -DWITH_WEBP=OFF -DWITH_PNG=OFF -DWITH_1394=OFF -DWITH_OPENJPEG=OFF -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_ITT=OFF  -Wno-dev

# 主要是手动指定DCMAKE_SIZEOF_VOID_P,opencv cmakelist会检查又获取不到
ninja -v -j$NUMBER_OF_PROCESSORS
```

error: unable to find dynamic system library 'pthread' using strategy 'paths_first'. searched paths: none

尝试失败https://github.com/ziglang/zig/issues/10989
zig不支持winpthread,怎么都找不到pthread.h
不折腾了zig cc -I /clang64/include -L/clang64/lib -lpthread pthread.c 可以编译成功