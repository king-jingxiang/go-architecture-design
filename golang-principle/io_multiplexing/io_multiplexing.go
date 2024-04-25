package io_multiplexing

import (
	"fmt"
	"net"
)

func CreateServer() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listen :8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server(conn)
	}

}

func server(conn net.Conn) {
	defer conn.Close()
	var buf [128]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		panic(err)
	}
	fmt.Println("Receive data: ", string(buf[:n]))

	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}
}
