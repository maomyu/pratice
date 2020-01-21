package main

import (
	"fmt"
	"sync"
)

type inter interface {
}
type set struct {
	m map[inter]bool
	sync.RWMutex
}

func New() *set {
	return &set{
		m: map[inter]bool{},
	}
}
func (s *set) Add(item inter) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}
func (s *set) Delete(item inter) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}
func main() {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Delete(4)
	fmt.Println(s.m)

}
