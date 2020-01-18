package main

import (
	"fmt"
	"io"
	"os"
)

func mergeFile() {
	f, _ := os.Open("file.txt")
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
func main() {
	f, _ := os.Open("file.txt")
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close file.txt err %v\n", err)
		} else {
			fmt.Println("file success close")
		}
	}(f)
	f, _ = os.Open("file2.txt")
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close file2.txt err %v\n", err)
		} else {
			fmt.Println("file2 success close")
		}
	}(f)
}
