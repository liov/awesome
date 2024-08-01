import cv2
import numpy as np

def order_points(pts):
    # 初始化点的顺序：左上、右上、右下和左下
    rect = np.zeros((4, 2), dtype="float32")

    # 从左到右排序点，根据它们的x坐标值
    s = pts.sum(axis=1)
    rect[0] = pts[np.argmin(s)]
    rect[2] = pts[np.argmax(s)]

    diff = np.diff(pts, axis=1)
    rect[1] = pts[np.argmin(diff)]
    rect[3] = pts[np.argmax(diff)]

    return rect

def four_point_transform(image, pts):
    # 获取输入坐标点并顺序排列
    rect = order_points(pts)
    (tl, tr, br, bl) = rect

    # 计算新的图像宽度和高度
    widthA = np.linalg.norm(br - bl)
    widthB = np.linalg.norm(tr - tl)
    maxWidth = max(int(widthA), int(widthB))

    heightA = np.linalg.norm(tr - br)
    heightB = np.linalg.norm(tl - bl)
    maxHeight = max(int(heightA), int(heightB))

    # 设置变换后的目标点
    dst = np.array([
        [0, 0],
        [maxWidth - 1, 0],
        [maxWidth - 1, maxHeight - 1],
        [0, maxHeight - 1]], dtype="float32")

    # 计算透视变换矩阵并应用变换
    M = cv2.getPerspectiveTransform(rect, dst)
    warped = cv2.warpPerspective(image, M, (maxWidth, maxHeight))

    return warped

# 读取输入图像
image = cv2.imread("D:\\work\\4.1-单板截图.jpg")

# 手动定义倾斜矩形的四个角点（根据你的实际图像调整）
pts = np.array([[4800, 3000], [6600,4800], [4800, 6522], [3100, 4800],], dtype="float32")

# 校正矩形
warped = four_point_transform(image, pts)

# 保存校正后的图像
cv2.imwrite('output.jpg', warped)

# 显示原始和校正后的图像（可选）
#cv2.imshow('Original Image', image)
cv2.imshow('Warped Image', warped)
cv2.waitKey(0)
cv2.destroyAllWindows()
