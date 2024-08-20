import numpy as np
import matplotlib.pyplot as plt

# 定义一个函数来绘制旋转矩形
def draw_rotated_rectangle(center, width, height, angle):
    # 将角度转换为弧度
    theta = np.radians(angle)

    # 计算矩形的四个顶点
    x = [center[0] - width/2, center[0] + width/2, center[0] + width/2, center[0] - width/2]
    y = [center[1] - height/2, center[1] - height/2, center[1] + height/2, center[1] + height/2]

    # 应用旋转变换
    x_rot = [center[0] + (xi - center[0]) * np.cos(theta) - (yi - center[1]) * np.sin(theta) for xi, yi in zip(x, y)]
    y_rot = [center[1] + (xi - center[0]) * np.sin(theta) + (yi - center[1]) * np.cos(theta) for xi, yi in zip(x, y)]

    # 绘制矩形
    plt.plot(x_rot + [x_rot[0]], y_rot + [y_rot[0]], 'r-')

    # 标出四个顶点的坐标
    for i in range(4):
        plt.text(x_rot[i], y_rot[i], f'({x_rot[i]:.2f}, {y_rot[i]:.2f})', fontsize=9, ha='center')

# 设置画布大小
plt.figure(figsize=(8, 6))

# 绘制旋转矩形
draw_rotated_rectangle((0.5, 0.5), 0.4, 0.2, 30)

# 在矩形内外生成一些点
np.random.seed(0)  # 设置随机种子以保证结果可重复
points_inside = np.random.uniform(low=[0.1, 0.1], high=[0.9, 0.9], size=(5, 2))
points_outside = np.random.uniform(low=[0, 0], high=[1, 1], size=(5, 2))

# 绘制点并标注坐标
for x, y in points_inside:
    plt.plot(x, y, 'bo', markersize=5)
    plt.text(x, y, f'({x:.2f}, {y:.2f})', fontsize=9, ha='right')

for x, y in points_outside:
    plt.plot(x, y, 'go', markersize=5)
    plt.text(x, y, f'({x:.2f}, {y:.2f})', fontsize=9, ha='right')

# 设置坐标轴范围和比例
plt.xlim(0, 1)
plt.ylim(0, 1)
plt.gca().set_aspect('equal', adjustable='box')

# 显示图形
plt.show()