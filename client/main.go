package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// ①実行の際に指定したhost:portでbind
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	fmt.Println(service)
	// ②ソケットの作成とIP:portに紐付け
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
	}
	// ③サーバ側に接続
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// ④ソケットにデータの書き込み
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	res := make([]byte, 1024)
	// ⑤ソケットからデータの読み込み
	len, err := conn.Read(res)
	fmt.Println("response:", string(res[:len]))
	// ⑥接続の切断
	conn.Close()
}