package main

import (
	"fmt"
	"net"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", ":9000")
	if err != nil {
		fmt.Println("地址解析失败!", err)
		return
	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("udp服务启动失败!", err)
		return
	}
	fmt.Println("udp服务启动成功")
	defer conn.Close()

	for {
		buf := make([]byte, 65535)
		num, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("接收数据失败!", err)
			continue
		}

		fmt.Println("客户端地址：", remoteAddr)
		fmt.Printf("接收客户端数据长度：%d, 数据：%s", num, buf)

		send := []byte("receive success!")

		_, err = conn.WriteToUDP(send, remoteAddr)
		if err != nil {
			fmt.Println("发送数据失败!", err)
			continue
		}
	}
}

// 作者：兔头咖啡
// 链接：http://www.imooc.com/article/265261
// 来源：慕课网
// 本文原创发布于慕课网 ，转载请注明出处，谢谢合作
