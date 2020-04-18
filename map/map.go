package hashmap

import (
	"errors"

	"github.com/go-ego/murmur"
)

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type HashMap struct {
	count   int // live cells == size of map. (used by len() builtin)
	size    uint32
	buckets []*Node
}

func NewHashMap(size uint32) (*HashMap, error) {
	if size < 0 {
		return nil, errors.New("Incorrect initial size for map")
	}
	mp := new(HashMap)
	mp.count = 0
	mp.size = size
	mp.buckets = make([]*Node, size)
	return mp, nil
}

// mapKeyToIndex maps a key to an index in the hash table
func hashKey(s string) uint32 {
	return murmur.Murmur3([]byte(s), 0)
}

func (mp *HashMap) Set(key string, value interface{}) error {
	mp.count++
	index := hashKey(key)
	B := mp.buckets[index%mp.size]
	if B == nil {
		mp.buckets[index%mp.size] = &Node{key, value, nil}
	} else {
		for B.next != nil {
			B = B.next
		}
		B.next = &Node{key, value, nil}
	}

	return nil
}

func (mp *HashMap) Get(key string) interface{} {
	index := hashKey(key)
	B := mp.buckets[index%mp.size]
	for B != nil {
		if B.key == key {
			return B.value
		}
		B = B.next
	}

	return nil
}

func (mp *HashMap) Del(key interface{}) error {
	return nil
}

func (mp *HashMap) Len() int {
	if mp == nil {
		return 0
	}

	return mp.count
}
