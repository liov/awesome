import itk
import vtk

# 读取 3D TIFF 图像
image = itk.imread(r'D:\work\hello.tiff', itk.F)


# 将 ITK 图像转换为 NumPy 数组
image_array = itk.GetArrayViewFromImage(image)

# 打印图像的形状以进行调试
print("Image shape:", image_array.shape)

# 检查图像的维度
if len(image_array.shape) == 3:
    depth, height, width = image_array.shape
    print(depth, height, width)
elif len(image_array.shape) == 2:
    height, width = image_array.shape
    depth = 1  # 如果是 2D 图像，则将深度设为 1
    print(depth, height, width)
else:
    raise ValueError("Unsupported image shape: {}".format(image_array.shape))

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
