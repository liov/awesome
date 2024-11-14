import cv2
import numpy as np



def calculate_contrast(path):
    # 读取图像，转换为灰度图
    image = cv2.imread(path, cv2.IMREAD_GRAYSCALE)

    # 计算图像像素亮度的标准差
    min_brightness = np.min(image)
    max_brightness = np.max(image)
    contrast = max_brightness - min_brightness
    print(f"图片对比度（标准差）：{contrast}")
    if (max_brightness + min_brightness) != 0:
        contrast = (max_brightness - min_brightness) / (max_brightness + min_brightness)
    else:
        contrast = 0
    print(f"图片对比度（Michelson）：{contrast}")

calculate_contrast("./images/lena.jpg")