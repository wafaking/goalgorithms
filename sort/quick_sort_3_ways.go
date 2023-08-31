package sort

import (
	"math/rand"
	"time"
)

// QuickSort3Ways 三路快排
func QuickSort3Ways(nums []int) {
	var partition func(l, r int)
	//2,1,4,3,7,8,5,6(p=4)
	//4,1,2,3,7,8,5,6
	//开始：idx=1,i==0,j=r+1=8,pivot=4(i:lEnd, j:rBegin)
	//4,1,2,3,7,8,5,6, 1:1,2:2,3:3交换, idx=4,i+3=3
	//4,1,2,3,6,8,5,7, 7<==>6交换，idx=4,j--=7
	//4,1,2,3,5,8,6,7, 6<==>5交换，idx=4,j--=6
	//4,1,2,3,8,5,6,7, 5<==>8交换，idx=4,j--=5
	//4,1,2,3,8,5,6,7, 8<==>8交换，idx=4,j--=4,结束
	//3,1,2,4,8,5,6,7, 4<==>3交换，pivot与nums[i]交换,得i=3,j=4

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
}
