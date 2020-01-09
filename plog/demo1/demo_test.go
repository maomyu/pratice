package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	now := time.Now().UnixNano()
	rand.Seed(now)
	fmt.Println("当前时间：", now)
	duration := time.Duration(rand.Int()%1500) * time.Millisecond
	time.Sleep(duration)
	fmt.Println(duration)
}
