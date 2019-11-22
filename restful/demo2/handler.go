/*
 * @Author: your name
 * @Date: 2019-11-22 13:43:28
 * @LastEditTime: 2019-11-22 13:43:55
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\restful\demo2\handler.go
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoId)
}
