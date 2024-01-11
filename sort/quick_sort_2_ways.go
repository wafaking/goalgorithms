package sort

import (
	"math/rand"
	"time"
)

// QuickSort2Ways 两路快排
func QuickSort2Ways(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	var (
		partition func(l, r int)
		swap      = func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		}
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

		// TODO
		if l >= r {
			return
		}

		// 随机索引
		rand.Seed(time.Now().UnixNano())
		swap(l, rand.Intn(r+1-l)+l)

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
}
