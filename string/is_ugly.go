package str

//丑数(leetcode-263)
//丑数就是只包含质因数2、3和5的正整数。
//给你一个整数n，请你判断n是否为丑数。如果是，返回true；否则，返回false。
//示例1：输入：n=6输出：true解释：6=2×3
//示例2：输入：n=1输出：true解释：1没有质因数，因此它的全部质因数是{2,3,5}的空集。习惯上将其视作第一个丑数。
//示例3：输入：n=14输出：false解释：14不是丑数，因为它包含了另外一个质因数7。
//提示：
//	-2^31 <= n <= 2^31 - 1
func isUgly1(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}
