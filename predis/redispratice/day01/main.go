package main

import (
	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/redisclient"
	"github.com/yuwe1/pratice/predis/redispratice/day01/demo"
)

func main() {
	config.Init()
	redisclient.Init()
	conn := redisclient.GetRedis()

	// demo.Schedule_row_cache(*conn, "277", 5)
	// demo.Schedule_row_cache(*conn, "294", 5)
	// demo.Schedule_row_cache(*conn, "279", 5)

	// result := conn.ZRangeByScoreWithScores("schedule:", redis.ZRangeBy{
	// 	Min:    "-inf",
	// 	Max:    "+inf",
	// 	Offset: 0,
	// 	Count:  0,
	// })
	// fmt.Println(result.Result())
	demo.Cache_rows(*conn)
}
