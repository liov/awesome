import cffi

ffi = cffi.FFI()

ffi.cdef("""
    int load_image_width(const char* image_path);
""")

# 加载Rust共享库
lib = ffi.dlopen("wrapper.dll")

# 定义add函数
@ffi.callback("int load_image_width(const char* image_path)")
def load_image_width(image_path):
    return lib.load_image_width(image_path)

load_image_width(r'D:\work\1--light1.jpg')