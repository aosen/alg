package tree

//各类树的接口, 开发者只要实现相应接口即可获得相应树

//树节点的接口, 无论实现什么树，该接口必须实现
/*
type Node interface {
	//获取节点中的值
	//两个节点比较大小
	//小于n  -1
	//大于n  1
	//等于n  0
	Compare(n Node) int
	//获取父节点
	GetParent() Node
	//获取左子数
	GetLeft() Node
	//获取右子数
	GetRight() Node
	//设置节点的父节点
	SetParent(Node)
	//设置节点的左节点
	SetLeft(Node)
	//设置节点的右节点
	SetRight(Node)
}
*/

type Element interface {
	//两个节点比较大小
	//小于el  -1
	//大于el  1
	//等于el  0
	Compare(el interface{}) int
}

type Node struct {
	Element Element
	Parent  *Node
	Left    *Node
	Right   *Node
}

//B 树
type BTree struct {
	Head *Node
	Size int
}

//构造一颗B tree
func NewBTree(n *Node) *BTree {
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
func (t *BTree) Insert(n *Node) {
	//如果是一个空树，直接将Head设置为n size＋＋
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

func (t *BTree) Search(n *Node) *Node {
	h := t.Head
	//如果n小于h h等于h的左节点，如果n大于h h等于h的右节点
	//如果n＝h 则返回查找到结果
	//直到h == nil 则没有搜索到结果
	for h != nil {
		switch h.Element.Compare(n.Element) {
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
