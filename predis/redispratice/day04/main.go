package main

import (
	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/rediscli/redispool"
	"github.com/yuwe1/pratice/predis/redispratice/day04/demo4"
)

func main() {
	config.Init()
	redispool.Init()
	f, _, p, c := redispool.NewSession()
	defer f.Relase(p, c)
	f.GetConn().Do("SELECT", 0)
	f.GetConn().Do("SET", "goods:item01", "item01")
	demo4.ExecSomthing(f.GetConn(), "item01")
}

// 112.124.31.82
