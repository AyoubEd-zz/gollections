package hashmap

type Node struct {
	Key   string
	value interface{}
	Next  *Node
}

type HashMap struct {
	size    int
	buckets []*Node
}

func NewHashMap(size int) (*HashMap, error) {
	return nil, nil
}

func (mp *HashMap) Get(key interface{}) (interface{}, error) {
	return nil, nil
}

func (mp *HashMap) Set(key, value interface{}) error {
	return nil
}

func (mp *HashMap) Del(key interface{}) error {
	return nil
}
