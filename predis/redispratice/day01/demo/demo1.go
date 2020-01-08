package demo

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// 负责调度和终止缓存
func Schedule_row_cache(conn redis.Client, row_id string, delay float64) {
	conn.ZAdd("delay:", redis.Z{
		Member: row_id,
		Score:  delay,
	})

	conn.ZAdd("schedule:", redis.Z{
		Member: row_id,
		Score:  float64(time.Now().Unix()),
	})
}

// 进行缓存
func Cache_rows(conn redis.Client) {

	for {
		// 获取下一个需要被缓存的数据
		result, _ := conn.ZRangeByScoreWithScores("schedule:", redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  0,
		}).Result()
		// 现在的时间
		now := float64(time.Now().Unix())
		if result[0].Score > now {
			fmt.Println("不用更新缓存")
			time.Sleep(1)
			continue
		}

		row_id := result[0].Member
		fmt.Println("开始设置缓存：", row_id)
		// 获取延迟值
		delay, _ := conn.ZScore("delay:", row_id.(string)).Result()
		if delay <= 0 {
			conn.ZRem("delay:", row_id.(string))
			conn.ZRem("delay:", row_id.(string))
			conn.Del("inv:" + row_id.(string))
			continue
		}
		fmt.Println("设置缓存")
		row := "aaaaaaa" + time.Now().String()
		conn.ZAdd("schedule:", redis.Z{
			Score:  now + delay,
			Member: row_id.(string),
		})
		conn.SAdd("inv:"+row_id.(string), row)
	}
}
