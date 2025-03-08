import cv2
import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# 读取JPG图像
image = cv2.imread(r"D:\work\imageBytes.jpg")
if image is None:
    print("Could not open or find the JPG image!")
    exit()

# 读取TIFF高度图
height_map = cv2.imread(r"D:\work\imageBytes3d.tiff", cv2.IMREAD_GRAYSCALE)
if height_map is None:
    print("Could not open or find the TIFF height map!")
    exit()

# 确保图像和高度图的尺寸相同
if image.shape[:2] != height_map.shape:
    print("The dimensions of the image and height map do not match!")
    exit()

# 归一化高度图到0-255范围
height_map_normalized = cv2.normalize(height_map, None, 0, 255, cv2.NORM_MINMAX)

# 创建3D坐标网格
x = np.arange(0, image.shape[1])
y = np.arange(0, image.shape[0])
x, y = np.meshgrid(x, y)

# 创建3D图形
fig = plt.figure(figsize=(10, 10))
ax = fig.add_subplot(111, projection='3d')

# 绘制3D表面
ax.plot_surface(x, y, height_map_normalized, facecolors=cv2.cvtColor(image, cv2.COLOR_BGR2RGB) / 255.0, rstride=1, cstride=1)

# 设置轴标签
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Height')

# 显示图形
plt.show()
