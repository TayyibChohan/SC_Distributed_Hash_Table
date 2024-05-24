package main

import (
	"fmt"
	"net"
	"os"

	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/shared/ProtocolBuffers"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("This program must be executed with one argument:\n" +
			"go run <go file> <server address>")
		return
	}

	// Server address
	serverAddr := args[0]

	// Create UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	//Create a message using protocol buffers
	buffer := make([]byte, constants.MAX_MESSAGE_SIZE)
	myKVrequest := &ProtocolBuffers.KVrequest{}

	KVrequest, err := ProtocolBuffers.Marshal

	// Send data to server
	message := []byte("Hello, server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending data to server:", err)
		os.Exit(1)
	}

	// Receive response from server
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error receiving data from server:", err)
		os.Exit(1)
	}

	// Print server response
	fmt.Println("Server response:", string(buffer[:n]))
}
