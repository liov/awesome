using OpenCvSharp;

namespace Lang.Opencv;

public class OverlayTool
{
    
    public static void Overlay()
    {
        // 读取两张图像（可以根据实际路径和文件名修改）
        Mat img1 = Cv2.ImRead("""D:\work\80.bmp""");
        Mat img2 = Cv2.ImRead("""D:\work\81.bmp""");

        if (img1.Empty() || img2.Empty())
        {
            Console.WriteLine("无法读取图像文件");
            return;
        }

        // 将两张图像叠加显示
        double alpha = 0.5; // 第一张图像的权重
        double beta = 1.0 - alpha; // 第二张图像的权重

        Mat result = new Mat();
        Cv2.AddWeighted(img1, alpha, img2, beta, 0, result);

        // 显示结果
        Cv2.NamedWindow("Overlay");
        Cv2.ImShow("Overlay", result);
        Cv2.WaitKey(0);
        Cv2.DestroyAllWindows();
    }


}