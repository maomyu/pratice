package main

import (
	"fmt"
)
/*
注意 recover 只有在 defer 函数中才有用，
在 defer 的函数调用的函数中 recover 不起作用，
如下实例代码不会 recover：
*/
func main() {

	f := func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}

	defer func(){
		f()
	}()
	panic("ok")
}
