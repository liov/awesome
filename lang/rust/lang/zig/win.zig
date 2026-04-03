const std = @import("std");

extern "c" fn cvNamedWindow(title: [*:0]const u8, flags: c_int) void;
extern "c" fn cvDestroyAllWindows() void;

extern "c" fn cvWaitKey(delay: c_int) c_int;

pub fn main() !void {
    cvNamedWindow("Example", 1);
    _ = cvWaitKey(0);
    cvDestroyAllWindows();
}