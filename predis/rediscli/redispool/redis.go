package redispool

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"github.com/garyburd/redigo/redis"
	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/logger"
)

var (
	inited bool
	p      *Pool
	m      sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()
	var err error
	if inited {
		err = fmt.Errorf("[init]redis已经被初始化")
		logger.Sugar.Fatalf(err.Error())
		return
	}
	fmt.Println(config.GetRedisConfig())
	if config.GetRedisConfig().GetEnabled() {
		initRedisPool()
	}
	inited = true
	fmt.Println("初始化成功")
}

// 初始化一个redis连接池
func initRedisPool() {
	fmt.Println("初始化redispool")
	p, _ = New(createConnec, config.GetRedisConfig().GetDBNum())
}

type redisConnection struct {
	conn *redis.Conn
	id   int32
}

var idCounter int32

// 实现相关接口
func (r *redisConnection) Close() error {
	logger.Sugar.Info("close redis pool conn", r.id)
	return nil
}

// 生成redis连接
func createConnec() (io.Closer, error) {
	fmt.Println("生成redis连接")
	id := atomic.AddInt32(&idCounter, 1)

	redisConn, _ := redis.Dial("tcp", config.GetRedisConfig().GetConn())
	redisConn.Send("auth", config.GetRedisConfig().GetPassword)

	return &redisConnection{
		id:   id,
		conn: &redisConn,
	}, nil
}
