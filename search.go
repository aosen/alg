package alg

//一系列排序算法实现

import ()

//二分查找
//二分查找又称折半查找，优点是比较次数少，查找速度快，
//平均性能好；其缺点是要求待查表为有序表，
//且插入删除困难。因此，折半查找方法适用于不经常变动而查找频繁的有序列表。
//首先，假设表中元素是按升序排列，
//将表中间位置记录的关键字与查找关键字比较，如果两者相等，则查找成功；
//否则利用中间位置记录将表分成前、后两个子表，
//如果中间位置记录的关键字大于查找关键字，则进一步查找前一子表，
//否则进一步查找后一子表。重复以上过程，
//直到找到满足条件的记录，使查找成功，或直到子表不存在为止，此时查找不成功。
type Searchable interface {
	Len() int
	//列表中index位置与v比较大小 大于返回1 等于返回0 小于返回-1
	//如果item判定出错返回－2
	Compare(index int, item interface{}) int
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l List) Compare(index int, item interface{}) int {
	if value, ok := item.(int); ok {
		if l[index] == value {
			return 0
		} else if l[index] > value {
			return 1
		} else {
			return -1
		}
	}
	return -2
}

//如果查找到返回查找到的位置，查找不到返回－1
func BinSearch(list List, item interface{}) int {
	startFlag := 0
	stopFlag := list.Len() - 1
	middleFlag := (startFlag + stopFlag) / 2

	for startFlag <= stopFlag {
		switch flag := list.Compare(middleFlag, item); flag {
		//无法比较直接返回－2
		case -2:
			return -1
		case -1:
			startFlag = middleFlag + 1
		case 0:
			return middleFlag
		case 1:
			stopFlag = middleFlag - 1
		}
		middleFlag = (startFlag + stopFlag) / 2
	}
	return -1
}
