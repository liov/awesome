import cv2
import numpy as np

# 自动获取电路板图片的主色
def get_dominant_color(image, k=2, resize_factor=0.1):
    # 将图片缩放到较小的分辨率
    small_image = cv2.resize(image, (0, 0), fx=resize_factor, fy=resize_factor, interpolation=cv2.INTER_AREA)
    # 将图片从 HxWxC 展平为 Nx3（每一行是一个像素的 RGB 值）
    pixels = small_image.reshape(-1, 3).astype(np.float32)

    # 定义 KMeans 参数
    criteria = (cv2.TERM_CRITERIA_EPS + cv2.TERM_CRITERIA_MAX_ITER, 100, 0.2)
    _, labels, centers = cv2.kmeans(pixels, k, None, criteria, 10, cv2.KMEANS_RANDOM_CENTERS)

    # 统计每个聚类的像素数，选择最大类别的中心作为主色
    dominant_color = centers[np.argmax(np.bincount(labels.flatten()))]
    return dominant_color.astype(np.uint8)

# 输入图片
image_path = r"D:\work.jpg"
image = cv2.imread(image_path)

# 将图片转换为 RGB 格式
image_rgb = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)

# 获取电路板的主色
dominant_color = get_dominant_color(image_rgb)
print(f"主色: {dominant_color}")

# 生成二值化图像，区分阻焊层和其他区域
# 计算与主色的差异
difference = cv2.absdiff(image_rgb, dominant_color)
gray_diff = cv2.cvtColor(difference, cv2.COLOR_RGB2GRAY)

# 自动计算阈值进行二值化
_, binary_mask = cv2.threshold(gray_diff, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)

# 形态学操作：腐蚀后膨胀
kernel1 = cv2.getStructuringElement(cv2.MORPH_RECT, (3, 3))  # 定义 3x3 的矩形核
kernel2 = cv2.getStructuringElement(cv2.MORPH_RECT, (3, 3))  # 定义 3x3 的矩形核
binary_mask_eroded = cv2.erode(binary_mask, kernel1, iterations=1)  # 腐蚀，减少噪点
binary_mask_dilated = cv2.dilate(binary_mask_eroded, kernel2, iterations=1)  # 膨胀，恢复形状

# 保存中间步骤的图像
cv2.imwrite("04_binary_mask.jpg", binary_mask)  # 二值化后的图像
cv2.imwrite("05_binary_eroded.jpg", binary_mask_eroded)  # 腐蚀后的图像
cv2.imwrite("06_binary_dilated.jpg", binary_mask_dilated)  # 膨胀后的图像
