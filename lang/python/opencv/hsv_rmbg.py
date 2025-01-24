import cv2
import numpy as np

def nothing(x):
    pass

# 创建一个黑色背景的窗口
cv2.namedWindow('Trackbars', cv2.WINDOW_AUTOSIZE)

# 创建滑动条用于调整HSV阈值
cv2.createTrackbar('H Low', 'Trackbars', 0, 180, nothing)
cv2.createTrackbar('H High', 'Trackbars', 180, 180, nothing)
cv2.createTrackbar('S Low', 'Trackbars', 0, 255, nothing)
cv2.createTrackbar('S High', 'Trackbars', 255, 255, nothing)
cv2.createTrackbar('V Low', 'Trackbars', 0, 255, nothing)
cv2.createTrackbar('V High', 'Trackbars', 255, 255, nothing)
cv2.createTrackbar('T Low', 'Trackbars', 50, 255, nothing)
image = cv2.imread(r"D:\work.jpg", cv2.IMREAD_COLOR)
while True:
    # 假设 image 是你的输入图像
    image = cv2.resize(image, (3928,2195), interpolation=cv2.INTER_AREA)
    hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)

    # 获取当前滑动条的位置
    h_low = cv2.getTrackbarPos('H Low', 'Trackbars')
    h_high = cv2.getTrackbarPos('H High', 'Trackbars')
    s_low = cv2.getTrackbarPos('S Low', 'Trackbars')
    s_high = cv2.getTrackbarPos('S High', 'Trackbars')
    v_low = cv2.getTrackbarPos('V Low', 'Trackbars')
    v_high = cv2.getTrackbarPos('V High', 'Trackbars')
    t_low = cv2.getTrackbarPos('T Low', 'Trackbars')

    # 定义HSV颜色范围
    lower_bound = np.array([h_low, s_low, v_low])
    upper_bound = np.array([h_high, s_high, v_high])

    # 根据颜色范围创建掩码
    mask = cv2.inRange(hsv, lower_bound, upper_bound)
    mask_inverted = cv2.bitwise_not(mask)  # 取反，保留非黑色部分

    # 应用掩码，去除底版黑色
    background_removed = cv2.bitwise_and(image, image, mask=mask_inverted)
    gray = cv2.cvtColor(background_removed, cv2.COLOR_BGR2GRAY)
    # 二值化
    _, binary = cv2.threshold(gray, t_low, 255, cv2.THRESH_BINARY)
    # 显示结果
    cv2.imshow('Mask', binary)

    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

cv2.destroyAllWindows()