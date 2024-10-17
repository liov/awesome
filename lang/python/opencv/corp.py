import cv2
import numpy as np

# 读取图片
image = cv2.imread(r"xxx")
if image is None:
    print("Could not open or find the image!")
    exit()

# 定义矩形框的参数
x, y, w, h = 5371, 5042, 5120, 5120  # (x, y, width, height)

x1 = int(x - w / 2)
y1 = int(y - h / 2)
x2 = int(x + w / 2)
y2 = int(y + h / 2)

roi = image[y1:y2, x1:x2]
# 显示图片
cv2.imwrite('rect.jpg', roi)