package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/upload", UploadFile)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/favicon.ico")
	})

	port := 9999
	fmt.Printf("浏览器运行 ht%s://127.0.0.1:%d", "tp", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}

func HomePage(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	_ = t.Execute(w, nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析上传的文件
	file, tempFile, err := parseAndSaveFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(file multipart.File) {
		_ = file.Close()
	}(file)
	defer func(tempFile *os.File) {
		_ = tempFile.Close()
	}(tempFile)

	// 检查签名
	signature, err := checkSignature(tempFile)
	if err != nil {
		http.Error(w, "Error Running Keytool Command", http.StatusInternalServerError)
		return
	}

	// 返回 JSON 数据
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, `{"signature": %q}`, signature)
}

func parseAndSaveFile(r *http.Request) (multipart.File, *os.File, error) {
	err := r.ParseMultipartForm(500 << 20)
	if err != nil {
		return nil, nil, fmt.Errorf("file too big: %w", err)
	} // 限制上传大小为500MB
	file, _, err := r.FormFile("apkfile")
	if err != nil {
		return nil, nil, fmt.Errorf("error retrieving the file")
	}

	tempFile, err := os.CreateTemp("", "upload-*.apk")
	if err != nil {
		return nil, nil, fmt.Errorf("error creating temp file")
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, nil, fmt.Errorf("error writing file content")
	}

	return file, tempFile, nil
}

func checkSignature(tempFile *os.File) (string, error) {
	cmd := exec.Command("keytool", "-printcert", "-jarfile", tempFile.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running keytool command")
	}

	// 将 GBK 编码转换为 UTF-8
	reader := transform.NewReader(&out, simplifiedchinese.GBK.NewDecoder())
	utf8Bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("error decoding GBK output: %v", err)
	}

	return string(utf8Bytes), nil
}
