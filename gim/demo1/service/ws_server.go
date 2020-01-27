package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/yuwe1/pratice/gim/demo1/gerrors"
	pb "github.com/yuwe1/pratice/gim/demo1/pb"
	"github.com/yuwe1/pratice/gim/demo1/ws"
	"github.com/yuwe1/pratice/predis/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	appId, _ := strconv.ParseInt(r.Header.Get(ws.CtxAppId), 10, 64)
	userId, _ := strconv.ParseInt(r.Header.Get(ws.CtxUserId), 10, 64)
	deviceId, _ := strconv.ParseInt(r.Header.Get(ws.CtxDeviceId), 10, 64)
	token := r.Header.Get(ws.CtxToken)
	requestId, _ := strconv.ParseInt(r.Header.Get(ws.CtxRequestId), 10, 64)

	if appId == 0 || userId == 0 || deviceId == 0 || token == "" {
		s, _ := status.FromError(gerrors.ErrUnauthorized)
		bytes, err := json.Marshal(s.Proto())
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		w.Write(bytes)
	}

	_, err := pb.LogicIntClient.SignIn(ws.ContextWithRequstId(context.TODO(), requestId), &pb.SignInReq{
		AppId:    appId,
		UserId:   userId,
		DeviceId: deviceId,
		Token:    token,
		ConnAddr: "config.WSConf.LocalAddr",
	})

	s, _ := status.FromError(err)
	if s.Code() != codes.OK {
		bytes, err := json.Marshal(s.Proto())
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		w.Write(bytes)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	// 断开这个设备之前的连接
	preCtx := ws.Load(deviceId)
	if preCtx != nil {
		preCtx.DeviceId = -1
	}

	ctx := ws.NewWSConnContext(conn, appId, userId, deviceId)
	ws.Store(deviceId, ctx)
	ctx.DoConn()
}

func StartWSServer(address string) {
	http.HandleFunc("/ws", wsHandler)
	logger.Logger.Info("websocket server stat")
	err := http.LisenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
