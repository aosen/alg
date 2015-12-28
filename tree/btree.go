package tree

//数据结构之B树

import ()

type Node struct {
	value  interface{}
	parent *Node
	left   *Node
	right  *Node
}

func NewNode(el interface{}) *Node {
	return &Node{value: el}
}

//两个节点比较大小
//小于m  -1
//大于m  1
//等于m  0
func (n *Node) Compare(m *Node) int {
	nv := n.value.(int)
	mv := n.value.(int)
	if nv < mv {
		return -1
	} else if nv > mv {
		return 1
	} else {
		return 0
	}
}

type Tree struct {
	head *Node
	size int
}

func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	} else {
		return &Tree{head: n, size: 1}
	}
}

//插入节点
//广度插入
//节点的左节点都小于该节点 节点的右节点都大于该节点
func (t *Tree) Insert(el interface{}) {
	n := NewNode(el)
	if t.head == nil {
		t.head = n
		t.size++
		return
	}
	h := t.head
	for {
		if n.Compare(h) == -1 {
			if h.Left == nil {
				h.Left = n
				n.Parent = h
				break
			} else {
				h = h.Left
			}
		} else {
			if h.Right == nil {
				h.Right = n
				n.Parent = h
				break
			} else {
				h = h.Right
			}
		}
	}
	t.size++
}

func (t *Tree) Search(el interface{}) *Node {
	h := t.head
	n := NewNode(el)
	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.right
		case 1:
			h = h.left
		case 0:
			return h
		}
	}
	return nil
}
