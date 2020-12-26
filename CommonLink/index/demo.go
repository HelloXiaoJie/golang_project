// 使用普通的dial链接服务器
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"project/errorDispose"
	"strings"
)

func main() {
	// 链接到服务器
	connDial, error := net.Dial("tcp", "127.0.0.1:7777")
	defer func() {
		_ = connDial.Close()
		fmt.Println("链接关闭")
	}()
	errorDispose.ErrorPrint(error, "链接错误")
	textBufio := bufio.NewReader(os.Stdin)
	for {
		// 读取信息 消除空格
		var text [1024]byte
		n, _ := textBufio.Read(text[:])
		// 消除空格
		textRemoveSpace := strings.TrimSpace(string(text[:n]))
		// 发送信息
		_, error = connDial.Write([]byte(textRemoveSpace))
		errorDispose.ErrorPrint(error, "发送失败")
	}
}
