import socket
import select
import socketserver  # 添加这行！
from socketserver import ThreadingMixIn, TCPServer, BaseRequestHandler

# 配置远端 HTTP 代理
HTTP_PROXY_HOST = "10.48.23.75"  # 远端 HTTP 代理地址
HTTP_PROXY_PORT = 3128                # 远端 HTTP 代理端口

class ThreadedTCPServer(ThreadingMixIn, socketserver.TCPServer):
    """多线程 SOCKS5 服务器"""
    pass

class SocksProxyHandler(socketserver.BaseRequestHandler):
    def handle(self):
        try:
            # SOCKS5 握手
            self.request.recv(1024)  # 读取客户端握手请求
            self.request.sendall(b"\x05\x00")  # 返回 SOCKS5 认证方式：无需认证

            # 读取客户端连接请求
            data = self.request.recv(1024)
            if data[0] != 0x05 or data[1] != 0x01:
                raise ValueError("非 SOCKS5 协议")

            # 解析目标地址（IP/域名 + 端口）
            addr_type = data[3]
            if addr_type == 0x01:  # IPv4
                target_host = socket.inet_ntoa(data[4:8])
                target_port = int.from_bytes(data[8:10], "big")
            elif addr_type == 0x03:  # 域名
                domain_len = data[4]
                target_host = data[5:5+domain_len].decode()
                target_port = int.from_bytes(data[5+domain_len:7+domain_len], "big")
            else:
                raise ValueError("不支持的地址类型")
            print(f"[*] 连接到: {target_host}:{target_port}")
            # 通知客户端连接成功
            self.request.sendall(b"\x05\x00\x00\x01\x00\x00\x00\x00\x00\x00")

            # 4. 通过 HTTP 代理建立隧道（HTTPS 必须用 CONNECT）
            proxy_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            proxy_sock.connect((HTTP_PROXY_HOST, HTTP_PROXY_PORT))

            # 发送 CONNECT 请求到 HTTP 代理
            proxy_sock.sendall(
                f"CONNECT {target_host}:{target_port} HTTP/1.1\r\n"
                f"Host: {target_host}:{target_port}\r\n\r\n".encode()
            )
            # 检查代理响应（必须返回 200）
            response = proxy_sock.recv(4096)
            if b"200 Connection established" not in response:
                raise ValueError(f"代理拒绝 CONNECT: {response.decode()}")
            # 双向转发数据（客户端 ↔ 远端 HTTP 代理）
            while True:
                r, _, _ = select.select([self.request, proxy_sock], [], [])
                if self.request in r:
                    data = self.request.recv(4096)
                    if not data:
                        break
                    proxy_sock.sendall(data)
                if proxy_sock in r:
                    data = proxy_sock.recv(4096)
                    if not data:
                        break
                    self.request.sendall(data)

        except Exception as e:
            print(f"代理错误: {e}")
        finally:
            self.request.close()

if __name__ == "__main__":
    # 启动本地 SOCKS5 代理服务器
    HOST, PORT = "127.0.0.1", 1080
    server = ThreadedTCPServer((HOST, PORT), SocksProxyHandler)
    print(f"[*] 本地 SOCKS5 代理已启动: socks5://{HOST}:{PORT}")
    print(f"[*] 所有流量将通过 HTTP 代理 ({HTTP_PROXY_HOST}:{HTTP_PROXY_PORT}) 转发")
    server.serve_forever()