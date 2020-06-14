package main

import (
	"fmt"
	"net"
	"strconv"
)

//UDP Client
func main()  {
	socket,err := net.DialUDP("udp",nil,&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:8888,
	})
	if err != nil{
		fmt.Println("连接服务器失败，错误：",err)
		return
	}
	defer socket.Close()

	packageNum := 100
	for i := 1; i <= packageNum; i++ {
		sendData := []byte(
			"总共发包:" + strconv.Itoa(packageNum) +
			" 当前包编号:" +  strconv.Itoa(i))
		fmt.Println("数据准备发送")
		_,err = socket.Write(sendData)
		if err != nil{
			fmt.Println("发送数据失败，错误：",err)
			return
		}
		fmt.Println("数据发送完毕")
		data := make([]byte,4096)
		n,remoteAddr,err := socket.ReadFromUDP(data)
		if err != nil{
			fmt.Println("接受数据失败，错误：",err)
			return
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
	}
	sendData := []byte("EOF")
	fmt.Println("准备发送结束指示包")
	_,err = socket.Write(sendData)
	if err != nil{
		fmt.Println("发送数据失败，错误：",err)
		return
	}
	fmt.Println("结束指示包已发送")
	data := make([]byte,4096)
	n,remoteAddr,err := socket.ReadFromUDP(data)
	if err != nil{
		fmt.Println("接受数据失败，错误：",err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v \n", string(data[:n]), remoteAddr, n)
}