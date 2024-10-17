import cv2

# 读取图片
image = cv2.imread(r"D:\xxx")
if image is None:
    print("Could not open or find the image!")
    exit()

# 定义矩形框的参数
x, y, w, h = 194-int(368/2), 1404-int(512/2), 368, 512  # (x, y, width, height)

# 在图片上绘制矩形框
cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 2)  # 绿色框，线宽为2
x, y, w, h = 4958-int(368/2), 1404-int(512/2), 368, 512  # (x, y, width, height)
cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 2)  # 绿色框，线宽为2

# 显示图片
cv2.imwrite('rect.jpg', image)