function update_commit_time_if_needed() {
  # 获取最后一次提交的哈希值
  last_commit_hash=$(git log -1 --format=%H)

  # 获取最后一次提交的时间
  last_commit_time=$(git show -s --format=%ci $last_commit_hash) # ci提交时间 ai 作者时间
  echo "最后一次提交时间: $last_commit_time"
  # 转换时间格式
  commit_time=$(date -d "$last_commit_time" +"%Y-%m-%d %H:%M:%S")
  current_time=$(date +"%Y-%m-%d %H:%M:%S")

  # 获取星期几，1-5 表示周一到周五，0和6表示周日和周六
  weekday=$(date -d "$commit_time" +%u)

  # 获取小时数
  hour=$(date -d "$commit_time" +%H)

  # 检查是否在周一到周五的9点到19点之间
  if [[ $weekday -ge 1 && $weekday -le 5 && $hour -ge 9 && $hour -lt 19 ]]; then
      # 计算当前时间前十个小时的时间
      new_time=$(date -d '-10 hours' '+%Y-%m-%d %H:%M:%S')

      # 修改最后一次提交的时间
      GIT_AUTHOR_DATE="$new_time" GIT_COMMITTER_DATE="$new_time" git commit --amend --no-edit --date "$new_time"
      echo "提交时间已修改为: $new_time"
  else
      echo "最后一次提交时间不在周一到周五的9-19点之间，无需修改。"
  fi
}

update_commit_time_if_needed