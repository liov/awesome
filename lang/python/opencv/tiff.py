import cv2

image= cv2.imread(r"xxx.tiff", cv2.IMREAD_UNCHANGED)
# 2. 打印图像的基本信息
print(f"图像形状: {image.shape}")
print(f"图像数据类型: {image.dtype}")

# 3. 遍历图像的每个像素，打印其位置和值
for y in range(image.shape[0]):
    for x in range(image.shape[1]):
        pixel_value = image[y, x]
        if pixel_value > 3000 :
            print(f"位置 ({y}, {x}): 像素值 {pixel_value}")
