# msys2有问题，编译成功执行报错
pacman -S mingw-w64-ucrt-x86_64-opencv
CGO_CXXFLAGS=--std=c++11 CGO_CPPFLAGS="-ID:\sdk\msys64\ucrt64\include\opencv4" CGO_LDFLAGS="-LD:\sdk\msys64\ucrt64\lib  -lopencv_core -lopencv_photo -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking -lopencv_img_hash -lopencv_calib3d"  go build -tags customenv -o D:/sdk/opencv/build/install/x64/mingw/bin/gocv.exe gocv.go
## error
gocv.exe: error while loading shared libraries: ?: cannot open shared object file: No such file or directory

## IDEA env
CGO_CXXFLAGS=--std=c++11;CGO_CPPFLAGS=-ID:\sdk\msys64\ucrt64\include\opencv4;CGO_LDFLAGS=-LD:\sdk\msys64\ucrt64\lib  -lopencv_core -lopencv_photo -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking -lopencv_img_hash -lopencv_calib3d
## 实测msys2的opencv动态库是可用的,可能与gocv写的包装不兼容
## 神奇的出现了,试了试又可以了,难道是缺了啥库?(mingw-w64-x86_64-ucrt-qt6-5compat (for the HighGUI module)
                           #mingw-w64-ucrt-x86_64-vtk (opencv_viz module)) ?
# build.cmd
msys2 mingw-w64 编译 or https://github.com/niXman/mingw-builds-binaries/releases posix-seh-msvcrt-rt_v11

CGO_CXXFLAGS=--std=c++11 CGO_CPPFLAGS="-ID:/sdk/opencv/build/install/include" CGO_LDFLAGS="-LD:/sdk/opencv/build/install/x64/mingw/bin -lopencv_core4100 -lopencv_photo4100 -lopencv_face4100 -lopencv_videoio4100 -lopencv_imgproc4100 -lopencv_highgui4100 -lopencv_imgcodecs4100 -lopencv_objdetect4100 -lopencv_features2d4100 -lopencv_video4100 -lopencv_dnn4100 -lopencv_xfeatures2d4100 -lopencv_plot4100 -lopencv_tracking4100 -lopencv_img_hash4100 -lopencv_calib3d4100 -lopencv_bgsegm4100 -lopencv_aruco4100 -lopencv_wechat_qrcode4100 -lopencv_ximgproc4100"  go build -tags customenv -o D:/sdk/opencv/build/install/x64/mingw/bin/gocv.exe gocv.go

## IDEA env
CGO_CXXFLAGS=--std=c++11;CGO_CPPFLAGS=-ID:/sdk/opencv/build/install/include;CGO_LDFLAGS=-LD:/sdk/opencv/build/install/x64/mingw/bin -lopencv_core4100 -lopencv_photo4100 -lopencv_face4100 -lopencv_videoio4100 -lopencv_imgproc4100 -lopencv_highgui4100 -lopencv_imgcodecs4100 -lopencv_objdetect4100 -lopencv_features2d4100 -lopencv_video4100 -lopencv_dnn4100 -lopencv_xfeatures2d4100 -lopencv_plot4100 -lopencv_tracking4100 -lopencv_img_hash4100 -lopencv_calib3d4100 -lopencv_bgsegm4100 -lopencv_aruco4100 -lopencv_wechat_qrcode4100 -lopencv_ximgproc4100

