package sort

import "goalgorithms/common"

// 递归
func mergeSort1(nums []int) []int {
	var (
		n = len(nums)
		// 用于盛放[l,r]之间排序后的元素
		ans = make([]int, n)
		// 用于将l,r位置中的分成两部分,并分别排序，放置到ans[l,r]中并排序
		mergeSortCore func(l, r int)
		// 合并nums数组以mid为分隔点的两个非递减数列
		merge func(nums []int, l, mid, r int)
	)

	// l, r范围：[l,r]
	mergeSortCore = func(l, r int) {
		if l > r {
			return
		} else if l == r { // 说明只有一个元素,直接放置到ans[l]中即可
			ans[l] = nums[l]
		} else { // 有多个元素

			//找到中间元素位置，继续拆分
			mid := (r + l) / 2

			// 将[l,mid]放入到ans中
			mergeSortCore(l, mid)

			// 将[mid+1,r]放入到ans中
			mergeSortCore(mid+1, r)

			// 将[l,mid]和[mid+1,r]两个递增数组合并并排列
			merge(ans, l, mid, r)
		}
	}

	merge = func(nums []int, l, mid, r int) {
		for i, j := mid, r; i >= l && j > i; {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				for k := i; k > l; k-- {
					if nums[k] >= nums[k-1] {
						break
					}
					nums[k], nums[k-1] = nums[k-1], nums[k]
				}
			}
			j--
		}
	}

	mergeSortCore(0, n-1)
	return ans
}

//迭代法
//① 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
//② 设定两个指针，最初位置分别为两个已经排序序列的起始位置
//③ 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
//④ 重复步骤③直到某一指针到达序列尾
//⑤ 将另一序列剩下的所有元素直接复制到合并序列尾
func mergeSort2(nums []int) []int {
	var n = len(nums)
	for seg := 1; seg < n; seg *= 2 { //以步长划分元素,1,2,4,8,16
		for start := 0; start < n; start += seg * 2 { // 将步长内的元素排序
			l := start
			mid := common.MinInTwo(start+seg, n)
			r := common.MinInTwo(start+seg*2, n)

			// l与r为左闭右开[l,r)
			// 将[l,r)之间的元素排序
			for i, j := mid-1, r-1; i >= l && j > i; {
				if nums[i] > nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					for k := i; k > l; k-- {
						if nums[k] >= nums[k-1] {
							break
						}
						nums[k], nums[k-1] = nums[k-1], nums[k]
					}
				}
				j--
			}

		}
	}
	return nums
}
