package array

import "goalgorithms/common"

//给定一个长度为n的整数数组height。有n条垂线，第i条线的两个端点是(i,0)和(i,height[i])。
//找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//返回容器可以储存的最大水量。
//输入：[1,8,6,2,5,4,8,3,7],输出：49
//解释：图中垂直线代表输入数组[1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
//示例2：输入：height=[1,1],输出：1

// 双指针
func maxArea1(height []int) int {
	var max int
	var start, end = 0, len(height) - 1

	for start < end {
		var minH int
		if height[end] < height[start] {
			minH = height[end]
		} else {
			minH = height[start]
		}

		if max < minH*(end-start) {
			max = minH * (end - start)
		}

		// 总是移动小的一条边
		// 如果移动大边，则新边可能变大也可能变小，但总体是变小的
		// 如果移动小边，则新边的高度可能变大，导致整体变大
		// 如果相同，则单移动一边整体还是变小，同时移动两边可能变大
		if height[start] > height[end] {
			end--
		} else if height[start] < height[end] {
			start++
		} else {
			start++
			end--
		}
	}
	return max
}

// 双指针
func maxArea2(height []int) int {
	var maxArea int
	for i, j := 0, len(height)-1; i < j; {
		area := common.MinInTwo(height[i], height[j]) * (j - i)
		maxArea = common.MaxInTwo(maxArea, area)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxArea
}
