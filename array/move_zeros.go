package array

//移动零(leetcode-283)
//给定一个数组nums，编写一个函数将所有0移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意，必须在不复制数组的情况下原地对数组进行操作。
//示例1:输入:nums=[0,1,0,3,12]输出:[1,3,12,0,0]
//示例2:输入:nums=[0]输出:[0]
//提示:
//	1<=nums.length<=104
//	-231<=nums[i]<=231-1

func moveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					if j == len(nums) {
						return
					}
					break
				}
				j++
				if j >= len(nums) {
					return
				}
			}
		}
	}
}

func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
