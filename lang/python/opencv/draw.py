import cv2

# 读取图片
image = cv2.imread(r"xxx")
if image is None:
    print("Could not open or find the image!")
    exit()

# 在图片上绘制矩形框
#cv2.rectangle(image, (10935, 32705), (11088,32855), (0, 255, 0), 2)  # 绿色框，线宽为2
#cv2.circle(image, (11011, 32780), 75,  (0, 255, 0), 2)

cv2.rectangle(image, (10789,5465), (11489,6165), (0, 255, 0), 2)
cv2.rectangle(image, (3327,4645), (4027,5345), (0, 255, 0), 2)
cv2.rectangle(image, (10761,16334), (11461,17034), (0, 255, 0), 2)
cv2.rectangle(image, (3300,17299), (4000,17999), (0, 255, 0), 2)
cv2.rectangle(image, (10719,19357), (11419,20057), (0, 255, 0), 2)
cv2.rectangle(image, (3254,20320), (3954,21020), (0, 255, 0), 2)
cv2.rectangle(image, (10687,32010), (11387,32710), (0, 255, 0), 2)
cv2.rectangle(image, (3228,31188), (3928,31888), (0, 255, 0), 2)
# 显示图片
cv2.imwrite('rect.jpg', image)