# msys2有问题，编译成功执行报错
pacman -S mingw-w64-x86_64-opencv
CGO_CXXFLAGS=--std=c++11 CGO_CPPFLAGS="-ID:\sdk\msys64\mingw64\include\opencv4" CGO_LDFLAGS="-LD:\sdk\msys64\mingw64\bin -LD:\sdk\msys64\mingw64\lib\opencv4\3rdparty -lopencv_core-410 -lopencv_photo-410 -lopencv_face-410 -lopencv_videoio-410 -lopencv_imgproc-410 -lopencv_highgui-410 -lopencv_imgcodecs-410 -lopencv_objdetect-410 -lopencv_features2d-410 -lopencv_video-410 -lopencv_dnn-410 -lopencv_xfeatures2d-410 -lopencv_plot-410 -lopencv_tracking-410 -lopencv_img_hash-410 -lopencv_calib3d-410"  go build -tags customenv gocv.go
## error
gocv.exe: error while loading shared libraries: ?: cannot open shared object file: No such file or directory

## IDEA env
CGO_CXXFLAGS=--std=c++11;CGO_CPPFLAGS=-ID:\sdk\msys64\mingw64\include\opencv4;CGO_LDFLAGS=-LD:\sdk\msys64\mingw64\bin -lopencv_core-410 -lopencv_photo-410 -lopencv_face-410 -lopencv_videoio-410 -lopencv_imgproc-410 -lopencv_highgui-410 -lopencv_imgcodecs-410 -lopencv_objdetect-410 -lopencv_features2d-410 -lopencv_video-410 -lopencv_dnn-410 -lopencv_xfeatures2d-410 -lopencv_plot-410 -lopencv_tracking-410 -lopencv_img_hash-410 -lopencv_calib3d-410

# build.cmd
msys2 mingw-w64 编译 or https://github.com/niXman/mingw-builds-binaries/releases posix-seh-msvcrt-rt_v11

pacman -S mmingw-w64-x86_64-opencv
CGO_CXXFLAGS=--std=c++11 CGO_CPPFLAGS="-ID:/sdk/opencv/build/install/include" CGO_LDFLAGS="-LD:/sdk/opencv/build/install/x64/mingw/bin -lopencv_core4100 -lopencv_photo4100 -lopencv_face4100 -lopencv_videoio4100 -lopencv_imgproc4100 -lopencv_highgui4100 -lopencv_imgcodecs4100 -lopencv_objdetect4100 -lopencv_features2d4100 -lopencv_video4100 -lopencv_dnn4100 -lopencv_xfeatures2d4100 -lopencv_plot4100 -lopencv_tracking4100 -lopencv_img_hash4100 -lopencv_calib3d4100 -lopencv_bgsegm4100 -lopencv_aruco4100 -lopencv_wechat_qrcode4100 -lopencv_ximgproc4100"  go build -tags customenv -o D:/sdk/opencv/build/install/x64/mingw/bin/gocv.exe gocv.go

## IDEA env
CGO_CXXFLAGS=--std=c++11;CGO_CPPFLAGS=-ID:/sdk/opencv/build/install/include;CGO_LDFLAGS=-LD:/sdk/opencv/build/install/x64/mingw/bin -lopencv_core4100 -lopencv_photo4100 -lopencv_face4100 -lopencv_videoio4100 -lopencv_imgproc4100 -lopencv_highgui4100 -lopencv_imgcodecs4100 -lopencv_objdetect4100 -lopencv_features2d4100 -lopencv_video4100 -lopencv_dnn4100 -lopencv_xfeatures2d4100 -lopencv_plot4100 -lopencv_tracking4100 -lopencv_img_hash4100 -lopencv_calib3d4100 -lopencv_bgsegm4100 -lopencv_aruco4100 -lopencv_wechat_qrcode4100 -lopencv_ximgproc4100

