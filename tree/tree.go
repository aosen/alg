package tree

//各类树的接口, 开发者只要实现相应接口即可获得相应树

//树节点的接口, 无论实现什么树，该接口必须实现

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

func NewNode(el Element) *Node {
	return &Node{Element: el}
}

//判断Btree的某个节点是否为整个树的根节点
//如果是根节点返回true
func IsRoot(n *Node) bool {
	return n.Parent == nil
}

//判断Btree的某个节点是否为叶子节点
//如果是则返回true
func IsLeaf(n *Node) bool {
	return n.Left == nil && n.Right == nil
}

//内部方法，删除页节点
//如果存在则删除 并返回节点
func deleteLeaf(n *Node) *Node {
	//如果是根节点直接返回
	if IsRoot(n) {
		return n
	} else {
		if n.Parent.Left == n {
			n.Parent.Left = nil
		} else {
			n.Parent.Right = nil
		}
		return n
	}
}
