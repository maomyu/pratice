package simplefactory

import "testing"

func TestNewProduct(t *testing.T) {
	huawei := NewProduct("huawei")
	haier := NewProduct("haier")
	huawei.play()
	haier.play()
}
