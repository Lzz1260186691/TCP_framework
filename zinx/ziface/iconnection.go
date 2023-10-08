package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConnID() uint32
}

type HandFunc func(*net.TCPConn, []byte, int) error
