using System.Runtime.InteropServices;

namespace Gerber
{
    public static class Gerbv
    {
        private const string LIBRARY_PATH = "libgerbv-1.dll"; 

        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_create_project")]
        public static extern IntPtr gerbv_create_project();

        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_destroy_project")]
        public static extern int gerbv_destroy_project(IntPtr project);

        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_open_layer_from_filename")]
        public static extern IntPtr gerbv_open_layer_from_filename(IntPtr project, string filename);


        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_export_png_file_from_project_autoscaled")]
        public static extern int gerbv_export_png_file_from_project_autoscaled(IntPtr project, int width, int height, string filename);


        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_export_svg_file_from_project_autoscaled")]
        public static extern int gerbv_export_svg_file_from_project_autoscaled(IntPtr project, string filename);

        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_export_autoscale_project")]
        public static extern IntPtr gerbv_export_autoscale_project(IntPtr project);

        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_export_svg_file_from_project")]
        public static extern int gerbv_export_svg_file_from_project(IntPtr project, IntPtr renderInfo, string filename);

        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_render_layer_to_cairo_target", CallingConvention = CallingConvention.Cdecl)]
        public static extern int gerbv_render_layer_to_cairo_target(IntPtr cairo, IntPtr fileInfo,ref gerbv_render_info_t renderInfo);
        
        [DllImport(LIBRARY_PATH, EntryPoint = "gerbv_render_all_layers_to_cairo_target_for_vector_output", CallingConvention = CallingConvention.Cdecl)]
        public static extern int gerbv_render_all_layers_to_cairo_target_for_vector_output(IntPtr project, IntPtr cairo,ref gerbv_render_info_t renderInfo);
        
    }
    
    [StructLayout(LayoutKind.Sequential)]
    public struct gerbv_render_info_t
    {
        public double scaleFactorX;
        public double scaleFactorY;
        public double lowerLeftX;
        public double lowerLeftY;
        public gerbv_render_types_t renderType;
        public int displayWidth;
        public int displayHeight;
        [MarshalAs(UnmanagedType.I1)]
        public bool show_cross_on_drill_holes;
    }

    public enum gerbv_render_types_t
    {
        GERBV_RENDER_TYPE_GDK,
        GERBV_RENDER_TYPE_GDK_XOR,
        GERBV_RENDER_TYPE_CAIRO_NORMAL,
        GERBV_RENDER_TYPE_CAIRO_HIGH_QUALITY,
        GERBV_RENDER_TYPE_MAX   
    }
}