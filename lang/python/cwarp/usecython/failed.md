反复尝试，失败，不搞了
Traceback (most recent call last):
File "D:\code\hopeio\hoper\awesome\lang\python\cwarp\usecython\test.py", line 5, in <module>
import wrapper
ImportError: DLL load failed while importing wrapper: 找不到指定的模块。
搜索不到pyd依赖的dll
python setup.py build_ext --inplace --compiler=mingw32
objdump -p wrapper.cp312-win_amd64.pyd|grep dll

神器https://github.com/lucasg/Dependencies/releases
N/A, 0 (0x00000000), QueryOOBESupport, ext-ms-win-oobe-query-l1-1-0.dll, False, None
https://github.com/ProarchwasTaken/pybind_test/blob/6bd59668cde03b58ccc65883f2111c5a6b4dc1b4/README.md
不搞了，https://github.com/AcademySoftwareFoundation/Imath/issues/238

思路错了啊
应该用rust+python啊
要不就是直接ffi
凡是需要手写动态库的都应该rust or zig 啊
mingw 就是配合cgo的啊