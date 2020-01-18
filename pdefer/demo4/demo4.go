package main

import "fmt"

func f1() {
	var n int
	defer fmt.Println(n)
	n = 3
	return
}

func f2() {
	var n int
	defer func() {
		fmt.Println(n)
	}()
	n = 3
	return
}

func f3() {
	var n int
	defer func(n int) {
		fmt.Println(n)
	}(n)
	n = 3

	return
}

func main() {
	f1()
	f2()
	f3()
}
