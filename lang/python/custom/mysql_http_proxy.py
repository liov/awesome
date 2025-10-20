import socket
import socks  # pip install pysocks

# 配置 HTTP 代理（需代理支持 SOCKS 转换或 CONNECT）
proxy_ip = "10.48.23.75"
proxy_port = 3128

# 将 HTTP 代理转为 SOCKS 协议（需代理支持）
socks.set_default_proxy(
    socks.HTTP,  # 使用 HTTP 代理模拟 SOCKS
    proxy_ip,
    proxy_port
)
socket.socket = socks.socksocket  # 全局替换 socket

# 尝试连接 MySQL
try:
    s = socket.socket()
    s.settimeout(10)
    s.connect(("8.222.139.120", 3306))  # 替换为你的 MySQL 地址
    greeting = s.recv(1024)
    print("MySQL Greeting:", greeting.hex())
    s.close()
except Exception as e:
    print("连接失败:", e)