using OpenCvSharp;

namespace Lang.Opencv;

public class ResizeTool
{
    
    public static Mat ResizeIfTooLarge(Mat image, int maxSize)
    {
        int width = image.Width;
        int height = image.Height;
        double scale = 1.0;

        if (width > maxSize || height > maxSize)
        {
            if (width >= height)
            {
                scale = (double)maxSize / width;
            }
            else
            {
                scale = (double)maxSize / height;
            }

            int newWidth = (int)(width * scale);
            int newHeight = (int)(height * scale);

            return image.Resize(new Size(newWidth, newHeight));
        }

        return image.Clone();
    }
    
    public static Mat Resize(Mat image, int zoom)
    {
        int width = image.Width;
        int height = image.Height;

        int newWidth = width / zoom;
        int newHeight = height / zoom;
        var newImage = image.Resize(new Size(newWidth, newHeight));
        image.Dispose();
        return newImage;
    }
}