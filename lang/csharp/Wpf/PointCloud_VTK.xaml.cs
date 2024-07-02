using System;
using System.Windows;
using Kitware.VTK;
using System.Windows.Controls;
using System.Windows.Forms.Integration;

namespace Wpf;

    public partial class PointCloud_VTK : Window
    {
        public PointCloud_VTK()
        {
            InitializeComponent();
            Render3DTiff();
        }

        private void Render3DTiff()
        {
            // 读取 TIFF 图像数据
            string tiffFilePath = "D:/work/pic/1.tiff";
            vtkTIFFReader reader = vtkTIFFReader.New();
            reader.SetFileName(tiffFilePath);
            reader.Update();

            // 将 TIFF 图像数据转换为 3D 模型
            vtkImageData imageData = reader.GetOutput();
            vtkMarchingCubes marchingCubes = vtkMarchingCubes.New();
            marchingCubes.SetInputConnection(imageData.GetProducerPort());
            marchingCubes.SetValue(0, 128); // 设置阈值
            marchingCubes.Update();

            // 创建渲染器和渲染窗口
            vtkRenderer renderer = vtkRenderer.New();
            vtkRenderWindow renderWindow = vtkRenderWindow.New();
            renderWindow.AddRenderer(renderer);

            // 创建交互器
            vtkRenderWindowInteractor renderWindowInteractor = vtkRenderWindowInteractor.New();
            renderWindowInteractor.SetRenderWindow(renderWindow);

            // 创建演员并添加到渲染器
            vtkPolyDataMapper mapper = vtkPolyDataMapper.New();
            mapper.SetInputConnection(marchingCubes.GetOutputPort());
            vtkActor actor = vtkActor.New();
            actor.SetMapper(mapper);
            renderer.AddActor(actor);
            renderer.SetBackground(0.1, 0.2, 0.4); // 设置背景颜色

            // 在 WPF 中显示
            WindowsFormsHost host = new WindowsFormsHost();
            System.Windows.Forms.Panel panel = new System.Windows.Forms.Panel();
            host.Child = panel;
            renderWindow.SetParentId(panel.Handle);
            renderWindow.SetSize(800, 600);

            // 添加到 WPF 窗口中
            this.Content = host;

            // 启动渲染和交互
            renderWindow.Render();
            renderWindowInteractor.Start();
        }
    }

