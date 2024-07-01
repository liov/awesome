from PIL import Image
import numpy as np
import vtk

# 读取 3D TIFF 图像r'D:\work\hello.tiff'
tiff_path =
tiff_img = Image.open(tiff_path)

# 检查图像的帧数（切片数）
depth = tiff_img.n_frames
print("Number of slices:", depth)
width, height = tiff_img.size

# 创建一个空的 NumPy 数组来存储 3D 图像数据
image_array = np.zeros((depth, height, width), dtype=np.float32)

# 逐帧读取图像数据
for i in range(depth):
    tiff_img.seek(i)
    image_array[i, :, :] = np.array(tiff_img)

# 打印图像的形状以确认是 3D 图像
print("Image shape:", image_array.shape)

# 将 NumPy 数组转换为 VTK 图像数据
vtk_image_data = vtk.vtkImageData()
vtk_image_data.SetDimensions(width, height, depth)
vtk_image_data.AllocateScalars(vtk.VTK_FLOAT, 1)

# 将 NumPy 数组数据复制到 VTK 图像数据
for z in range(depth):
    for y in range(height):
        for x in range(width):
            vtk_image_data.SetScalarComponentFromFloat(x, y, z, 0, image_array[z, y, x])

# 创建 VTK 渲染器
renderer = vtk.vtkRenderer()

# 创建 VTK 渲染窗口
render_window = vtk.vtkRenderWindow()
render_window.AddRenderer(renderer)

# 创建 VTK 渲染窗口交互器
render_window_interactor = vtk.vtkRenderWindowInteractor()
render_window_interactor.SetRenderWindow(render_window)

# 创建 VTK 数据映射器
mapper = vtk.vtkDataSetMapper()
mapper.SetInputData(vtk_image_data)

# 创建 VTK 演员
actor = vtk.vtkActor()
actor.SetMapper(mapper)

# 将演员添加到渲染器
renderer.AddActor(actor)

# 开始渲染循环
render_window.Render()
render_window_interactor.Start()
