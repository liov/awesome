from setuptools import Extension, setup
from Cython.Build import cythonize
from setuptools.command.build_ext import build_ext

# 使用 pkg-config 获取库和头文件路径
import subprocess

def get_pkg_config_output(pkg_name):
    try:
        output = subprocess.check_output(["pkg-config", "--cflags", "--libs", pkg_name])
        return output.decode("utf-8").strip()
    except subprocess.CalledProcessError:
        return ""

pkg_config_output = get_pkg_config_output("opencv4").split()
if pkg_config_output:
    include_dirs = [pkg_config_output[0].strip("-I")]
    library_dirs = [pkg_config_output[1].strip("-L")]
    libraries =  [item[2:] if item.startswith("-l") else item for item in pkg_config_output[2:]]
    print(include_dirs,library_dirs,libraries)
else:
    include_dirs = []
    library_dirs = []
    libraries = []


extensions = [
    Extension("wrapper", ["wrapper.pyx"],
              include_dirs=include_dirs,
              libraries=libraries,
              library_dirs=library_dirs,
              extra_compile_args=['-O3', '-ffast-math'],  # 编译选项
     ),

]

link_args = ['-static-libgcc',
             '-static-libstdc++',
             '-Wl,-Bstatic,--whole-archive',
             '-lwinpthread',
             '-Wl,--no-whole-archive']
class Build(build_ext):
    def build_extensions(self):
        if self.compiler.compiler_type == 'mingw32':
            for e in self.extensions:
                e.extra_link_args += link_args
        super(Build, self).build_extensions()

# python setup.py build_ext --inplace --compiler=mingw32
setup(
    name="wrapper",
    version="1.0",
    description="My Python module",
    ext_modules=cythonize(extensions, language_level=3),
    cmdclass={'build_ext': Build},
    zip_safe=False,
)