package dynamics

// (多重背包问题)有n件物品和一个最多能背重量为w的背包。第i件物品的重量是weight[i]，得到的价值是value[i]。
// 每件物品只能用times[i]次，求解将哪些物品装入背包里物品价值总和最大。

// 动态规划，将每个物品可以使用的次数分拆，视为最多可以一个个不同的物品，每个物品仅可以使用1次
func weightBag21(weight, value, times []int, bagWeight int) int {
	// 物品，分别是数量，体积和单位价格
	//items = [(10, 3, 5), (5, 6, 3), (2, 2, 4)]
	//volume = 20
	// 1. 计算weight及value对应的长度
	var length int
	for _, v := range times {
		length += v
	}
	// 2. 生成新的weight和value对照表
	var (
		newWeight = make([]int, 0, length)
		newValue  = make([]int, 0, length)
	)
	for i := 0; i < len(weight); i++ {
		for j := 0; j < times[i]; j++ {
			newWeight = append(newWeight, weight[i])
			newValue = append(newValue, value[i])
		}
	}
	// 3.按照0/1背包问题求解(使用滚动数组)
	var (
		pre = make([]int, bagWeight+1)
		cur = make([]int, bagWeight+1)
	)

	for i := 1; i < length+1; i++ {
		for j := 1; j < bagWeight+1; j++ {
			if j < newWeight[i-1] { // 大于背包重量，装不下
				cur[j] = pre[j]
			} else {
				cur[j] = MaxInTwo(pre[j], newValue[i-1]+pre[j-newWeight[i-1]])
			}
		}
		copy(pre, cur)
	}

	return cur[len(cur)-1]
}

// weightBag22 TODO 动态规划，使用二进制方法
func weightBag22(weight, value, times []int, bagWeight int) int {
	return 0
}
