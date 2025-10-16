#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
HEIF图像二值化处理，将白色区域变为透明
"""

import cv2
import numpy as np
from PIL import Image
import pillow_heif

def binarize_heif_image(input_path, output_path, threshold=200):
    """
    将HEIF图像进行二值化处理，并将白色区域变为透明
    
    Args:
        input_path (str): 输入HEIF图像路径
        output_path (str): 输出图像路径（PNG格式）
        threshold (int): 二值化阈值，默认200
    """
    # 注册HEIF格式
    pillow_heif.register_heif_opener()
    
    # 读取HEIF图像
    image = Image.open(input_path)
    
    # 转换为OpenCV格式
    # 如果图像有透明通道，需要特殊处理
    if image.mode == 'RGBA':
        # 分离RGBA通道
        r, g, b, a = image.split()
        # 合并RGB通道
        rgb_image = Image.merge('RGB', (r, g, b))
        # 转换为OpenCV格式
        cv_image = cv2.cvtColor(np.array(rgb_image), cv2.COLOR_RGB2BGR)
        # 保存alpha通道用于后续处理
        alpha_channel = np.array(a)
    else:
        # 转换为OpenCV格式
        cv_image = cv2.cvtColor(np.array(image), cv2.COLOR_RGB2BGR)
        alpha_channel = None
    
    # 转换为灰度图
    gray = cv2.cvtColor(cv_image, cv2.COLOR_BGR2GRAY)
    
    # 应用阈值进行二值化处理
    _, binary = cv2.threshold(gray, threshold, 255, cv2.THRESH_BINARY)
    
    # 创建透明通道，将白色区域变为透明
    # 白色区域在二值图像中为255，需要变为透明(0)
    # 黑色区域在二值图像中为0，需要保持不透明(255)
    alpha = cv2.bitwise_not(binary)
    
    # 如果原图像有alpha通道，需要与新alpha通道结合
    if alpha_channel is not None:
        # 结合两个alpha通道
        alpha = cv2.bitwise_and(alpha, alpha_channel)
    
    # 将二值化结果转换为三通道
    binary_bgr = cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR)
    
    # 添加alpha通道
    b, g, r = cv2.split(binary_bgr)
    rgba = cv2.merge([b, g, r, alpha])
    
    # 保存结果
    result_image = Image.fromarray(rgba)
    result_image.save(output_path, 'PNG')
    
    print(f"处理完成，结果已保存到: {output_path}")

def binarize_heif_image_advanced(input_path, output_path, threshold=200, kernel_size=3):
    """
    高级版本的HEIF图像二值化处理，包含形态学操作
    
    Args:
        input_path (str): 输入HEIF图像路径
        output_path (str): 输出图像路径（PNG格式）
        threshold (int): 二值化阈值，默认200
        kernel_size (int): 形态学操作的核大小，默认3
    """
    # 注册HEIF格式
    pillow_heif.register_heif_opener()
    
    # 读取HEIF图像
    image = Image.open(input_path)
    
    # 转换为OpenCV格式
    if image.mode == 'RGBA':
        # 分离RGBA通道
        r, g, b, a = image.split()
        # 合并RGB通道
        rgb_image = Image.merge('RGB', (r, g, b))
        # 转换为OpenCV格式
        cv_image = cv2.cvtColor(np.array(rgb_image), cv2.COLOR_RGB2BGR)
        # 保存alpha通道用于后续处理
        alpha_channel = np.array(a)
    else:
        # 转换为OpenCV格式
        cv_image = cv2.cvtColor(np.array(image), cv2.COLOR_RGB2BGR)
        alpha_channel = None
    
    # 转换为灰度图
    gray = cv2.cvtColor(cv_image, cv2.COLOR_BGR2GRAY)
    
    # 应用阈值进行二值化处理
    _, binary = cv2.threshold(gray, threshold, 255, cv2.THRESH_BINARY)
    
    # 创建形态学操作的核
    kernel = np.ones((kernel_size, kernel_size), np.uint8)
    
    # 进行形态学操作，去除噪声
    # 先腐蚀再膨胀（开运算）
    binary = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)
    # 先膨胀再腐蚀（闭运算）
    binary = cv2.morphologyEx(binary, cv2.MORPH_CLOSE, kernel)
    
    # 创建透明通道，将白色区域变为透明
    alpha = cv2.bitwise_not(binary)
    
    # 如果原图像有alpha通道，需要与新alpha通道结合
    if alpha_channel is not None:
        # 结合两个alpha通道
        alpha = cv2.bitwise_and(alpha, alpha_channel)
    
    # 将二值化结果转换为三通道
    binary_bgr = cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR)
    
    # 添加alpha通道
    b, g, r = cv2.split(binary_bgr)
    rgba = cv2.merge([b, g, r, alpha])
    
    # 保存结果
    result_image = Image.fromarray(rgba)
    result_image.save(output_path, 'PNG')
    
    print(f"高级处理完成，结果已保存到: {output_path}")

# 使用示例
if __name__ == "__main__":
    # 示例用法
    binarize_heif_image_advanced("/Users/jyb/Downloads/IMG_20251016_120019.HEIF", "/Users/jyb/Downloads/IMG_20251016_120019.png")
    # binarize_heif_image_advanced("input.heic", "output_advanced.png")
    pass