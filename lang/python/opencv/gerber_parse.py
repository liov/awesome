import cv2
import numpy as np

# 读取图像并转换为灰度图
image = cv2.imread(r"D:\work\1a25d9a8d9b9ef124847610c7debb1e7.png")
gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# 二值化图像
_, binary = cv2.threshold(gray, 1, 255, cv2.THRESH_BINARY)

# 查找轮廓（联通区域）
contours, _ = cv2.findContours(binary, cv2.RETR_LIST, cv2.CHAIN_APPROX_SIMPLE)
print("Found %d objects" % len(contours))
# 初始化变量保存最大的矩形框
max_area = 0
max_rect = None
second_rect = None
# 在原图上绘制矩形框
for contour in contours:
    # 计算轮廓的外接矩形
    x, y, w, h = cv2.boundingRect(contour)
    area = w * h
    if area > max_area:
        max_area = area
        second_rect = max_rect
        max_rect = (x, y, w, h)
    # 提取矩形区域的像素
    roi = binary[y:y+h, x:x+w]
    nonzero_pixels = cv2.countNonZero(roi)
    # 计算像素比值
    fill_ratio = nonzero_pixels / area
    if area < 10000 and w != 1 and h != 1 :
        # 绘制矩形框（绿色，线宽为2）
        cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 1)
print(f"second {second_rect}")
h,w = image.shape[:2]
x, y, w, h = second_rect
print(x, y, w, h)
cropped_image = image[y:y+h, x:x+w]
target=binary[y:y+h, x:x+w]
# 高斯模糊，减少噪声
blurred = cv2.GaussianBlur(target, (9, 9), 0)

# 使用 Hough Circle Transform 检测圆
circles = cv2.HoughCircles(
    blurred,
    cv2.HOUGH_GRADIENT,  # 检测方法
    dp=1,                # 累加器分辨率与图像分辨率的反比
    minDist=200,          # 圆之间的最小距离
    param1=30,           # 边缘检测的高阈值（Canny 的参数）
    param2=30,           # 累加器阈值（越小检测越多假阳性）
    minRadius=10,         # 最小半径
    maxRadius=10         # 最大半径
)
print(f"Found {len(circles[0, :])} circles")
if circles is not None:
    circles = np.int32(np.around(circles))  # 四舍五入并转为整数
    for circle in circles[0, :]:
        x, y, radius = circle
        print(f"Circle at ({x}, {y}) with radius {radius}")

        roi = target[max(y-radius,0):y+radius, max(x-radius,0):x+radius]

        total_pixels = (radius*2) * (radius*2)
        nonzero_pixels = cv2.countNonZero(roi)
         # 计算像素比值
        fill_ratio = nonzero_pixels / total_pixels
        print(fill_ratio)
        if fill_ratio >= 0.7:
            cv2.circle(cropped_image, (x, y), 1, (255, 0, 0), 1)  # 绘制圆
            cv2.circle(cropped_image, (x, y), radius, (0, 0, 255), 1) # 绘制圆心

cv2.imwrite("output.png", cropped_image)

