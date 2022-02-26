package list

import (
	"log"
	"sort"
)

//15. 三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
//请你找出所有和为 0 且不重复的三元组。

//注意：答案中不可以包含重复的三元组。

//示例 1：
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]

//示例 2：
//输入：nums = []
//输出：[]

//示例 3：
//输入：nums = [0]
//输出：[]

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序保证a<=b<=c
	//log.Println("nums: ", nums)
	length := len(nums)
	res := make([][]int, 0)

	// 枚举第一个指针
	for first := 0; first < length; first++ {
		// 保证枚举的和上一次的不同，此处要注意a要从第二个数开始
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		// 枚举第二个指针
		for second := first + 1; second < length; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			target := -1 * nums[first] // 取对等反向值
			third := length - 1        // c 对应的指针初始指向数组的最右端

			for second < third && nums[second]+nums[third] > target { // 大了，将第三个数的指针向前移
				//log.Println("second: ******: ", nums[second], nums[third], target)
				third--
			}

			if second == third { // 防止c跑到b的前面
				break
			}

			// 注：此处排序后如果相等，b++向后移，会忽略相等值
			if nums[second]+nums[third] == target { // 只有相等了再添加
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	log.Println("nums: ", nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			log.Println("*****second*****: ", nums[second], nums[third])
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
