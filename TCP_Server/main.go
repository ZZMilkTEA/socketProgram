package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//TCP server端

func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	fmt.Println("准备接收数据……")
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("连接客户端失败,错误信息：", err)
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端信息:  ", recvStr)
		if strings.ToUpper(recvStr) == "Q" {
			fmt.Println("客户退出，与客户断开连接")
			return
		}
		conn.Write([]byte("服务端刚刚收到您的信息为:  " + recvStr)) //发送数据
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败，错误：", err)
		return
	}
	fmt.Println("监听中……")
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("建立连接失败，错误：", err)
			continue
		}
		fmt.Println("已经与客户端连接")
		go process(conn) //启动一个goroutine处理连接,使服务器具备并发处理请求的能力
	}
}
