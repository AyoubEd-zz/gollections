package stack

import (
	"fmt"
)

type stack struct {
	m []interface{}
}

func New(vals ...interface{}) *stack {
	st := &stack{[]interface{}{}}

	for _, v := range vals {
		st.Insert(v)
	}
	
	return st
}

func (s *stack) Insert(val interface{}) {
	s.m = append(s.m, val)
}

func (s *stack) Remove() {
	s.m = s.m[:len(s.m) - 1]
}

func (s *stack) Size() int {
	return len(s.m)
}

func (s *stack) Show() string {
	str := ""

	for _, v := range s.m {
		str += fmt.Sprintf("%v ", v)
	}

	fmt.Println(str)

	return str[:len(str) - 1]
}
