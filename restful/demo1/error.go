/*
 * @Author: your name
 * @Date: 2019-11-20 19:05:02
 * @LastEditTime: 2019-11-20 19:05:13
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \GoWebd:\pratice\restful\demo1\error.go
 */
package main

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
