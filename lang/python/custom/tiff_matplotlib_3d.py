import numpy as np
from PIL import Image
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# 读取灰度 TIFF 图像
tiff_path = r'D:\work\hello.tiff'
tiff_img = Image.open(tiff_path).convert('L')
image_array = np.array(tiff_img)

# 正常化高度值
z = image_array / np.max(image_array) * 100  # 将灰度值映射到 0-100 的范围

# 创建网格
x = np.linspace(0, image_array.shape[1], image_array.shape[1])
y = np.linspace(0, image_array.shape[0], image_array.shape[0])
x, y = np.meshgrid(x, y)

# 创建 3D 图
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

# 调整视角
ax.view_init(elev=60, azim=30)

# 绘制 3D 表面
surface = ax.plot_surface(x, y, z, cmap='viridis')

# 添加颜色条
fig.colorbar(surface, shrink=0.5, aspect=5)

# 显示图像
plt.show()
