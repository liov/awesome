# 方法1：使用 shell 参数扩展去除 .conf 后缀
for conf in room.cmd.*.conf; do
    if [ -f "$conf" ]; then
        process_name="${conf%.conf}"
        supervisorctl stop "$process_name"
    fi
done

# 方法2：使用 basename 命令
for conf in room.cmd.*.conf; do
    if [ -f "$conf" ]; then
        process_name=$(basename "$conf" .conf)
        supervisorctl stop "$process_name"
    fi
done

# 方法3：一行命令批量处理
supervisorctl stop $(ls room.cmd.*.conf 2>/dev/null | sed 's/\.conf$//')
