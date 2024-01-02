package dynamics

import (
	"goalgorithms/common"
	"math"
)

//地下城游戏(leetcode-174)
//恶魔们抓住了公主并将她关在了地下城dungeon的右下角。地下城是由mxn个房间组成的二维网格。
//我们英勇的骑士最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
//骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至0或以下，他会立即死亡。
//有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；
//其他房间要么是空的（房间里的值为0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
//为了尽快解救公主，骑士决定每次只向右或向下移动一步。
//返回确保骑士能够拯救到公主所需的最低初始健康点数。
//注意：任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
//示例1：输入：dungeon=[[-2,-3,3],[-5,-10,1],[10,30,-5]]输出：7
//解释：如果骑士遵循最佳路径：右->右->下->下，则骑士的初始健康点数至少为7。
//示例2：输入：dungeon=[[0]]输出：1

//动态规划(从右下角向左上角遍历)
//如果按照从左上往右下的顺序进行动态规划，我们无法直接确定到达(1,2)(1,2)(1,2)的方案，
//因为有两个重要程度相同的参数同时影响后续的决策。也就是说，这样的动态规划是不满足「无后效性」的。
func calculateMinimumHP1(dungeon [][]int) int {
	var (
		m, n = len(dungeon), len(dungeon[0])
		//dp[i][j]表示到达dungeon位置需要的最小初始值
		dp = make([][]int, m+1)
	)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			// 初始值默认最大
			dp[i][j] = math.MaxInt32
		}
	}

	// 初始值
	dp[m][n-1] = 1
	dp[m-1][n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//选择当前位置的下和右所需的最小的值
			min := common.MinInTwo(dp[i+1][j], dp[i][j+1])
			//当前位置需要的值与1取较大的
			dp[i][j] = common.MaxInTwo(min-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

// 动态规划(从左上向右下遍历)(TODO wrong)
func calculateMinimumHP2(dungeon [][]int) int {
	var (
		// 由题已知m>0,n>0
		m = len(dungeon)
		n = len(dungeon[0])
		//
		dp = make([]int, m*n)
	)

	// 初始值默认是1
	dp[0] = -1*dungeon[0][0] + 1
	for i := 1; i < n; i++ { // 第一行
		cur := dungeon[0][i]
		dungeon[0][i] += dungeon[0][i-1]

		// cur>=0,直接取前面的最大值
		if cur > 0 {
			dp[i] = dp[i-1]
			continue
		}
		// cur<0，取决于sum<0
		if dungeon[0][i] < 0 {
			dp[i] = dp[i-1] + (-1 * cur)
		} else {
			dp[i] = dp[i-1]
		}
	}

	for i := 1; i < m; i++ { // 第一列
		cur := dungeon[i][0]
		dungeon[i][0] += dungeon[i-1][0]
		if cur > 0 {
			dp[i*n] = dp[(i-1)*n]
			continue
		}

		if dungeon[i][0] >= 0 { // sum>0
			dp[i*n] = dp[(i-1)*n]
		} else {
			dp[i*n] = dp[(i-1)*n] + (-1 * cur)
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			cur := dungeon[i][j]

			if dp[i*n+j-1] < dp[(i-1)*n+j] { // 选左
				dungeon[i][j] += dungeon[i][j-1]
				if cur > 0 {
					dp[i*n+j] = dp[i*n+j-1]
				} else {
					if dungeon[i][j] >= 0 { // sum>0
						dp[i*n+j] = dp[i*n+j-1]
					} else {
						//dp[i*n+j] = dp[i*n+j-1] + (-1 * cur)
						dp[i*n+j] = common.MinInTwo(dp[i*n+j-1]+(-1*cur), -1*dungeon[i][j]+1)
					}
				}

			} else { // 选上
				dungeon[i][j] += dungeon[i-1][j]
				if cur > 0 {
					dp[i*n+j] = dp[(i-1)*n+j]
				} else {
					if dungeon[i][j] > 0 {
						dp[i*n+j] = dp[(i-1)*n+j]
					} else {
						//dp[i*n+j] = dp[(i-1)*n+j] + (-1 * cur)
						dp[i*n+j] = common.MinInTwo(dp[(i-1)*n+j]+(-1*cur), -1*dungeon[i][j]+1)
					}
				}
			}
		}
	}
	return dp[m*n-1]
}
