package simplefactory

import "fmt"

type Haier struct {
}

func (h *Haier) play() {
	fmt.Println("haier")
}
