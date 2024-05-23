package server

import (
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
)

type Server struct {
	// Fields
	port int
	ip   string

	possibleNodes []*nodes.ServerNode
}

// NewServer creates a new server with the given port and possible nodes
func NewServer(port int, possibleNodes []*nodes.ServerNode) *Server {
	return &Server{
		port:          port,
		ip:            "localhost",
		possibleNodes: possibleNodes,
	}
}

// run starts the server
func (s *Server) Run() {
	// Placeholder code
}
