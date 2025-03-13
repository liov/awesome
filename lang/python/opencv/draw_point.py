import cv2
import json

# 读取图片
image = cv2.imread(r"D:\xxx.jpg")
if image is None:
    print("Could not open or find the image!")
    exit()

with open(r"D:\work\test.json", 'r', encoding='utf-8') as file:
    data = json.load(file)

for item in data['components']:
    x = int(item['x'])
    y = int(item['y'])
    cv2.circle(image, (x, y), 1, (0, 255, 0), cv2.FILLED)

# 显示图片
cv2.imwrite('rect.jpg', image)