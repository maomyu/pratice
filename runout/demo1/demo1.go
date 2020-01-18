package main

func fool() *int {
	t := 3
	return &t
}

func main() {
	fool()
	// fmt.Println(*x)
}

// go build -gcflags '-m -l' demo1.go

// 反汇编
// go tool compile -S demo1.go
