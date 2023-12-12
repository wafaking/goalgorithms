package str

//快乐数(leetcode-202)
//编写一个算法来判断一个数n是不是快乐数。
//「快乐数」定义为：
//	对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
//	然后重复这个过程直到这个数变为1，也可能是无限循环但始终变不到1。
//	如果这个过程结果为1，那么这个数就是快乐数。
//	如果n是快乐数就返回true；不是，则返回false。
//示例1：输入：n=19输出：true解释：
//	1^2+9^2=82
//	8^2+2^2=68
//	62+82=100
//	12+02+02=1
//示例2：输入：n=2输出：false
//注：1<=n<=2^31-1

//哈希集合检测循环
func isHappy1(n int) bool {
	var (
		f func(n int) int
		m = make(map[int]bool, 0)
	)
	f = func(n int) int {
		if n == 1 {
			return 1
		}
		t := n % 10
		t *= t
		if n < 10 {
			return t
		}
		t += f(n / 10)
		return t
	}

	for n != 1 && !m[n] {
		m[n] = true
		n = f(n)
	}

	return n == 1
}

// 快慢指针
func isHappy2(n int) bool {
	var (
		f func(n int) int
	)
	f = func(n int) int {
		var sum int
		for n > 0 {
			t := n % 10
			sum += t * t
			n /= 10
		}
		return sum
	}
	slow, fast := n, f(n)
	for slow != fast && fast != 1 {
		slow = f(slow)
		fast = f(f(fast))
	}
	return fast == 1
}

// 数学(TODO)
func isHappy3(n int) bool {
	var step = func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		return sum
	}

	cycle := map[int]bool{4: true, 6: true, 37: true, 58: true, 89: true, 145: true, 42: true, 20: true}
	for n != 1 && !cycle[n] {
		n = step(n)
	}
	return n == 1
}
