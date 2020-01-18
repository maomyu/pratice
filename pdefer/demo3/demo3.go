package main

import "fmt"

func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 对上个函数的拆解,因此会打印5
func f1() (r int) {
	t := 5
	r = t
	defer func() {
		t = t + 5
	}()
	return
}

func f3() (r int) {
	t := 5

	// 这里改变的r是之前传值传进去的r，如果不用作参数，返回10
	defer func(r int) {
		r = t + 5
	}(r)
	return t
}

func f4() (r int) {
	t := 5

	// 这里改变的r是之前传值传进去的r，如果不用作参数，返回10
	defer func() {
		r = t + 5
	}()
	return t
}
func main() {
	fmt.Println(f4())
}
