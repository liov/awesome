using System.Runtime.InteropServices;

namespace Gerber
{

    public static class Cairo
    {
        private const string
            LIBRARY_PATH = "libcairo-2.dll"; // Change this to the path of the Gerbv library on your system

        [DllImport(LIBRARY_PATH)]
        public static extern IntPtr cairo_create(IntPtr surface);

        [DllImport(LIBRARY_PATH)]
        public static extern void cairo_set_source_rgb(IntPtr cr, double red, double green, double blue);

        [DllImport(LIBRARY_PATH)]
        public static extern void cairo_rectangle(IntPtr cr, double x, double y, double width, double height);

        [DllImport(LIBRARY_PATH)]
        public static extern void cairo_fill(IntPtr cr);

        [DllImport(LIBRARY_PATH)]
        public static extern void cairo_destroy(IntPtr cr);

        [DllImport(LIBRARY_PATH)]
        public static extern IntPtr cairo_image_surface_create(int format, int width, int height);

        [DllImport(LIBRARY_PATH)]
        public static extern IntPtr cairo_image_surface_get_data(IntPtr surface);

        [DllImport(LIBRARY_PATH)]
        public static extern void cairo_surface_destroy(IntPtr surface);

        [DllImport(LIBRARY_PATH)]
        public static extern void cairo_surface_flush(IntPtr surface);

        [DllImport(LIBRARY_PATH, EntryPoint = "cairo_surface_write_to_png_stream")]
        public static extern int cairo_surface_write_to_png_stream(IntPtr surface, IntPtr write_func, IntPtr closure);

    }
}