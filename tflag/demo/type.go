/*
 * @Author: your name
 * @Date: 2019-11-22 15:38:48
 * @LastEditTime: 2019-11-22 19:19:56
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\tflag\demo\type.go
 */
package main

import "fmt"

import "errors"

import "strings"

type self struct {
	Address string
	Area    string
}

func (s *self) String() string {
	// fmt.Printf("%p\n", s)
	*s = self{"Beingjing", "China"}
	// s = &self{"Beingjing", "China"}
	// fmt.Printf("%p", s)
	return fmt.Sprintf("%s", *s)
}
func (s *self) Set(value string) error {
	if len(value) <= 0 {
		return errors.New("参数为空，请输入正确的参数")
	}
	data := strings.Split(value, " ")
	*s = self{
		data[0],
		data[1],
	}
	fmt.Println(*s)
	return nil
}
