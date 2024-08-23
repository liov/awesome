package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// 文件路径，请替换为实际文件路径
	filePath := "/path/to/your/file.txt"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	buffer, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Read file error: %v", err)
	}
	// 创建HTTP POST请求的Body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("fileToUpload", filePath) // "fileToUpload" 应与服务器端接收的字段名匹配
	if err != nil {
		log.Fatalf("Create form file error: %v", err)
	}
	_, err = io.Copy(part, bytes.NewReader(buffer))
	if err != nil {
		log.Fatalf("Copy file to buffer error: %v", err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatalf("Close writer error: %v", err)
	}

	// 发送POST请求
	url := "http://example.com/upload" // 替换为实际的上传URL
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("Create request error: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Send request error: %v", err)
	}
	defer resp.Body.Close()

	// 打印服务器响应
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read response error: %v", err)
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(responseBody))
}
