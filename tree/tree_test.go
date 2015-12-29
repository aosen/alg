package tree

import (
	"log"
	"testing"
)

//数据结构之B树

type EL struct {
	V int
}

func (self *EL) Compare(el interface{}) int {
	e := el.(*EL)
	if self.V < e.V {
		return -1
	} else if self.V > e.V {
		return 1
	} else {
		return 0
	}
}

func NewNode(v int) *Node {
	element := &EL{V: v}
	node := new(Node)
	node.Element = element
	return node
}

func TestBTree(t *testing.T) {
	bt := NewBTree(NewNode(0))
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		bt.Insert(NewNode(v))
	}
	log.Println(bt.Search(NewNode(6)).Element.(*EL).V)
}
