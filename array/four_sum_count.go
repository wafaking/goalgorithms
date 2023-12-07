package array

// 四数相加II(leetcode-454)
// 给定四个整数数组nums1、nums2、nums3和nums4,数组长度都是n，
// 计算使得nums1[i]+nums2[j]+nums3[k]+nums4[l]==0的所有i,j,k,l组合总数；
//示例1： 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
//输出：2，如：
//	nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
//  nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
//示例2： 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0], 输出：1

// fourSumCount11 考虑数组内数值重复的情况
func fourSumCount11(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var (
		f   func(numsA, numsB []int) map[int]int
		ans int
	)
	f = func(numsA, numsB []int) map[int]int {
		var (
			m       = make(map[int]int, 0)
			usedMap = make(map[[2]int]struct{})
		)

		for i, v1 := range numsA {
			for j, v2 := range numsB {
				if _, ok := usedMap[[2]int{i, j}]; ok {
					continue
				}
				usedMap[[2]int{i, j}] = struct{}{}
				m[v1+v2]++
			}
		}
		return m
	}

	var (
		mapA = f(nums1, nums2)
		mapB = f(nums3, nums4)
	)

	for i, listA := range mapA {
		if listB, ok := mapB[-1*i]; ok {
			ans += listA * listB
		}
	}

	return ans
}

// fourSumCount12 使用相反数（不考虑数组内有重复数值的情况，因本题中的i,j,k,l为数组下标）
func fourSumCount12(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var (
		ans int
		m   = make(map[int]int)
	)

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}

	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			ans += m[-1*(v1+v2)]
		}
	}
	return ans
}
