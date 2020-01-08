package main

import (
	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/redisclient"
)

func main() {
	config.Init()
	redisclient.Init()
	redisclient.GetRedis()
}
