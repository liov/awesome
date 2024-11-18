import cv2
import numpy as np
import json
import easyocr
reader = easyocr.Reader(['en']) # this needs to run only once to load the model into memory

def calculate_sharpness_brightness_contrast(image):
    # 转换为灰度图
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    # 计算对比度（灰度值的标准差）
    contrast = np.std(gray)
    # 使用拉普拉斯变换检测清晰度（方差越大越清晰）
    sharpness = cv2.Laplacian(gray, cv2.CV_32F).var()
    return sharpness, contrast

def find_best_image(image_files):
    best_image = None
    best_score = -1

    for file in image_files:
        # 读取图片
        image = cv2.imread(file)

        # 计算亮度和对比度
        sharpness,brightness, contrast = calculate_sharpness_brightness_contrast(image)
        # 计算评分：可以调整权重
        score = sharpness * 0.7 + contrast * 0.2 + brightness * 0.1

        print(f"{file}: 清晰度={sharpness:.2f}, 亮度={brightness:.2f}, 对比度={contrast:.2f}, 总评分={score:.2f}")

        # 更新最佳图像
        if score > best_score:
            best_score = score
            best_image = image

    return best_image

dir= r"xxx"
with open(dir+r"\board\board.json", 'r') as file:
    data = json.load(file)
    for  key, value in data['inlineWinsMap'].items():
        if not key.startswith('C'):
            continue
        for item in value:
            if item['winType'] == 'silk':
                x1, y1 = int(item['xPx']) - int(item['lengthPx']/2), int(item['yPx'])- int(item['widthPx']/2)
                x2, y2 = int(item['xPx']) + int(item['lengthPx']/2), int(item['yPx'])+ int(item['widthPx']/2)
                image1 = cv2.imread(dir+rf"\board\outline-win\{key}--light2.jpg", cv2.IMREAD_COLOR)
                sharpness, contrast = calculate_sharpness_brightness_contrast(image1[y1:y2, x1:x2])
                # 计算评分：可以调整权重
                score = sharpness * 0.7 + contrast * 0.2
                print(f"{key} light2  清晰度={sharpness:.2f}, 对比度={contrast:.2f}, 总评分={score:.2f}")
                result = reader.readtext(image1)
                print([ {'text':item[1],'confident':item[2]} for item in result])
                image2 = cv2.imread(dir+rf"\board\outline-win\{key}--light3.jpg", cv2.IMREAD_COLOR)
                sharpness, contrast  = calculate_sharpness_brightness_contrast(image2[y1:y2, x1:x2])
                score = sharpness * 0.7 + contrast * 0.2
                print(f"{key} light3  清晰度={sharpness:.2f},  对比度={contrast:.2f}, 总评分={score:.2f}")
                result = reader.readtext(image2)
                print([ {'text':item[1],'confident':item[2]} for item in result])
                break

