package znet

import (
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPversion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s,Port %d, is starting \n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println("listen", s.IPversion, "err", err)
			return
		}
		fmt.Println("start Zinx server", s.Name, "succ,now listenning...")

		for {
			conn, err := listenner.AcceptTCP()
		}

	}()
}
