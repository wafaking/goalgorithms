package str

import (
	"strconv"
)

// 回文数(leetcode-9)
// 判断给定整数x ，是不是回文整数，是返回true；否则，返回false。

// 示例 1： 输入：x=121 输出：true
// 示例 2： 输入：x=-121 输出：false, 即（-121与121-不同）
// 示例 3： 输入：x = 10 输出：false

//isPalindrome1 转成字符串
func isPalindrome1(x int) bool {
	var str = strconv.Itoa(x)
	//for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// isPalindrome21 翻转一半整数后对比是否相等(需要考虑数值的位数是奇数还是偶数)
func isPalindrome21(x int) bool {
	// 注意x=10的情况
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var num int
	for x > num {
		mod := x % 10
		num = num*10 + mod
		if x == num { // 数值长度为奇数位的情况
			break
		}
		x /= 10
	}
	return x == num
}

// isPalindrome21 翻转一半数字
func isPalindrome22(x int) bool {
	// 特殊情况：
	// 如上所述，当x< 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

// isPalindrome3 分别对比前后位数字
func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}

	// 获取最大mod,如1221：div=1000
	var div = 1
	for ; x/div >= 10; div *= 10 {
	}

	for x > 0 {
		if x/div != x%10 { // 判断前后首位是否相等
			return false
		}
		x = (x % div) / 10
		div /= 100
	}
	return true
}
