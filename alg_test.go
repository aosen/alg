package alg

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

func Test_InsertSort(t *testing.T) {
	vector := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(InsertSort(vector))
}

func Test_BubbleSort(t *testing.T) {
	vector := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(BubbleSort(vector))
}

func Test_SelectSort(t *testing.T) {
	vector := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	log.Println(SelectSort(vector))
}

func Test_BinSearch(t *testing.T) {
	list := List{1, 1, 2, 3, 5, 9, 10, 11, 12, 13}

	index := BinSearch(list, 3)
	log.Printf("The index of 3 in the list is %d\n", index)

	index = BinSearch(list, 4)
	log.Printf("The index of 4 in the list is %d\n", index)
}

func Test_SearchMaxSubArray(t *testing.T) {
	//list := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	list := []int{-13, -3, -25, -20, -3, -16, -23, -18, -20, -7, -12, -5, -22, -15, -4, -7}
	s, e, sum := DCSearchMaxSubArray(list, 0, len(list)-1)
	log.Printf("The max subarray in the list is %d-%d-%d\n", s, e, sum)
}
