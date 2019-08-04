package models

import "net"

type Server struct {
	UDPServer *net.UDPConn
	Rooms [] room
}
