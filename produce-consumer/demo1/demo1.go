package main

import (
	"fmt"
	"time"
)

var (
	ch   = make(chan int, 5)
	data int
	max  int
	min  int
)

const (
	step = 10
)

func reload() {
	min = data
	max = min + step
	data = max
}

// 生产者
func producer() {
	for {
		if min >= max {
			reload()
		}
		min++
		fmt.Println("正在生产：", min)
		ch <- min
		time.Sleep(1 * time.Second)
	}
}

// 消费者
func consumer() {
	select {
	case uid := <-ch:
		fmt.Println("正在消费：", uid)
	}
}

func main() {
	go producer()
	for {
		consumer()
		time.Sleep(1 * time.Second)
	}
}
