import cv2
import numpy as np
import glob

def calculate_sharpness(image):
    # 转换为灰度图
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    # 使用拉普拉斯变换检测清晰度（方差越大越清晰）
    laplacian_var = cv2.Laplacian(gray, cv2.CV_64F).var()
    return laplacian_var

def calculate_brightness_contrast(image):
    # 转换为灰度图
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    # 计算亮度均值
    brightness = np.mean(gray)
    # 计算对比度（灰度值的标准差）
    contrast = np.std(gray)
    return brightness, contrast

def find_best_image(image_files):
    best_image = None
    best_score = -1

    for file in image_files:
        # 读取图片
        image = cv2.imread(file)

        # 计算清晰度
        sharpness = calculate_sharpness(image)
        # 计算亮度和对比度
        brightness, contrast = calculate_brightness_contrast(image)

        # 计算评分：可以调整权重
        score = sharpness * 0.7 + contrast * 0.2 + brightness * 0.1

        print(f"{file}: 清晰度={sharpness:.2f}, 亮度={brightness:.2f}, 对比度={contrast:.2f}, 总评分={score:.2f}")

        # 更新最佳图像
        if score > best_score:
            best_score = score
            best_image = image

    return best_image

# 获取所有图像文件（修改为图像所在路径）
image_files = glob.glob("images/*.jpg")

# 找到最佳图像
best_image = find_best_image(image_files)

if best_image is not None:
    cv2.imshow("Best Image", best_image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()
else:
    print("未找到最佳图像")
