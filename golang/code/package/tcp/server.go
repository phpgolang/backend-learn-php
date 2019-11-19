package main

import (
	"fmt"
	"net"
)

func main() {
	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9900")
	if err != nil {
		println("地址解析失败", err)
		return
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		println("tcp服务启动失败", err)
	}
	println("tcp服务启动成功")
	defer listener.Close()
	for {
		//listener对象的Accept方法会直接阻塞，直到一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			println("建立连接失败", err)
		}
		go handleTcpClient(conn)
	}
}

func handleTcpClient(conn *net.TCPConn) {
	buf := make([]byte, 4096)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			println("接收数据失败", err)
		}
		println("客户端地址：", conn.RemoteAddr())
		fmt.Printf("接收客户端数据长度：%d, 数据：%s", num, buf[:num])

		if fmt.Sprintf("%s", buf[:num]) == "ByeBye" {
			println("客户端离线")
			return
		}

		send := []byte("receive success")
		_, err = conn.Write(send)
		if err != nil {
			println("发送数据失败", err)
		}
	}
}

// 作者：兔头咖啡
// 链接：http://www.imooc.com/article/265261
// 来源：慕课网
// 本文原创发布于慕课网 ，转载请注明出处，谢谢合作
