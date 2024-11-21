import cv2
import numpy as np

# 读取图片
image = cv2.imread(r"xxx.jpg", cv2.IMREAD_GRAYSCALE)

# Sobel 算子计算梯度
sobel_x = cv2.Sobel(image, cv2.CV_64F, 1, 0, ksize=3)
sobel_y = cv2.Sobel(image, cv2.CV_64F, 0, 1, ksize=3)

# 计算梯度方向
angles = np.arctan2(sobel_y, sobel_x) * 180 / np.pi

# 过滤有效的角度值
valid_angles = angles[np.abs(angles) < 45]  # 只考虑小倾斜角度范围
average_angle = np.mean(valid_angles)

print(f"检测到的倾斜角度: {average_angle:.6f}°")
