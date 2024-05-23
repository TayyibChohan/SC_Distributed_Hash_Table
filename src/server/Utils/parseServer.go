package Utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes"
)

func ReadServers(filename string) ([]*nodes.ServerNode, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var servers []*nodes.ServerNode
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid server info: ", line)
			continue
		}
		address := net.ParseIP(parts[0])
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid port number: ", parts[1])
			continue
		}
		servers = append(servers, nodes.NewServerNode(address, port))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading server info file: ", err)
	}

	return servers, nil
}
