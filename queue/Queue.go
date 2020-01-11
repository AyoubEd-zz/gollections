package queue

import (
	"fmt"
)

type queue struct {
	m []interface{}
}

func New(vals ...interface{}) *queue {
	q := &queue{[]interface{}{}}

	for _, v := range vals {
		q.Insert(v)
	}

	return q
}

func (q *queue) Insert(val interface{}) {
	q.m = append([]interface{}{val}, q.m...)
}

func (q *queue) Pop() interface{} {
	if len(q.m) == 0 {
		return nil
	}

	last_element := q.m[len(q.m) - 1]
	q.m = q.m[:len(q.m) - 1]

	return last_element
}

func (q *queue) Size() int {
	return len(q.m)
}

func (q *queue) Show(vals ...int) string {
	str := ""

	for _,v := range q.m {
		str += fmt.Sprintf("%v ", v)
	}
	
	if len(vals) == 0 || vals[0] != 0 {
		fmt.Printf("%v\n", str)
	}

	return str
}
