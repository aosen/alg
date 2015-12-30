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

func NewEL(v int) *EL {
	return &EL{V: v}
}

func TestBTree(t *testing.T) {
	bt := NewBTree(NewEL(0))
	for _, v := range []int{1, 4, 2, 3, 5, 7, 9, 8, 6} {
		bt.Insert(NewEL(v))
	}
	log.Println(bt.Search(NewEL(6)).Element.(*EL).V)
	//寻找最大元素
	log.Println(BTreeMax(bt.Head).Element.(*EL).V)
	//最小元素
	log.Println(BTreeMin(bt.Head).Element.(*EL).V)
	//判断是否为根节点
	log.Println(IsRoot(bt.Head))
	//判断以某个节点为根节点的BTree 是否存在某个元素
	log.Println(BTreeSearch(bt.Head, NewEL(9)).Element.(*EL).V)
	//删除Btree中的某个节点
	log.Println(BTreeDelete(bt.Head).(*EL).V)
	log.Println(bt.Search(NewEL(0)))
	//打印头节点元素值
	log.Println(bt.Head.Element.(*EL).V)
	//删除树的元素
	log.Println(bt.Delete(NewEL(9)))
	//打印树结构
	bt.PrintTree(func(s []*Node) {
		l := []int{}
		for _, v := range s {
			l = append(l, v.Element.(*EL).V)
		}
		log.Println(l)
	})
}
