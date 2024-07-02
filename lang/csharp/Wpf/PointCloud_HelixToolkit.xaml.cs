using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Media.Media3D;
using HelixToolkit.Wpf;
using OpenCvSharp;
using Window = System.Windows.Window;
using Point = System.Windows.Point;

namespace Wpf;

public partial class PointCloud_HelixToolkit : Window
{
    public PointCloud_HelixToolkit()
    {
        InitializeComponent();
        Create3DSceneFromTiff("D:/work/pic/1.tiff", "D:/work/pic/1_R.png");
    }


    private void Create3DSceneFromTiff(string filePath, string imgPath)
    {
        // 读取 TIFF 文件
        var image = Cv2.ImRead(filePath, ImreadModes.Grayscale | ImreadModes.AnyDepth);

        var positions = new Point3DCollection();
        var triangleIndices = new Int32Collection();
        var textureCoordinates = new PointCollection();
        var normals = new Vector3DCollection();

        var offsetX = image.Width / 2;
        var offsetY = image.Height / 2;
        var zoom = 10;
        // 遍历图像的每个像素，将灰度值映射到高度
        for (var y = 0; y < image.Height; y += zoom)
        {
            for (var x = 0; x < image.Width; x += zoom)
            {
                var pixelValue = image.At<float>(y, x);

                // 添加顶点，将灰度值映射到高度范围 0 到 100
                var heightValue = Math.Floor(((3 - pixelValue) * 200.0));

                // 添加顶点
                positions.Add(new Point3D(x - offsetX, y - offsetY, heightValue));
                normals.Add(new Vector3D(0, 0, 1)); // 朝向正Z轴的法线向量
                // 添加贴图坐标
                double textureX = x / (double)image.Width;
                double textureY = y / (double)image.Height;
                textureCoordinates.Add(new Point(textureX, textureY));
                // 生成 MeshGeometry3D 中的三角形索引
                if (x >= image.Width - zoom || y >= image.Height - zoom) continue;
                var topLeft = (x / zoom) + (y / zoom) * (image.Width / zoom);
                var topRight = topLeft + 1;
                var bottomLeft = topLeft + (image.Width / zoom);
                var bottomRight = bottomLeft + 1;

                // 添加三角形索引
                triangleIndices.Add(topLeft);
                triangleIndices.Add(bottomLeft);
                triangleIndices.Add(topRight);

                triangleIndices.Add(topRight);
                triangleIndices.Add(bottomLeft);
                triangleIndices.Add(bottomRight);
            }
        }

        Console.WriteLine($"points Length:{positions.Count},triangleIndices Length:{triangleIndices.Count}");

        // Create a 3D mesh
        var mesh = new MeshGeometry3D
        {
            Positions = positions,
            TriangleIndices = triangleIndices,
            Normals = normals,
            TextureCoordinates = textureCoordinates
        };
        var textureBitmap = new BitmapImage(new Uri(imgPath));
        var textureBrush = new ImageBrush(textureBitmap);
        var front = new DiffuseMaterial(textureBrush);
        var back = new DiffuseMaterial(textureBrush);
        // Create a model and add it to the viewport
        //var model = new GeometryModel3D(mesh, material);
        //var model = new GeometryModel3D(mesh, Materials.Gray);
        var modelVisual = new MeshGeometryVisual3D {  
            MeshGeometry = mesh,
            Material = front ,BackMaterial = back };
        

            var lights = new Model3DGroup();
            
            // 添加环境光
            lights.Children.Add(new AmbientLight
            {
                Color = Colors.White
            });

            // 创建 ModelVisual3D 对象，并将灯光设置为其内容
            var lightVisual = new ModelVisual3D
            {
                Content = lights
            };
        helixViewport.Children.Add(modelVisual);
        helixViewport.Children.Add(lightVisual);
        image.Dispose();
    }
    
}