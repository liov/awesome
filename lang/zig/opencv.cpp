#include <opencv2/opencv.hpp>

int main() {
    cv::Mat image = cv::imread("test.jpg", cv::IMREAD_COLOR);
    if (image.empty()) {
        std::cerr << "Could not open or find the image!" << std::endl;
        return -1;
    }
    cv::imshow("Window", image);
    cv::waitKey(0);
    return 0;
}
