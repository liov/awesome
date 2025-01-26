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

def process(image_rgb,remove_color,origin,size,op):

    # 生成二值化图像，区分阻焊层和其他区域
    # 计算与主色的差异
    difference = cv2.absdiff(image_rgb, remove_color)
    gray_diff = cv2.cvtColor(difference, cv2.COLOR_RGB2GRAY)

    # 自动计算阈值进行二值化
    _, binary = cv2.threshold(gray_diff, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)
    # 定义膨胀核
    kernel_size = size  # 核的大小可以根据实际情况调整
    kernel = cv2.getStructuringElement(cv2.MORPH_RECT, kernel_size)

    # 可选：平滑处理（腐蚀操作）
    #opening = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)
    binary = cv2.morphologyEx(binary,op , kernel)
    cv2.imwrite("binary.jpg", binary)

    # 查找轮廓（联通区域）
    contours, _ = cv2.findContours(binary, cv2.RETR_LIST, cv2.CHAIN_APPROX_SIMPLE)
    print("Found %d objects" % len(contours))
    # 初始化变量保存最大的矩形框

    # 在原图上绘制矩形框
    for contour in contours:
        # 计算轮廓的外接矩形
        x, y, w, h = cv2.boundingRect(contour)
        rect = cv2.minAreaRect(contour)
        area = w * h
        # 提取矩形区域的像素
        nonzero_pixels = cv2.countNonZero(binary[y:y+h, x:x+w])
        # 计算像素比值
        box = cv2.boxPoints(rect)
        box = np.intp(box)  # 将浮点数转换为整数
        fill_ratio = float(nonzero_pixels) / float(area)

        if 35000 > area > 225 and fill_ratio > 0.5:
            #print(f"Found rectangle at ({x}, {y}) with area {area} and fill ratio {fill_ratio}")
            # 绘制矩形框（绿色，线宽为2）
            #cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 1)
            cv2.polylines(origin, [box], isClosed=True, color=(255, 0, 0), thickness=2)
            cv2.fillPoly(image_rgb, [box], 0)
    return image

# 输入图片
image = cv2.imread(r"D:\work.jpg", cv2.IMREAD_COLOR)

# 将图片转换为 RGB 格式
image_rgb = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)

# 获取电路板的主色
dominant_color = get_dominant_color(image_rgb)
print(f"主色: {dominant_color}")
process(image_rgb,dominant_color,image,(5,5),cv2.MORPH_CLOSE)
cv2.imwrite("result.jpg", image)