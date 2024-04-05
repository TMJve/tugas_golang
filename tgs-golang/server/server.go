package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleIncomingConnection(clientConn)

	}
}

func handleIncomingConnection(client net.Conn) {
	defer client.Close()

	// Set a read deadline for the connection
	err := client.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		panic(err)
	}

	var size uint32
	err = binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	bytMsg := make([]byte, size)
	_, err = client.Read(bytMsg)
	if err != nil {
		panic(err)
	}

	strMsg := string(bytMsg)
	fmt.Printf("Received message: %s\n", strMsg)

	var reply string
	if strings.HasSuffix(strMsg, ".zip") {
		reply = "zip file has been received"
	} else {
		reply = "message has been received"
	}

	// Set a write deadline for the connection
	err = client.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		panic(err)
	}

	err = binary.Write(client, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}

	_, err = client.Write([]byte(reply))
	if err != nil {
		panic(err)
	}
}

// var reply string

// if strings.HasSuffix(strMsg, ".zip") {
// 	reply = "File has been received"
// } else if strings.Contains(strMsg, ".") {
// 	reply = "only for uploading zip"
// } else {
// 	reply = "Message have been received"
// }

// err = binary.Write(client, binary.LittleEndian, uint32(len(reply)))
// if err != nil {
// 	panic(err)
// }

// _, err = client.Write([]byte(reply))
// if err != nil {
// 	panic(err)
// }
