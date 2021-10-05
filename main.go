package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"tcp/server/Entities"
)

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	fmt.Println(conn)

	num := 0

	for {
		num++
		fmt.Println(num)

		user_accept := new(Entities.User)

		n, err := conn.Read(buf)

		data := buf[:n]

		string_data := string(data)

		fmt.Println(string_data)

		err = json.Unmarshal([]byte(string_data), user_accept)

		fmt.Println(user_accept)
		if err != nil {
			if ne, ok := err.(net.Error); ok {
				switch {
				case ne.Temporary():
					continue
				}
			}
			log.Println("Read", err)
			return
		}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println("Write", err)
			return
		}
	}
}

func handleListener(l *net.TCPListener) error {
	for {
		conn, err := l.AcceptTCP()

		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
			return err
		}

		go handleConnection(conn)
	}
}

func main()  {
	// net.ResolveTCPAddr（network, address string) (*net.TCPAddr, error)は
	// addressに与えたIPアドレスの形式がnetworkの形式にそった形かを判定・解決し、大丈夫そうなら*net.TCPAddrを返します。
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("ResolveTCPAddr", err)
		return
	}

	// 待ち受け開始
	fmt.Println("待受開始...")
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println("ListenTCP", err)
		return
	}

	err = handleListener(l)
	if err != nil {
		log.Println("handleListener", err)
	}

	fmt.Println("終了")
}
