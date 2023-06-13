package dynamics

// 爬楼梯(leetcode-70)
// 假设你正在爬楼梯。需要n阶你才能到达楼顶。
// 每次你可以爬1或2个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 示例1：输入：n=2,输出：2 ,解释：有两种方法可以爬到楼顶, 1+1, 2
// 示例2：输入：n=3, 输出：3, 有三种方法可以爬到楼顶: 1+1+1, 1+2, 2+1

// climbStairs1 递归，从大到小
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

// climbStairs2 从小到大
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	var a, b, c = 1, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return b
}
