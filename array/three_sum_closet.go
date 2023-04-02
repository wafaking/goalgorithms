package array

import (
	"sort"
)

// 最接近的三数之和（leetcode-16）
// 给定长度为n的整数数组nums和目标值target。请从nums中选出三个整数，使它们的和与target最接近。
// 返回这三个数的和。 假定每组输入只存在恰好一个解。
//示例1： 输入：nums=[-1,2,1,-4], target=1, 输出：2(与1最接近的和是 2 (-1 + 2 + 1 = 2)) 。
//示例2： 输入：nums = [0,0,0], target = 1, 输出：0

// threeSumClosest1 排序+三次遍历
func threeSumClosest1(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)
	// nums=-5,-5,-4,0,0,3,3,4,5, target=-1
	var getDistance func(a, b int) int
	getDistance = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	var (
		initDistance = getDistance(nums[0]+nums[1]+nums[2], target)
		ans          = nums[0] + nums[1] + nums[2]
	)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				var temp = getDistance(nums[i]+nums[j]+nums[k], target)
				if temp < initDistance {
					ans = nums[i] + nums[j] + nums[k]
					initDistance = temp
					//log.Println("nums: ", nums[i], nums[j], nums[k])
				}
			}
		}
	}
	return ans
}

// threeSumClosest2 排序+双指针
func threeSumClosest2(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)

	var (
		ans      = nums[0] + nums[1] + nums[2]
		distance = ans - target
	)

	if distance < 0 {
		distance *= -1
	}

	// nums=-5,-5,-4,0,0,3,3,4,5, target=-1
	// --- -2 --- -1 --- 0 --- 1 --- 2 --- 3 --- 4 ---> x-axis
	for i := 0; i < len(nums); i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum > target { // sum is to left of the target on x-axis
				if (sum - target) < distance {
					distance = sum - target
					ans = sum
				}
				k--
			} else {
				if (target - sum) < distance {
					distance = target - sum
					ans = sum
				}
				j++
			}
		}
	}
	return ans
}
