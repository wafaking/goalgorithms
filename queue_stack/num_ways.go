package queue_stack

// 青蛙跳台阶问题(sword-10)
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个n级的台阶总共有多少种跳法。
// 当n=0时，取1；
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

func numWays1(n int) int {
	if n <= 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	var (
		fib0, fib1 = 1, 1
		fibN       int
	)
	for i := 2; i <= n; i++ {
		fibN = (fib0 + fib1) % (1e9 + 7)
		fib0, fib1 = fib1, fibN
	}
	return fibN
}

func numWays2(n int) int {
	var (
		fibA, fibB = 1, 1
	)
	for i := 2; i <= n; i++ {
		fibA, fibB = fibB, (fibA+fibB)%(1e9+7)
	}
	return fibB
}
