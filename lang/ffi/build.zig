const std = @import("std");

pub fn build(b: *std.Build) void {
    // 标准构建目标
    const target = b.standardTargetOptions(.{});
    // 标准构建模式
    const optimize = b.standardOptimizeOption(.{
        .preferred_optimize_mode = .ReleaseFast,
    });
    const exe = b.addExecutable(.{
        .name = "zig_hello",
        .root_source_file = b.path("hello.zig"),
        .target = target,
        .optimize = optimize,
    });
    exe.linkLibC();
    exe.addIncludePath(b.path("."));
    exe.addCSourceFiles(.{
        .files = &.{"newplus/plus.c"},
        .flags = &.{},
    });
    exe.addRPath(b.path("$ORIGIN"));
    b.default_step.dependOn(&exe.step);
    b.installArtifact(exe);
}
