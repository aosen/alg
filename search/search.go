package search

import "container/list"

//一系列排序算法实现

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

//只要开发者实现上面接口，如下 就可以正常使用BinSearch
/*
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
*/

//如果查找到返回查找到的位置，查找不到返回－1
//时间复杂度O(log(n))
func BinSearch(list Searchable, item interface{}) int {
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

//使用分治策略求解一个包含负数数组的最大子数组
//只有当数组中包涵负数求解最大子数组才有意义，否则最大子数组就是其本身
//返回最大子数组的开始和结束 及和
//数组中一定是即包涵整数也包涵负数，否则如果都是正数，最大子数组为本身，如果都为负数，最大子数组为数组中的最大值
//时间复杂度：
//O(n∗log n)

func DCSearchMaxSubArray(l []int, low, high int) (max_left, max_right, sum int) {
	if high == low {
		max_left = low
		max_right = low
		sum = l[low]
		return
	} else {
		mid := (low + high) / 2
		l_low, l_high, l_sum := DCSearchMaxSubArray(l, low, mid)
		r_low, r_high, r_sum := DCSearchMaxSubArray(l, mid+1, high)
		c_low, c_high, c_sum := dcFindMaxCrossingSubArray(l, low, mid, high)
		if l_sum >= r_sum && l_sum >= c_sum {
			max_left = l_low
			max_right = l_high
			sum = l_sum
		} else if r_sum >= l_sum && r_sum >= c_sum {
			max_left = r_low
			max_right = r_high
			sum = r_sum
		} else {
			max_left = c_low
			max_right = c_high
			sum = c_sum
		}
		return
	}
}

func dcFindMaxCrossingSubArray(l []int, low, mid, high int) (max_left, max_right, sum int) {
	left_sum := 0
	tmp := 0
	for i := mid; i >= low; i-- {
		tmp += l[i]
		if tmp > left_sum {
			left_sum = tmp
			max_left = i
		}
	}
	right_sum := 0
	tmp = 0
	for i := mid + 1; i <= high; i++ {
		tmp += l[i]
		if tmp > right_sum {
			right_sum = tmp
			max_right = i
		}
	}
	sum = left_sum + right_sum
	return
}

/*
暴力方法实现最大子数组
暴力求解法

该方法的思想非常简单，先找出从第1个元素开始的最大子数组，而后再从第2个元素开始找出从第2个元素开始的最大子数组，依次类推，比较得出最大的子数组。实现代码如下：

/*
常规方法,时间复杂度O（n*n）
先从第一个元素开始向后累加，
每次累加后与之前的和比较，保留最大值，
再从第二个元素开始向后累加，以此类推。
*/
/*
int MaxSubSum1(int *arr,int len)
{
    int i,j;
    int MaxSum = 0;
    //每次开始累加的起始位置的循环
    for(i=0;i<len;i++)
    {
        int CurSum = 0;
        //向后累加的循环
        for(j=i;j<len;j++)
        {
            CurSum += arr[j];
            if(CurSum > MaxSum)
            MaxSum = CurSum;
        }
    }
    return MaxSum;
}
*/

/*
时间复杂度：n*n
求最大子数组
用空间换时间的一个算法
此方法比递归方法更好理解
*/
func MaxSubSum(l []int) (maxsum int) {
	queue := list.New()
	length := len(l)
	for i := 0; i < length; i++ {
		cursum := 0
		if queue.Len() == 0 {
			for j := i; j < length; j++ {
				cursum += l[j]
				if cursum > maxsum {
					maxsum = cursum
				}
				queue.PushBack(cursum - l[i])
			}
		} else {
			for k := 0; k < i; k++ {
				el := queue.Back()
				cursum := queue.Remove(el).(int)
				if cursum > maxsum {
					maxsum = cursum
				}
				queue.PushBack(cursum - l[i])
			}
		}
	}
	return
}

/*
动态规划，时间复杂度 n
*/
func DPMaxSubSum(l []int) (maxsum int) {
	length := len(l)
	dp := []int{}
	dp[0] = l[0]
	maxsum = dp[0]
	tmp := 0 //记录起始位置
	start := 0
	end := 0
	for i := 1; i < length; i++ {
		if dp[i-1] <= 0 {
			dp[i] = l[i]
			tmp = i
		} else {
			dp[i] = l[i] + dp[i-1] //向后求和
		}
		if dp[i] > maxsum {
			maxsum = dp[i]
			start = tmp
			end = i
		}
	}
	return
}
