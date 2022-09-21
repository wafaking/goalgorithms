package list

import "sort"

// 含有重复元素集合的组合(sword2-82)
// 给定一个可能有重复数字的整数数组candidates和目标数target，找出数组中所有可以使数字和为target的组合。
// 数组中的数字在每个组合中只能使用一次，解集不能包含重复的组合。
// 示例1: 输入: candidates=[10,1,2,7,6,1,5], target=8, 输出: [[1,1,6],[1,2,5],[1,7],[2,6]]
// 示例2: 输入: candidates=[2,5,2,1,2], target=5, 输出: [[1,2,2],[5]]

// todo -------------
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	// 2,2,2,4,6 == 8
	// freq = [2, 1]
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
