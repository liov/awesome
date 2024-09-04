import cffi

ffi = cffi.FFI()

ffi.cdef("""
    int load_image_width(const char* image_path);
""")

# 加载Rust共享库
lib = ffi.dlopen("wrapper.dll")

# 定义add函数
def load_image_width(image_path:str):
    return lib.load_image_width(image_path.encode("utf-8"))

print(load_image_width(r'D:\work\1--light1.jpg'))