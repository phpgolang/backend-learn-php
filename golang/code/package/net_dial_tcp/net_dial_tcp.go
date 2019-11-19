package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	url := "www.baidu.com:80"
	// 将一个host地址转换为TCPAddr。host=ip:port
	pRemoteTCPAddr, err := net.ResolveTCPAddr("tcp4", url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	// pLocalTCPAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7070")

	pTCPConn, err := net.DialTCP("tcp", nil /*pLocalTCPAddr*/, pRemoteTCPAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	n, err := pTCPConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}
	defer pTCPConn.Close()

	fmt.Fprintf(os.Stdout, "writed: %d\n", n)

	buf, err := ioutil.ReadAll(pTCPConn)
	r := len(buf)
	fmt.Fprintf(os.Stdout, string(buf[:r]))
	fmt.Fprintf(os.Stdout, "readed: %d\n", r)

}
