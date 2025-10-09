const std = @import("std");

extern "c" fn cvNamedWindow(title: [*:0]const u8, flags: c_int) void;
extern "c" fn cvDestroyAllWindows() void;

pub fn main() !void {
    cvNamedWindow("Example", 1);
    // 在这里添加更多的 OpenCV 调用
    cvDestroyAllWindows();
}