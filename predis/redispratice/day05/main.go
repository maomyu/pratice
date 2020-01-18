package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/rediscli/redispool"
)

func Del(conn redis.Conn, item string) {
	ditem := "goods:" + item
	item = "lock:" + item
	i := 1
	for {
		if result, _ := redis.String(conn.Do("GET", item)); len(result) > 0 {
			fmt.Println("该商品目前正在锁定", i)
			i++
		} else {
			fmt.Println("该商品目前没有锁定，可以删除", i)
			if result, _ := redis.Int(conn.Do("DEL", ditem)); result == 1 {
				fmt.Println("成功删除")
				break
			} else {
				fmt.Println("删除失败")
			}
			break
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	config.Init()
	redispool.Init()
	f, _, p, c := redispool.NewSession()
	defer f.Relase(p, c)
	f.GetConn().Do("SELECT", 0)
	Del(f.GetConn(), "item01")
	// fmt.Println(redis.String(f.GetConn().Do("GET", "item01")))
}
