import cv2
import numpy as np
import json
import easyocr
reader = easyocr.Reader(['ch_sim','en']) # this needs to run only once to load the model into memory

dir= r"xxx"
def calculate_contrast(image):
    # 读取图像，转换为灰度图
    gary = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    # 计算图像像素亮度的标准差
    contrast = np.std(gary)

    print(f"图片对比度（标准差）：{contrast}")



with open(dir+r"\board\board.json", 'r') as file:
    data = json.load(file)
    for  key, value in data['inlineWinsMap'].items():
        for item in value:
            if item['winType'] == 'silk':
                x1, y1 = item['xPx'] - item['lengthPx']/2, item['yPx']- item['widthPx']/2
                x2, y2 = item['xPx'] + item['lengthPx']/2, item['yPx']+ item['widthPx']/2
                image1 = cv2.imread(dir+rf"\board\outline-win\{key}--light1.jpg")
                calculate_contrast(image1[y1:y2, x1:x2])
                result = reader.readtext(image1)

