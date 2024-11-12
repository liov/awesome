const std = @import("std");
const Builder = std.build.Builder;

pub fn build(b: *Builder) void {
    const target = b.standardTargetOptions();
    const mode = b.standardReleaseOptions();
    const exe = b.addExecutable("opencv_test", "main.cpp");

    exe.setBuildMode(mode);
    exe.setTarget(target);

    // 设置 OpenCV 路径
    const opencv_include_dir = "path/to/opencv/include";
    const opencv_lib_dir = "path/to/opencv/lib";
    exe.addIncludeDir(opencv_include_dir);
    exe.addLibDir(opencv_lib_dir);

    // 添加 OpenCV 库
    exe.addLib("opencv_core");
    exe.addLib("opencv_imgproc");
    exe.addLib("opencv_highgui");

    // 编译
    b.installExecutable(exe);
}
