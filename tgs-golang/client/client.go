package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Send a message to server.")
		fmt.Println("2. Exit.")
		fmt.Print("Choice >> ")
		scanner.Scan()

		ch := scanner.Text()

		if ch == "1" {
			sendMessageMenu()
		} else if ch == "2" {
			fmt.Println("Thankyou for using.")
			break
		}
	}
}

func sendMessageMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	var msg string
	for {
		fmt.Print("Type message to send : ")
		scanner.Scan()
		msg = scanner.Text()

		if len(msg) < 10 {
			fmt.Println("message must be more than 10 characters")
		} else if strings.Contains(msg, "kasar") {
			fmt.Println("Message cannot contain kata kasar")
		} else {
			break
		}
	}

	sendMessageToServer(msg)
}

func sendMessageToServer(message string) {
	serverConn, err := net.DialTimeout("tcp", "localhost:3000", 4*time.Second)
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()

	// Set a write deadline for the connection
	err = serverConn.SetWriteDeadline(time.Now().Add(4 * time.Second))
	if err != nil {
		panic(err)
	}

	err = binary.Write(serverConn, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}

	_, err = serverConn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	var size uint32
	binary.Read(serverConn, binary.LittleEndian, &size)
	replyBytMsg := make([]byte, size)

	_, err = serverConn.Read(replyBytMsg)
	if err != nil {
		panic(err)
	}
	replyMsg := string(replyBytMsg)

	fmt.Printf("Received: %s\n", replyMsg)
}

func main() {
	menu()
}
