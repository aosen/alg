package sort

//各种排序算法实现

import ()

//开发者只需要实现以下接口，即可使用以下排序方法, 详情见sort_test.go
type Vector interface {
	Len() int
	//如果第i个元素小于第j个元素则返回True : False
	Less(i, j int) bool
	//交换
	Swap(i, j int)
}

//插入排序法
func InsertSort(vector Vector) Vector {
	l := vector.Len() - 1
	for i := 1; i <= l; i++ {
		// 每一趟不满足条件就选择i为哨兵保存，将哨兵插入0~i-1有序序列（0~i-1始终是有序的）
		for j := i; j > 0 && vector.Less(j, j-1); j-- {
			vector.Swap(j, j-1)
		}
	}
	return vector
}

//冒泡排序法
func BubbleSort(vector Vector) Vector {
	l := vector.Len() - 1
	for i := 0; i < l; i++ {
		// 每一趟将最大的数冒泡
		for j := l; j > i; j-- {
			if vector.Less(j, j-1) {
				vector.Swap(j, j-1)
			}
		}
	}
	return vector
}

/*选择排序
时间复杂度
排序算法复杂度对比 lgn = log2n
排序算法复杂度对比 lgn = log2n
选择排序的交换操作介于 0 和 (n - 1） 次之间。
选择排序的比较操作为 n (n - 1） / 2 次之间。选择排序的赋值操作介于 0 和 3 (n - 1） 次之间。
比较次数O(n^2），比较次数与关键字的初始状态无关，总的比较次数N=(n-1）+(n-2）+...+1=n*(n-1）/2。
交换次数O(n），最好情况是，已经有序，交换0次；最坏情况交换n-1次，逆序交换n/2次。
交换次数比冒泡排序少多了，由于交换所需CPU时间比比较所需的CPU时间多，n值较小时，选择排序比冒泡排序快。*/
func SelectSort(vector Vector) Vector {
	l := vector.Len() - 1
	for i := 0; i < l; i++ {
		// 选择最小的元素
		min := i
		for j := i + 1; j <= l; j++ {
			if vector.Less(j, min) {
				min = j
			}
		}
		// 交换
		vector.Swap(i, min)
	}
	return vector
}

//将排序好的数据进行翻转
func Reverse(vector Vector) Vector {
	l := vector.Len() - 1
	r := vector.Len() / 2
	for i := 0; i < r; i++ {
		vector.Swap(i, l-i)
	}
	return vector
}
