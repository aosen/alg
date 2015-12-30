package tree

import "container/list"

//Btree的实现，开发者只需要实现Element接口，即可使用Btree

//B 树
type BTree struct {
	Head *Node
	Size int
}

//构造一颗B tree
func NewBTree(el Element) *BTree {
	n := NewNode(el)
	if n == nil {
		return &BTree{}
	} else {
		return &BTree{
			Head: n,
			Size: 1,
		}
	}
}

//插入节点
//节点的左节点都小于该节点 节点的右节点都大于该节点
func (t *BTree) Insert(el Element) {
	//如果是一个空树，直接将Head设置为n size＋＋
	n := NewNode(el)
	if t.Head == nil {
		t.Head = n
		t.Size++
		return
	}

	h := t.Head
	//n 与 h进行比较 如果小于h，并且h的左节点为空，则将n填充子节点
	//如果h的左节点不为空，则将h移动到左节点，等待下次循环
	//右节点比较同理
	for {
		if n.Element.Compare(h.Element) == -1 {
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
	t.Size++
}

//搜索B树中以某个节点为根节点的子树中是否存在某个节点
func (t *BTree) Search(el Element) *Node {
	h := t.Head
	//如果n小于h h等于h的左节点，如果n大于h h等于h的右节点
	//如果n＝h 则返回查找到结果
	//直到h == nil 则没有搜索到结果
	for h != nil {
		switch h.Element.Compare(el) {
		case -1:
			h = h.Right
		case 1:
			h = h.Left
		case 0:
			return h
		}
	}
	return nil
}

//打印树结构
//遍历
func (t *BTree) PrintTree(f func([]*Node)) {
	//存储每一行的节点位置
	store := list.New()
	store.PushBack(t.Head)
	l := []*Node{}
	l = append(l, t.Head)
	for store.Len() != 0 {
		element := store.Back()
		n := store.Remove(element).(*Node)
		if n.Left != nil {
			l = append(l, n.Left)
			store.PushBack(n.Left)
		}
		if n.Right != nil {
			l = append(l, n.Right)
			store.PushBack(n.Right)
		}
	}
	f(l)
}

//删除树中的某个节点，如果这个节点在这个树中则删除并返回true
//如果不在这个树中则返回false
func (t *BTree) Delete(el Element) bool {
	n := t.Search(el)
	if n != nil {
		BTreeDelete(n)
		return true
	} else {
		return false
	}
}

//B树中寻找以某个节点为根节点的最大节点
func BTreeMax(n *Node) *Node {
	h := n
	if h == nil {
		return nil
	} else {
		for h.Right != nil {
			h = h.Right
		}
		return h
	}
}

//B树中以某个节点为根的最小节点
func BTreeMin(n *Node) *Node {
	h := n
	if h == nil {
		return nil
	} else {
		for h.Left != nil {
			h = h.Left
		}
		return h
	}
}

//搜索以某个节点为根节点的BTree中是否存在某个元素，
//如果存在返回相应节点
//不存在返回nil
func BTreeSearch(h *Node, el Element) *Node {
	tmp := h
	for tmp != nil {
		switch tmp.Element.Compare(el) {
		case -1:
			tmp = tmp.Right
		case 1:
			tmp = tmp.Left
		case 0:
			return tmp
		}
	}
	return nil
}

//删除BTree中的某个节点
func BTreeDelete(n *Node) Element {
	//如果是叶子节点直接删除
	if IsLeaf(n) {
		DeleteLeaf(n)
		return n.Element
	} else {
		//如果不是叶子节点，找一个左子数中最大的替换，如果左子树为空
		//找右子树中最小的替换，被替换的位置递归删除
		var replace *Node
		if n.Left != nil {
			replace = BTreeMax(n.Left)
		} else {
			replace = BTreeMin(n.Right)
		}
		el := n.Element
		n.Element = BTreeDelete(replace)
		return el
	}
}
