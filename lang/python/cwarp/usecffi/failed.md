OSError: cannot load library 'wrapper.dll': error 0x7e.  Additionally, ctypes.util.find_library() did not manage to locate a library called 'wrapper.dll'

看来就是运行时问题
msvc python 加载不了 mingw32 的库