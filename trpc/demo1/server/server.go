package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/yuwe1/pratice/trpc/public"
)

type ConnectRPCServer struct{}

var ConnectRPCServerIP = "127.0.0.1:60001"

func (s *ConnectRPCServer) Message(req public.MessageReq, resp *public.MessageAckResp) error {
	fmt.Println(req)
	fmt.Println(resp)
	return nil
}

// 启动rpc服务
func StartRPCServer() {
	rpc.Register(new(ConnectRPCServer))
	tcpAddr, err := net.ResolveTCPAddr("tcp", ConnectRPCServerIP)
	if err != nil {
		log.Println(err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
