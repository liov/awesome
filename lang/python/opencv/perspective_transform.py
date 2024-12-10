import cv2
import numpy as np

# 读取图片
image = cv2.imread(r'D:\1--light1.jpg')

# 定义源点和目标点
src_points = np.float32([[309, 378], [4110, 391], [301, 2981], [4103, 3153]])  # 四个顶点的坐标
dst_points = np.float32([[310, 378], [4110, 378],[310, 2972],  [4110, 3132]])  # 目标四个顶点
# 计算透视变换矩阵
matrix = cv2.getPerspectiveTransform(src_points, dst_points)

# 应用透视变换
warped_image = cv2.warpPerspective(image, matrix, (4600, 3800))

cv2.imwrite(r'D:\1--light1-r.jpg', warped_image)
