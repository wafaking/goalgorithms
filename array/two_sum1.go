package array

// 两数之和(leetcode-1)
// 给定整数数组nums和整数目标值target，请在数组中找出和为目标值的那两个整数的下标。
// 示例1： 输入：nums=[2,7,11,15], target=9, 输出：[0,1]
// 示例2： 输入：nums=[3,2,4], target=6，输出：[1,2]
// 示例3： 输入：nums=[3,3], target=6， 输出：[0,1]

func twoSum11(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func twoSum12(nums []int, target int) []int {
	var (
		res     []int
		visited = make(map[int]int) // map[值]索引
	)
	for k, v := range nums {
		if idx, ok := visited[target-v]; ok {
			res = append(res, idx, k)
			break
		}
		visited[v] = k
	}
	return res
}
