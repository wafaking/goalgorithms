package array

// 不修改数组寻找重复数(leetcode-287)
// 给定一个包含n+1个整数的数组nums，其数字都在[1,n]范围内（包括1和n），仅有一个数重复，但重复次数未知。
// 返回这个重复的数。条件：不修改数组nums，空间复杂度为 O(1) 。
//示例 1：输入：nums = [1,3,4,2,2], 输出：2
//示例 2： 输入：nums = [3,1,3,4,2], 输出：3

// findDuplicate1 暴力解法(双循环，空间复杂度为O(1)，时间复杂度为O(n^2))
func findDuplicate1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

// 解题思路：
// 示例：[1,2,3,3,4](为方便解释，已排序)
// 1.因取值范围在[1,n]且元素个数为n+1,因此一定有重复元素；
// 2.将元素根据mid=2分为两部分[1,2]和(2,4](即[3,4])
// 	即将元素分为[1,2]和[3,3,4]
// 3.如果被分隔的元素个数大于原数组的一半，则重复的元素一定在此数组中
// 	即重复的元素一定在[3,3,4]中，以此循环，直至仅有一个数
// 4.时间复杂度，空间复杂度
// findDuplicate2 使用二分查找法
func findDuplicate21(nums []int) int {
	//1,2,3,3,4,5,4 n=6, len=7
	var l, r = 1, len(nums) - 1
	for l <= r {
		var mid = (r-l)>>1 + l
		var count int
		for _, v := range nums {
			if v >= l && v <= mid {
				count++
			}
		}

		if l == r { // 左右重合
			if count >= 1 { // 有满足条件的
				return l
			} else {
				break
			}
		}

		if count > mid-l+1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

// findDuplicate22 二分法查找法
func findDuplicate22(nums []int) int {
	//1,2,3,3,3,5,4 n=6, len=7
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1 // 取中间位
		cnt := 0
		for i := 0; i < n; i++ { // 获取<=mid的数值的个数
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt > mid { // <=mid小的数占多数，右边界缩小
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

//如果被替换的数第i 位为 1，且 target 第 ii 位为 1：x 不变，满足 x>y。
//如果被替换的数第i 位为 0，且 target 第 ii 位为 1：x 加一，满足 x>y。
//如果被替换的数第i 位为 1，且 target 第 ii 位为 0：x 减一，满足 x≤y。
//如果被替换的数第i 位为 0，且 target 第 ii 位为 0：x 不变，满足 x≤y。
// 1.计nums二进制第i位等于1的所有数的和为x，[1,n]二进制所有第i位为1的和为y
// 2.如果数字重复了两次，则除去一位重复的数字，nums和[1,n]是相等的,如：
//	[1,2,2,3]和[1,2,3](注：nums比[1,n]多一个数字)，则当x>y时的i组成的一定为重复的数值;
// 3.如果数字重复了超过两次，则可以看成[1,n]中部分整数被替换成了重复的数字，对于nums中
//	i位为1的每一位，替换之前有x>y,替换之后仍有x>y

// findDuplicate3 使用二进制法
func findDuplicate3(nums []int) int {
	var res int
	for i := 0; i < 32; i++ {
		var x, y int // nums中第i位为1的数的个数
		for _, v := range nums {
			if (v & (1 << i)) > 0 { // 判断v第i位是否为1
				x++
			}
		}

		for j := 1; j < len(nums); j++ { // [1,n]中第i位为1的数的个数
			if (j & (1 << i)) > 0 {
				y++
			}
		}

		if x > y { // 满足条件的
			res |= 1 << i // 将i位为0的二进制拼起来即为结果
		}
	}
	return res
}

// 1.类似于判断链表中是否有环一样，慢指针每次走一步，快指针每次走两步
// 2.将数组构造成链表，起始节点i=0，以num[i]为下一个节点，nums[nums[i]]为下下个节点
// 	即将nums[i]当作慢指针(每次走一步)，将nums[nums[i]]当作快指针(每次走两步,即相当于在nums[i]的基础上再走一步)
// 3.当快慢指针相遇时，将快指针重新放到头部，此后快慢指针每次都仅走一步，直至两者相遇节点即重复节点
// 4.此题隐含条件，nums[i]及nums[nums[i]]一定不会越界
// findDuplicate4 使用快慢指针
func findDuplicate4(nums []int) int {
	// {1, 3, 4, 2, 2}
	var slow, fast = 0, 0 // 两者都置于起始点,相当于虚拟一个头节点0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	fast = 0           // 头节点应从0开始，而不是nums[0]
	for slow != fast { // 此后两者每次都只走一步
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
