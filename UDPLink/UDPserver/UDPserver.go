package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = conn.Close()
	}()

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, 10)
		//_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		n, addr, _ := conn.ReadFrom(data)
		fmt.Println(n, addr.String())

		strData := string(data)
		fmt.Println("Received:", strData)

		upper := strings.ToUpper(strData)
		//n, err := conn.WriteToUDP([]byte(upper), rAddr)
		n, err = conn.WriteTo([]byte(upper), addr) // addr
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(n)

		fmt.Println("Send:", upper)
	}
}
