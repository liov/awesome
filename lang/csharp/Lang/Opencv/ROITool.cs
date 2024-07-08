using OpenCvSharp;
using Range = OpenCvSharp.Range;

namespace Lang.Opencv;

public class ROITool
{
    public static void ROI()
    {
        var image = Cv2.ImRead("""D:\work\WXWork\1688858122545615\Cache\File\2024-07\0.96\18.bmp""");
        // image = new Mat(image, rowRange:new Range(4413, 4955),new Range(625, 794)); //17
        image = new Mat(image, rowRange:new Range(0, 800),new Range(0, 1800));
        Cv2.ImShow("image",image);
        Cv2.WaitKey();
    }
}