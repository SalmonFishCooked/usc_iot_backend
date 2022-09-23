package tcp

import (
	"bufio"
	"fmt"
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
		rec := string(buf[:n])
		//发送HTTP请求
		sendHttpRequest.SendHttpRequest(rec)

		//fmt.Printf("接收到的数据: %v\n", rec)
		//conn.Write([]byte("ok"))
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
