#!/bin/bash

cd $1

# 获取最后一次提交的日期时间戳
last_commit_time=$(git log -1 --format=%ct)

# 将时间戳转换为日期时间格式（这里使用UTC时间）
formatted_last_commit_time=$(date -u -d @$last_commit_time +"%Y-%m-%d %H:%M:%S")

echo "Last commit was at: $formatted_last_commit_time UTC"

# 计算上次提交是否在周一至周五的9:30到18:30之间
day_of_week=$(( ($(date -d @$last_commit_time +%u) % 7) ))
hour=$(date -d @$last_commit_time +%H)

if [[ $day_of_week -ge 1 && $day_of_week -le 5 ]] && [[ $hour -ge 9 && $hour -lt 18 ]]; then
    echo "The last commit was within the target time range."

    # 如果满足条件，将时间戳减去10小时
    adjusted_time="$(date -d '-10 hours' '+%Y-%m-%d %H:%M:%S')"

    # 注意：下面的步骤不建议在真实环境中执行，因为这会修改Git历史
    # 若要演示，可以打印出将要执行的命令而非真正执行
    echo "Would run: git commit --amend --date=$adjusted_time --no-edit"
    # 实际执行命令应替换上一行注释的命令如下
    git commit --amend --date=$adjusted_time --no-edit
else
    echo "The last commit was outside the target time range."
fi