import cv2
import numpy as np

# 读取图像
image = cv2.imread(r"D:\微信图片_20241209171356.jpg")

regionCenterX, regionCenterY = (4308,820)
regionWidth = 100

image = image[regionCenterY-regionWidth:regionCenterY+regionWidth, regionCenterX-regionWidth:regionCenterX+regionWidth]

# 转换为灰度图
gray_image = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# 使用高斯模糊平滑图像
blurred = cv2.GaussianBlur(gray_image, (9, 9), 0.0)


# 使用HoughCircles检测圆形
circles = cv2.HoughCircles(blurred, cv2.HOUGH_GRADIENT, dp=1, minDist=50, param1=30, param2=30, minRadius=1,
                           maxRadius=300)

# 确保检测到了圆形
if circles is not None:
    # 转换圆心和半径为整数
    circles = np.round(circles[0, :]).astype("int")

    # 逐个处理检测到的圆形
    for (x, y, r) in circles:
        print(f"Circle center: ({regionCenterX-regionWidth+x}, {regionCenterY-regionWidth+y}), radius: {r}")
        # 绘制外圆
        cv2.circle(image, (x, y), r, (0, 255, 0), 2)
        # 绘制圆心
        cv2.circle(image, (x, y), 1, (0, 0, 255), 1)

# 保存结果图像
cv2.imwrite('filtered_circles.jpg', image)
