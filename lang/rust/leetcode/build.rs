extern crate napi_build;

fn main() {
    napi_build::setup();
    println!("build successfully");
    //build具体到bin中的某个文件无效，猜测可能是库crate才有效
}