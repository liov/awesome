import cv2

# 读取TIFF图像
image_path = "D:/result.tiff"
img = cv2.imread(image_path, cv2.IMREAD_ANYDEPTH)

# 检查图像是否正确加载
if img is not None:
    # 打印第一行的灰度值
    print(img[1245][2225:2250])
else:
    print("无法读取图像，请检查图像路径是否正确。")

image_path = r"xxx"
img = cv2.imread(image_path, cv2.IMREAD_ANYDEPTH)

# 检查图像是否正确加载
if img is not None:
    # 打印第一行的灰度值
    first_row_gray_values = img[1245][2225:2250]
    print(img[1245][2225:2250])
else:
    print("无法读取图像，请检查图像路径是否正确。")