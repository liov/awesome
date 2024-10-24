import cv2

# 读取图片
image = cv2.imread(r"xxx")
if image is None:
    print("Could not open or find the image!")
    exit()

# 在图片上绘制矩形框
#cv2.rectangle(image, (10935, 32705), (11088,32855), (0, 255, 0), 2)  # 绿色框，线宽为2
cv2.circle(image, (11011, 32780), 75,  (0, 255, 0), 2)

# 显示图片
cv2.imwrite('rect.jpg', image)