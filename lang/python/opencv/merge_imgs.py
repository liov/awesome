import cv2
import numpy as np

horizontalOverlaps=[]
verticalOverlaps=[]

def merge_img(imgs,width,height,horizontalOverlaps,verticalOverlaps,dst):
    resultWidth,resultHeight = 0,0
    horizontalPixels,verticalPixels = [],[]
    for i in range(imgs[0]):
        resultWidth += width
        if i < len(horizontalOverlaps):
            resultWidth -= horizontalOverlaps[i]
        horizontalPixels.append(resultWidth)

    for i in range(imgs):
        resultHeight += height
        if i < len(verticalOverlaps):
            resultHeight -= verticalOverlaps[i]
        verticalPixels.append(resultHeight)

    img0 = cv2.imread(imgs[0], cv2.IMREAD_ANYCOLOR|cv2.IMREAD_ANYDEPTH)
    original_shape = img0.shape
    original_channels = original_shape[2] if len(original_shape) == 3 else 1
    result = np.zeros((resultHeight, resultWidth, original_channels), dtype=img0.dtype)

