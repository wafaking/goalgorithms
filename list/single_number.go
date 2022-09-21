package list

// 只出现一次的数字(leetcode-136)
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//示例 1: 输入: [2,2,1],输出: 1
//示例 2: 输入: [4,1,2,1,2],输出: 4

// singleNumber 位运算(相同的数值异或运算结果为0)
func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

// 只出现一次的数字II(137)
// 给定整数数组nums，除某个元素仅出现一次外，其余每个元素都恰出现三次。请找出只出现了一次的元素。
//示例 1： 输入：nums = [2,2,3,2],输出：3
//示例 2： 输入：nums = [0,1,0,1,0,1,99],输出：99

// singleNumber2 使用位运算(分别将每一位分别相加，不能被3整除的位组成的数就是要找)
func singleNumber21(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ { // 每个数字的
		var sum int32
		// 错误方法
		for _, v := range nums {
			sum += int32(v) & (1 << i)
		}
		//for _, v := range nums { // 此处可能为负数，因此右移
		//	sum += (int32(v) >> i) & 1
		//}

		if sum%3 > 0 {
			res |= 1 << i
		}
	}
	return int(res)
}

//解法1：三进制异或
// singleNumber22 三进制异或（自定义三进制异或运算：b1^b2 = (b1+b2)%3，这样3个同样的bit异或后就会归0）
func singleNumber22(nums []int) int {
	res1, res2 := 0, 0
	for _, num := range nums {
		if num >= 0 {
			res1 = TernaryBitXor(res1, num)
		} else {
			res2 = TernaryBitXor(res2, num)
		}
	}
	if res1 != 0 {
		return res1
	} else {
		return res2
	}
}

func TernaryBitXor(num1, num2 int) int {
	res := 0
	t := 1
	for ; num1 != 0 && num2 != 0; t *= 3 {
		b := (num1 + num2) % 3
		res += b * t
		num1 /= 3
		num2 /= 3
	}
	return res + (num1+num2)*t
}

func singleNumber23(nums []int) int {
	result := 0
	flag := uint64(1)
	for flag != 0 {
		count := 0
		for _, v := range nums {
			temp := v & int(flag)
			if temp == int(flag) {
				count++
			}
		}
		if count%3 == 1 {
			result |= int(flag)
		}
		flag <<= 1
	}
	return result
}

// 只出现一次的数字III(leetcode-260)
// 给定整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次,找出只出现一次的那两个元素。
//示例 1： 输入：nums = [1,2,1,3,2,5], 输出：[3,5]
//示例 2： 输入：nums = [-1,0], 输出：[-1,0]
//示例 3： 输入：nums = [0,1], 输出：[1,0]

func singleNumber31(nums []int) []int {
	var (
		res  = make([]int, 0, 2)
		sign = 1
		temp int
	)
	for _, v := range nums { // 所有数据异或后，temp是仅出现一次两个数的异或结果
		temp ^= v
	}
	// 获取temp的第1个1的位置，并返回数值，如1000
	for i := 0; ; i++ {
		if temp&sign != 0 {
			break
		}
		sign = sign << i
	}

	var t1, t2 int
	for _, v := range nums {
		if v&sign == 0 {
			t1 ^= v
		} else {
			t2 ^= v
		}
	}

	res = append(res, t1, t2)
	return res
}

func singleNumber32(nums []int) []int {
	var (
		res          = make([]int, 0, 2)
		temp, t1, t2 int
	)
	for _, v := range nums { // 所有数据异或后，temp是仅出现一次两个数的异或结果
		temp ^= v
	}
	//log.Printf("temp: %d: %v", temp, temp&-temp)
	// 获取temp的第1个1的位置，并返回数值，如1000
	temp &= -temp // 返回第1个1所在位置构成的数值

	for _, v := range nums {
		if v&temp == 0 {
			t1 ^= v
		} else {
			t2 ^= v
		}
	}

	res = append(res, t1, t2)

	return res
}
