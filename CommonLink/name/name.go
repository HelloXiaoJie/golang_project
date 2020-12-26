// 使用普通的listen链接
package main

import (
	"fmt"
	"net"
	"time"
	"project/errorDispose"
)

func main() {
	Listen, error := net.Listen("tcp", "127.0.0.1:7777")
	defer func() {
		_ = Listen.Close()
	}()
	errorDispose.ErrorPrint(error, "监听错误")
	// 监听端口信息
	fmt.Println(Listen.Addr().String())
	fmt.Println(Listen.Addr().Network())
	fmt.Println("等待链接:")
	for {
		conn, error := Listen.Accept()
		errorDispose.ErrorPrint(error, "连接错误")
		fmt.Println(conn.RemoteAddr().String(), "链接成功")
		go ConnLink(conn)
	}
}


func ConnLink(conn net.Conn) {
	defer func() {
		_ = conn.Close()
		fmt.Println(conn.RemoteAddr().String(), "链接关闭")
	}()
	TimeContent := time.Now().Format("2006-01-02 15:04:05")
	_, error := conn.Write([]byte("和服务器链接时间:->" + TimeContent))
	errorDispose.ErrorPrint(error, "发送时间错误")
	// 等待客户端信息
	for {
		text := new([1024]byte)
		n, error := conn.Read(text[:])
		if error != nil {
			fmt.Println(error, "接受客户端数据错误")
			return
		}
		fmt.Println(conn.RemoteAddr().String(), "->", string(text[:n]))
	}
}
