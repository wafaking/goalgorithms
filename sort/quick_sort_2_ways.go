package sort

import (
	"math/rand"
	"time"
)

// QuickSort2Ways1 两路快排
func QuickSort2Ways1(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		partition func(l, r int)
		swap      = func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
	)

	// 左闭右闭
	partition = func(l, r int) {
		// 将数组分为左、右两部分
		//2, 1, 4, 3, 7, 8, 5, 6
		//(选7) 7,1,4,3,2,(8),5,(6)--->l=5:8,r=7:6位置交换得
		// 7,1,4,3,2,6,(5),8 -->l,r分别向右、左移[l=r=6,即数值5处]-->继续-->
		// l=r不满足条件结束 ---> r=6
		// 5,1,4,3,2,6,(7),8 (即7与5交换)

		// 7,1,4,3,2,6,(9),(5),8 [l=6,r=7,即数值9,5处]-->继续-->l,r分别向右、左移
		// 7,1,4,3,2,6,5,9,8 -->l=7,r=6,即数据9,5，结束-->r=6
		// 将r:5与7交换-->5,1,4,3,2,6,7,9,8

		if l >= r {
			return
		}

		// 随机索引
		rand.Seed(time.Now().UnixNano())
		// 将随机数值移到最l位置
		swap(l, rand.Intn(r+1-l)+l)

		// 将随机数与第一个元素l交换位置
		// l从1开始，位置1空出用于放置nums[randIdx]
		i, j := l+1, r
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

			//交换l与r
			nums[i], nums[j] = nums[j], nums[i]

			// l,r都各自归到左右部分，因此继续其余元素
			i++
			j--
		}

		//注：此处将l与j交换,因为for中用的i<=j,且i是从l+1开始的，因此j位置的才是边界;
		//示例如：1,3,2中选中pivot=1, 则需要l与1交换，此时j=0
		nums[l], nums[j] = nums[j], nums[l]
		partition(l, j-1)
		partition(j+1, r)
	}
	partition(0, len(nums)-1)
}

func QuickSort2Ways2(nums []int) {
	var (
		// 交换并找出划分子数组的位置
		partition      func(l, r int) int
		quickSort2Ways func(l, r int)
	)

	// 左闭右闭
	partition = func(l, r int) int {
		pivot := nums[l]
		for l < r {
			// 右指针左移
			for l < r && nums[r] >= pivot {
				r--
			}
			// 将nums[r]小于pivot的放到首位
			// 此时r位置的数值小于pivot
			nums[l] = nums[r]

			// 左指针右移
			for l < r && nums[l] <= pivot {
				l++
			}
			// 此时l位置的数值大于pivot

			// 将l位置较大的值放到r位置较大的值处
			nums[r] = nums[l]
			// 继续找出l位置，将其与pivot交换(l位置元素一直是个临时值)
			//此时r位置的元素一定大于l处，因此for循环只需l<r即可
		}
		nums[l] = pivot
		return l
	}

	quickSort2Ways = func(l, r int) {
		if l < r {
			pivot := partition(l, r)
			quickSort2Ways(0, pivot-1)
			quickSort2Ways(pivot+1, r)
		}
	}
	quickSort2Ways(0, len(nums)-1)
}
