node-gyp rebuild --GTK_Root=C:\somewhere\GTK --jpeg_root=C:\somewhere\libjpeg-turbo
node-gyp rebuild --GTK_Root=D:\sdk\msys64\ucrt64 --jpeg_root=D:\sdk\msys64\ucrt64
搞了半天这个必须要在要安装的包目录下执行
直接改包目录下的binding.gyp文件，改相应的参数，比如GTK_Root
msys2不兼容，老老实实下载https://iso.mirrors.ustc.edu.cn/gnome/binaries/win64/gtk+/2.22/gtk+-bundle_2.22.1-20101229_win64.zip
编辑node_modules\canvas\binding.gyp 的GTK_Root 为解压目录