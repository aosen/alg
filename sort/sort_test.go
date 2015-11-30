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

func TestReverse(t *testing.T) {
	v1 := List{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	v2 := List{9, 8, 7, 6, 5, 4, 3, 2, 1}
	log.Println(Reverse(v1))
	log.Println(Reverse(v2))
}
