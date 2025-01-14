import easyocr
import cv2

reader = easyocr.Reader(['ch_sim','en']) # this needs to run only once to load the model into memory
image = cv2.imread('/home/jyb/Pictures/20250114182855.jpg', cv2.IMREAD_GRAYSCALE)
result = reader.readtext(image)
print(result)