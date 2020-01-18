package main

import (
	"fmt"
	"io"
	"os"
)

func mergeFile() {
	f, _ := os.Open("file.txt")
	// f作为值进行赋值引用
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file.txt err %v\n", err)
			} else {
				fmt.Println("file success close")
			}
		}(f)
	}

	f, _ = os.Open("file2.txt")
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file2.txt err %v\n", err)
			} else {
				fmt.Println("file2 success close")
			}
		}(f)
	}
}

// 错误用法,运行到file会报错
func errMeraggfile() {
	f, _ := os.Open("file.txt")
	defer func() {
		// fmt.Println(f)
		if err := f.Close(); err != nil {
			fmt.Printf("defer close file.txt err %v\n", err)
		} else {
			fmt.Println("file success close")
		}
	}()
	f, _ = os.Open("file2.txt")
	defer func() {
		// fmt.Println(f)
		if err := f.Close(); err != nil {
			fmt.Printf("defer close file2.txt err %v\n", err)
		} else {
			fmt.Println("file2 success close")
		}
	}()
}
func main() {

}
