package skiplist

import (
	"log"
	"testing"
)

type El struct {
	Value int
}

func NewEl(v int) *El {
	return &El{
		Value: v,
	}
}

func TestRandomLevel(t *testing.T) {
	sl := NewSkipList(4)
	for i := 0; i < 100; i++ {
		log.Println(sl.randomLevel())
	}
}
