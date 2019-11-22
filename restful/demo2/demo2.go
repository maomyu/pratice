/*
 * @Author: your name
 * @Date: 2019-11-22 13:34:13
 * @LastEditTime: 2019-11-22 13:52:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\restful\demo2\demo2.go
 */
package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
