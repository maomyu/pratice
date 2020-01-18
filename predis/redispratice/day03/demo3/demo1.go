package demo3

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

func Add_update_contact(conn redis.Conn, user string, contact string) {
	// 键名
	ac_list := "recent:" + user
	// 使用事务
	conn.Do("MULTI")
	// 移除列表中所有与contact相等的值
	_, err := conn.Do("lrem", ac_list, 0, contact)
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Do("lpush", ac_list, contact)
	if err != nil {
		fmt.Println(err)
	}

	conn.Do("ltrim", ac_list, 0, 99)
	conn.Do("exec")
}

// 获取所有的最近联系人
func Get_Contact(conn redis.Conn, user string) []string {
	ac_list := "recent:" + user
	reply, err := conn.Do("lrange", ac_list, 0, -1)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	result := get_changeto_list(reply)
	// result, _ = redis.Strings(reply, nil)
	return result
}
func get_changeto_list(reply interface{}) []string {
	strs := reply.([]interface{})

	result := []string{}
	for i := 0; i < len(strs); i++ {
		result = append(result, string(strs[i].([]uint8)))
	}

	return result
}

// 匹配联系人
func Fetch_autocomplete_list(conn redis.Conn, user, prefix string) []string {
	// 获取自动补全的列表
	licontact, err := conn.Do("lrange", "recent:"+user, 0, -1)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	match := []string{}
	result := get_changeto_list(licontact)
	for _, v := range result {
		if strings.HasPrefix(v, prefix) {
			match = append(match, v)
		}
	}
	return match
}
