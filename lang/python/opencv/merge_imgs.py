import cv2
import numpy as np

horizontalOverlaps=[]
verticalOverlaps=[]

dir = r"xxx"
def merge_img(imgs, horizontal_overlaps, vertical_overlaps, dst):
    result_width,result_height = 0,0
    horizontal_pixels,vertical_pixels = [0],[0]
    img0 = cv2.imread(dir+str(imgs[0][0])+"--light2.jpg", cv2.IMREAD_UNCHANGED)
    height,width = img0.shape[:2]
    for i in range(len(imgs[0])):
        result_width += width
        if i < len(horizontal_overlaps):
            result_width -= horizontal_overlaps[i]
            horizontal_pixels.append(result_width)

    for i in range(len(imgs)):
        result_height += height
        if i < len(vertical_overlaps):
            result_height -= vertical_overlaps[i]
            vertical_pixels.append(result_height)

    print(result_height, result_width,horizontal_pixels, vertical_pixels)

    original_channels = img0.shape[2] if len(img0.shape) == 3 else 1
    result = np.zeros((result_height, result_width, original_channels), dtype=img0.dtype)
    for i in range(len(imgs)):
        for j in range(len(imgs[i])):
            img = cv2.imread(dir+str(imgs[i][j])+"--light2.jpg", cv2.IMREAD_UNCHANGED)
            result[vertical_pixels[i]:vertical_pixels[i]+height, horizontal_pixels[j]:horizontal_pixels[j]+width] = img
    cv2.imwrite(dir+dst, result)

merge_img([[
    0,
    1
],
    [
        3,
        2
    ],
    [
        4,
        5
    ]],[3770],[357, 358], 'result.jpg')