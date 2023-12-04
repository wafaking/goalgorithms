package array

import "sort"

//有序数组的平方(leetcode-977)
//给你一个按非递减顺序排序的整数数组nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
//示例1：输入：nums=[-4,-1,0,3,10]输出：[0,1,9,16,100]
//解释：平方后，数组变为[16,1,0,9,100]排序后，数组变为[0,1,9,16,100]
//示例2：输入：nums=[-7,-3,2,3,11]输出：[4,9,9,49,121]

func sortedSquares1(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// 双指针：时间复杂度为O(n)
func sortedSquares2(nums []int) []int {
	var (
		n   = len(nums)
		ans = make([]int, 0, n)
	)
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	for l, r := 0, len(nums)-1; l <= r; {
		if nums[l] > nums[r] {
			ans = append(ans, nums[l])
			l++
		} else {
			ans = append(ans, nums[r])
			r--
		}
	}
	//[-4,-1,-1,0,3,3,10]
	//16, 1, 1, 0, 9,9, 100
	for i := 0; i <= (n-1)>>1; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}

	return ans
}

// 双指针：从两头到中间，时间复杂度为O(n)
func sortedSquares3(nums []int) []int {
	var (
		n   = len(nums)
		ans = make([]int, len(nums))
	)
	for l, r := 0, n-1; l <= r; {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			ans[n-1] = nums[l] * nums[l]
			l++
		} else {
			ans[n-1] = nums[r] * nums[r]
			r--
		}
		n--
	}
	return ans
}

// 双指针，从中间到两端，时间复杂度为O(n)
func sortedSquares4(nums []int) []int {
	//[-4,-1,-1,0,3,3,10]
	// 16, 1, 1, 0, 9,9, 100
	var (
		pivot = -1
		n     = len(nums)
		ans   = make([]int, 0, n)
	)

	// 找出正负值的分界点
	for i := 0; i < n && nums[i] < 0; i++ {
		pivot++
	}
	for l, r := pivot, pivot+1; l >= 0 || r < n; {
		if l < 0 { // 将右侧都添加进去
			ans = append(ans, nums[r]*nums[r])
			r++
		} else if r >= n { // 将左侧的都添加进去
			ans = append(ans, nums[l]*nums[l])
			l--
		} else if nums[l]*nums[l] > nums[r]*nums[r] {
			ans = append(ans, nums[r]*nums[r])
			r++
		} else {
			ans = append(ans, nums[l]*nums[l])
			l--
		}
	}
	return ans
}
