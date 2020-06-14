package main

import (
	"fmt"
	"net"
)

// UDP Sever
func main()  {
	fmt.Println("准备监听")
	listen,err := net.ListenUDP("udp",&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:8888,
	})
	successNum := 0
	if err != nil{
		fmt.Println("监听失败，错误：",err)
		return
	}
	fmt.Println("正在监听")
	defer listen.Close()
	for {
		var data [1024]byte
		n,addr,err := listen.ReadFromUDP(data[:])
		if err != nil{
			fmt.Println("接收udp数据失败，错误：",err)
			continue
		}
		if string(data[:n]) == "EOF"{
			packetLossRate := (100 - successNum)/100
			response := fmt.Sprintf("丢包率为%v%%", packetLossRate)
			fmt.Println(response)
			_ ,err = listen.WriteToUDP([]byte(response),addr)   //发送数据
			if err != nil{
				fmt.Println("发送数据失败，错误：",err)
			}
			successNum = 0
			continue
		}
		//stringSplits := strings.Fields(string(data[:n]))
		//successNumString := strings.Split(stringSplits[1],":")
		//successNum, _ := strconv.Atoi(successNumString[1])
		//successNum++
		//sum = successNum
		successNum++
		fmt.Println("成功数量:",successNum)
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		_ ,err = listen.WriteToUDP(data[:n],addr)   //发送数据
		if err != nil{
			fmt.Println("发送数据失败，错误：",err)
			continue
		}
	}
}