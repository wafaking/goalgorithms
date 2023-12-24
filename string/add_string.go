package str

import (
	"strconv"
	"strings"
)

//字符串相加(leetcode-415)
//求给定两个字符串形式的非负整数num1和num2的和;
//示例1：输入：num1="11",num2="123"输出："134"
//示例2：输入：num1="456",num2="77"输出："533"
//示例3：输入：num1="0",num2="0"输出："0"
//1<=num1.length,num2.length<=10^4
//num1和num2都只包含数字0-9
//num1和num2都不包含任何前导零

func addStrings1(num1 string, num2 string) string {
	// 1. 特殊情况
	for _, v := range num1 {
		if v-'0' < 0 || v-'9' > 9 {
			return ""
		}
	}
	for _, v := range num2 {
		if v-'0' < 0 || v-'9' > 9 {
			return ""
		}
	}

	// 2. 拼凑长度使得m=n
	if len(num1) > len(num2) {
		num2 = strings.Repeat("0", len(num1)-len(num2)) + num2
	} else {
		num1 = strings.Repeat("0", len(num2)-len(num1)) + num1
	}

	// 3.求和并拼接
	var (
		level uint8
		res   string
	)
	for i := len(num1) - 1; i >= 0; i-- {
		temp := num1[i] - '0' + num2[i] - '0' + level
		if temp > 9 {
			level = 1
			temp %= 10
		} else {
			level = 0
		}
		res = strconv.FormatUint(uint64(temp), 10) + res
	}
	if level == 1 {
		res = "1" + res
	}

	return res
}

func addStrings2(num1 string, num2 string) string {
	var sign uint8
	var res string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || sign > 0; i, j = i-1, j-1 {
		var x, y uint8
		if i >= 0 {
			x = num1[i] - '0'
			if x > 9 || x < 0 {
				return ""
			}
		}
		if j >= 0 {
			y = num2[j] - '0'
			if y > 9 || y < 0 {
				return ""
			}
		}
		temp := x + y + sign
		if temp > 9 {
			sign = 1
			temp %= 10
		} else {
			sign = 0
		}
		res = strconv.Itoa(int(temp)) + res
	}
	return res
}

/*
func addTwoStrings(num1 string, num2 string) string {
	var (
		ans    string                         // 用来存储结果
		sign   uint8                          // 用来表示是否有进位
		n1, n2 = len(num1) - 1, len(num2) - 1 // 分别表示字符串的长度
	)

	//示例："128"+"35"
	// 从末尾开始， 先是8+5+sign=13+0=13,
	//	13>9,因此sign=1, sum=13%10=3
	//	将结果的字符串拼接起来，结果是：ans = sum + sum = "3" + ""= 3

	// 倒数第二位是：sum=2+3+sign=5+1=6
	//	因 6>9不成立，因此清空sign, sign=0
	//	拼接结果， 结果是：ans = sum + ans = "6"+"3"= "63"

	// 倒数第三位是， sum = 1+0+sign= 1+0= 1
	// 	注：因为第二个字符串的没有倒数第三位，因此默认为0，这也是下面的循环结束条件是i<=n1,i<=n2的原因，
	//	如果还理解不了，你可以把两个字符串补成长度一样，如："1238"+"12"可以把写成"1238"+"0012"
	//	因1>9不成立，因此清空sign，sign=0
	// 拼接结果： ans = sum + ans = "1"+"63"="163"

	// 只要sign，str1和str2的最高位三者中有一个不为0，都继续计算，因此循环结束的条件是i<=n1||i<=n2||sign>0
	// 从0开始累计，如果
	for i := 0; i <= n1 || i <= n2 || sign != 0; i++ {
		var a, b uint8
		if n1 >= i {
			a = num1[n1-i] - '0'
		}
		if n2 >= i {
			b = num2[n2-i] - '0'
		}
		sum := sign + a + b
		if sum > 9 {
			sign = 1
			sum %= 10
		} else {
			sign = 0
		}
		ans = strconv.Itoa(int(sum)) + ans
	}
	return ans
}

*/
