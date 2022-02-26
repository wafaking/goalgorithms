package list

func maxArea(height []int) int {
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
