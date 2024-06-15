package main

import (
	"fmt"
	"os"
	"strconv"

	server "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Server"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Utils"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("This program must be executed with two arguments:\n" +
			"go run <go file> <port> <path to server.txt>")
		return
	}

	port, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error parsing port number:", err)
		return
	}

	possibleNodes, err := Utils.ReadServers(args[1])
	if err != nil {
		fmt.Println("Error reading servers:", err)
		return
	}

	server := server.NewServer(port, possibleNodes)
	server.Run()
}
