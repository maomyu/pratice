/*
 * @Author: your name
 * @Date: 2019-11-23 10:25:34
 * @LastEditTime: 2019-11-23 10:26:50
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\trpc\app\connect\main.go
 */
package main

import "github.com/yuwe1/pratice/trpc/demo1/server"

func main() {
	go server.StartRPCServer()
	select {}
}
