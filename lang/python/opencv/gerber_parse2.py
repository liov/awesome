import cv2
import numpy as np

def preprocess(image):
    # 转换为 HSV 色彩空间
    hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)
    cv2.imwrite("hsv.png", hsv)
    # 定义黑色背景的颜色范围
    lower_black = np.array([0, 0, 0])  # HSV下限
    upper_black = np.array([180, 255, 213])  # HSV上限

    # 创建掩码，去除黑色背景
    mask = cv2.inRange(hsv, lower_black, upper_black)
    mask_inverted = cv2.bitwise_not(mask)  # 取反，保留非黑色部分

    # 应用掩码，去除底版黑色
    background_removed = cv2.bitwise_and(image, image, mask=mask_inverted)

    # 转换为灰度图
    gray = cv2.cvtColor(background_removed, cv2.COLOR_BGR2GRAY)
    cv2.imwrite("gray.png", gray)
    # 二值化
    _, binary = cv2.threshold(gray, 50, 255, cv2.THRESH_BINARY)
    # 定义膨胀核
    kernel_size = (5, 5)  # 核的大小可以根据实际情况调整
    kernel = cv2.getStructuringElement(cv2.MORPH_RECT, kernel_size)

    # 可选：平滑处理（腐蚀操作）
    opening = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel,)
    closing = cv2.morphologyEx(opening, cv2.MORPH_CLOSE, kernel, iterations=1)
    cv2.imwrite("smoothed.png", closing)
    return closing

# 读取图像并转换为灰度图
image = cv2.imread(r"D:\work.jpg", cv2.IMREAD_COLOR)
smoothed=preprocess(image)
h,w = image.shape[:2]
x, y = 0,0
print(x, y, w, h)

# 查找轮廓（联通区域）
contours, _ = cv2.findContours(smoothed, cv2.RETR_LIST, cv2.CHAIN_APPROX_SIMPLE)
print("Found %d objects" % len(contours))
# 初始化变量保存最大的矩形框

# 在原图上绘制矩形框
for contour in contours:
    # 计算轮廓的外接矩形
    x, y, w, h = cv2.boundingRect(contour)
    rect = cv2.minAreaRect(contour)
    area = w * h
    # 提取矩形区域的像素
    nonzero_pixels = cv2.countNonZero(smoothed[y:y+h, x:x+w])
    # 计算像素比值
    box = cv2.boxPoints(rect)
    box = np.intp(box)  # 将浮点数转换为整数
    fill_ratio = float(nonzero_pixels) / float(area)

    if 35000 > area > 225 and fill_ratio > 0.3:
        #print(f"Found rectangle at ({x}, {y}) with area {area} and fill ratio {fill_ratio}")
        # 绘制矩形框（绿色，线宽为2）
        #cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 1)
        cv2.polylines(image, [box], isClosed=True, color=(255, 0, 0), thickness=2)
# 高斯模糊，减少噪声
blurred = cv2.GaussianBlur(smoothed, (9, 9), 0)

# 使用 Hough Circle Transform 检测圆
circles = cv2.HoughCircles(
    blurred,
    cv2.HOUGH_GRADIENT,  # 检测方法
    dp=1,                # 累加器分辨率与图像分辨率的反比
    minDist=200,          # 圆之间的最小距离
    param1=30,           # 边缘检测的高阈值（Canny 的参数）
    param2=30,           # 累加器阈值（越小检测越多假阳性）
    minRadius=49,         # 最小半径
    maxRadius=51         # 最大半径
)
#print(f"Found {len(circles)} circles")
if circles is not None:
    circles = np.int32(np.around(circles))  # 四舍五入并转为整数
    for circle in circles[0, :]:
        x, y, radius = circle
        print(f"Circle at ({x}, {y}) with radius {radius}")

        roi = smoothed[max(y-radius,0):y+radius, max(x-radius,0):x+radius]

        total_pixels = (radius*2) * (radius*2)
        nonzero_pixels = cv2.countNonZero(roi)
        # 计算像素比值
        fill_ratio = nonzero_pixels / total_pixels
        print(fill_ratio)
        if fill_ratio >= 0.7:
            cv2.circle(image, (x, y), 1, (255, 0, 0), 1)  # 绘制圆
            cv2.circle(image, (x, y), radius, (0, 0, 255), 1) # 绘制圆心

cv2.imwrite("output.png", image)

