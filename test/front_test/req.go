package front_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseUrl = "http://127.0.0.1:8080/front/"

func Req(url string, method string, body io.Reader) {
	req, err := http.NewRequest(method, baseUrl+url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应的状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Status code %d\n", resp.StatusCode)
		return
	}

	// 直接使用os.Stdout来将响应体写入标准输出
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
}

//写一个函数用来发送POST请求
