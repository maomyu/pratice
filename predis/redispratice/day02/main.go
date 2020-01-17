package main

import (
	"fmt"
	"path/filepath"

	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/rediscli/redispool"
	"github.com/yuwe1/pratice/predis/redispratice/day02/ip"
)

func main() {
	config.Init()
	// redisclient.Ini()
	redispool.Init()
	f, _, p, c := redispool.NewSession()
	defer f.Relase(p, c)

	fmt.Println(ip.IpToScore("192.179.11.1"))
	// 获取文件路径
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(".", "/")))
	pr := filepath.Join(appPath, "file/GeoLiteCity-Blocks.csv")
	ip.ImportIpToRedis(f.GetConn(), pr)
}
