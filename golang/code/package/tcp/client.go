package main

import (
	"fmt"
	"net"
)

func main() {
	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9900")
	if err != nil {
		fmt.Println("地址解析失败！", err)
	}
	conn, err := net.DialTCP("tcp", nil, laddr)
	if err != nil {
		fmt.Println("连接服务失败！", err)
		return
	}
	send := []byte("hello world!")
	buf := make([]byte, 65535)
	for i := 0; i < 10; i++ {
		_, err = conn.Write(send)
		if err != nil {
			println("发送数据失败", err)
			return
		}
		num, err := conn.Read(buf)
		if err != nil {
			println("数据接收失败", err)
			return
		}

		println("服务端地址：", conn.RemoteAddr())
		fmt.Printf("接收服务端数据长度：%d, 数据：%s", num, buf[:num])
	}

	bye := []byte("ByeBye")
	_, err = conn.Write(bye)
	if err != nil {
		println("发送数据失败", err)
		return
	}

}

// 作者：兔头咖啡
// 链接：http://www.imooc.com/article/265261
// 来源：慕课网
// 本文原创发布于慕课网 ，转载请注明出处，谢谢合作
