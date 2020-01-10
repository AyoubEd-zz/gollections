package set 

import (
	"fmt"
)

type set struct {
	m map[interface{}]struct{}
}

func New(vals ...interface{}) *set{
	s := &set{make(map[interface{}]struct{})}

	for _, val := range vals {
		s.Insert(val)
	}

	return s
}

func (s *set) Insert(val interface{}) {
	s.m[val] = struct{}{}
}

func (s *set) Remove(val interface{}) {
	delete(s.m, val)
}

func (s *set) Contains(val interface{}) bool {
	if _, ok := s.m[val]; ok {
		return true
	}
	return false
}

func (s *set) Size() int {
	return len(s.m)
}

func (s *set) Show() {
	mp := s.m

	for key, _ := range mp {
		fmt.Printf("%v ", key)	
	}
	fmt.Printf("\n")
}
