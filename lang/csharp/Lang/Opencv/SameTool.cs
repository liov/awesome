using OpenCvSharp.Features2D;

namespace Lang.Opencv;

using System;
using OpenCvSharp;

public class SameTool
{
    
    public static void Same()
    {
        // 读取两张图像
        Mat img1 = Cv2.ImRead("""D:\work\80.bmp""");
        Mat img2 = Cv2.ImRead("""D:\work\81.bmp""");

        if (img1.Empty() || img2.Empty())
        {
            Console.WriteLine("无法读取图像文件");
            return;
        }
        
        // 初始化特征检测器和描述子
        var detector = SIFT.Create();
        var matcher = new BFMatcher();

        // 检测关键点和计算描述子
        KeyPoint[] keypoints1, keypoints2;
        Mat descriptors1 = new Mat(), descriptors2 = new Mat();
        detector.DetectAndCompute(img1, null, out keypoints1, descriptors1);
        detector.DetectAndCompute(img2, null, out keypoints2, descriptors2);

        // 使用描述子进行匹配
        DMatch[] matches = matcher.Match(descriptors1, descriptors2);

        // 找到水平匹配的点对
        DMatch[] horizontalMatches = GetHorizontalMatches(matches, keypoints1, keypoints2);
        
        // 提取最佳匹配的关键点
        List<Point2f> srcPoints = new List<Point2f>();
        List<Point2f> dstPoints = new List<Point2f>();
        foreach (var match in horizontalMatches)
        {
            srcPoints.Add(keypoints1[match.QueryIdx].Pt);
            dstPoints.Add(keypoints2[match.TrainIdx].Pt);
        }
        
        // 显示最佳匹配的区域
        Mat resultImg = DrawMatches(img1, img2, srcPoints, dstPoints);

        // 显示结果
        Cv2.NamedWindow("Matches");
        Cv2.ImShow("Matches", resultImg);
        Cv2.WaitKey(0);
    }
    // 如果图像尺寸超过指定的最大尺寸，则等比例缩小图像
  
    
    // 获取水平匹配的点对
    public static DMatch[] GetHorizontalMatches(DMatch[] matches, KeyPoint[] keypoints1, KeyPoint[] keypoints2)
    {
        // 根据距离排序
        Array.Sort(matches, (a, b) => a.Distance.CompareTo(b.Distance));
        double maxDist = matches[^1].Distance;
        // 选择最佳匹配
        List<DMatch> bestMatches = new List<DMatch>();
        

        foreach (var match in matches)
        {
            Point2f pt1 = keypoints1[match.QueryIdx].Pt;
            Point2f pt2 = keypoints2[match.TrainIdx].Pt;

            // 计算线段的斜率
            double deltaY = Math.Abs(pt2.Y - pt1.Y);
            double deltaX = Math.Abs(pt2.X - pt1.X);
            double slope = deltaY / deltaX;
            // 排除斜率大于阈值的线段
            if (match.Distance < 0.1 * maxDist && slope < 0.01)
            {
                bestMatches.Add(match);
            }
        }

        return bestMatches.ToArray();
    }

    
    // 绘制匹配的线段
    static Mat DrawMatches(Mat img1, Mat img2, List<Point2f> srcPoints, List<Point2f> dstPoints)
    {
        // 创建结果图像
        Mat resultImg = new Mat();
        Cv2.HConcat(img1, img2, resultImg);

        // 绘制匹配的线段（水平）
        Scalar color = Scalar.Green;
        for (int i = 0; i < srcPoints.Count; i++)
        {
            Point pt1 = new Point((int)srcPoints[i].X, (int)srcPoints[i].Y);
            Point pt2 = new Point((int)dstPoints[i].X + img1.Width, (int)dstPoints[i].Y);

            Cv2.Line(resultImg, pt1, pt2, color, 2);
            Cv2.Circle(resultImg, pt1, 5, color, -1);
            Cv2.Circle(resultImg, pt2, 5, color, -1);
        }

        return resultImg;
    }
}