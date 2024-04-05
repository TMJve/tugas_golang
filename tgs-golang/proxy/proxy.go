package main

import (
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleProxyConnection(clientConn)
	}
}

func handleProxyConnection(client net.Conn) {
	defer client.Close()

	serverConn, err := net.Dial("tcp", "localhost:3000")

	if err != nil {
		panic(err)
	}
	defer serverConn.Close()

	go func() {
		io.Copy(serverConn, client)
	}()
	io.Copy(client, serverConn)
}
