package main

import (
	"context"
	"net"

	pb "github.com/yuwe1/pratice/gim/demo1/pb"
	"github.com/yuwe1/pratice/gim/demo1/ws"
	"github.com/yuwe1/pratice/predis/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ConnIntServer struct{}

// UnaryServerInterceptor 服务器端的单向调用的拦截器
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	logger.Logger.Debug("interceptor", zap.Any("info", info), zap.Any("req", req), zap.Any("resp", resp))
	return resp, err
}

// 投递消息
func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	return &pb.DeliverMessageResp{}, ws.DeliverMessage(ctx, req)
}

// 启动一个rpc
func StartRpcServer() {
	listener, err := net.Listen("tcp", ":60000")
	if err != nil {
		panic(err)
	}
	// 拦截器
	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor))
	pb.RegisterConnIntServer(server, &ConnIntServer{})
	logger.Logger.Debug("rpc服务已经开启")
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
func main() {

}
