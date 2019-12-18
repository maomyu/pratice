package main

import (
	"github.com/yuwe1/pratice/trpc/demo1/client"
	"github.com/yuwe1/pratice/trpc/demo1/server"
)

func main() {
	go server.StartRPCServer()
	client.InitRpcClient()
	client.ConnectRpcClient.Message(12)
}
