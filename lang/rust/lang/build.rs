extern crate cc;

fn main() {
    println!("cargo:rerun-if-changed=c/*");
    println!("cargo:rerun-if-changed=build.rs");
    build();
    println!("build successfully");
    //build具体到bin中的某个文件无效，猜测可能是库crate才有效
}

#[cfg(windows)]
fn build(){
    //gcc -shared -O2 -o clib.dll clib.c
    cc::Build::new()
        .file("c/clib.c")
        .define("FOO", Some("bar"))
        .include("src")
        .opt_level(2)
        .shared_flag(true)
        .static_flag(true)
        .compile("clib");
}