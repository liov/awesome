const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{
        .preferred_optimize_mode = .ReleaseFast,
    });
    const mod = b.createModule(.{
        .root_source_file = b.path("win.zig"),
        .target = target,
        .optimize = optimize,
    });
    const exe = b.addExecutable(.{
        .name = "win",
        .root_module = mod,
    });

    exe.linkLibC();
    exe.linkLibCpp();
    exe.addIncludePath(.{ .cwd_relative = "/opt/homebrew/opt/opencv/include/opencv4" });
    exe.addLibraryPath(.{ .cwd_relative = "/opt/homebrew/opt/opencv/lib" });
    exe.linkSystemLibrary("opencv_highgui");
    exe.linkSystemLibrary("opencv_core");

    b.installArtifact(exe);
}
