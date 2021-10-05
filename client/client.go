package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func makeJson() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "OK"}`)
		return
	})
}

func main() {
	// tcp://127.0.0.1:8888に接続する
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("tcp://127.0.0.1:8888に接続できませんでした")
	}

	// connからレスポンスを標準出力にだす
	defer conn.Close()
	sendMessage(conn)
	//user := Entities.User{
	//	UserId: 13,
	//	UserName: "菊池",
	//	UserRank: 78,
	//}
	//
	//fmt.Println(user)
	//
	//json_data, err := json.Marshal(user)
	//
	//string_json_data := string(json_data)
	//
	//user_accept := new(Entities.User)
	//
	//err = json.Unmarshal([]byte(string_json_data), user_accept)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%+v\n", user_accept.UserId)
	//
	//fmt.Println(string_json_data)
	//
	//conn.Write(json_data)
	//
	//fmt.Println("メッセージを送信")
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	//for  {
	//	// メッセージを受信する
	//	res := make([]byte, 1024)
	//	_, err := conn.Read(res)
	//
	//	if err != nil {
	//		switch err {
	//			case io.EOF:
	//			fmt.Println("ServerとのSocketが切断されています。")
	//			fmt.Println("相手が切断したソケットに対してreadメソッドを呼び出したために発生。")
	//			fmt.Println(err)
	//			return
	//		}
	//		break
	//	}
	//
	//	fmt.Println("メッセージを受信")
	//	fmt.Printf("Server> %s \n", bytes.NewBuffer(res))
	//
	//	// キーボード入力
	//	stdin := bufio.NewScanner(os.Stdin)
	//
	//	if stdin.Scan() == false {
	//		fmt.Println("Ciao ciao!")
	//		return
	//	}
	//
	//	_, error := conn.Write([]byte(stdin.Text()))
	//
	//	if error != nil {
	//		switch error {
	//		case syscall.EPIPE:
	//			fmt.Println("ServerとのSocketが切断されています。")
	//			fmt.Println("相手が切断したソケットに対してwriteメソッドを呼び出すと発生")
	//			panic(err)
	//			return
	//		}
	//		break
	//	}
	//}
}

func sendMessage(connection net.Conn) {
	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("Ciao ciao!")
		return
	}

	_, error := connection.Write([]byte(stdin.Text()))

	if error != nil {
		panic(error)
	}

	var response = make([]byte, 4*1024)
	_, error = connection.Read(response)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Server> %s \n", response)

	sendMessage(connection)
}
