package main

import (
	"fmt"

	"github.com/xtaci/kcp-go"
)

func main() {
	//不加密，不设置
	listener, err := kcp.ListenWithOptions(":9900", nil, 0, 0)

	if err != nil {
		fmt.Println("udp服务启动失败", err)
		return
	}
	println("kcp服务启动成功")
	defer listener.Close()
	for {
		conn, err := listener.AcceptKCP()
		if err != nil {
			fmt.Println("建立连接失败!", err)
			continue
		}
		fmt.Println("客户端地址：", conn.RemoteAddr())
		go handleKcpClient(conn)
	}
}
func handleKcpClient(conn *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接收数据失败!", err)
			return
		}
		fmt.Printf("接收客户端数据长度：%d, 数据：%s\n", num, buf[:num])

		if fmt.Sprintf("%s", buf[:num]) == "ByeBye" {
			// 接收到 ByeBye 就断线
			fmt.Println("客户端离线!")
			return
		}
		send := []byte("receive success!")
		_, err = conn.Write(send)
		if err != nil {
			fmt.Println("发送数据失败!", err)
			return
		}
	}
}

// 作者：兔头咖啡
// 链接：http://www.imooc.com/article/265261
// 来源：慕课网
// 本文原创发布于慕课网 ，转载请注明出处，谢谢合作
