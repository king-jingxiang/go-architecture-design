package io_multiplexing

import (
	"fmt"
	"net"
	"testing"
)

func TestCreateServer(t *testing.T) {
	CreateServer()
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer conn.Close()                // 关闭TCP连接
	_, err = conn.Write([]byte("Hi")) // 发送数据
	if err != nil {
		return
	}
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	}
	fmt.Println(string(buf[:n]))

}
