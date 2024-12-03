import cv2
import numpy as np

# 读取图像并转换为灰度图
image = cv2.imread(r"D:\Download\gerbv-2.10.0-win64\bin\gerber1.png")
gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# 二值化图像
_, binary = cv2.threshold(gray, 1, 255, cv2.THRESH_BINARY)

# 查找轮廓（联通区域）
contours, hierarchy = cv2.findContours(binary, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
print("Found %d objects" % len(contours))

# 在原图上绘制矩形框
for i, contour in enumerate(contours):
    rect = cv2.minAreaRect(contour)
    # 获取矩形的四个顶点
    box = cv2.boxPoints(rect)
    box = np.intp(box)
    # 绘制斜矩形
    cv2.drawContours(image, [box], 0, (0, 255, 0), 2)
    # cv2.drawContours(image, contours, i, (0, 255, 0), 1)

cv2.imwrite("output.png", image)

