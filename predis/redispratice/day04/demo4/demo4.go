package demo4

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"time"
)

func Acquire_lock(conn redis.Conn, lockname string, acquire_timeout int64) string {
	identifier := uuid.NewV4().String()

	i := 0
	end := time.Now().Unix() + acquire_timeout
	for time.Now().Unix() < end {
		// 尝试获得锁
		reply, err := conn.Do("SETNX", "lock:"+lockname, identifier)
		fmt.Println(reply)
		if result, _ := redis.Bool(reply, err); result {
			return identifier
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("正在获取锁,第%d秒\n", i)
		i++
	}
	return ""
}

func ExecSomthing(conn redis.Conn, item string) {
	locked := Acquire_lock(conn, item, 10)
	defer func(locked string) {
		if len(locked) != 0 {
			Rlease_lock(conn, item, locked)
		}
	}(locked)
	if len(locked) == 0 {
		fmt.Println("加锁失败")
		return
	}
	fmt.Println("正在对", item, "进行加锁")
	i := 1
	end := time.Now().Unix() + 10
	for time.Now().Unix() < end {
		fmt.Printf("正在执行该商品,第%d秒\n", i)
		time.Sleep(1 * time.Second)
		i++
	}
}

// / 释放锁
func Rlease_lock(conn redis.Conn, lock_name, identifier string) {
	lock_name = "lock:" + lock_name
	if result, _ := redis.String(conn.Do("GET", lock_name)); result == identifier {
		conn.Do("DEL", lock_name)
	}
	fmt.Println("释放锁", lock_name, "成功")
}
