import socket
import time

# 不知道为啥,nebula就是连不上 mysql
# 手动完成TCP+MySQL握手
s = socket.socket()
s.connect(('8.222.139.120', 3306))


# 设置非阻塞或设置超时
s.settimeout(5.0)  # 设置5秒超时

try:
    greeting = s.recv(1024)
    print("Received:", greeting.hex())

    # 保持连接观察
    time.sleep(10)

except socket.timeout:
    print("Timeout waiting for data")
finally:
    s.close()