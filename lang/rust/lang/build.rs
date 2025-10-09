extern crate cc;
use std::env;
use std::path::PathBuf;
use cxx_build::CFG;

fn main() {
    println!("cargo:rerun-if-changed=c/*");
    println!("cargo:rerun-if-changed=build.rs");
    build();
    rustbind();
    opencv();
    println!("build successfully");
    //build具体到bin中的某个文件无效，猜测可能是库crate才有效
}

#[cfg(windows)]
fn build(){
    //gcc -shared -O2 -o clib.dll clib.c
    cc::Build::new()
        .file("c/clib.c")
        .define("FOO", Some("bar"))
        .include("c")
        .opt_level(2)
        .shared_flag(true)
        .static_flag(true)
        .compile("clib");
}

fn rustbind(){
     // Tell cargo to look for shared libraries in the specified directory
        println!("cargo:rustc-link-search=c");

        // Tell cargo to tell rustc to link the system bzip2
        // shared library.
        println!("cargo:rustc-link-lib=bz2");

        // The bindgen::Builder is the main entry point
        // to bindgen, and lets you build up options for
        // the resulting bindings.
        let bindings = bindgen::Builder::default()
            // The input header we would like to generate
            // bindings for.
            .header("c/wrapper.h")
            // Tell cargo to invalidate the built crate whenever any of the
            // included header files changed.
            .parse_callbacks(Box::new(bindgen::CargoCallbacks::new()))
            // Finish the builder and generate the bindings.
            .generate()
            // Unwrap the Result and panic on failure.
            .expect("Unable to generate bindings");

        // Write the bindings to the $OUT_DIR/bindings.rs file.
        let out_path = PathBuf::from(env::var("OUT_DIR").unwrap());
        bindings
            .write_to_file(out_path.join("bindings.rs"))
            .expect("Couldn't write bindings!");
}

fn opencv(){
     CFG.include_prefix = "";
        let opencv = pkg_config::probe_library("opencv4").unwrap();
        cxx_build::bridge("src/bin/opencv.rs")
            .file("cxx/resize.cpp")
            .flag_if_supported("-std=c++11")
            .include("cxx")
            .includes(opencv.include_paths)
            .opt_level(2)
            .compile("resize");
        for link_path in opencv.link_paths {
            println!("cargo:rustc-link-search={}", link_path.to_str().unwrap());
        }
        for lib in opencv.libs {
            println!("cargo:rustc-link-lib={}", lib);
        }
    }