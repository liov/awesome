pacman -S mingw-w64-ucrt-x86_64-unicorn

如果提示版本不对
// https://github.com/unicorn-engine/unicorn/releases
// 复制windows_mingw64-shared.7z /bin/libunicorn.dll 到执行目录下
然后会报错
Exception 0xc0000005 0x1 0x18ad53a0000 0x7ff85dfc9f95
PC=0x7ff85dfc9f95
signal arrived during external code execution