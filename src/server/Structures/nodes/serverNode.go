package nodes

import (
	"net"
	"strconv"
)

type ServerNode struct {
	Host net.IP
	Port int
}

func NewServerNode(host net.IP, port int) *ServerNode {
	return &ServerNode{Host: host, Port: port}
}

func (s *ServerNode) String() string {
	return s.Host.String() + ":" + strconv.Itoa(s.Port)
}

func (s *ServerNode) Equal(other *ServerNode) bool {
	return s.Host.Equal(other.Host) && s.Port == other.Port
}

func (s *ServerNode) CompareTo(other *ServerNode) int {
	if s.Host.Equal(other.Host) {
		return s.Port - other.Port
	} else {
		hostBytes := s.Host.To4()
		otherHostBytes := other.Host.To4()

		for i := 3; i >= 0; i-- {
			if hostBytes[i] != otherHostBytes[i] {
				return int(hostBytes[i]) - int(otherHostBytes[i])
			}
		}
	}

	return 0
}
