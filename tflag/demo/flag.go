/*
 * @Author: your name
 * @Date: 2019-11-22 15:17:53
 * @LastEditTime: 2019-11-22 18:59:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\tflag\demo\flag.go
 */
package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数对应的变量，这三个变量都是指针类型
var cliName = flag.String("name", "nick", "Input Your Name")
var cliAge = flag.Int("age", 28, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")

// 定义一个值类型的命令行参数变量，在 Init() 函数中对其初始化
// 因此，命令行参数对应变量的定义和初始化是可以分开的
var cliFlag int
var selfFlag self

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
	flag.Var(&selfFlag, "self", "please input address,area eg: \"beijing china\"")
}
func main() {
	// 初始化变量 cliFlag
	Init()
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	// 输出命令行参数
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("flagname=", cliFlag)
	fmt.Println("selfFlag=", selfFlag)
}
