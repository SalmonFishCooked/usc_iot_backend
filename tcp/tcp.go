package tcp

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [2048]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err: %v\n", err)
			break
		}
		rec := string(buf[:n])
		fmt.Printf("接收到的数据: %v\n", rec)
		conn.Write([]byte("ok"))
	}
}

func CreateTCPServer() {
	listen, err := net.Listen("tcp", "192.168.3.55:20000")
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
