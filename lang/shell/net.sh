# 端口占用
netstat -ntulp
netstat -ntulp | grep 9090
lsof -i:9090