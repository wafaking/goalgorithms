package tree

import (
	"log"
	"sort"
)

// func main() {
// 	log.Println(permute([]int{1, 2, 3}))
// }

// func permute(nums []int) [][]int {
// 	res := [][]int{}
// 	visited := map[int]bool{}

// 	var dfs func(path []int)
// 	dfs = func(path []int) {

// 		if len(path) == len(nums) {
// 			temp := make([]int, len(path))
// 			copy(temp, path)
// 			res = append(res, temp)
// 			return
// 		}

// 		for _, n := range nums {
// 			if visited[n] {
// 				continue
// 			}
// 			path = append(path, n)
// 			visited[n] = true
// 			dfs(path)
// 			path = path[:len(path)-1]
// 			visited[n] = false
// 		}
// 	}

// 	dfs([]int{})
// 	return res
// }

// -----------leetcode 46
var result [][]int

func permute(nums []int) [][]int {
	result = [][]int{}
	trace := []int{}

	backtrack(nums, trace)
	return result
}

func backtrack(nums []int, trace []int) {
	if len(trace) == len(nums) {
		len2 := len(nums)
		tmp := make([]int, len2)
		copy(tmp, trace)
		result = append(result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		var exist = false
		for _, v := range trace {
			if v == nums[i] {
				exist = true
				break
			}
		}
		if exist {
			continue
		}
		trace = append(trace, nums[i])
		backtrack(nums, trace)
		trace = trace[:len(trace)-1]
	}

}

// ------------leetcode 47
func PermuteUnique(nums []int) [][]int {
	var result [][]int
	visited := make([]bool, len(nums))
	path := make([]int, 0, 0)
	// 排序保证后续的剪枝
	sort.Ints(nums)
	dfs(&result, nums, visited, &path)
	return result
}
func dfs(result *[][]int, nums []int, visited []bool, path *[]int) {
	// path的长度 等于 nums元素的个数，将path复制到result中
	log.Println("-------------: ", *result, nums, visited, *path)

	if len(*path) == len(nums) {
		temp := make([]int, len(nums), len(nums))
		copy(temp, *path)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 如果nums的元素被标记为true，证明已经被使用了，需要跳过
		if visited[i] {
			log.Println("yes1, visited: ", i, visited)
			continue
		}
		// 剪枝，保证nums的中相同的元素按顺序组合，就可以避免重复
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
			log.Println("yes2, visited: ", i, visited)
			continue
		}
		// 添加至path，更改visited的bool值
		*path = append(*path, nums[i])
		visited[i] = true
		// 递归
		log.Println("====++++++", i, *path)
		dfs(result, nums, visited, path)
		log.Println("++++++++++")
		// 回退
		visited[i] = false
		*path = (*path)[:len(*path)-1]
	}
}
