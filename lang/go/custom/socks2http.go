package main

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// SOCKS5 认证方法
const (
	MethodNoAuth       byte = 0x00
	MethodGSSAPI       byte = 0x01
	MethodUsername     byte = 0x02
	MethodNoAcceptable byte = 0xFF
)

// SOCKS5 命令
const (
	CmdConnect      byte = 0x01
	CmdBind         byte = 0x02
	CmdUDPAssociate byte = 0x03
)

// SOCKS5 地址类型
const (
	ATYPIPv4   byte = 0x01
	ATYPDomain byte = 0x03
	ATYPIPv6   byte = 0x04
)

// SOCKS5 回复状态
const (
	StatusSuccess byte = 0x00
	StatusFailure byte = 0x01
)

type Config struct {
	ListenAddr string
	HTTPProxy  string
	Username   string
	Password   string
	Timeout    time.Duration
}

type SOCKS5Proxy struct {
	config     *Config
	httpClient *http.Client
}

func NewSOCKS5Proxy(config *Config) *SOCKS5Proxy {
	// 创建 HTTP 传输层，使用指定的 HTTP 代理
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(config.HTTPProxy)
		},
		DialContext: (&net.Dialer{
			Timeout:   config.Timeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &SOCKS5Proxy{
		config: config,
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   config.Timeout,
		},
	}
}

func (s *SOCKS5Proxy) Start() error {
	listener, err := net.Listen("tcp", s.config.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", s.config.ListenAddr, err)
	}
	defer listener.Close()

	log.Printf("SOCKS5 proxy started on %s, forwarding to HTTP proxy: %s", s.config.ListenAddr, s.config.HTTPProxy)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *SOCKS5Proxy) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 设置超时
	conn.SetDeadline(time.Now().Add(s.config.Timeout))

	// SOCKS5 握手
	if err := s.handshake(conn); err != nil {
		log.Printf("Handshake failed: %v", err)
		return
	}

	// 读取请求
	targetAddr, err := s.readRequest(conn)
	if err != nil {
		log.Printf("Read request failed: %v", err)
		return
	}

	// 连接到目标（通过 HTTP 代理）
	targetConn, err := s.connectThroughHTTPProxy(targetAddr)
	if err != nil {
		log.Printf("Connect to %s failed: %v", targetAddr, err)
		s.sendReply(conn, StatusFailure, nil)
		return
	}
	defer targetConn.Close()

	// 发送成功响应
	if err := s.sendReply(conn, StatusSuccess, nil); err != nil {
		log.Printf("Send reply failed: %v", err)
		return
	}

	// 开始数据转发
	s.relay(conn, targetConn)
}

func (s *SOCKS5Proxy) handshake(conn net.Conn) error {
	// 读取客户端支持的认证方法
	buf := make([]byte, 257) // 最大 255 个方法
	if _, err := io.ReadFull(conn, buf[:2]); err != nil {
		return err
	}

	if buf[0] != 0x05 {
		return errors.New("unsupported SOCKS version")
	}

	nmethods := int(buf[1])
	if _, err := io.ReadFull(conn, buf[:nmethods]); err != nil {
		return err
	}

	// 选择认证方法（这里只支持无认证）
	var method byte = MethodNoAcceptable
	for i := 0; i < nmethods; i++ {
		if buf[i] == MethodNoAuth {
			method = MethodNoAuth
			break
		}
	}

	// 发送选择的认证方法
	response := []byte{0x05, method}
	if _, err := conn.Write(response); err != nil {
		return err
	}

	if method == MethodNoAcceptable {
		return errors.New("no acceptable authentication method")
	}

	return nil
}

func (s *SOCKS5Proxy) readRequest(conn net.Conn) (string, error) {
	buf := make([]byte, 262) // 最大地址长度 255 + 其他字段

	// 读取请求头
	if _, err := io.ReadFull(conn, buf[:4]); err != nil {
		return "", err
	}

	if buf[0] != 0x05 {
		return "", errors.New("unsupported SOCKS version")
	}

	cmd := buf[1]
	if cmd != CmdConnect {
		return "", errors.New("only CONNECT command is supported")
	}

	// 读取目标地址
	atyp := buf[3]
	var host string
	var port uint16

	switch atyp {
	case ATYPIPv4:
		if _, err := io.ReadFull(conn, buf[:6]); err != nil {
			return "", err
		}
		host = net.IPv4(buf[0], buf[1], buf[2], buf[3]).String()
		port = binary.BigEndian.Uint16(buf[4:6])

	case ATYPDomain:
		if _, err := io.ReadFull(conn, buf[:1]); err != nil {
			return "", err
		}
		domainLen := int(buf[0])
		if _, err := io.ReadFull(conn, buf[:domainLen+2]); err != nil {
			return "", err
		}
		host = string(buf[:domainLen])
		port = binary.BigEndian.Uint16(buf[domainLen : domainLen+2])

	case ATYPIPv6:
		if _, err := io.ReadFull(conn, buf[:18]); err != nil {
			return "", err
		}
		host = net.IP(buf[:16]).String()
		port = binary.BigEndian.Uint16(buf[16:18])

	default:
		return "", errors.New("unsupported address type")
	}

	return net.JoinHostPort(host, strconv.Itoa(int(port))), nil
}

func (s *SOCKS5Proxy) connectThroughHTTPProxy(targetAddr string) (net.Conn, error) {
	// 使用 HTTP CONNECT 方法通过 HTTP 代理建立连接
	req, err := http.NewRequest(http.MethodConnect, "http://"+targetAddr, nil)
	if err != nil {
		return nil, err
	}

	// 设置代理认证（如果需要）
	if s.config.Username != "" && s.config.Password != "" {
		req.SetBasicAuth(s.config.Username, s.config.Password)
	}

	// 直接使用 DialContext 建立连接，而不是发送 HTTP 请求
	transport := s.httpClient.Transport.(*http.Transport)
	return transport.DialContext(context.Background(), "tcp", targetAddr)
}

func (s *SOCKS5Proxy) sendReply(conn net.Conn, status byte, bindAddr net.Addr) error {
	reply := make([]byte, 10)
	reply[0] = 0x05 // SOCKS version
	reply[1] = status
	reply[2] = 0x00 // RSV
	reply[3] = 0x01 // IPv4

	// 对于成功连接，我们不需要绑定特定地址
	if bindAddr == nil {
		copy(reply[4:8], []byte{0, 0, 0, 0}) // 0.0.0.0
		copy(reply[8:10], []byte{0, 0})      // 端口 0
	} else {
		// 这里可以处理实际的绑定地址
	}

	_, err := conn.Write(reply[:10])
	return err
}

func (s *SOCKS5Proxy) relay(client, target net.Conn) {
	// 重置超时设置
	client.SetDeadline(time.Time{})
	target.SetDeadline(time.Time{})

	errCh := make(chan error, 2)

	// 客户端 -> 目标
	go func() {
		_, err := io.Copy(target, client)
		errCh <- err
	}()

	// 目标 -> 客户端
	go func() {
		_, err := io.Copy(client, target)
		errCh <- err
	}()

	// 等待任一方向出错
	err := <-errCh
	log.Println("relay error:", err)
}

func main() {
	config := &Config{
		ListenAddr: ":1080",
		HTTPProxy:  "http://your-http-proxy:8080", // 替换为你的 HTTP 代理地址
		Username:   "",                            // HTTP 代理用户名（如果需要）
		Password:   "",                            // HTTP 代理密码（如果需要）
		Timeout:    30 * time.Second,
	}

	proxy := NewSOCKS5Proxy(config)

	log.Printf("Starting SOCKS5 to HTTP proxy server...")
	if err := proxy.Start(); err != nil {
		log.Fatalf("Failed to start proxy: %v", err)
	}
}
