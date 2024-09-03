import cv2

# 读取 TIFF 文件
tiff_image = cv2.imread(r'D:\work\scan-panel-0826\1.tiff', cv2.IMREAD_GRAYSCALE)
resized_image = cv2.resize(tiff_image,(512,512))
# 检查是否成功读取图像
if tiff_image is None:
    print("Failed to read TIFF file.")
else:
    # 显示图像
    cv2.namedWindow("TIFF Image", cv2.WINDOW_NORMAL)
    cv2.resizeWindow("TIFF Image", 512, 512)
    cv2.imshow("TIFF Image", tiff_image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()