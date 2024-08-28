import cv2
import numpy as np

# 读取图像
image = cv2.imread(r'D:\work\1--light1.jpg')
image = image[4816:5344, 3263:4162]

# 获取图像的宽度和高度
height, width = image.shape[:2]

# 转换为灰度图
gray_image = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# 使用高斯模糊平滑图像
blurred = cv2.GaussianBlur(gray_image, (9, 9), 2)

# 使用Canny边缘检测
edges = cv2.Canny(blurred, 50, 150)

# 使用HoughCircles检测圆形
circles = cv2.HoughCircles(edges, cv2.HOUGH_GRADIENT, dp=1.2, minDist=300, param1=200, param2=80, minRadius=10,
                           maxRadius=500)

# 确保检测到了圆形
if circles is not None:
    # 转换圆心和半径为整数
    circles = np.round(circles[0, :]).astype("int")

    # 逐个处理检测到的圆形
    for (x, y, r) in circles:
        # 检查圆是否完整，即圆的边缘不会超出图像边界
        if (x - r) > 0 and (x + r) < width and (y - r) > 0 and (y + r) < height:
            # 获取圆形边缘区域
            edge_mask = np.zeros_like(gray_image)
            cv2.circle(edge_mask, (x, y), r, 255, 2)  # 边缘厚度为2像素

            # 计算边缘区域的平均亮度
            mean_edge = cv2.mean(gray_image, mask=edge_mask)[0]

            # 计算圆外区域的平均亮度
            outer_mask = np.zeros_like(gray_image)
            cv2.circle(outer_mask, (x, y), r + 10, 255, -1)  # 比圆稍大一点的区域
            outer_mask = cv2.subtract(outer_mask, edge_mask)
            mean_outside = cv2.mean(gray_image, mask=outer_mask)[0]

            # 保留边缘和外部对比度高的圆
            if abs(mean_edge - mean_outside) > 0:  # 根据需要调整此阈值
                # 绘制外圆
                cv2.circle(image, (x, y), r, (0, 255, 0), 2)
                # 绘制圆心
                cv2.circle(image, (x, y), 2, (0, 0, 255), 3)

# 保存结果图像
cv2.imwrite('filtered_circles.jpg', image)
