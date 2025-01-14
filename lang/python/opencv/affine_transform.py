import cv2
import numpy as np

# 读取图像
#image = cv2.imread(r'D:\1--light2.jpg')
# 定义源点和目标点
src_points = np.float32([[128.08327894389708, 13.295278943897083], [123.16627894389707, 24.473278943897085], [110.63627894389708, 17.256278943897083]])  # 四个顶点的坐标
dst_points = np.float32([[26.525, (10.375+9.95)/2], [21.475, (21.425+21.175)/2],[(8.825+9.175)/2, 14]])  # 目标四个顶点

# 计算仿射变换矩阵
matrix = cv2.getAffineTransform(src_points, dst_points)
print(matrix)
# 定义要变换的点 (例如 [100, 150])
point = np.array([128.08327894389708, 13.295278943897083, 1], dtype=np.float32).reshape(-1, 1)

# 应用仿射变换矩阵到点
transformed_point = matrix @ point

# 将齐次坐标转换回普通坐标
transformed_point = transformed_point.flatten()[:2]
print(transformed_point)
# 应用仿射变换
#warped_image = cv2.warpAffine(image, matrix, (image.shape[1], image.shape[0]))

#cv2.imwrite(r'D:\1--light2-r.jpg', warped_image)
