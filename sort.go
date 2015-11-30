package alg

//插入排序法
func InsertSort(vector []int) []int {
	for i := 1; i < len(vector); i++ {
		// 每一趟不满足条件就选择i为哨兵保存，将哨兵插入0~i-1有序序列（0~i-1始终是有序的）
		if vector[i] < vector[i-1] {
			temp := vector[i]
			//后移直到找到哨兵合适的位置
			j := i - 1
			for ; j >= 0 && vector[j] > temp; j-- {
				vector[j+1] = vector[j]
			}
			//插入位置前后都是有序的，最后也是有序的
			vector[j+1] = temp
		}
	}
	return vector
}

//冒泡排序法
func BubbleSort(vector []int) []int {
	for i := 0; i < len(vector); i++ {
		// 每一趟将最大的数冒泡
		for j := 0; j < len(vector)-i-1; j++ {
			if vector[j] > vector[j+1] {
				temp := vector[j]
				vector[j] = vector[j+1]
				vector[j+1] = temp
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
func SelectSort(vector []int) []int {
	for i := 0; i < len(vector); i++ {
		// 选择最小的元素
		k := i
		for j := i + 1; j < len(vector); j++ {
			if vector[k] > vector[j] {
				k = j
			}
		}
		// 交换
		if k != i {
			vector[i], vector[k] = vector[k], vector[i]
		}
	}
	return vector
}
