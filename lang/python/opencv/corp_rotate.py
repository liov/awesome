import cv2
import numpy as np

# 读取图片
image = cv2.imread(r"D:\xxx")
if image is None:
    print("Could not open or find the image!")
    exit()

# 定义矩形框的参数
x, y, w, h = 11028-int(368/2), 17375-int(512/2), 368, 512  # (x, y, width, height)

# 在图片上绘制矩形框
cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 2)  # 绿色框，线宽为2
x, y, w, h = 6264-int(368/2), 17375-int(512/2), 368, 512  # (x, y, width, height)
cv2.rectangle(image, (x, y), (x + w, y + h), (0, 255, 0), 2)  # 绿色框，线宽为2
x, y, w, h = 8646, 17943, 5152, 1672  # (x, y, width, height)
x1 = int(x - w / 2)
y1 = int(y - h / 2)
x2 = int(x + w / 2)
y2 = int(y + h / 2)

roi = image[y1:y2, x1:x2]
# 显示图片
cv2.imwrite('rect.jpg', roi)