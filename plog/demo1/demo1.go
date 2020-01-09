package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Connection struct {
	url    string
	logger *log.Logger
}

func (v *Connection) Process(ctx context.Context) {
	go checkRequest(ctx, v.url)
	// 换算成毫秒单位
	duration := time.Duration(rand.Int()%1500) * time.Millisecond
	time.Sleep(duration)
	v.logger.Println(fmt.Sprintf("Process connection ok, url=%v, duration=%v", v.url, duration))
}

func checkRequest(ctx context.Context, url string) {
	duration := time.Duration(rand.Int()%1500) * time.Millisecond
	time.Sleep(duration)
	logger.Println(fmt.Sprintf("Check request ok, url=%v, duration=%v", url, duration))
}

var logger *log.Logger

func main() {
	ctx := context.Background()
	// 使用给定的seed来初始化生成器到一个确定的状态。否则每次运行程序生成的随机数和上次都是一样的
	// UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）
	rand.Seed(time.Now().UnixNano())
	logger = log.New(os.Stdout, "", log.LstdFlags)
	for i := 0; i < 5; i++ {
		go func(url string) {
			// fmt.Println(url)
			connection := &Connection{}
			connection.url = url
			connection.logger = logger
			connection.Process(ctx)
		}(fmt.Sprintf("url #%v", i))
	}
	time.Sleep(3 * time.Second)
}
