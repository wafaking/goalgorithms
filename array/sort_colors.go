package array

//颜色分类(leetcode-75)
//给定一个包含红色、白色和蓝色、共n个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//我们使用整数0、1和2分别表示红色、白色和蓝色。
//不使用库内置的sort函数。
//示例1：输入：nums=[2,0,2,1,1,0],输出：[0,0,1,1,2,2]
//示例2：输入：nums=[2,0,1],输出：[0,1,2]

// 单指针(两次遍历，分别交换0，1)
func sortColors1(nums []int) {
	var traverseFunc func(pos, i int) int
	traverseFunc = func(pos, targetNum int) int {
		for i := pos; i < len(nums); i++ {
			if nums[i] == targetNum {
				nums[i], nums[pos] = nums[pos], nums[i]
				pos++
			}
		}
		return pos
	}

	pos := traverseFunc(0, 0)
	_ = traverseFunc(pos, 1)
}

// 双指针（三路快排）
func sortColors2(nums []int) {
	var l, r = -1, len(nums)
	// (-1,l]是0,(l, r)是1，[r,)是2
	for i := 0; i < r; {
		switch nums[i] {
		case 0:
			l++
			nums[l], nums[i] = nums[i], nums[l]
			i++
		case 1:
			i++
		case 2:
			r--
			nums[r], nums[i] = nums[i], nums[r]
		}
	}
}
