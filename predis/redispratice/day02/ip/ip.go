package ip

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
)

// 将IP进行整数的转换
func IpToScore(ip string) int {
	score := 0
	for _, v := range strings.Split(ip, ".") {
		num, _ := strconv.Atoi(v)
		score = score*256 + num
	}
	return score
}
func ImportIpToRedis(conn *redis.Conn, filename string) error {
	// csv_file :=csv.NewReader(os.OpenFile())
	csv_file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	r2 := csv.NewReader(strings.NewReader(string(csv_file)))

	ss, _ := r2.Read()
	fmt.Println(ss)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(ss[i])
	// }
	fmt.Println(ss)
	return nil
}

// 创建IP地址和ID的映射
