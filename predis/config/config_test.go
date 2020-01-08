package config

import (
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

type AStruct struct {
	A string `yaml:"a"`
}
type BStruct struct {
	AStruct `yaml:",inline"`
	C       AStruct
	B       string `yaml:"b"`
}

var data = `
a: a string from struct A
b: a string from struct B
c:
  a: astring from c.a
`

func TestConfig(t *testing.T) {
	// 读取配置文件
	// redisfile, _ := ioutil.ReadFile("application-redis.yaml")
	var b BStruct
	//开始加载配置文件
	// yaml.Unmarshal(redisfile, &redisConfig)
	yaml.Unmarshal([]byte(data), &b)
	fmt.Println(b.B)
	fmt.Println(b.A)
	fmt.Println(b.C.A)
}
