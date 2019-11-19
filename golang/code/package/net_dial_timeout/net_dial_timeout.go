package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	//url := "baidu.com:80"
	url := "facebook.com:80"
	conn, error := net.DialTimeout("tcp", url, time.Second*10) // 带有超时限制的conn
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", error.Error())
		return
	}

	n, error := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", error.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "writed: %d", n)

	buf := make([]byte, 2048)
	n, error = conn.Read(buf)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", error.Error())
		return
	}

	fmt.Fprintf(os.Stdout, string(buf[:n]))
}
