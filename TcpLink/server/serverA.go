package main

import (
	"fmt"
	"net"
	"time"
)
import "project/errorDispose"

func main() {
	// 使用tcp
	TcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7777")
	TcpLink, error := net.ListenTCP("tcp", TcpAddr)
	defer func() {
		_ = TcpLink.Close()
	}()
	errorDispose.ErrorPrint(error, "监听错误")
	fmt.Println(TcpLink.Addr().String())
	fmt.Println(TcpLink.Addr().Network())
	fmt.Println("等待链接:")
	for {
		acceptTcp, _ := TcpLink.AcceptTCP()
		fmt.Println(acceptTcp.RemoteAddr())
		fmt.Println(acceptTcp.LocalAddr())
		//conn, error := TcpLink.Accept()
		//errorDispose.ErrorPrint(error, "连接错误")
		//fmt.Println(conn.RemoteAddr().String(), "链接成功")
		//go ConnLink(conn)
		go ConnLink1(acceptTcp)
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

func ConnLink1(acceptTcp *net.TCPConn) {
	defer func() {
		_ = acceptTcp.Close()
		fmt.Println(acceptTcp.RemoteAddr().String(), "链接关闭")
	}()
	TimeContent := time.Now().Format("2006-01-02 15:04:05")
	_, error := acceptTcp.Write([]byte("和服务器链接时间:->" + TimeContent))
	errorDispose.ErrorPrint(error, "发送时间错误")
	// 等待客户端信息
	for {
		text := new([1024]byte)
		n, error := acceptTcp.Read(text[:])
		if error != nil {
			fmt.Println(error, "接受客户端数据错误")
			return
		}
		fmt.Println(acceptTcp.RemoteAddr().String(), "->", string(text[:n]))
	}
}
