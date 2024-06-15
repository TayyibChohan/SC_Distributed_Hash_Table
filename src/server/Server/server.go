package server

import (
	"fmt"
	"net"
	"strconv"

	ProtocolBuffers "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers"
	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Hasher"
	rr "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/RequestReply"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/KVStore"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Utils"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	// Fields
	port          int
	ip            string
	udpSocket     net.PacketConn
	possibleNodes []*nodes.ServerNode
	rr            *rr.RequestReply
}

// NewServer creates a new server with the given port and possible nodes
func NewServer(port int, possibleNodes []*nodes.ServerNode) *Server {
	//Create udp socket
	hasher := Hasher.NewHasher("SHA1") // TODO: Actually make the hasher
	primaryStore := KVStore.NewKVStore()
	secondaryStore := KVStore.NewKVStore()
	rr := rr.NewRequestReply(hasher, primaryStore, secondaryStore)
	udpSocket, err := net.ListenPacket("udp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		return nil
	}

	return &Server{
		port:          port,
		ip:            constants.LOCALHOST,
		udpSocket:     udpSocket,
		possibleNodes: possibleNodes,
		rr:            rr,
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
	message := &ProtocolBuffers.Msg{}
	err := proto.Unmarshal(buffer, message)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return
	}
	myKVrequest := &ProtocolBuffers.KVRequest{}
	err = proto.Unmarshal(message.Payload, myKVrequest)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return
	}

	myResponse, err := s.rr.HandleRequest(myKVrequest)
	if err != nil {
		fmt.Println("Error handling request:", err)
		return
	}
	payload, err := proto.Marshal(myResponse)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	reply := Utils.SerializeMessage(message.MessageID, payload)
	packet, err := proto.Marshal(reply)
	if err != nil {
		fmt.Println("Error marshalling reply:", err)
		return
	}

	// Send response
	go func() {
		_, err := s.udpSocket.WriteTo(packet, addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}()

}
