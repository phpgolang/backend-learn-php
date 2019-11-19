package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	pTCPAddr, error := net.ResolveTCPAddr("tcp4", ":7070")
	if error != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", error.Error())
		return
	}

	pTCPListener, error := net.ListenTCP("tcp4", pTCPAddr)
	if error != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", error.Error())
		return
	}
	defer pTCPListener.Close()

	for {
		pTCPConn, error := pTCPListener.AcceptTCP()
		if error != nil {
			fmt.Fprintf(os.Stdout, "Error: %s", error.Error())
			continue
		}
		go connHandler(pTCPConn)
	}
}

func connHandler(conn *net.TCPConn) {
	defer conn.Close()
	now := time.Now()
	conn.Write([]byte(now.String() + "\n"))
}
