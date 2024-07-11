@echo off

set cvdir=D:\sdk\opencv

if not exist %cvdir% mkdir %cvdir%
if not exist %cvdir%\build mkdir %cvdir%\build

echo Downloading OpenCV sources
echo.
echo For monitoring the download progress please check the D:\sdk\opencv directory.
echo.

REM This is why there is no progress bar:
REM https://github.com/PowerShell/PowerShell/issues/2138

echo Downloading: opencv-4.10.0.zip [91MB]
if not exist %cvdir%\opencv-4.10.0.zip powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri https://github.com/opencv/opencv/archive/4.10.0.zip -OutFile %cvdir%\opencv-4.10.0.zip"
echo Extracting...
if not exist %cvdir%\opencv-4.10.0 powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path %cvdir%\opencv-4.10.0.zip -DestinationPath %cvdir%"

echo Downloading: opencv_contrib-4.10.0.zip [58MB]
if not exist %cvdir%\opencv_contrib-4.10.0.zip powershell -command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; Invoke-WebRequest -Uri https://github.com/opencv/opencv_contrib/archive/4.10.0.zip -OutFile %cvdir%\opencv_contrib-4.10.0.zip"
echo Extracting...
if not exist %cvdir%\opencv_contrib-4.10.0 powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path %cvdir%\opencv_contrib-4.10.0.zip -DestinationPath %cvdir%"

echo Done with downloading and extracting sources.
echo.

@echo on

cd /D %cvdir%\build
:: set PATH=%PATH%;CMake\bin;mingw-w64\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\bin
if [%1]==[static] (
  echo Build static opencv
  set enable_shared=OFF
) else (
  set enable_shared=ON
)
cmake %cvdir%\opencv-4.10.0 -G "MinGW Makefiles" -B%cvdir%\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=%cvdir%\opencv_contrib-4.10.0\modules -DBUILD_SHARED_LIBS=%enable_shared% -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=ON -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=ON -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -Wno-dev
mingw32-make -j%NUMBER_OF_PROCESSORS%
mingw32-make install

rmdir %cvdir%\opencv-4.10.0 /s /q
rmdir %cvdir%\opencv_contrib-4.10.0 /s /q

@echo off
echo rm soucre[Y/N]:
echo.
set /p user_input=

if "%user_input%"=="Y" (
  del %cvdir%\opencv-4.10.0.zip /q
  del %cvdir%\opencv_contrib-4.10.0.zip /q
)


