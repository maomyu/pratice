package main

import (
	"fmt"

	"github.com/yuwe1/pratice/predis/config"
	"github.com/yuwe1/pratice/predis/rediscli/redispool"
	"github.com/yuwe1/pratice/predis/redispratice/day03/demo3"
)

func main() {
	config.Init()
	// redisclient.Ini()
	redispool.Init()
	f, _, p, c := redispool.NewSession()
	defer f.Relase(p, c)

	result := demo3.Get_Contact(f.GetConn(), "陈威")
	fmt.Println(result)
	result = demo3.Fetch_autocomplete_list(f.GetConn(), "陈威", "徐")
	fmt.Println(result)
}
