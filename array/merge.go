package array

import "sort"

//合并两个有序数组(leetcode-88)
//给你两个有序整数数组nums1和nums2，请你将nums2合并到nums1中，使nums1成为一个有序数组。
//初始化nums1和nums2的元素数量分别为m和n。你可以假设nums1的空间大小等于m+n，
//这样它就有足够的空间保存来自nums2的元素。
//示例1：输入：nums1=[1,2,3,0,0,0],m=3,nums2=[2,5,6],n=3,输出：[1,2,2,3,5,6]
//示例2：输入：nums1=[1],m=1,nums2=[],n=0,输出：[1]

// 法一:添加到队列，并排序
func merge1(nums1 []int, m int, nums2 []int, n int) {
	if m+n != len(nums1) {
		return
	}
	for i, _ := range nums2 {
		nums1[m] = nums2[i]
		m++
	}
	sort.Ints(nums1)
	return
}

// 法二：新建列表，依次遍历比较并添加
func merge2(nums1 []int, m int, nums2 []int, n int) {
	tail := m + n
	numsNew := make([]int, tail)
	for {
		if m <= 0 || n <= 0 {
			for i := m; i > 0; i-- {
				numsNew[tail-1] = nums1[i-1]
				tail--
			}
			for i := n; i > 0; i-- {
				numsNew[tail-1] = nums2[i-1]
				tail--
			}
			break
		}

		if nums1[m-1] > nums2[n-1] {
			numsNew[tail-1] = nums1[m-1]
			m--
		} else if nums1[m-1] < nums2[n-1] {
			numsNew[tail-1] = nums2[n-1]
			n--
		} else {
			numsNew[tail-1] = nums1[m-1]
			m--
			tail--
			numsNew[tail-1] = nums2[n-1]
			n--
		}
		tail--
	}

	copy(nums1, numsNew)
}

//法三
func merge3(nums1 []int, m int, nums2 []int, n int) {
	if m+n != len(nums1) { // 保证长度
		return
	}

	tail := len(nums1) - 1
	for n > 0 { // 保证nums2遍历完成
		if m <= 0 { // nums1遍历完了，直接将nums2剩余的添加到nums1中
			for n > 0 {
				nums1[tail] = nums2[n-1]
				tail--
				n--
			}
			return
		}

		if nums2[n-1] > nums1[m-1] {
			nums1[tail] = nums2[n-1]
			n--
		} else if nums2[n-1] < nums1[m-1] {
			nums1[tail] = nums1[m-1]
			m--
		} else { // 相等
			nums1[tail] = nums2[n-1]
			n--
			tail--
			nums1[tail] = nums1[m-1]
			m--
		}

		tail-- // 注意，每次添加完成都要移动尾指针
	}
}
