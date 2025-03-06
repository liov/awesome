import cv2
import numpy as np

# 读取图片
image = cv2.imread(r"D:\work\test03052_top\gerber\gerber.jpg")
if image is None:
    print("Could not open or find the image!")
    exit()

# 定义矩形框的参数
x, y, w, h = 2208, 2933, 5120, 5120  # (x, y, width, height)

x1 = int(x - w // 2)
y1 = int(y - h // 2)
x2 = int(x + w // 2)
y2 = int(y + h // 2)

# 确保坐标在图像范围内
x1 = max(0, x1)
y1 = max(0, y1)
x2 = min(image.shape[1], x2)
y2 = min(image.shape[0], y2)
# 显示图片
cv2.imwrite('rect.jpg', image[y1:y2, x1:x2])