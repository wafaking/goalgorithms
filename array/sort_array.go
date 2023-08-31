package array

import (
	"math/rand"
	"time"
)

//排序数组(leetcode-912)
//给你一个整数数组nums，请你将该数组升序排列。
//示例1：输入：nums=[5,2,3,1],输出：[1,2,3,5]
//示例2：输入：nums=[5,1,1,2,0,0],输出：[0,0,1,1,2,5]

// 使用双路快排(参考QuickSort2Ways,超时)
func sortArray1(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	var partition func(l, r int)
	partition = func(l, r int) {
		if l >= r {
			return
		}

		// 随机索引
		rand.Seed(time.Now().UnixNano())
		randIdx := rand.Intn(r+1-l) + l
		nums[randIdx], nums[l] = nums[l], nums[randIdx]

		// 将随机数与第一个元素l交换位置
		// l从1开始，位置1空出用于放置nums[randIdx]
		i := l + 1
		j := r

		for {
			// l一直向右移动,找到左部分大于nums[randIdx]的数
			for i <= j && nums[i] < nums[l] {
				i++
			}

			//找到部分大于nums[randIdx]的数
			for i <= j && nums[j] > nums[l] {
				j--
			}

			if i >= j {
				break
			}

			//交的l与r
			nums[i], nums[j] = nums[j], nums[i]

			// l,r都各自归到左右部分，因此继续其余元素
			i++
			j--
		}

		nums[l], nums[j] = nums[j], nums[l]
		partition(l, j-1)
		partition(j+1, r)
	}
	partition(0, len(nums)-1)
	return nums
}

// 使用三路快排
func sortArray2(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	//2,1,4,3,7,8,5,6(p=4)
	//4,1,2,3,7,8,5,6
	//开始：idx=1,i==0,j=r+1=8,pivot=4(i:lEnd, j:rBegin)
	//4,1,2,3,7,8,5,6, 1:1,2:2,3:3交换, idx=4,i+3=3
	//4,1,2,3,6,8,5,7, 7<==>6交换，idx=4,j--=7
	//4,1,2,3,5,8,6,7, 6<==>5交换，idx=4,j--=6
	//4,1,2,3,8,5,6,7, 5<==>8交换，idx=4,j--=5
	//4,1,2,3,8,5,6,7, 8<==>8交换，idx=4,j--=4,结束
	//3,1,2,4,8,5,6,7, 4<==>3交换，pivot与nums[i]交换,得i=3,j=4

	var partition func(l, r int)
	partition = func(l, r int) {
		if l >= r {
			return
		}
		// 随机选取一个元素与第一个元素交换
		rand.Seed(time.Now().UnixNano())
		var randIdx = rand.Intn(r+1-l) + l
		nums[l], nums[randIdx] = nums[randIdx], nums[l]

		var (
			idx    = l + 1 // 选取基准元素后面的元素作为扫描点
			lEnd   = l     // 将扫描点的前一个元素作为小于部分的尾部，此时该部分是空的
			rBegin = r + 1 // 将数组末尾的元素作为大于部分的头部，此时该部分也是空的
		)

		// 遍历结束条件：当大于或等于大于部分的头部元素后结束
		for idx < rBegin {
			//当前元素小于基准元素时，将其与等于部分的头部元素交换,小于部分的尾部后移
			if nums[idx] < nums[l] {
				nums[idx], nums[lEnd+1] = nums[lEnd+1], nums[idx]
				lEnd++
			} else if nums[idx] > nums[l] {
				//当前元素大于基准元素时，将其于等于部分的尾部元素元素交换,大于部分的头部前移
				//注：此时遍历点不向后移，因为交换后的元素未处理过
				nums[idx], nums[rBegin-1] = nums[rBegin-1], nums[idx]
				rBegin--
				continue
			}
			idx++
		}
		nums[l], nums[lEnd] = nums[lEnd], nums[l]

		partition(l, lEnd-1)
		partition(rBegin, r)
	}

	partition(0, len(nums)-1)
	return nums
}

// TODO
func sortArray3(nums []int) []int {
	quick(&nums, 0, len(nums)-1)
	return nums
}

func quick(arr *[]int, i, j int) {
	if i >= j {
		return
	}

	var partition func(arr *[]int, i int, j int) int
	partition = func(arr *[]int, i int, j int) int {
		p := rand.Intn(j-i+1) + i // 随机选取“支点”
		nums := *arr
		nums[i], nums[p] = nums[p], nums[i]
		for i < j {
			// 修改原来的 nums[j] >= nums[i]，增加交换频率
			for nums[i] < nums[j] && i < j {
				j--
			}

			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}

			// 修改原来的 nums[j] >= nums[i]，增加交换频率
			for nums[i] < nums[j] && i < j {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				j--
			}
		}
		return i
	}

	mid := partition(arr, i, j)
	quick(arr, i, mid-1)
	quick(arr, mid+1, j)
}
