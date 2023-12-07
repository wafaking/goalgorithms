package array

//只出现一次的数字(leetcode-136)
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//示例1:输入:[2,2,1],输出:1
//示例2:输入:[4,1,2,1,2],输出:4

// singleNumber 位运算(相同的数值异或运算结果为0)
func singleNumber1(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

//只出现一次的数字II(137)
//给定整数数组nums，除某个元素仅出现一次外，其余每个元素都恰出现三次。请找出只出现了一次的元素。
//示例1：输入：nums=[2,2,3,2],输出：3
//示例2：输入：nums=[0,1,0,1,0,1,99],输出：99

// singleNumber2 使用位运算(分别将每一位分别相加，不能被3整除的位组成的数就是要找)
// 注：与时不会考虑符号位，因此不能左移
func singleNumber21(nums []int) int {
	var (
		res1 int32
		res2 int32
	)
	for i := 0; i < 32; i++ { // 每个数字的
		var sum1 int32
		var sum2 int32

		// 错误方法(示例数组：1, 1, 1, -4）,此处用8位计算表示
		//因为当i=31时,v=-4时int32(-4)用补码表示为1111 1100
		//用补码与(1<<31)相与计算结果为：1000 0000，转换为十进制结果为128,
		//因int8表示范围为-128~127，因此会溢出，结果为-128，导致后面sum多次溢出，结果sum%3不准确，因此右移是准确的
		// 注：与运算结果一定是正数，因此累加后导致溢出，sum变负数
		// i越大时，1<<i结果也越大，导致int32(v) & (1 << i)越大。
		for _, v := range nums {
			sum1 += int32(v) & (1 << i) // 错误写法
		}
		// 因为任何数与1作与运算都为1或0，因此不会溢出，除非nums数组的长度特别长，导致sum超出int32范围
		for _, v := range nums { // 此处可能为负数，因此右移
			sum2 += (int32(v) >> i) & 1
		}

		if sum1%3 != 0 { // 注：此处不能用sum1%3>0，因为可能sum1结果可能会溢出，导致为负数
			res1 |= 1 << i
		}
		if sum2%3 != 0 {
			res2 |= 1 << i
		}
	}
	return int(res2)
}

// TODO 解法1：三进制异或
// singleNumber22 三进制异或（自定义三进制异或运算：b1^b2=(b1+b2)%3，这样3个同样的bit异或后就会归0）
func singleNumber22(nums []int) int {
	var TernaryBitXor = func(num1, num2 int) int {
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

// TODO
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

// TODO 数字电路设计
func singleNumber24(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a, b = b&^a&num|a&^b&^num, (b^num)&^a
	}
	return b
}

// TODO 数字电路设计优化
func singleNumber25(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}

//只出现一次的数字III(leetcode-260)
//给定整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次,找出只出现一次的那两个元素。
//你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
//示例1：输入：nums=[1,2,1,3,2,5],输出：[3,5]
//示例2：输入：nums=[-1,0],输出：[-1,0]
//示例3：输入：nums=[0,1],输出：[1,0]

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

//位运算
func singleNumber32(nums []int) []int {

	// 所有元素异或结果
	var (
		xorSum     int
		num1, num2 int
	)
	for _, v := range nums {
		xorSum ^= v
	}
	// -xorSum&xorSum可得出二进制的xorSum从右到左第一位是1的值
	// 由于异或得1的一定是0和1,因此可以将nums中的数值分成第一位是1和非1的两队
	// 再分别将两队的数字各自异或，则可以得到各自的数值
	xorSum &= -xorSum
	for _, v := range nums {
		if xorSum&v == 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}

	return []int{num1, num2}
}
