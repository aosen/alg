package tree

//各类树的接口, 开发者只要实现相应接口即可获得相应树

import ()

//树节点的接口, 无论实现什么树，该接口必须实现
type Node interface {
	//两个节点的比较
	Compare(*Node) int
}

//B树 实现举例: btree.go
type Btree interface {
	//树的插入
	Insert(el interface{})
	//树的删除 删除成功返回True : False
	Delete(el interface{}) bool
	//树的搜索
	Search(el interface{}) *Node
}
