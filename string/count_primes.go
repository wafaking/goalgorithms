package str

//计数质数(leetcode-204)
//给定整数n，返回所有小于非负整数n的质数的数量。
//质数(素数): 指在大于1的自然数中，除了1和它本身以外不再有其他因数的自然数。
//示例1：输入：n=10输出：4解释：小于10的质数一共有4个,它们是2,3,5,7。
//示例2：输入：n=0输出：0
//示例3：输入：n=1输出：0

// 暴力枚举(易超出时间限制)
func countPrimes1(n int) int {
	var (
		ans         int
		notPrimeMap = make(map[int]bool, 0)
		isPrime     func(n int) bool
	)
	isPrime = func(n int) bool {
		if notPrimeMap[n] {
			return false
		}

		if n == 2 {
			return true
		} else if n&1 == 0 {
			notPrimeMap[n] = true
			return false
		}

		for i := 1; i < n; i++ {
			temp := 2*i + 1 // 只枚举奇数
			if temp >= n {
				break
			}
			if n%temp == 0 {
				notPrimeMap[n] = true
				return false
			}
		}
		return true
	}

	for i := n - 1; i > 1; i-- {
		if isPrime(i) {
			ans++
		}
	}

	return ans
}

// 埃氏筛
func countPrimes2(n int) int {
	var (
		isPrime = make([]bool, n)
		ans     int
	)
	for k := range isPrime {
		isPrime[k] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return ans
}

// 线性筛(TODO)
func countPrimes3(n int) int {
	var (
		primes  = make([]int, 0)
		isPrime = make([]bool, n)
	)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
