# cgo 不能直接引用c++的头文件#include<opencv2/opencv.hpp>

In file included from D:/sdk/msys64/ucrt64/include/opencv4/opencv2/opencv.hpp:52,
from .\opencv.go:8:
D:/sdk/msys64/ucrt64/include/opencv4/opencv2/core.hpp:49:4: error: #error core.hpp header must be compiled as C++
49 | #  error core.hpp header must be compiled as C++
|    ^~~~~
In file included from D:/sdk/msys64/ucrt64/include/opencv4/opencv2/core.hpp:53:
D:/sdk/msys64/ucrt64/include/opencv4/opencv2/core/base.hpp:49:4: error: #error base.hpp header must be compiled as C++
49 | #  error base.hpp header must be compiled as C++
|    ^~~~~
D:/sdk/msys64/ucrt64/include/opencv4/opencv2/core/base.hpp:54:10: fatal error: climits: No such file or directory
54 | #include <climits>
|          ^~~~~~~~~
compilation terminated.

opencv only C++ API. no C API.
cgo 本身不能直接调用 C++ API 接口的头文件，因为它主要用于在 Go 代码中调用 C 语言编写的库

要在 Go 代码中使用 C++ API，你需要创建一个 C 语言的包装器，该包装器将调用 C++ API 并在 C++ 代码中封装它们。然后，你可以在 Go 代码中使用 cgo 调用这个 C 包装器。
这个包装怎么写呢,C的函数签名,实现调用原来的c++api
## wrap 文件名的问题
当一个go文件名字为a.go时，cgo会自动寻找a.c或者a.cc或者a.cpp文件，如果找不到，则报错。undefined reference to `xxx'
因此请确保你的实现包装和go文件同名