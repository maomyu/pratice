package main

import (
  "fmt"
)

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println(r)
    }
  }()

  b := []int{0, 1}
  fmt.Println("Hello, playground", b[2])
}