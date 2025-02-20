
ASM="zig cc" \
CC="zig cc" \
CXX="zig c++" \
cmake \
-DCMAKE_SYSTEM_NAME="Windows" \
-DCMAKE_SYSTEM_PROCESSOR="x86_64" \
-DCMAKE_ASM_COMPILER_TARGET="x86_64-windows-gnu" \
-DCMAKE_C_COMPILER_TARGET="x86_64-windows-gnu" \
-DCMAKE_CXX_COMPILER_TARGET="x86_64-windows-gnu" \
-DCMAKE_AR="$PWD/zig-ar" \
-DCMAKE_RANLIB="$PWD/zig-ranlib" \
-B build