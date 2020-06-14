package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//客户端

func main()  {
	conn ,err := net.Dial("tcp","127.0.0.1:8888")
	fmt.Println("准备与服务端连接")
	if err != nil {
		fmt.Println("连接失败，错误: ",err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功，请随便输入些什么,输入q退出")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')    //读取用户输入
		inputInfo := strings.Trim(input,"\r\n")
		if strings.ToUpper(inputInfo) == "Q"{
			return  //如果输入q就退出
		}
		_,err = conn.Write([]byte(inputInfo))   //发送数据
		if err != nil{
			return
		}
		buf := [512]byte{}
		n,err := conn.Read(buf[:])
		if err != nil{
			fmt.Println("接受失败，错误: ",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}