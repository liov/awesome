const crypto = require('crypto');
const fs = require('fs');

async function calculateFileHash(filePath) {
    try {
        // 读取文件内容
        const fileContent = await fs.promises.readFile(filePath);

        // 创建一个SHA-256哈希对象
        const hash = crypto.createHash('sha256');

        // 更新哈希对象的内容
        hash.update(fileContent);

        // 计算哈希值
        const hashValue = hash.digest('hex');

        console.log(`File ${filePath} hash: ${hashValue}`);
    } catch (error) {
        console.error('Error calculating file hash:', error);
    }
}

// 调用函数，传入文件路径
calculateFileHash("D:\\SDK\\msys64\\ucrt64\\x86_64-w64-mingw32\\bin\\ld.exe");
calculateFileHash("D:\\SDK\\msys64\\ucrt64\\x86_64-w64-mingw32\\bin\\nm.exe");
calculateFileHash("D:\\SDK\\msys64\\ucrt64\\bin\\ld.exe");