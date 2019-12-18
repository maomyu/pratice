/*
 * @Author: your name
 * @Date: 2019-11-22 14:02:48
 * @LastEditTime: 2019-11-22 14:08:22
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\tcp\demo1\server\server.go
 */
package main

import "fmt"

import "net"

func main() {
	fmt.Println("Starting the server ...")
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}
	// 监听客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
