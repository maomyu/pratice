package ws

import (
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/yuwe1/pratice/gim/demo1/pb"
	"github.com/yuwe1/pratice/predis/logger"
	"google.golang.org/grpc/status"
)

var manager sync.Map

type WSConnContext struct {
	Conn     *websocket.Conn // websocket连接
	AppId    int64           // AppId
	DeviceId int64           // 设备id
	UserId   int64           // 用户id
}

func (c *WSConnContext) Output(pt pb.PackageType, requestId int64, err error, message proto.Message) {
	var output = pb.Output{
		Type:      pt,
		RequestId: requestId,
	}
	if err != nil {
		status, _ := status.FromError(err)
		output.Code = int32(status.Code())
		output.Message = status.Message()
	}
	if message != nil {
		msgBytes, err := proto.Marshal(message)
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		output.Data = msgBytes
	}
}

// 存储
func store(deviceId int64, ctx *WSConnContext) {
	manager.Store(deviceId, ctx)
}

// 获取
func load(deviceId int64) *WSConnContext {
	if value, ok := manager.Load(deviceId); ok {
		return value.(*WSConnContext)
	}
	return nil
}

// 删除
func delete(deviceId int64) {
	manager.Delete(deviceId)
}
