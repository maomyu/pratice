package main

import (
	"fmt"
	"reflect"
)

type A struct {
	a int
	b string
}

func main() {
	aa := &A{a: 2, b: "test1"}
	bb := &A{a: 2, b: "test1"}
	cc := aa
	fmt.Println(reflect.DeepEqual(bb, cc))

	fmt.Println(aa == bb)
	fmt.Println(cc == bb)
}
