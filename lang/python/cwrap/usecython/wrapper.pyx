# cython: language_level = 3
# distutils: language = c++
# distutils: sources = ../opencv.cpp

cdef extern from "../wrapper.h":
    int load_image_width(const char *image_path)

def py_load_image_width(image_path):
    return load_image_width(image_path.encode("utf-8"))