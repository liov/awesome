import cv2

def load_image(image_path):
    """加载图片并返回OpenCV图像对象"""
    image = cv2.imread(image_path, cv2.IMREAD_GRAYSCALE)
    return image

def calculate_similarity(overlap_region1, overlap_region2):
    """计算两个重合区域的相似度"""
    # 计算像素值差异的平方和
    diff = cv2.absdiff(overlap_region1, overlap_region2)
    similarity = cv2.mean(diff)[0]  # 返回平均差异
    return similarity

def find_overlap(image1, image2,row=False,min_overlap=1, max_overlap=100):
    """自动检测两张图片的纵向重合区域"""
    height1, width1 = image1.shape[:2]
    height2, width2 = image2.shape[:2]

    # 确保两张图片的宽度相同
    if width1 != width2:
        raise ValueError("两张图片的宽度必须相同")

    best_overlap = 0
    best_similarity = float('inf')

    for overlap in range(min_overlap, max_overlap + 1):
        # 提取重合区域
        if row:
            overlap_region1 = image1[-overlap:, :]
            overlap_region2 = image2[:overlap, :]
        else:
            overlap_region1 = image1[:, -overlap:]
            overlap_region2 = image2[:, :overlap]

        # 计算相似度
        similarity = calculate_similarity(overlap_region1, overlap_region2)

        # 更新最佳重合高度
        if similarity < best_similarity:
            best_similarity = similarity
            best_overlap = overlap

    return best_overlap, best_similarity

dir= r"xxx"
def col_overlap(col):
    overlaps = []
    for i in range(len(col)-1):
        # 图片路径
        image1_path = dir+rf"{col[i]}--light1.jpg"
        image2_path = dir+rf"{col[i+1]}--light1.jpg"

        # 加载图片
        image1 = load_image(image1_path)
        image2 = load_image(image2_path)

        best_overlap, best_similarity = find_overlap(image1, image2, True,300,2000)
        overlaps.append(best_overlap)
    print(f"重合度: {overlaps}")

def row_overlap(row):
    overlaps = []
    for i in range(len(row)-1):
        # 图片路径
        image1_path = dir+rf"{row[i]}--light1.jpg"
        image2_path = dir+rf"{row[i+1]}--light1.jpg"

        # 加载图片
        image1 = load_image(image1_path)
        image2 = load_image(image2_path)

        best_overlap, best_similarity = find_overlap(image1, image2, False,2000,4000)
        overlaps.append(best_overlap)
    print(f"重合度: {overlaps}")

col = [0,3,4,7]
col_overlap(col)
row=[0,1]
row_overlap(row)