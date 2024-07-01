import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

def plot_pointcloud_gray(points):
    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')

    x = points[:, 0]
    y = points[:, 1]
    z = points[:, 2]
    gray = points[:, 3]

    # Normalize the gray scale values to [0, 1] for visualization
    gray_normalized = (gray - gray.min()) / (gray.max() - gray.min())

    # Plot the point cloud
    scatter = ax.scatter(x, y, z, c=gray_normalized, cmap='gray', s=50)

    # Create a colorbar
    cbar = plt.colorbar(scatter)
    cbar.set_label('Gray Scale')

    ax.set_xlabel('X')
    ax.set_ylabel('Y')
    ax.set_zlabel('Z')

    plt.show()

def parse_pointcloud_gray_file(file_path):
    points = []

    with open(file_path, 'r') as file:
        # 读取文件头信息（在此示例中，我们暂时忽略它）
        while True:
            line = file.readline()
            if not line:
                break
            if line.startswith('size'):
                key, value = line.strip().split(':')
                num_points = int(value.strip())
                points = np.empty((num_points, 4))
            if line.startswith('data'):
                break

        # 读取数据部分
        idx = 0
        for line in file:
            data = line.strip().split(' ')
            x, y, z, gray = map(float, data)
            points[idx] = [x, y, z, gray]
            idx += 1

    return np.array(points)

# Example usage with dummy data
points = parse_pointcloud_gray_file("D:/work/sdk_save_cloud_csharp.pcp")  # Generate random points in the range [0, 100]
plot_pointcloud_gray(points)