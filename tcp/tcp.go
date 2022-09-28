package tcp

import (
	"bufio"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net"
	"usc_iot_backend/utils/sendHttpRequest"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [2048]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("设备已经断开连接, err: %v\n", err)
			break
		}
		//rec := string(buf[:n])
		json, _ := simplejson.NewJson([]byte(buf[:n]))

		//利用t的值判断json传过来的数据
		tValue, _ := json.Get("t").Int()
		switch tValue {
		case 1:
			fmt.Println("收到t为1的设备信息值")
		case 3:
			//发送HTTP请求
			sendHttpRequest.SendHttpRequest(json, 3)
		}

		//fmt.Printf("接收到的数据: %v\n", rec)
		//conn.Write([]byte("teeeeeeeeee"))
	}
}

func CreateTCPServer() {
	listen, err := net.Listen("tcp", "localhost:20000")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	for {
		fmt.Println("[TCP Server]: Waiting a new connection...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			continue
		}
		fmt.Println("[TCP Server]: TCP Server: A new connection found.")
		go process(conn)
	}
}
