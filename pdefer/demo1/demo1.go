package main

import (
	"fmt"
)

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

func main() {
	var n number
	// n作为值，打印最初的值0
	defer n.print()
	defer func() {
		// 此时n为引用，最终打印3
		n.pprint()
	}()
	n = 3
}
