using System.Drawing.Imaging;
using System.Windows;
using System.Windows.Media.Imaging;
using System.Windows.Threading;
using Lang.Opencv;
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
                    double sharpness = SharpnessTool.CalculateSharpness(frame);
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