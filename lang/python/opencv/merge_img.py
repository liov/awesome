import cv2
import numpy as np
dir=r"xxx"
overlap_width =2845
 # 1. 读取图像
img1 = cv2.imread(dir+r"0--light1.jpg")
img2 = cv2.imread(dir+r"1--light1.jpg")

# 拼合
result = np.hstack((img1[:, :-overlap_width], img2))
cv2.imwrite(dir+r"merge.jpg", result)