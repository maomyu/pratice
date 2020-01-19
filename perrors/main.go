package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func a() error {
	e := errors.New("standard error")
	return errors.Wrap(e, "faffa")
}
func b() error {
	return errors.Wrap(a(), "hahahah")
}
func main() {
	fmt.Println(b())
}
