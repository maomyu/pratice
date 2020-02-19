package simplefactory

// 工厂类
type Factory struct {
}

func NewProduct(str string) TV {
	if str == "haier" {
		return &Haier{}
	} else if str == "huawei" {
		return &HuaWei{}
	}
	return nil
}
