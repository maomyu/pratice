package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	redisConfig             defaultRedisConfig
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {
		fmt.Println("配置文件已经被初始化过.....")
		return
	}
	// 获取当前文件的绝对路径
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
	// 获得conf路径
	pr := filepath.Join(appPath, "/../conf")

	// 读取配置文件
	redisfile, _ := ioutil.ReadFile(pr + sp + defaultConfigFilePrefix + "redis.yaml")

	//开始加载配置文件
	yaml.Unmarshal(redisfile, &redisConfig)
	// fmt.Println(redisConfig)
	inited = true
}

// 获取
func GetRedisConfig() RedisConfig {
	return redisConfig
}
