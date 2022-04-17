package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// ソケットの作成 & 接続待ちのためのIPアドレスとポート番号を設定
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":10000")
	if err != nil {
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		// ③接続の受信
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// ④ソケットの読み込み
		req := make([]byte, 1024)
		len, err := conn.Read(req)
		fmt.Println("request:", string(req[:len]))
		// ⑤ソケットの書き込み
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		// ⑥接続の切断
		conn.Close()
	}
}