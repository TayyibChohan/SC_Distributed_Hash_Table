package server

import (
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	"net"
	"strconv"
)

type Server struct {
	// Fields
	port int
	ip   string
	updSocket net.PacketConn
	possibleNodes []*nodes.ServerNode
}

// NewServer creates a new server with the given port and possible nodes
func NewServer(port int, possibleNodes []*nodes.ServerNode) *Server {
	//Create udp socket
	updSocket, err := net.ListenPacket("udp", "localhost:" + strconv.Itoa(port))
	if err != nil {
		return nil
	}

	return &Server{
		port:          port,
		ip:          constants.LOCALHOST,
		updSocket:    updSocket,
		possibleNodes: possibleNodes,
	}
}

// run starts the server
func (s *Server) Run() {
	defer s.updSocket.Close()

	buffer := make([]byte, constants.MAX_MESSAGE_SIZE)
	for {
		_, addr, err := s.updSocket.ReadFrom(buffer)
		if err != nil {
			continue
		}

		go s.handleRequest(buffer, addr)
	}
}

// handleRequest handles the request
func (s *Server) handleRequest(buffer []byte, addr net.Addr) {
	print("Received request from: " + addr.String())
	// Placeholder code
}

		
	

