package array

// 组合(leetcode-77)
// 给定两个整数n和k，返回范围[1,n]中所有可能的k个数的组合。
// 示例1： 输入：n=4, k=2,输出：[2,4], [3,4], [2,3], [1,2], [1,3], [1,4]]
// 示例2： 输入：n=1, k=1,输出：[[1]]

// combine1 全排列
func combine1(n int, k int) [][]int {
	var (
		ans  [][]int
		path = make([]int, 0, k)
		dfs  func(idx int)
	)
	dfs = func(idx int) {
		// 剪枝，当temp的长度+剩余元素([idx,n])的数量<k时，直接结束
		// 如：n=4,k=5，不可能有答案
		if len(path)+(n-idx+1) < k {
			return
		}
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := idx; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}

	}
	dfs(1)
	return ans
}

// combine2 二进制子集枚举
func combine2(n int, k int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	for mask := 0; mask < 1<<n; mask++ {
		path = nil
		for i := 0; i < n; i++ {
			if 1<<i&mask > 0 {
				path = append(path, i+1)
			}
		}
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
		}
	}

	return ans
}

// combine3 使用二进制的思想，参数官方思路（todo-----）
func combine3(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	var temp []int
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}