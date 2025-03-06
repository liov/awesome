import cv2
import json

# 读取图片
image = cv2.imread(r"D:\xxx.jpg")
if image is None:
    print("Could not open or find the image!")
    exit()

with open(r"D:\xxx.json", 'r', encoding='utf-8') as file:
    data = json.load(file)

for item in data['rect']:
    x = int(item['xPx'])
    y = int(item['yPx'])
    h = int(item['heightPx'])
    w = int(item['widthPx'])
    angle = float(item['rotation'])
    # 绘制旋转矩形
    cv2.ellipse(image, (x, y), (w // 2, h // 2), angle, 0, 360, (0, 255, 0), 2)

# 在图片上绘制矩形框
#cv2.rectangle(image, (10935, 32705), (11088,32855), (0, 255, 0), 2)  # 绿色框，线宽为2
#cv2.circle(image, (11011, 32780), 75,  (0, 255, 0), 2)

# 显示图片
cv2.imwrite('rect.jpg', image)