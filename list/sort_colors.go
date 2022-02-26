package list

import (
	"errors"
	"fmt"
)

func sortColors(nums []int) {
	var start, end int = 0, len(nums) - 1 // 设置起、始指针

	// 遍历数组，当值为0时，与起始指针值交换，
	for i := 0; i <= end; {
		switch nums[i] {
		case 0:
			nums[i], nums[start] = nums[start], nums[i]
			start++ // 起始指针向后移
			i++     // 游标后移
		case 1:
			i++ // 不作交换，游标后移
		case 2:
			nums[i], nums[end] = nums[end], nums[i] // 游标元素与尾指针元素交换
			end--                                   // 尾指针前移
		default:
			panic(errors.New("invalid given slice"))
			return
		}
	}
	fmt.Println("nums: ", nums)
}
