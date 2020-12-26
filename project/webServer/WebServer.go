package main

import (
	"fmt"
	"net/http"
)


func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}
func main() {
	// 注册
	http.HandleFunc("/index/", Index)
	// 开启server
	error := http.ListenAndServe("127.0.0.1:8000", nil)
	if error != nil {
		fmt.Println("server开启失败", error)
		return
	}
}
