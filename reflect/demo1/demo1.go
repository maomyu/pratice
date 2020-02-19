package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Student struct {
	ID   string
	Name string
	age  int
}

// 根据字段名称获取实例获取值
func GetField(i interface{}, parm string) (interface{}, error) {
	// 判断字段的是否存在
	if !IsHasField(i, []reflect.Kind{reflect.Struct, reflect.Ptr}) {
		return nil, errors.New("interface{} is not struct or ptr")
	}
	value := reflect.ValueOf(i)
	pvalue := value.FieldByName(parm)
	fmt.Println(pvalue)
	return pvalue, nil
}

// 判断是否是结构体
func IsHasField(i interface{}, k []reflect.Kind) bool {
	for _, v := range k {
		if reflect.TypeOf(i).Kind() == v {
			return true
		}
	}
	return false
}
func main() {
	s := Student{"asdflkdjkls", "yuwei", 12}
	value, _ := GetField(s, "ID")
	fmt.Println(value)
}
