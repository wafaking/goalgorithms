package str

//找出数组的最大公约数(leetcode-1979)
//给你一个整数数组nums，返回数组中最大数和最小数的最大公约数。
//两个数的最大公约数是能够被两个数整除的最大正整数。
//示例1：输入：nums=[2,5,6,9,10]输出：2
//解释：nums中最小的数是2nums中最大的数是102和10的最大公约数是2
//示例2：输入：nums=[7,5,6,8,3]输出：1解释：
//nums中最小的数是3,nums中最大的数是8,3和8的最大公约数是1
//示例3：输入：nums=[3,3]输出：3解释：
//nums中最小的数是3,nums中最大的数是3,3和3的最大公约数是3
//注：1<=nums[i]<=1000（详解参照漫画算法）

// 暴力求解
func findGCD1(nums []int) int {
	// 由题可知1<=nums[i]<=1000
	if len(nums) == 0 {
		return -1
	}
	var minNum, maxNum = nums[0], nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		} else if v > maxNum {
			maxNum = v
		}
	}

	if maxNum%minNum == 0 {
		return minNum
	}

	for i := minNum / 2; i > 1; i-- {
		if minNum%i == 0 && maxNum%i == 0 {
			return i
		}
	}
	return 1
}

//欧几里得算法
//两个正整数a和b（a＞b），它们的最大公约数等于a除以b的余数c和b之间的最大公约数。
//例如10和25，25除以10商2余5，那么10和25的最大公约数，等同于10和5的最大公约数。
func findGCD2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var minNum, maxNum = nums[0], nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		} else if v > maxNum {
			maxNum = v
		}
	}

	// 注：当两个数较大时，取模运算性能会比较差
	for maxNum%minNum != 0 {
		minNum, maxNum = maxNum%minNum, minNum
	}
	return minNum
}

//更相减损术
//两个正整数a和b（a＞b），它们的最大公约数等于a-b的差值c和较小数b的最大公约数。
//例如10和25，25减10的差是15，那么10和25的最大公约数，等同于10和15的最大公约数;
func findGCD3(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var minNum, maxNum = nums[0], nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		} else if v > maxNum {
			maxNum = v
		}
	}

	for minNum != maxNum {
		if maxNum-minNum < minNum {
			maxNum, minNum = minNum, maxNum-minNum
		} else {
			maxNum = maxNum - minNum
		}
	}
	return minNum
}

//更相减损术与移位相结合
//当a和b均为偶数时，gcd（a，b）=2×gcd（a/2，b/2）=2×gcd（a＞＞1，b＞＞1）。
//当a为偶数，b为奇数时，gcd（a，b）=gcd（a/2，b）=gcd（a＞＞1，b）。
//当a为奇数，b为偶数时，gcd（a，b）=gcd（a，b/2）=gcd（a，b＞＞1）。
//当a和b均为奇数时，先利用更相减损术运算一次，gcd（a，b）=gcd（b，a-b），此时a-b必然是偶数，然后又可以继续进行移位运算。
func findGCD4(nums []int) int {
	var minNum, maxNum = nums[0], nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		} else if v > maxNum {
			maxNum = v
		}
	}

	// greatest common divisor
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if a == b {
			return a
		}
		if a&1 == 0 && b&1 == 0 {
			//return 2 * gcd(a>>1, b>>1)// 这个也是正确的
			return gcd(a>>1, b>>1) << 1
		} else if a&1 == 0 && b&1 != 0 {
			return gcd(a>>1, b)
		} else if a&1 != 0 && b&1 == 0 {
			return gcd(a, b>>1)
		} else { // 都为奇数
			// 此处默认a>b
			if a < b {
				a, b = b, a
			}
			if a-b > b {
				return gcd(a-b, b)
			}
			return gcd(b, a-b)
		}
	}

	return gcd(maxNum, minNum)
}
