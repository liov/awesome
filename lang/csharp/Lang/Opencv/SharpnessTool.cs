using OpenCvSharp;
using Range = OpenCvSharp.Range;

namespace Lang.Opencv;

public class SharpnessTool
{
    public static void Clarity()
    {
        for (int i = 17; i < 26; i++)
        {
            Console.Write($"\x1b[1;32m{i}\x1b[0m:{{");
            foreach (var z in new[] { "0.96", "1.72", "2.8", "3.52", "4.64" })
            {
                var image = Cv2.ImRead($"""D:\work\z\{z}\{i}.bmp""",ImreadModes.Grayscale);
                if (i == 17)
                {
                    // 截取区域并创建新图像
                    image = new Mat(image, rowRange: new Range(4413, 4955), new Range(625, 794));
                }

                if (i == 18)
                {
                    // 截取区域并创建新图像
                    image = new Mat(image, rowRange: new Range(0, 800), new Range(0, 1800));
                }

                if (i is > 18 and < 22)
                {
                    // 截取区域并创建新图像
                    image = new Mat(image, rowRange: new Range(0, 800), new Range(0, 5120));
                }

                if (i == 24)
                {
                    image = new Mat(image, rowRange: new Range(2200, 4300), new Range(2800, 5120));
                }

                Console.Write($"\x1b[1;33m{z}\x1b[0m:{CalculateSharpness(image)} ");
            }

            Console.WriteLine("}");
        }
    }

    public static double CalculateSharpness(Mat gray)
    {
        // 计算拉普拉斯变换
        Mat laplacian = new Mat();
        Cv2.Laplacian(gray, laplacian, MatType.CV_64F);

        // 计算方差
        Scalar mean, stddev;
        Cv2.MeanStdDev(laplacian, out mean, out stddev);

        double sharpness = stddev.Val0; // 方差作为清晰度度量

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