package main

import (
	"fmt"
	"net"
)

func main() {

	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("地址解析失败!", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, laddr)
	if err != nil {
		fmt.Println("连接服务失败!", err)
		return
	}

	send := []byte("hello world!")

	_, err = conn.Write(send)
	if err != nil {
		fmt.Println("发送数据失败!", err)
		return
	}
	buf := make([]byte, 65535)
	num, remoteAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("接收数据失败!", err)
		return
	}

	fmt.Println("服务端地址：", remoteAddr)
	fmt.Printf("接收服务端数据长度：%d, 数据：%s", num, buf)
}

// 作者：兔头咖啡
// 链接：http://www.imooc.com/article/265261
// 来源：慕课网
// 本文原创发布于慕课网 ，转载请注明出处，谢谢合作
