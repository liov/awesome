@echo off

if not exist "D:\sdk\opencv" mkdir "D:\sdk\opencv"
if not exist "D:\sdk\opencv\build" mkdir "D:\sdk\opencv\build"

echo Downloading OpenCV sources
echo.
echo For monitoring the download progress please check the D:\sdk\opencv directory.
echo.

REM This is why there is no progress bar:
REM https://github.com/PowerShell/PowerShell/issues/2138

echo Downloading: opencv-4.10.0.zip [91MB]
powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri https://github.com/opencv/opencv/archive/4.10.0.zip -OutFile D:\sdk\opencv\opencv-4.10.0.zip"
echo Extracting...
powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path D:\sdk\opencv\opencv-4.10.0.zip -DestinationPath D:\sdk\opencv"
del D:\sdk\opencv\opencv-4.10.0.zip /q
echo.

echo Downloading: opencv_contrib-4.10.0.zip [58MB]
powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri https://github.com/opencv/opencv_contrib/archive/4.10.0.zip -OutFile D:\sdk\opencv\opencv_contrib-4.10.0.zip"
echo Extracting...
powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path D:\sdk\opencv\opencv_contrib-4.10.0.zip -DestinationPath D:\sdk\opencv"
del D:\sdk\opencv\opencv_contrib-4.10.0.zip /q
echo.

echo Done with downloading and extracting sources.
echo.

@echo on

cd /D D:\sdk\opencv\build
:: set PATH=%PATH%;CMake\bin;mingw-w64\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\bin
if [%1]==[static] (
  echo Build static opencv
  set enable_shared=OFF
) else (
  set enable_shared=ON
)
cmake D:\sdk\opencv\opencv-4.10.0 -G "MinGW Makefiles" -BD:\sdk\opencv\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=D:\sdk\opencv\opencv_contrib-4.10.0\modules -DBUILD_SHARED_LIBS=%enable_shared% -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=ON -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=ON -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -Wno-dev
mingw32-make -j%NUMBER_OF_PROCESSORS%
mingw32-make install
rmdir D:\sdk\opencv\opencv-4.10.0 /s /q
rmdir D:\sdk\opencv\opencv_contrib-4.10.0 /s /q

