package array

// 组合总和III(leetcode-216)
// 找出所有相加之和为n的k个数的组合，且满足下列条件：只使用数字1到9, 每个数字最多使用一次
// 返回所有可能的有效组合的列表。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
// 示例1: 输入: k = 3, n = 7 输出: [[1,2,4]] 解释: 1 + 2 + 4 = 7
// 示例2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]
// 示例3: 输入: k = 4, n = 1, 输出: []

func combinationSum31(k int, n int) [][]int {
	var (
		ans [][]int
		min = (k + 1) * (k / 2)
		max = (19 - k) * (k / 2) // max=(9+(9-k+1))*k/2
	)

	if k&1 != 0 {
		min += (k + 1) / 2
		max += (19 - k) / 2
	}
	if n < min || n > max {
		return ans
	}

	var (
		dfs  func(min int)
		path []int
	)
	dfs = func(min int) {
		if n == 0 && k == 0 {
			ans = append(ans, append([]int(nil), path...))
		} else if n < 0 || k <= 0 {
			return
		}

		for i := min; i < 10; i++ {
			k--
			n -= i
			path = append(path, i)
			dfs(i + 1)
			k++
			n += i
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}

// 子集枚举(枚举所有子集，将sum=n的添加到结果集中)
// 组合中只允许含有1-9的正整数，并且每种组合中不存在重复的数字,意味着这个组合中最多包含9个数字。我们可以把原问题转化成集合S={1,2,3,4,5,6,7,8,9}，我们要找出S的当中满足如下条件的子集：
//	1. 大小为k
//	2.集合中元素的和为n
// 因此我们可以用子集枚举的方法来做这道题。即原序列中有9个数，每个数都有两种状态，「被选择到子集中」和「不被选择到子集中」，所以状态的总数为 2^9。
// 我们用一个9位二进制数mask来记录当前所有位置的状态，从低到高第i位为0表示i不被选择到子集中，为1表示i被选择到子集中。当我们按顺序枚举 [0, 2^9-1]
// 中的所有整数的时候，就可以不重不漏地把每个状态枚举到，对于一个状态mask，我们可以用位运算的方法得到对应的子集序列，然后再判断是否满足上面的两个条件，如果满足，就记录答案。
// 如何通过位运算来得到mask各个位置的信息？对于第i个位置我们可以判断(1 << i) & mask 是否为0，如果不为0则说明i在子集当中。当然，这里要注意的是，一个9位二进制数i的范围是[0, 8],
// 而可选择的数字是[1,9]，所以我们需要做一个映射，最简单的办法就是当我们知道i位置不为0的时候将i+1加入子集。
// combinationSum32 子集枚举(枚举所有子集，将sum=n的添加到结果集中, 参考leetcode-77/78)
func combinationSum32(k int, n int) [][]int {
	var (
		ans   [][]int
		temp  []int
		check func(mask int) bool
	)
	check = func(mask int) bool {
		temp = nil
		sum := 0
		for i := 0; i < 9; i++ { // 注：此处取值应为[0:8]
			if 1<<i&mask > 0 {
				// 添加的值应为[1,9]
				temp = append(temp, i+1)
				sum += i + 1
			}
		}
		return len(temp) == k && sum == n
	}

	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			ans = append(ans, append([]int(nil), temp...))
		}
	}
	return ans
}
