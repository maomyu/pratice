package ws

import (
	"context"
	"strconv"

	"github.com/yuwe1/pratice/gim/demo1/pb"
	"github.com/yuwe1/pratice/predis/logger"
	"google.golang.org/grpc/metadata"
)

const (
	CtxAppId     = "app_id"
	CtxUserId    = "user_id"
	CtxDeviceId  = "device_id"
	CtxToken     = "token"
	CtxRequestId = "request_id"
)

// GetCtxAppId 获取ctx的app_id
func GetCtxRequstId(ctx context.Context) int64 {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}

	requstIds, ok := md[CtxRequestId]
	if !ok && len(requstIds) == 0 {
		return 0
	}
	requstId, err := strconv.ParseInt(requstIds[0], 10, 64)
	if err != nil {
		return 0
	}
	return requstId
}
func DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) error {
	// 获取连接
	conn := load(req.DeviceeId)
	if conn == nil {
		logger.Sugar.Warn("ctx id nil")
		return nil
	}

	// 发送消息
	conn.Output(pb.PackageType_PT_MESSAGE, GetCtxRequstId(ctx), nil, req.Message)
	return nil
}
