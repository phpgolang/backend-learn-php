package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xtaci/kcp-go"
)

var data = []byte("hello world!")

func main() {

	conn, err := kcp.DialWithOptions("127.0.0.1:9900", nil, 0, 0)
	if err != nil {
		fmt.Println("连接服务失败!", err)
		return
	}
	fmt.Println("服务端地址：", conn.RemoteAddr())

	go handleKcpConn(conn)

	select {}
}

func handleKcpConn(conn *kcp.UDPSession) {
	buf := make([]byte, 65535)

	for {
		send := append(data, strconv.FormatInt(time.Now().Unix(), 10)...)
		num, err := conn.Write(send)
		if err != nil {
			fmt.Println("发送数据失败!", err)
			return
		}
		num, err = conn.Read(buf)
		if err != nil {
			fmt.Println("接收数据失败!", err)
			return
		}
		fmt.Printf("接收服务端数据长度：%d, 数据：%s\n", num, buf[:num])

		time.Sleep(time.Second)
	}
}

// 作者：兔头咖啡
// 链接：http://www.imooc.com/article/265261
// 来源：慕课网
// 本文原创发布于慕课网 ，转载请注明出处，谢谢合作
