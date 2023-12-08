package array

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"goalgorithms/common"
	"sort"
)

// 含有重复元素集合的组合(sword2-82/leetcode-40)
// 给定一个可能有重复数字的整数数组candidates和目标数target，找出数组中所有可以使数字和为target的组合。
// 数组中的数字在每个组合中只能使用一次，解集不能包含重复的组合。
// 示例1: 输入: candidates=[10,1,2,7,6,1,5], target=8, 输出: [[1,1,6],[1,2,5],[1,7],[2,6]]
// 示例2: 输入: candidates=[2,5,2,1,2], target=5, 输出: [[1,2,2],[5]]

// combinationSum21 法一：借助数字的全排列的思想并通过序列化去重
func combinationSum21(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var (
		dfs     func(nums []int, path []int, sum int)
		visited = make(map[string]struct{})
	)

	dfs = func(nums []int, path []int, sum int) {
		if sum == target {
			var temp = make([]int, len(path))
			copy(temp, path)
			sort.Ints(temp) // 排序
			tempByte, err := json.Marshal(temp)
			if err != nil {
				panic(err)
			}
			tempStr := fmt.Sprintf("%x", md5.Sum(tempByte))
			if _, ok := visited[tempStr]; !ok {
				ans = append(ans, temp)
				visited[tempStr] = struct{}{}
			}
			return
		}

		var m = make(map[int]struct{}, len(nums))
		for i := 0; i < len(nums); i++ {
			if _, ok := m[nums[i]]; ok {
				continue
			}
			m[nums[i]] = struct{}{}

			// 记录元素
			var cur = nums[i]
			sum += cur
			path = append(path, cur)
			nums = append(nums[:i], nums[i+1:]...)
			// 递归
			dfs(nums, path, sum)

			// 回溯
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			sum -= cur
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, []int{}, 0)
	return ans
}

// 借助全排列思想（推荐此方法）
func combinationSum22(candidates []int, target int) [][]int {
	var (
		ans = make([][]int, 0)
		dfs func(idx int, path []int)
		n   = len(candidates)
	)
	dfs = func(idx int, path []int) {
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		// 每一层不能重复
		// 记录每层已遍历过的数值
		var m = make(map[int]struct{})
		for i := idx; i < n; i++ {
			if _, ok := m[candidates[i]]; ok {
				continue
			}
			m[candidates[i]] = struct{}{}
			target -= candidates[i]
			path = append(path, candidates[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
			target += candidates[i]
			for i+1 < n && candidates[i] == candidates[i+1] {
				i++
			}
		}
	}
	sort.Ints(candidates)
	dfs(0, []int{})
	return ans
}

// combinationSum23 借助全排列思想(推荐,效率高)
func combinationSum23(candidates []int, target int) (ans [][]int) {

	var (
		dfs  func(nums []int, sum int)
		path []int
	)
	dfs = func(nums []int, sum int) {
		// len(path)>0 用于过滤target=0的情况
		if sum == target && len(path) > 0 {
			ans = append(ans, append([]int{}, path...))
			// 如果考虑全为数组全为0，或数组内有0的情况，可以去掉return
			return
		} else if target < sum {
			return
		}

		// 记录已遍历过的数值
		for i := 0; i < len(nums); i++ {
			var cur = nums[i]
			sum += cur
			path = append(path, cur)
			// 此处与全排列不一样，仅使用nums[i+1:]，不再使用cur前面的数字
			// 不再遍历
			dfs(nums[i+1:], sum)
			sum -= cur
			path = path[:len(path)-1]

			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
				continue
			}
		}
	}
	sort.Ints(candidates) // 先排序
	dfs(candidates, 0)
	return ans
}

// combinationSum24 todo（没理解太明白）
func combinationSum24(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	// candidates:2,2,2,4,6  target: 8
	// 用于记录每个数字出现的次数
	var freq [][2]int
	// freq = [ [2, 4], [4, 1], [6, 1] ]
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var (
		sequence []int
		dfs      func(pos, rest int)
	)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		// 即rest还能对当前的数可以再累加多少次
		most := common.MinInTwo(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}
