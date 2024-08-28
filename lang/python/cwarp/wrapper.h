#ifndef OPENCV_WRAPPER_H
#define OPENCV_WRAPPER_H

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

// 示例：一个简单的函数，加载图像并返回其宽度
int load_image_width(const char* image_path);

#ifdef __cplusplus
}
#endif

#endif // OPENCV_WRAPPER_H
