import cv2
import numpy as np

# 加载你的图像
image = cv2.imread(r"D:\work.jpg", cv2.IMREAD_COLOR)
w, h, _ = image.shape
size = w * h
# 将图像从BGR颜色空间转换为HSV颜色空间
hsv_image = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)

# 分离H,S,V通道，这里我们只关注V(亮度)通道
_, _, v_channel = cv2.split(hsv_image)

# 计算V通道的平均值
average_v = np.mean(v_channel)
print(f"平均V值: {average_v}")

# 统计每个V值的数量，并按数量从大到小排序
v_counts = np.bincount(v_channel.ravel())
sorted_v_counts = np.argsort(-v_counts)

# 输出前几个V值及其出现次数
for v_value in sorted_v_counts[:255]:  # 只输出最常见的10个V值
    print(f"V值: {v_value}, 数量: {v_counts[v_value]}, 占比: {v_counts[v_value] / size * 100:.2f}%")


def find_pads_color(hsv_image):

    # 分离H,S,V通道
    h, s, v = cv2.split(hsv_image)

    # 使用自适应阈值方法创建二值图像，以分离出高亮区域
    _, binary_v = cv2.threshold(v, 200, 255, cv2.THRESH_BINARY)  # 调整阈值参数以适应你的图像

    # 形态学操作 - 膨胀和腐蚀，去除小的噪点
    kernel = np.ones((5, 5), np.uint8)
    binary_v = cv2.dilate(binary_v, kernel, iterations=1)
    binary_v = cv2.erode(binary_v, kernel, iterations=1)

    # 根据二值图像掩码提取焊盘区域的颜色
    pads_mask = binary_v == 255
    pads_colors = hsv_image[pads_mask]

    if pads_colors.size == 0:
        print("没有检测到焊盘区域")
        return None

    # 将颜色数据调整为适合KMeans输入的格式
    reshaped_colors = pads_colors.reshape((-1, 3)).astype(np.float32)

    # 应用KMeans聚类找到焊盘的主要颜色
    criteria = (cv2.TERM_CRITERIA_EPS + cv2.TERM_CRITERIA_MAX_ITER, 10, 1.0)
    _, labels, centers = cv2.kmeans(reshaped_colors, K=1, bestLabels=None, criteria=criteria, attempts=10, flags=cv2.KMEANS_RANDOM_CENTERS)

    dominant_color_hsv = centers[0].astype(int)
    # 如果需要，可以将主色转换回BGR进行可视化
    dominant_color_bgr = cv2.cvtColor(np.array([[dominant_color_hsv]], dtype=np.uint8), cv2.COLOR_HSV2BGR)[0][0]
    print(f"焊盘颜色(HSV): {dominant_color_hsv}")
    print(f"焊盘颜色(BGR): {dominant_color_bgr}")

    # 可视化结果
    output_image = image.copy()
    output_image[~pads_mask] = [0, 0, 0]  # 非焊盘区域设为黑色

    return dominant_color_hsv, dominant_color_bgr

# 使用函数获取焊盘颜色
find_pads_color(hsv_image)