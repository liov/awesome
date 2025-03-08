import cv2
import numpy as np
import pyvista as pv

# 读取JPG图像
image = cv2.imread(r"D:\work\memoImage3.png")
if image is None:
    print("Could not open or find the JPG image!")
    exit()

# 读取TIFF高度图
height_map = cv2.imread(r"D:\work\memoTiff3.tiff", cv2.IMREAD_GRAYSCALE)
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

# 创建3D网格
z = height_map_normalized
grid = pv.StructuredGrid(x, y, z)

# 设置颜色映射
colors = cv2.cvtColor(image, cv2.COLOR_BGR2RGB).reshape(-1, 3)
grid.point_data['colors'] = colors

# 创建3D图形
plotter = pv.Plotter()
plotter.add_mesh(grid, scalars='colors', rgb=True, interpolate_before_map=True)
plotter.show()
