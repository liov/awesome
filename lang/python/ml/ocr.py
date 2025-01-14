import easyocr
import cv2
reader = easyocr.Reader(['en']) # this needs to run only once to load the model into memory

image = cv2.imread(dir, cv2.IMREAD_GARY)
result = reader.readtext(image)