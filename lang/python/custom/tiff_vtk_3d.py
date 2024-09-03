import numpy as np
from PIL import Image
import vtk
from vtkmodules.util import numpy_support

# 读取灰度 TIFF 图像
tiff_path = r'D:\result.tiff'
tiff_img = Image.open(tiff_path)

image_array = np.array(tiff_img)

# 使用 numpy 的 unique 函数获取唯一值
unique_count = len(np.unique(image_array))
print(unique_count)
# 计算灰度图像的最小值和最大值
min_value = np.nanmin(image_array)
max_value = np.nanmax(image_array)
print(min_value, max_value)

gary_range = 500
# 正常化灰度值到 [0, 1000]
normalized_image = ((image_array - min_value) / (max_value - min_value)) * gary_range
print(np.nanmax(normalized_image))
# 获取图像的尺寸
height, width = image_array.shape

# 将 NumPy 数组转换为 VTK 数组
vtk_data_array = numpy_support.numpy_to_vtk(num_array=normalized_image.ravel(), deep=True, array_type=vtk.VTK_FLOAT)

# 创建 VTK 图像数据对象
vtk_image_data = vtk.vtkImageData()
vtk_image_data.SetDimensions(width, height, 1)
vtk_image_data.SetSpacing(1, 1, 1)
vtk_image_data.GetPointData().SetScalars(vtk_data_array)

# 创建 VTK 高度图
surface = vtk.vtkWarpScalar()
surface.SetInputData(vtk_image_data)
surface.Update()

# 创建 VTK 等高线过滤器
contour_filter = vtk.vtkContourFilter()
contour_filter.SetInputConnection(surface.GetOutputPort())
contour_filter.GenerateValues(gary_range, 0,  gary_range)  # 生成20个等高线，范围从0.1到1.0

# 创建颜色映射器
color_range = 128
color_mapper = vtk.vtkLookupTable()
color_mapper.SetNumberOfTableValues(256)
color_mapper.SetRange(0, color_range)  # 设置颜色映射器的范围为 [0, 100]
color_mapper.Build()
for i in range(256):
    gray_value = i / 255.0 * color_range
    color_mapper.SetTableValue(i, gray_value, gray_value, gray_value, 0.5)


# 创建 VTK 映射器
mapper = vtk.vtkDataSetMapper()
mapper.SetInputConnection(surface.GetOutputPort())
mapper.SetLookupTable(color_mapper)
# mapper.SetScalarRange(0, 0)  # 设置灰度值范围

# 创建 VTK 演员
actor = vtk.vtkActor()
actor.SetMapper(mapper)

# 创建 VTK 渲染器
renderer = vtk.vtkRenderer()
renderer.AddActor(actor)
renderer.SetBackground(0.1, 0.2, 0.4)

# 创建 VTK 渲染窗口
render_window = vtk.vtkRenderWindow()
render_window.AddRenderer(renderer)

# 创建 VTK 渲染窗口交互器
render_window_interactor = vtk.vtkRenderWindowInteractor()
render_window_interactor.SetRenderWindow(render_window)

# 启动渲染
render_window.Render()
render_window_interactor.Start()
