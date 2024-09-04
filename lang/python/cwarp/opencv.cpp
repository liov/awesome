#include "wrapper.h"

int load_image_width(const char* image_path) {
    cv::Mat image = cv::imread(image_path, cv::IMREAD_COLOR);
    if (image.empty()) {
        return -1; // 加载失败
    }
    return image.cols;
}
// 神奇,真坑
// 失败：g++ -shared -fPIC `pkg-config --cflags --libs opencv4` -o wrapper.dll opencv.cpp
// cpp:(.text+0x5f): undefined reference to `cv::imread(std::__cxx11::basic_string<char, std::char_traits<char>, std::allocator<char> > const&, int)'

// 成功  g++ -shared -fPIC -o wrapper.dll opencv.cpp  `pkg-config --cflags --libs opencv4`