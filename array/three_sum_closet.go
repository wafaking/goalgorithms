package array

import (
	"math"
	"sort"
)

//最接近的三数之和（leetcode-16）
//给定长度为n的整数数组nums和目标值target。请从nums中选出三个整数，使它们的和与target最接近。
//返回这三个数的和。假定每组输入只存在恰好一个解。
//示例1：输入：nums=[-1,2,1,-4],target=1,输出：2(与1最接近的和是2(-1+2+1=2))。
//示例2：输入：nums=[0,0,0],target=1,输出：0

// 排序+三次遍历
func threeSumClosest1(nums []int, target int) int {
	var (
		n           = len(nums)
		ans         = -1
		distance    = math.MaxInt64 // 注：题中-1000<nums[i]<1000
		getDistance = func(num1, num2 int) int {
			if num1 > num2 {
				return num1 - num2
			}
			return num2 - num1
		}
	)
	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				sum := nums[i] + nums[j] + nums[k]
				dis := getDistance(sum, target)
				if distance > dis {
					ans = sum
					distance = dis
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
