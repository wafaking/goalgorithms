package list

import "sort"

// 四数之和(leetcode-18)
// 给定数组nums和目标值target。请找出其中的四个数使得nums[a]+nums[b]+nums[c]+nums[d]==target
// 示例1： 输入：nums=[1,0,-1,0,-2,2], target=0, 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 示例2： 输入：nums = [2,2,2,2,2], target=8, 输出：[[2,2,2,2]]

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] { // a与上一个a不重复
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ { // b从a的后面一个位置开始
			if j > i+1 && nums[j] == nums[j-1] { // b与上一个b不重复
				continue
			}
			for k, l := j+1, len(nums)-1; k < l; { // 双指针c,d: c从b后面第一位开始，d从最后一位开始
				var temp = nums[i] + nums[j] + nums[k] + nums[l]
				if temp > target { // d向前移
					l--
				} else if temp < target { // c向后移
					k++
				} else { // equal
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					k++                                            // 让c向后移一位
					for k > j+1 && k < l && nums[k-1] == nums[k] { // 将c后移至与前一位不重复
						k++
					}
					for k < l-1 && nums[l] == nums[l-1] { // 将d前移至与后一位不重复
						l--
					}
				}
			}
		}
	}

	return res
}
