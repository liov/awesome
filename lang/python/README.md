# python
别用msys2的python，别找麻烦
尽管msys2很好，但是windows环境，很多软件默认的C编译器的msvc
依赖会找mscv去编译,比如matplotlib
https://github.com/msys2/MINGW-packages/issues/18192
https://www.msys2.org/docs/python/#known-issues
当然你可以使用pacman -S mingw-w64-ucrt-x86_64-python-matplotlib

只是所有涉及编译的依赖都要这样，会很麻烦
