/*
Author: Aosen
Data: 2016-01-05
Contact: QQ 316052486
Desc:
跳表(skip List)是一种随机化的数据结构，基于并联的链表，实现简单，插入、删除、查找的复杂度均为O(logN)。

链表的一种，只不过它在链表的基础上增加了跳跃功能，正是这个跳跃的功能，使得在查找元素时，跳表能够提供O(log n)的时间复杂
度。红黑树等这样的平衡数据结构查找的时间复杂度也是O(log n)，但是要实现像红黑树这样的数据结构并非易事,但是只要你熟悉链表
的基本操作,再加之对跳表原理的理解，实现一个跳表数据结构就是一个很自然的事情了。
此外，跳表在当前热门的开源项目中也有很多应用，比如LevelDB的核心数据结构memtable是用跳表实现的，redis的sorted set数据结构
也是有跳表实现的。
详细说明参考：http://blog.csdn.net/daniel_ustc/article/details/20218489?utm_source=tuicool&utm_medium=referral
*/

package skiplist

import "math/rand"

/*
开发者只要实现以下接口，即可使用skiplist
*/
type Element interface {
	//两个节点比较大小
	//小于el  -1
	//大于el  1
	//等于el  0
	Compare(el interface{}) int
}

//跳表节点结构体
type SkipListNode struct {
	Level   int //当前所在层级
	Element Element
	Right   *SkipListNode //跳表的右边节点
}

//新建一个节点
func NewSkipListNode(el Element) *SkipListNode {
	return &SkipListNode{
		Element: el,
	}
}

//跳表结构体
type SkipList struct {
	maxLevel int             // 跳表的最大层级数
	length   []int           //每一层的数据量
	head     []*SkipListNode //指向头结点
}

//新建一个跳表
func NewSkipList(level int) *SkipList {
	sl := new(SkipList)
	sl.maxLevel = level
	for i := 0; i < level; i++ {
		sl.head[i] = &SkipListNode{}
		sl.head[i].Right = nil
		sl.length[i] = 0
	}
	return sl
}

//我们知道跳表是一种随机化数据结构，
//其随机化体现在插入元素的时候元素所占有的层数完全是随机的，
//层数是通过随机算法产生的
func (self *SkipList) randomLevel() (level int) {
	level = 1
	for rand.Int31()%2 != 0 {
		level++
	}
	if level > self.maxLevel {
		level = self.maxLevel
	}
	return
}

//获取跳表查询的起始位置
func (self *SkipList) StartPos() *SkipListNode {
	return self.head[self.maxLevel-1]
}

//获取跳表的数据量
func (self *SkipList) Length(level int) int {
	return self.length[level]
}

func (self *SkipList) Insert(el Element) {
	level := self.randomLevel()
	nodelist := make([]SkipListNode, level)
	//由高层像低层插入
	for i := self.maxLevel - 1; i >= 0; i-- {
		tmp := self.head[i].Right
		//如果本层中没有数据则直接插入
		for {
			//如果已经到达末尾，则在末尾插入
			if tmp.Right == nil {
				nodelist[i].Right = tmp.Right
				tmp.Right = &nodelist[i]
			} else {
				if tmp.Right.Element.Compare(nodelist[i].Element) < -1 {
					tmp = tmp.Right
				} else {
					nodelist[i].Right = tmp.Right
					tmp.Right = &nodelist[i]
				}
			}
		}
	}
}
