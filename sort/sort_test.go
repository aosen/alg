package sort

/*
测试文件
*/

import (
	"log"
	"testing"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func Test_InsertSort(t *testing.T) {
	vector := List{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(InsertSort(vector))
}

func Test_BubbleSort(t *testing.T) {
	vector := List{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(BubbleSort(vector))
}

func Test_SelectSort(t *testing.T) {
	vector := List{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(SelectSort(vector))
}

//快速排序法
func Test_QuickSort(t *testing.T) {
	log.Println("快速排序法结果:")
	vector := List{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(QuickSort(vector))
}

func TestReverse(t *testing.T) {
	v1 := List{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(Reverse(v1))
}
