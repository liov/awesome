import cv2
import numpy as np

# 读取图像
image = cv2.imread(r'D:\1--light2.jpg')

# 定义源点和目标点
src_points = np.float32([[309, 378], [4110, 391], [301, 2981]])  # 四个顶点的坐标
dst_points = np.float32([[310, 378], [4110, 378],[310, 2972]])  # 目标四个顶点

# 计算仿射变换矩阵
matrix = cv2.getAffineTransform(src_points, dst_points)

# 应用仿射变换
warped_image = cv2.warpAffine(image, matrix, (image.shape[1], image.shape[0]))

cv2.imwrite(r'D:\1--light2-r.jpg', warped_image)
