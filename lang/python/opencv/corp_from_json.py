# 打开并读取JSON文件
import json
import cv2

with open(r"xx.json", 'r') as file:
    data = json.load(file)
