package str

//  质数因子(HJ6-https://www.nowcoder.com/practice/196534628ca6490ebce2e336b47b3607?tpId=37&tags=&title=&difficulty=0&judgeStatus=0&rp=1&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&dayCountBigMember=365%E5%A4%A9)
// 输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）
// 示例1
//   输入：180, 输出： 2 2 3 3 5

func primeFactor(num int) []int {
	var res = make([]int, 0)
	if num == 1 {
		res = append(res, num)
		return res
	}

	for i := 2; i*i <= num; {
		if num%i == 0 {
			res = append(res, i)
			num /= i
		} else {
			i++
		}
	}
	if num >= 2 {
		res = append(res, num)
	}
	return res
}
