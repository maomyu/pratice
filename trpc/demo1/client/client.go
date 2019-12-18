package client

import (
	"log"
	"net/rpc"
	"time"

	"github.com/yuwe1/pratice/trpc/public"
)

var cli *rpc.Client
var ConnectRPCServerIP = "127.0.0.1:60001"

func InitRpcClient() {
	var err error
	for {
		cli, err = rpc.Dial("tcp", ConnectRPCServerIP)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

type connectRpcClient struct{}

var ConnectRpcClient = new(connectRpcClient)

// Message 消息投递
func (connectRpcClient) Message(deviceId int64) (*public.MessageResp, error) {

	req := public.MessageReq{
		DeviceId: deviceId,
		// Bytes:    bytes,
	}
	var resp = new(public.MessageResp)
	err := cli.Call("ConnectRPCServer.Message", req, &resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
