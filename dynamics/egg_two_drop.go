package dynamics

import "goalgorithms/common"

//鸡蛋掉落-两枚鸡蛋(leetcode-1884)
//给你2枚相同的鸡蛋和一栋从第1层到第n层共有n层楼的建筑。
//已知存在楼层f，满足0<=f<=n，任何从高于f的楼层落下的鸡蛋都会碎，从f楼层或比它低的楼层落下的鸡蛋都不会碎。
//每次你可以取一枚没有碎的鸡蛋并把它从楼层x扔下（1<=x<=n）:
//	1.如果鸡蛋碎了，你就不能再次使用它。
//	2.如果鸡蛋没有摔碎，则可以重复使用。
//请你计算并返回要确定f确切的值的最小操作次数是多少？

//示例1：输入：n=2,输出：2;解释：
//  我们可以将第一枚鸡蛋从1楼扔下，然后将第二枚从2楼扔下。
//	如果第一枚鸡蛋碎了，可知f=0；
//	如果第二枚鸡蛋碎了，但第一枚没碎，可知f=1；
//	否则，当两个鸡蛋都没碎时，可知f=2。

//示例2：输入：n=100,输出：14,解释：
//	一种最优的策略是：
//	-将第一枚鸡蛋从9楼扔下。如果碎了，那么f在0和8之间。将第二枚从1楼扔下，然后每扔一次上一层楼，在8次内找到f。总操作次数=1+8=9。
//	-如果第一枚鸡蛋没有碎，那么再把第一枚鸡蛋从22层扔下。如果碎了，那么f在9和21之间。将第二枚鸡蛋从10楼扔下，然后每扔一次上一层楼，在12次内找到f。总操作次数=2+12=14。
//	-如果第一枚鸡蛋没有再次碎掉，则按照类似的方法从34,45,55,64,72,79,85,90,94,97,99和100楼分别扔下第一枚鸡蛋。
//	不管结果如何，最多需要扔14次来确定f。

// 动态规划+二维数组(O(n^3)容易超时)
func twoEggDrop1(n int) int {
	// dp[i][j]表示剩余i个鸡蛋在总层数为j时需要的最小操作数
	var dp = make([][]int, 3)
	for i := 0; i <= 2; i++ {
		dp[i] = make([]int, n+1)
	}

	if n == 0 {
		return 0
	}

	// 若只有1个鸡蛋，则n层楼最少需要n次操作
	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}

	for i := 2; i < 2+1; i++ {
		for j := 1; j <= n; j++ {
			if i > j {
				// 如果剩余鸡蛋数比楼层数多，则说明不需要这么多鸡蛋，
				// 至少少一个鸡蛋的情况前面已经处理过，直接取dp[i-1][j]
				dp[i][j] = dp[i-1][j]
				continue
			}

			//最小操作次数一定是小于楼层数的，因此初始化dp[i][j]为j
			dp[i][j] = j
			// 可以有j种方法，即从1,2,3...k
			for m := 1; m <= j; m++ {
				// 分鸡蛋碎了和没碎两种情况
				//1.碎了(鸡蛋数-1，楼层数-1),dp[i][j] = dp[i-1][j-1]
				//2.没碎(鸡蛋数不变，楼层数-m),dp[i][j] = dp[i][j-m]
				dp[i][j] = common.MinInTwo(1+common.MaxInTwo(dp[i-1][m-1], dp[i][j-m]), dp[i][j])
			}
		}
	}

	return dp[2][n]
}

// 动态规划+二维数组(二分查找)
func twoEggDrop2(n int) int {
	// dp[i][j]表示剩余i个鸡蛋在总层数为j时需要的最小操作数
	var dp = make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
	}

	if n == 0 {
		return 0
	}

	// 若只有1个鸡蛋，则n层楼最少需要n次操作
	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}

	for i := 2; i <= 2; i++ {
		for j := 1; j <= n; j++ {
			if i > j {
				// 如果剩余鸡蛋数比楼层数多，则说明不需要这么多鸡蛋，
				// 至少少一个鸡蛋的情况前面已经处理过，直接取dp[i-1][j]
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = j

			var l, r = 1, j
			for l < r {
				mid := (r-l)>>1 + l
				// 碎了与没碎的操作次数比较
				if dp[i-1][mid-1] < dp[i][j-mid] { //碎了的数小,则向右移
					l = mid + 1
				} else { //否则向左移
					r = mid
				}
			}

			// 分鸡蛋碎了和没碎两种情况
			//1.碎了(鸡蛋数-1，楼层数-1),dp[i][j] = dp[i-1][j-1]
			//2.没碎(鸡蛋数不变，楼层数-m),dp[i][j] = dp[i][j-m]
			dp[i][j] = common.MinInTwo(1+common.MaxInTwo(dp[i-1][l-1], dp[i][j-l]), dp[i][j])
		}
	}

	return dp[2][n]
}

// 动态规划(二分查找+滚动数组)
func twoEggDrop3(n int) int {

	if n == 0 {
		return 0
	}

	// cur[i]表示i层时需要的最小操作数
	var odd, even = make([]int, n+1), make([]int, n+1)

	// 若只有1个鸡蛋，则n层楼最少需要n次操作
	for i := 1; i <= n; i++ {
		odd[i] = i
	}

	for i := 2; i <= 2; i++ {
		for j := 1; j <= n; j++ {
			if i > j {
				// 如果剩余鸡蛋数比楼层数多，则说明不需要这么多鸡蛋，
				// 至少少一个鸡蛋的情况前面已经处理过，直接取dp[i-1][j]
				even[j] = odd[j]
				continue
			}

			even[j] = j

			var l, r = 1, j
			for l < r {
				mid := (r-l)>>1 + l
				// 碎了与没碎的操作次数比较
				if odd[mid-1] < even[j-mid] { //碎了的数小,则向右移
					l = mid + 1
				} else { //否则向左移
					r = mid
				}
			}

			// 分鸡蛋碎了和没碎两种情况
			//1.碎了(鸡蛋数-1，楼层数-1),dp[i][j] = dp[i-1][j-1]
			//2.没碎(鸡蛋数不变，楼层数-m),dp[i][j] = dp[i][j-m]
			even[j] = common.MinInTwo(1+common.MaxInTwo(odd[l-1], even[j-l]), even[j])
		}
		copy(odd, even)
	}

	return odd[n]
}

// 动态规划
// 参考：https://leetcode.cn/problems/super-egg-drop/solutions/2231940/go-chao-duan-dai-ma-ti-mu-li-jie-qiao-mi-pkxc/
func twoEggDrop4(n int) int {
	//dp[i][j]表示i个鸡蛋最少扔j次所能探索到的最大楼层数，即鸡蛋不会烂的最大楼层数
	// 总有k个鸡蛋，有k行,有n层最多扔n次
	var dp = make([][]int, 3)
	for i := 0; i <= 2; i++ {
		dp[i] = make([]int, n+1)
	}

	//总楼层数=之下楼层数+之上楼层数+1（本层）
	//即dp[k][m] = dp[k][m - 1] + dp[k - 1][m - 1] + 1
	//按照上述的搜索流程，k个鸡蛋，扔m次，扔完第一次，可能上楼搜索，也可能下楼搜索。
	//	没碎，上楼搜索的最大可能搜索范围为dp[k][m-1]。（k个鸡蛋，还能扔m-1次）
	//	碎了，下楼搜索的最大可能搜索范围为dp[k-1][m-1]。（k-1个鸡蛋，还能扔m-1次）

	var m int
	for dp[2][m] < n {
		m++
		for i := 1; i <= 2; i++ {
			dp[i][m] = 1 + dp[i][m-1] + dp[i-1][m-1]
		}
	}

	return m
}
