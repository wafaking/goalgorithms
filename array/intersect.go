package array

import "sort"

//两个数组的交集II(leetcode-350)
//给你两个整数数组nums1和nums2，请你以数组形式返回两数组的交集。
//返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//示例1：输入：nums1=[1,2,2,1],nums2=[2,2]输出：[2,2]
//示例2:输入：nums1=[4,9,5],nums2=[9,4,9,8,4]输出：[4,9]
//提示：
//	1<=nums1.length,nums2.length<=1000
//	0<=nums1[i],nums2[i]<=1000
//进阶：
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果nums1的大小比nums2小，哪种方法更优？
//如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

// 使用map
func intersect1(nums1 []int, nums2 []int) []int {
	var (
		ans = make([]int, 0)
		m   = make(map[int]int, 0)
	)
	for k := range nums1 {
		m[nums1[k]]++
	}

	for k := range nums2 {
		if v, ok := m[nums2[k]]; ok && v > 0 {
			ans = append(ans, nums2[k])
			m[nums2[k]]--
		}
	}
	return ans
}

// 使用位数组(由题知：0<=num[i]<=1000)
func intersect2(nums1 []int, nums2 []int) []int {
	var (
		ans    = make([]int, 0)
		bitMap = make([]int, 1001)
	)
	for k := range nums1 {
		bitMap[nums1[k]]++
	}

	for k := range nums2 {
		if bitMap[nums2[k]] > 0 {
			ans = append(ans, nums2[k])
			bitMap[nums2[k]]--
		}
	}
	return ans
}

// 排序+双指针
//如果nums2的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中那么就无法高效地对nums2进行排序，
//因此推荐使用方法一。在方法一中，mums只关系到查询操作，因此每次读取nums2中的一部分数据，并进行处理即可。
func intersect3(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var ans = make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return ans
}
