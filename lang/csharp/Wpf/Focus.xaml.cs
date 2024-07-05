using System.Drawing.Imaging;
using System.Windows;
using System.Windows.Media.Imaging;
using System.Windows.Threading;
using OpenCvSharp;
using Application = System.Windows.Application;
using MessageBox = System.Windows.MessageBox;
using Window = System.Windows.Window;

namespace Wpf;

public partial class Focus : Window
{
    
    private VideoCapture capture;
    private Mat frame;
    private BitmapImage image;
    private bool isCameraRunning = false;
    
    public Focus()
    {
        InitializeComponent();
        StartCamera();
    }
    
    private void StartCamera()
    {
        capture = new VideoCapture(0);
        capture.Open(0);
        if (!capture.IsOpened())
        {
            MessageBox.Show("Failed to open camera.");
            return;
        }

        frame = new Mat();
        isCameraRunning = true;

        // 尝试禁用自动对焦
        // 请注意，不同的摄像头可能使用不同的属性进行自动对焦控制
        const int CAP_PROP_AUTOFOCUS = 39; // 39 是常用的自动对焦属性，但可能因设备而异
        capture.Set(CAP_PROP_AUTOFOCUS, 0);
        
        Task.Run(() =>
        {
            while (isCameraRunning)
            {
                capture.Read(frame);
                if (!frame.Empty())
                {
                    double sharpness = CalculateSharpness(frame);
                    Console.WriteLine(sharpness);
                    image = BitmapSourceConvert.ToBitmapImage(frame);
                    image.Freeze();
                    Application.Current.Dispatcher.Invoke(() => videoDisplay.Source = image);
                }
            }
        });
    }

    private void StopCamera()
    {
        isCameraRunning = false;
        capture.Release();
    }

    protected override void OnClosed(EventArgs e)
    {
        StopCamera();
        base.OnClosed(e);
    }
    
    static double CalculateSharpness(Mat image)
    {
        Mat gray = new Mat();
        Cv2.CvtColor(image, gray, ColorConversionCodes.BGR2GRAY);

        // 计算拉普拉斯变换
        Mat laplacian = new Mat();
        Cv2.Laplacian(gray, laplacian, MatType.CV_64F);

        // 计算方差
        Scalar mean, stddev;
        Cv2.MeanStdDev(laplacian, out mean, out stddev);

        double sharpness = stddev.Val0 ; // 方差作为清晰度度量
        
        // 使用 Sobel 算子计算图像的梯度
        Mat sobelX = new Mat(), sobelY = new Mat();
        Cv2.Sobel(gray, sobelX, MatType.CV_64F, 1, 0);
        Cv2.Sobel(gray, sobelY, MatType.CV_64F, 0, 1);

        // 计算梯度幅值的平均值
        Mat magnitude = new Mat();
        Cv2.Magnitude(sobelX, sobelY, magnitude);
        
        Cv2.MeanStdDev(magnitude, out mean, out stddev);

        // 返回平均梯度幅值作为图像清晰度指标
        return sharpness + mean.Val0;
    }
    
}

public static class BitmapSourceConvert
{
    public static BitmapImage ToBitmapImage(Mat image)
    {
        using (var stream = image.ToMemoryStream())
        {
            var bitmapImage = new BitmapImage();
            bitmapImage.BeginInit();
            bitmapImage.CacheOption = BitmapCacheOption.OnLoad;
            bitmapImage.StreamSource = stream;
            bitmapImage.EndInit();
            return bitmapImage;
        }
    }
}