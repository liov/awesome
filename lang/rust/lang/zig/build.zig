const std = @import("std");
// 跑不了，别搞zig了,坑多资料少, opencv去lib下找opencv_core.dll不找libopencv_core.dll
pub fn build(b: *std.Build) void {
    // 标准构建目标
    const target = b.standardTargetOptions(.{});
    // 标准构建模式
    const optimize = b.standardOptimizeOption(.{
        .preferred_optimize_mode = .ReleaseFast,
    });
    const exe = b.addExecutable(.{
        .name = "opencv_test",
        .root_source_file = b.path("win.zig"),
        .target = target,
        .optimize = optimize,
    });

    exe.linkLibC();
    exe.linkLibCpp();
    // 添加 OpenCV 库
    exe.linkSystemLibrary("opencv4");

    // 编译
    exe.addRPath(b.path("$ORIGIN"));
    b.default_step.dependOn(&exe.step);
    b.installArtifact(exe);
}
