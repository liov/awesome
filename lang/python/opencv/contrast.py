import cv2
import numpy as np
import json
import easyocr
reader = easyocr.Reader(['en']) # this needs to run only once to load the model into memory


def calculate_contrast(image):
    # 读取图像，转换为灰度图
    gary = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    # 计算图像像素亮度的标准差
    return np.std(gary)



dir= r"xxx"
with open(dir+r"\board\board.json", 'r') as file:
    data = json.load(file)
    for  key, value in data['inlineWinsMap'].items():
        for item in value:
            if item['winType'] == 'silk':
                x1, y1 = int(item['xPx']) - int(item['lengthPx']/2), int(item['yPx'])- int(item['widthPx']/2)
                x2, y2 = int(item['xPx']) + int(item['lengthPx']/2), int(item['yPx'])+ int(item['widthPx']/2)
                image1 = cv2.imread(dir+rf"\board\outline-win\{key}--light2.jpg", cv2.IMREAD_COLOR)
                contrast  = calculate_contrast(image1[y1:y2, x1:x2])
                print(f"{key} light2 图片对比度（标准差）：{contrast}")
                result = reader.readtext(image1)
                print([ {'text':item[1],'confident':item[2]} for item in result])
                image2 = cv2.imread(dir+rf"\board\outline-win\{key}--light3.jpg", cv2.IMREAD_COLOR)
                contrast  = calculate_contrast(image2[y1:y2, x1:x2])
                print(f"{key} light3 图片对比度（标准差）：{contrast}")
                result = reader.readtext(image2)
                print([ {'text':item[1],'confident':item[2]} for item in result])

