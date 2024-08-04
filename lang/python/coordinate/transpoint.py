import matplotlib.pyplot as plt
import numpy as np

# 定义一个函数来画箭头表示坐标轴
def draw_arrow(ax, start, end, color):
    ax.annotate(
        '', xy=end, xytext=start,
        arrowprops=dict(facecolor=color, edgecolor=color, arrowstyle='->', lw=2)
    )

# 创建图形和坐标轴
fig, ax = plt.subplots(figsize=(8, 8))
ax.set_xlim(0, 15)
ax.set_ylim(0, 15)
ax.set_aspect('equal', 'box')
ax.grid(True, which='both')
ax.invert_yaxis()  # 反转y轴，使其向下

# 坐标系a1
origin_a1 = np.array([0, 0])
x_axis_a1 = np.array([3, 0])
y_axis_a1 = np.array([0, 3])

# 坐标系a2
origin_a2 = np.array([5, 3])
angle = np.deg2rad(30)  # 顺时针旋转30度
x_axis_a2 = np.array([np.cos(angle), np.sin(angle)])
y_axis_a2 = np.array([-np.sin(angle), np.cos(angle)])

# 画出坐标系a1
draw_arrow(ax, origin_a1, origin_a1 + 5 * x_axis_a1, 'blue')
draw_arrow(ax, origin_a1, origin_a1 + 5 * y_axis_a1, 'blue')
ax.text(15, 0, 'X1', color='blue', fontsize=12)
ax.text(0, 15, 'Y1', color='blue', fontsize=12)
ax.text(0, 0, 'O1', color='blue', fontsize=12, ha='right')

# 画出坐标系a2
draw_arrow(ax, origin_a2, origin_a2 + 10 * x_axis_a2, 'red')
draw_arrow(ax, origin_a2, origin_a2 + 10 * y_axis_a2, 'red')
ax.text(origin_a2[0] + 10 * x_axis_a2[0], origin_a2[1] + 10 * x_axis_a2[1], 'X2', color='red', fontsize=12)
ax.text(origin_a2[0] + 10 * y_axis_a2[0], origin_a2[1] + 10 * y_axis_a2[1], 'Y2', color='red', fontsize=12)
ax.text(origin_a2[0], origin_a2[1], 'O2', color='red', fontsize=12, ha='right')

def draw_point(ax,x,y):
# 画出在a2中的点和转换后的在a1中的点
    point_a2 = np.array([x, y])
    point_a1 = origin_a2 + point_a2[0] * x_axis_a2 + point_a2[1] * y_axis_a2

    ax.plot([origin_a2[0], point_a1[0]], [origin_a2[1], point_a1[1]], 'k--')
    ax.plot(point_a1[0], point_a1[1], 'ro')
    ax.text(point_a1[0], point_a1[1], f'({point_a2[0]:.2f}, {point_a2[1]:.2f})', color='red', fontsize=12, ha='right')
    ax.text(point_a1[0], point_a1[1], f'({point_a1[0]:.2f}, {point_a1[1]:.2f})', color='blue', fontsize=12, ha='left')

draw_point(ax,4,2)
draw_point(ax,5,3)
# 显示图形
plt.title('Transformation between Coordinate Systems a1 and a2')
plt.xlabel('X')
plt.ylabel('Y')
plt.axhline(0, color='black',linewidth=0.5)
plt.axvline(0, color='black',linewidth=0.5)
plt.grid(color = 'gray', linestyle = '--', linewidth = 0.5)
plt.show()
