clang -S -c xxx.c -o xxx.asm -mno-red-zone, -mstackrealign, -mllvm, -inline-threshold=1000,
-fno-asynchronous-unwind-tables, -fno-exceptions, -fno-rtti

clang -c xxx.c -o xxx.o -mno-red-zone, -mstackrealign, -mllvm, -inline-threshold=1000,
-fno-asynchronous-unwind-tables, -fno-exceptions, -fno-rtti

objdump -d xxx.o --insn-width 16
