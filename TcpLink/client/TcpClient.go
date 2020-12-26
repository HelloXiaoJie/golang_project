package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
import "project/errorDispose"

func main() {
	// 使用tcp链接服务器
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":7777")
	tcpConn, error := net.DialTCP("tcp", nil, tcpAddr)
	defer func() {
		_ = tcpConn.Close()
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
		_, error = tcpConn.Write([]byte(textRemoveSpace))
		//errorDispose.ErrorPrint(error, "发送失败")
		if error != nil {
			fmt.Println("发送失败")
			return
		}
	}
}
