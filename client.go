package main

import (
	"fmt"
	"net"
	"os"
)

// go run client.go 127.0.0.1:10000

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	// ソケットの作成 & 接続相手のIPアドレスとポート番号を設定
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
	}
	// 接続の要求を行う
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// ソケットにデータの書き込み
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	res := make([]byte, 1024)
	// ソケットからデータの読み込み
	len, err := conn.Read(res)
	fmt.Println("response:", string(res[:len]))
	// ソケットを閉じる
	conn.Close()
}