package link

//数据结构之链表
//链表格式:
/*
   ++++++++++
   + length +       ++++++++++++++++
   + head   +       +              +
   +        + ----> +              +
   + tail   +       +              +
   + lock   +       ++++++++++++++++
   ++++++++++
*/

import (
	"sync"
)

type List struct {
	length int
	head   *Node
	lock   *sync.Mutex
}

func NewLink() *List {
	l := new(List)
	l.length = 0
	l.lock = new(sync.Mutex)
	return l
}

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

//获取链表的长度
func (l *List) Len() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.length
}

//判断链表是否为空
func (l *List) IsEmpty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.length == 0
}

//头插法
func (l *List) HeadInsert(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node := NewNode(value)
	node.next = l.head
	l.head = node
	l.length++
}

//尾插法
func (l *List) TailInsert(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node := NewNode(value)
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
	l.length++
}

//后续增加其它方法 由于Go中用到链表的机会很少，所以暂时先写到这里
