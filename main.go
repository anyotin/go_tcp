package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, error := conn.Read(buf)

	log.Printf("connection : %p \n", conn)
	fmt.Printf("Client> %s \n", buf[:n])

	if error != nil {
		if error == io.EOF {
			return
		} else {
			panic(error)
		}
	}

	n, error = conn.Write(buf[:n])
	if error != nil {
		panic(error)
	}

	handleConnection(conn)
}

func handleListener(l *net.TCPListener) error {
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
			return err
		}

		// 1対1のユニキャスト通信であればgoルーチンは必要にないが、複数のクライアントからの接続を可能にするためにはgoルーチンが必要になる。
		// goルーチンを使用しないと各接続のインスタンス？的な繋がりがなくなってしまう。
		// 仮にAという繋がりができて通信をしあっていて、新たにBという繋がりを形成しようとしてもhandleConnection関数で無限に処理が実行されているため次の繋がりの処理ができない。
		// そこでgoルーチンを使用することにより並行処理を可能にする。
		go handleConnection(conn)
	}
}

func main() {
	// ソケットの生成・接続
	// net.ResolveTCPAddr（network, address string) (*net.TCPAddr, error)は
	// addressに与えたIPアドレスの形式がnetworkの形式にそった形かを判定・解決し、大丈夫そうなら*net.TCPAddrを返します。
	log.Println("ソケットの生成・接続")
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("ResolveTCPAddr", err)
		return
	}

	// 接続準備
	log.Println("接続準備")
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println("ListenTCP", err)
		return
	}

	// 待ち受け開始
	log.Println("接続待機中....")
	err = handleListener(l)
	if err != nil {
		log.Println("handleListener", err)
	}

	fmt.Println("終了")
}
