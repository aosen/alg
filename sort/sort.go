package sort

//各种排序算法实现

//开发者只需要实现以下接口，即可使用以下排序方法, 详情见sort_test.go
type Vector interface {
	Len() int
	//如果第i个元素小于第j个元素则返回True : False
	Less(i, j int) bool
	//交换
	Swap(i, j int)
}

//插入排序法
//每次处理就是将无序数列的第一个元素与有序数列的元素从后往前逐个进行比较，
//找出插入位置，将该元素插入到有序数列的合适位置中。
//假设在一个无序的数组中，要将该数组中的数按插入排序的方法从小到大排序。
//假设啊a[]={3,5,2,1,4};插入排序的思想就是比大小，满足条件交换位置，
//一开始会像冒泡排序一样，但会比冒泡多一步就是交换后（a[i]=a[i+1]后）原位置（a[i]）
//会继续和前面的数比较满足条件交换，直到a[i+1]前面的数组是有序的。
//比如在第二次比较后数组变成a[]={2,3,5,1,4};
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

/*
快速排序法
设要排序的数组是A[0]……A[N-1]，首先任意选取一个数据（通常选用数组的第一个数）作为关键数据，然后将所有比它小的数都放到它前面，所有比它大的数都放到它后面，这个过程称为一趟快速排序。值得注意的是，快速排序不是一种稳定的排序算法，也就是说，多个相同的值的相对位置也许会在算法结束时产生变动。
一趟快速排序的算法是：
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]和A[i]互换；
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]互换；
5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，
i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。
*/

func QuickSort(vector Vector) Vector {
	quicksort(vector, 0, vector.Len()-1)
	return vector
}

func quicksort(vector Vector, l int, r int) {
	if l >= r {
		return
	}

	mid := l
	i := l + 1

	for j := l; j <= r; j++ {
		if vector.Less(j, mid) {
			vector.Swap(i, j)
			i++
		}
	}

	vector.Swap(l, i-1)

	quicksort(vector, l, i-2)
	quicksort(vector, i, r)
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
