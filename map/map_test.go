package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMap(t *testing.T) {
	mp, err := NewHashMap(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, mp.size, uint32(1))
	assert.Equal(t, mp.count, 0)
}

func TestUseMapNoResize(t *testing.T) {
	mp, err := NewHashMap(8)
	assert.Equal(t, err, nil)

	//1. Test setting and getting keys with different hashes
	mp.Set("a", 1)
	mp.Set("b", 0)

	assert.Equal(t, mp.Get("a"), 1)
	assert.Equal(t, mp.Get("b"), 0)

	//2. Test setting with keys that have equal hashes
	mp, _ = NewHashMap(8)

	mp.Set("aaa", 1)
	mp.Set("bb", 0)

	assert.Equal(t, mp.Get("aaa"), 1)
	assert.Equal(t, mp.Get("bb"), 0)
}

func getFromMap(t *testing.T) {

}

func deleteFromMap(t *testing.T) {

}
