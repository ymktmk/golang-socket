package main

import (
	"fmt"
	"net"
)

func main() {
	// ソケットの作成 & 接続待ちのためのIPアドレスとポート番号を設定
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":10000")
	if err != nil {
		fmt.Println(err)
	}
	// 接続を待つ
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		// 接続を受け付ける
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// ソケットの読み込み
		req := make([]byte, 1024)
		len, err := conn.Read(req)
		fmt.Println("request:", string(req[:len]))
		// ソケットの書き込み
		conn.Write([]byte(req[:len]))
		// ソケットを閉じる
		conn.Close()
	}
}