import cv2
import numpy as np

image = cv2.imread(r"D:\result.jpg")

gray = image[:, :, 0]
_, binary = cv2.threshold(gray, 40, 255, cv2.THRESH_BINARY)
blurred = cv2.GaussianBlur(binary, (5, 5), 0)

# 边缘检测
edged = cv2.Canny(blurred, 50, 150)

# 寻找轮廓
contours, _ = cv2.findContours(edged, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

# 筛选矩形
rectangles = []
for contour in contours:
    # 近似轮廓
    epsilon = 0.02 * cv2.arcLength(contour, True)
    approx = cv2.approxPolyDP(contour, epsilon, True)
    rectangles.append(approx)

for rect in rectangles:
    cv2.drawContours(image, [rect], -1, (0, 255, 0), 2)

cv2.imwrite("gray.jpg", gray)
cv2.imwrite("output.jpg", image)
cv2.imwrite("blurred.jpg", blurred)
cv2.imwrite("binary.jpg", binary)