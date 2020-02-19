package simplefactory

import "testing"

func TestNewProduct(t *testing.T) {
	tv := NewProduct(structConfig.Name)
	tv.play()
}
