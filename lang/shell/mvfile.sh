#!/bin/bash
# 功能：每月1日将源目录中5个月前的文件转移到按年月命名的目标子目录
# 使用方式：需配置crontab定时执行（见第二部分）

# ---------- 配置区 ----------
SOURCE_DIR="/data/source_files"   # 源目录（需替换为实际路径）
TARGET_PARENT="/backup/old_files" # 目标父目录（需替换为实际路径）

# ---------- 核心逻辑 ----------
# 1. 计算5个月前的年份和月份（格式：YYYYMM）
FIVE_MONTHS_AGO=$(date -d "5 months ago" +%Y%m)

# 2. 创建目标目录（格式：/backup/old_files/YYYYMM）
TARGET_DIR="${TARGET_PARENT}/${FIVE_MONTHS_AGO}"
mkdir -p "${TARGET_DIR}" || exit 1  # 目录创建失败则退出

END_DATE=$(date -d "5 months ago" +\%Y-\%m-\%d)

# 4. 移动符合条件的文件
find "${SOURCE_DIR}" -type f \
    ! -newermt "${END_DATE}" \
    -exec mv -v {} "${TARGET_DIR}" \;

# 5. 记录日志（可选）
echo "[$(date +'%F %T')] 转移完成：${FIVE_MONTHS_AGO}月的文件已移至 ${TARGET_DIR}" >> /var/log/file_transfer.log

#crontab -e
#0 0 1 * * /bin/bash /path/to/your_script.sh >> /var/log/move_old_files.log 2>&1