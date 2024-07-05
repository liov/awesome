using OpenCvSharp;
using OpenCvSharp.Features2D;

namespace Lang.Opencv;

public class PuzzleTool
{

    public static void BFMatcher()
    {
        BFMatcher("""D:\work\0704\80.bmp""", """D:\work\0704\81.bmp""",true);
    }
    
    public static void BatchMatcher()
    {
        for (int i = 0; i < 100; i++)
        {
            var img1 = $"""D:\work\0704\{i}.bmp""";
            var img2 = $"""D:\work\0704\{i+1}.bmp""";
            Console.WriteLine(i);
            BFMatcher(img1, img2);
        }
    }
    
    public static void BFMatcher(String file1,String file2,Boolean show =false)
    {
        // 读取两张图像（可以根据实际路径和文件名修改）
        Mat img1 = Cv2.ImRead(file1);
        Mat img2 = Cv2.ImRead(file2);

        if (img1.Empty() || img2.Empty())
        {
            Console.WriteLine("无法读取图像文件");
            return;
        }

        // 调整图像大小（可选）
        /*int zoom = 4;
        img1 = ResizeTool.Resize(img1, zoom);
        img2 = ResizeTool.Resize(img2, zoom);*/

        // 转换为灰度图像
        Mat gray1 = new Mat();
        Mat gray2 = new Mat();
        Cv2.CvtColor(img1, gray1, ColorConversionCodes.BGR2GRAY);
        Cv2.CvtColor(img2, gray2, ColorConversionCodes.BGR2GRAY);

        // 初始化特征检测器和描述符
        ORB orb = ORB.Create();
        KeyPoint[] keypoints1, keypoints2;
        Mat descriptors1 = new Mat(), descriptors2 = new Mat();

        // 检测特征点和计算描述符
        orb.DetectAndCompute(gray1, null, out keypoints1, descriptors1);
        orb.DetectAndCompute(gray2, null, out keypoints2, descriptors2);

        // 匹配特征点
        BFMatcher matcher = new BFMatcher(NormTypes.Hamming);
        DMatch[] matches = matcher.Match(descriptors1, descriptors2);

        double minDist = matches.Min(m => m.Distance);
        var goodMatches = matches.Where(m => m.Distance < 3 * minDist);
        // 筛选匹配点并计算重叠区域
        List<Point2f> points1 = new List<Point2f>();
        List<Point2f> points2 = new List<Point2f>();
        foreach (DMatch match in goodMatches)
        {
            points1.Add(keypoints1[match.QueryIdx].Pt);
            points2.Add(keypoints2[match.TrainIdx].Pt);
        }

        // 使用RANSAC算法估计重叠区域的宽度
        int overlapWidth = CalculateOverlapWidthRANSAC(points1,points2);

        // 展示重叠区域宽度
        Console.WriteLine($"Overlap width: {overlapWidth} pixels");
       
        if (show)
        {
            using ( Mat overlappedImage = new Mat())
            {
                // 显示或保存重叠区域的图像
                Cv2.HConcat(img1, img2.ColRange(overlapWidth, img2.Width), overlappedImage);
                Cv2.ImShow("Overlapped Image", overlappedImage);
                Cv2.WaitKey(0);
            }
       
        }
        // 释放资源
        img1.Dispose();
        img2.Dispose();
        gray1.Dispose();
        gray2.Dispose();
    }
    
    static int CalculateOverlapWidthRANSAC(List<Point2f> points1, List<Point2f> points2)
    {
        // 使用RANSAC算法估计重叠区域的宽度
        const int iterations = 1000;
        const double threshold = 5.0; // 阈值，匹配点的距离小于此阈值才认为是内点

        int bestOverlapWidth = 0;
        int bestInliersCount = 0;

        Random random = new Random();
        for (int i = 0; i < iterations; i++)
        {
            // 随机选择一组匹配点
            int idx = random.Next(points1.Count);
            Point2f p1 = points1[idx];
            Point2f p2 = points2[idx];

            // 计算平移距离
            int overlapWidth = (int)Math.Round(p1.X-p2.X);

            // 计算内点数
            int inliersCount = CountInliers(points1, points2, overlapWidth, threshold);

            // 更新最佳结果
            if (inliersCount > bestInliersCount)
            {
                bestInliersCount = inliersCount;
                bestOverlapWidth = overlapWidth;
            }
        }

        return bestOverlapWidth;
    }

    static int CountInliers(List<Point2f> points1, List<Point2f> points2, int overlapWidth, double threshold)
    {
        // 计算在给定平移距离下的内点数
        int inliersCount = 0;
        for (int i = 0; i < points1.Count; i++)
        {
            float dx =  points1[i].X - points2[i].X;
            if (Math.Abs(dx - overlapWidth) < threshold)
            {
                inliersCount++;
            }
        }
        return inliersCount;
    }
    
}