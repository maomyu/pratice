package main

import "fmt"

// 定义一个最小的容量
const minQueueLen = 16

// 定义队列的结构体
type Queue struct {
	buf               []interface{}
	head, tail, count int
}

// 返回一个新的队列
func New() *Queue {
	return &Queue{
		buf: make([]interface{}, minQueueLen),
	}
}

// 获得队列的长度
func (q *Queue) Length() int {
	return q.count
}

// 重置队列的大小
func (q *Queue) resize() {
	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}
	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

// 添加一个元素
func (q *Queue) Add(elem interface{}) {
	if q.count == len(q.buf) {
		q.resize()
	}
	q.buf[q.tail] = elem
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

// 返回一个队列的头元素
func (q *Queue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue:peek()called on empty queue")
	}
	return q.buf[q.head]
}

// 获取一个元素
func (q *Queue) Get(i int) interface{} {
	if i < 0 {
		i += q.count
	}

	if i < 0 || i >= q.count {
		panic("queue:Get()called on empty queue")
	}
	return q.buf[i]
}

// 移除一个元素
func (q *Queue) Remove() interface{} {
	if q.count <= 0 {
		panic("queue:Remove() called on empty queue")
	}
	ret := q.buf[q.tail-1]
	q.buf[q.tail-1] = nil
	q.count--
	q.tail--
	return ret
}
func main() {
	q := New()
	q.Add("first")
	q.Add("second")
	q.Add("third")
	q.Add("firt1")
	q.Add("fite3")
	q.Add("feere0")
	q.Add("fdsfd3")
	q.Add("third4")
	q.Add("third54")
	q.Add("third6")
	q.Add("third7")
	q.Add("third8s")
	fmt.Println(q.Get(0))

	q.Remove()
	fmt.Println(q.Get(-1))
}
