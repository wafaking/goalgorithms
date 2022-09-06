package list

import (
	"sort"
)

//15. 三数之和(leetcode-15)
//给你一个包含n个整数的数组nums，判断nums中是否存在三个元素a，b，c ，使得a+b+c=0？
//示例1： 输入：nums = [-1,0,1,2,-1,-4], 输出：[[-1,-1,2],[-1,0,1]]
//示例2： 输入：nums = [], 输出：[]
//示例3： 输入：nums = [0], 输出：[]

// threeSum1，排序后拆分成a=-(b+c),a固定，b+c用双指针确定
func threeSum1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	// 先枚举a
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // a与前一个元素不重复
			continue
		}
		// 使用双指针j,k
		for j, k := i+1, len(nums)-1; j < k; {
			var temp = nums[i] + nums[j] + nums[k]
			if temp > 0 {
				k-- // 右指针前移
			} else if temp < 0 {
				j++ // 左指针后移
			} else {
				// -4, -1, -1, -1, -1, 0, 1, 2, 2
				res = append(res, []int{nums[i], nums[j], nums[k]})
				k-- // 右指针前移
				// 因双指针两个元素已使用过，也可同时右指针前移并左指针后移
				for j < k && nums[k] == nums[k+1] { // 右指针元素与后面的不重复
					k--
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums) // 排序保证a<=b<=c
	res := make([][]int, 0)

	// 枚举第一个指针
	for a := 0; a < len(nums); a++ {
		// 保证枚举的和上一次的不同，此处要注意a要从第二个数开始
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		// 枚举第二个指针
		for b := a + 1; b < len(nums); b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			target := -1 * nums[a] // 取对等反向值
			c := len(nums) - 1     // c 对应的指针初始指向数组的最右端

			for b < c && nums[b]+nums[c] > target { // 大了，将第三个数的指针向前移
				c--
			}

			if b == c { // 防止c跑到b的前面
				break
			}

			// 注：此处排序后如果相等，b++向后移，会忽略相等值
			if nums[b]+nums[c] == target { // 只有相等了再添加
				res = append(res, []int{nums[a], nums[b], nums[c]})
			}
		}
	}

	return res
}
