package server

import (
	"fmt"
	"net"
	"strconv"

	ProtocolBuffers "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers"
	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes"
	"google.golang.org/protobuf/proto"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Utils"
)

type Server struct {
	// Fields
	port          int
	ip            string
	udpSocket     net.PacketConn
	possibleNodes []*nodes.ServerNode
}

// NewServer creates a new server with the given port and possible nodes
func NewServer(port int, possibleNodes []*nodes.ServerNode) *Server {
	//Create udp socket
	udpSocket, err := net.ListenPacket("udp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		return nil
	}

	return &Server{
		port:          port,
		ip:            constants.LOCALHOST,
		udpSocket:     udpSocket,
		possibleNodes: possibleNodes,
	}
}

// run starts the server
func (s *Server) Run() {
	defer s.udpSocket.Close()

	for {
		buffer := make([]byte, constants.MAX_MESSAGE_SIZE)
		_, addr, err := s.udpSocket.ReadFrom(buffer)
		if err != nil {
			continue
		}

		bufferCopy := make([]byte, len(buffer))
		go s.handleRequest(bufferCopy, addr)
	}
}

// handleRequest handles the request
func (s *Server) handleRequest(buffer []byte, addr net.Addr) {
	print("Received request from: " + addr.String())
	// Placeholder code
	myKVrequest := &ProtocolBuffers.KVRequest{}
	err := proto.Unmarshal(buffer, myKVrequest)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return
	}

	myResponse := &ProtocolBuffers.KVResponse{} //TODO: Implement this
	payload, err := proto.Marshal(myResponse)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	checkSum := Utils.Checksum(myResponse)
	id := Utils.CreateID()

	message:= &ProtocolBuffers.Msg{
		MessageID: id,
		Payload: payload,
		CheckSum: checkSum,
	}

	// Send response
	go func() {
		_, err := s.udpSocket.WriteTo(payload, addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}()

}
