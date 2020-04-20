package hashmap

import (
	"errors"
	"math"

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

// hashKey uses Murmur hashing algorithm and returns a strings hash
func hashKey(s string) uint32 {
	return murmur.Murmur3([]byte(s), 0)
}

func getHashTableIndex(key string, size uint32) uint32 {
	h := hashKey(key)
	mask := uint32(math.Log2(float64(size))+1) / 2
	return h & (1<<mask - 1)
}

func (mp *HashMap) Set(key string, value interface{}) error {
	mp.count++
	index := getHashTableIndex(key, mp.size)
	B := mp.buckets[index]

	if B == nil {
		mp.buckets[index] = &Node{key, value, nil}
	} else {
		for B.next != nil {
			B = B.next
		}
		B.next = &Node{key, value, nil}
	}

	return nil
}

func (mp *HashMap) Get(key string) interface{} {
	index := getHashTableIndex(key, mp.size)
	B := mp.buckets[index]

	for B != nil {
		if B.key == key {
			return B.value
		}
		B = B.next
	}

	return nil
}

func (mp *HashMap) Del(key string) error {
	mp.count--
	index := getHashTableIndex(key, mp.size)
	B := mp.buckets[index]
	nB := &Node{"", nil, nil}
	head := nB

	for B != nil {
		if B.key != key {
			nB.next = &Node{B.key, B.value, nil}
			nB = nB.next
		}
		B = B.next
	}

	mp.buckets[index] = head.next
	return nil
}

func (mp *HashMap) Len() int {
	if mp == nil {
		return 0
	}

	return mp.count
}
