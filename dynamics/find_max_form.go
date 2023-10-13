package dynamics

import (
	"goalgorithms/common"
	"strings"
)

//一和零(leetcode-474)
//给你一个二进制字符串数组strs和两个整数m和n。
//请你找出并返回strs的最大子集的长度，该子集中最多有m个0和n个1。
//如果x的所有元素也是y的元素，集合x是集合y的子集。

//示例1：输入：strs=["10","0001","111001","1","0"],m=5,n=3,输出：4;
//解释：最多有5个0和3个1的最大子集是{"10","0001","1","0"}，因此答案是4。
//其他满足题意但较小的子集包括{"0001","1"}和{"10","1","0"}。{"111001"}不满足题意，因为它含4个1，大于n的值3。
//示例2：输入：strs=["10","0","1"],m=1,n=1; 输出：2;解释：最大的子集是{"0","1"}，所以答案是2。

// 动态规划+三维数组
func findMaxForm1(strs []string, m int, n int) int {
	// dp[i][j][k]表示选择前i个字符串，所能满足子集中最多j个0，k个1的最大子集长度;
	var dp = make([][][]int, len(strs)+1)
	for i := 0; i < len(strs)+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i < len(dp); i++ {
		cnt0 := strings.Count(strs[i-1], "0")
		cnt1 := len(strs[i-1]) - cnt0
		for j := 0; j < m+1; j++ { // j与k都从0开始，可能m=0或n=0
			for k := 0; k < n+1; k++ {
				if cnt0 <= j && cnt1 <= k {
					// 选择此字符串与不选的最大值
					// 不选：dp[i-1][j][k]
					// 选：dp[i-1][j-cnt0][k-cnt1]+1
					dp[i][j][k] = common.MaxInTwo(dp[i-1][j][k], dp[i-1][j-cnt0][k-cnt1]+1)
				} else { // 不满足条件,则不选此元素
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[len(dp)-1][m][n]
}
