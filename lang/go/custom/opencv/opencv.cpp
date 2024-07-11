#include "wrapper.h"

int load_image_width(const char* image_path) {
    cv::Mat image = cv::imread(image_path, cv::IMREAD_COLOR);
    if (image.empty()) {
        return -1; // 加载失败
    }
    return image.cols;
}
