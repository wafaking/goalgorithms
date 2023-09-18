package dynamics

import (
	"testing"
)

func TestCuttingRope(t *testing.T) {
	var sampleMap = map[int]int{
		2:  1,
		3:  2,
		5:  6,
		6:  9,
		7:  12, //  3;4
		10: 36, // 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
	}
	for n, expected := range sampleMap {
		// res := cuttingRope1(n)
		// res := cuttingRope2(n)
		res := cuttingRope3(n)
		t.Logf("res: %t, n:%d, res:%d, expected:%d", res == expected, n, res, expected)
	}
}

func TestFindContentChildren(t *testing.T) {
	var list = []item1{
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 1},
			expected: 1,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{1, 2, 3},
			expected: 2,
		},
		{
			nums1:    []int{10, 9, 8, 7},
			nums2:    []int{5, 6, 7, 8},
			expected: 2,
		},
	}

	for _, item := range list {
		res := findContentChildren(item.nums1, item.nums2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestClimbStairs(t *testing.T) {
	var list = []item2{
		{
			num:      2,
			expected: 2,
		},
		{
			num:      3,
			expected: 3,
		},
		{
			num:      4,
			expected: 5,
		},
	}

	for _, item := range list {
		res := climbStairs1(item.num)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = climbStairs2(item.num)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	var list = []item3{
		{
			nums:     []int{10, 15, 20},
			expected: 15,
		},
		{
			nums:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, item := range list {
		res := minCostClimbingStairs1(item.nums)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = minCostClimbingStairs2(item.nums)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePaths(t *testing.T) {
	var list = []item4{
		{
			num1:     3,
			num2:     7,
			expected: 28,
		},
		{
			num1:     3,
			num2:     2,
			expected: 3,
		},
		{
			num1:     7,
			num2:     3,
			expected: 28,
		},
		{
			num1:     3,
			num2:     3,
			expected: 6,
		},
		{
			num1:     1,
			num2:     2,
			expected: 1,
		},
		{
			num1:     2,
			num2:     1,
			expected: 1,
		},
	}

	for _, item := range list {
		res := uniquePaths1(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = uniquePaths2(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = uniquePaths3(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = uniquePaths4(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	var list = []item5{
		{
			grid:     [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expected: 2,
		},
		{
			grid:     [][]int{{0, 1}, {0, 0}},
			expected: 1,
		},
		{
			grid:     [][]int{{0, 0}},
			expected: 1,
		},
		{
			grid:     [][]int{{1, 0}},
			expected: 0,
		},
		{
			grid:     [][]int{{0}, {0}},
			expected: 1,
		},
		{
			grid:     [][]int{{1}, {0}},
			expected: 0,
		},
		{
			grid:     [][]int{{0, 1, 0}},
			expected: 0,
		},
		{
			grid:     [][]int{{0, 0}, {1, 1}, {0, 0}},
			expected: 0,
		},
		{
			grid:     [][]int{{0, 0}, {0, 0}, {0, 1}},
			expected: 0,
		},
		{
			grid:     [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
			expected: 0,
		},
		{
			grid:     [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			expected: 0,
		},
		{
			grid:     [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}},
			expected: 10,
		},
		{
			grid: [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
				{0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {1, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {0, 1}, {0, 0}, {0, 0},
				{1, 0}, {0, 0}, {0, 0}, {0, 1}, {0, 0},
				{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 1}, {0, 0}, {0, 0}, {0, 0},
				{0, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0},
				{0, 0}},
			expected: 0,
		},
	}

	for _, item := range list {
		res := uniquePathsWithObstacles1(item.grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = uniquePathsWithObstacles2(item.grid)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = uniquePathsWithObstacles3(item.grid)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIntegerBreak(t *testing.T) {
	var list = []item2{
		{
			num:      2,
			expected: 1,
		},
		{
			num:      3,
			expected: 2,
		},
		{
			num:      4,
			expected: 4,
		},
		{
			num:      5,
			expected: 6,
		},
		{
			num:      6,
			expected: 9,
		},
		{
			num:      7,
			expected: 12,
		},
		{
			num:      8,
			expected: 18,
		},
		{
			num:      9,
			expected: 27,
		},
		{
			num:      10,
			expected: 36,
		},
		{
			num:      11,
			expected: 54,
		},
		{
			num:      12,
			expected: 81,
		},
	}

	for _, item := range list {
		res := integerBreak1(item.num)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob(t *testing.T) {
	var list = []item3{
		//{nums: []int{}, expected: 0},
		//{nums: []int{1}, expected: 1},
		//{nums: []int{3, 1}, expected: 3},
		//{nums: []int{1, 2, 4, 3, 0, 9}, expected: 14},
		//{nums: []int{1, 2, 3, 1}, expected: 4},
		{nums: []int{2, 7, 9, 3, 1}, expected: 12},
		{nums: []int{2, 1, 1, 2}, expected: 4},
	}
	for _, item := range list {
		res := rob11(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob12(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob13(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob2(t *testing.T) {
	var list = []item3{
		{nums: []int{}, expected: 0},
		{nums: []int{1}, expected: 1},
		{nums: []int{1, 2}, expected: 2},
		{nums: []int{3, 1}, expected: 3},
		{nums: []int{2, 3, 2}, expected: 3},
		{nums: []int{1, 2, 3}, expected: 3},
		{nums: []int{1, 2, 1, 1}, expected: 3},
		{nums: []int{1, 2, 3, 1}, expected: 4},
		{nums: []int{2, 7, 9, 3, 1}, expected: 11},
		{nums: []int{2, 1, 1, 2}, expected: 3},
		//例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob3(t *testing.T) {
	var list = []item3{
		{nums: []int{}, expected: 0},
		{nums: []int{1}, expected: 1},
		{nums: []int{1, 2}, expected: 2},
		{nums: []int{3, 1}, expected: 3},
		{nums: []int{2, 3, 2}, expected: 3},
		{nums: []int{1, 2, 3}, expected: 3},
		{nums: []int{1, 2, 1, 1}, expected: 3},
		{nums: []int{1, 2, 3, 1}, expected: 4},
		{nums: []int{2, 7, 9, 3, 1}, expected: 11},
		{nums: []int{2, 1, 1, 2}, expected: 3},
		//例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxValue(t *testing.T) {
	var list = []item5{
		{
			grid: [][]int{
				{4, 2, 1},
			},
			expected: 7,
		},
		{
			grid: [][]int{
				{1},
				{1},
				{4},
			},
			expected: 6,
		},
		{
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 12,
		},
	}
	for _, item := range list {
		res := maxValue12(item.grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = maxValue11(item.grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}

}

func TestLengthOfLIS11(t *testing.T) {
	var list = []item3{
		{
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
	}
	for _, item := range list {
		res := lengthOfLIS11(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS12(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS13(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS14(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS15(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")

	}
}

func TestWeightBag1(t *testing.T) {
	var list = []item6{
		{
			weight:    []int{3, 2, 5, 5},
			value:     []int{2, 8, 1, 3},
			bagWeight: 6,
			expected:  10,
		},
		{
			weight:    []int{1, 4, 4, 3},
			value:     []int{2, 1, 3, 1},
			bagWeight: 6,
			expected:  5,
		},
		{
			weight:    []int{2, 3, 3, 1},
			value:     []int{6, 8, 1, 9},
			bagWeight: 6,
			expected:  23,
		},
		{
			weight:    []int{3, 3, 2, 2, 5, 5, 5, 5},
			value:     []int{2, 2, 8, 8, 1, 1, 3, 3},
			bagWeight: 6,
			expected:  16,
		},
		{
			weight:    []int{1, 1, 4, 4, 4, 4, 3, 3},
			value:     []int{2, 2, 1, 1, 3, 3, 1, 1},
			bagWeight: 6,
			expected:  7,
		},
		{
			weight:    []int{2, 2, 3, 3, 3, 3, 1, 1},
			value:     []int{6, 6, 8, 8, 1, 1, 9, 9},
			bagWeight: 6,
			expected:  30,
		},
	}
	var res int
	for _, item := range list {
		res = weightBag11(item.weight, item.value, item.bagWeight)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = weightBag13(item.weight, item.value, item.bagWeight)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestWeightBag2(t *testing.T) {
	var list = []item7{
		{
			weight:    []int{3, 2, 5, 5},
			value:     []int{2, 8, 1, 3},
			times:     []int{2, 2, 2, 2},
			bagWeight: 6,
			expected:  16,
		},
		{
			weight:    []int{1, 4, 4, 3},
			value:     []int{2, 1, 3, 1},
			times:     []int{2, 2, 2, 2},
			bagWeight: 6,
			expected:  7,
		},
		{
			weight:    []int{2, 3, 3, 1},
			value:     []int{6, 8, 1, 9},
			times:     []int{2, 2, 2, 2},
			bagWeight: 6,
			expected:  30,
		},
	}
	for _, item := range list {
		res := weightBag21(item.weight, item.value, item.times, item.bagWeight)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsMatchRegx(t *testing.T) {
	var list = []item8{
		{
			s:        "abc",
			p:        ".*c",
			expected: true,
		},
		{
			s:        "abc",
			p:        "*.c",
			expected: false,
		},
		{
			s:        "AAAA",
			p:        "BAAC*.",
			expected: false,
		},
		{
			s:        "ACAB",
			p:        "A.B*AB",
			expected: true,
		},
		{
			s:        "ABAA",
			p:        "A*CA.C",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			s:        "ab",
			p:        "..",
			expected: true,
		},
		{
			s:        "ab",
			p:        "**",
			expected: false,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isMatchRegx1(item.s, item.p)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsMatchWildCard(t *testing.T) {
	var list = []item8{
		{
			s:        "abc",
			p:        "?*c",
			expected: true,
		},
		{
			s:        "abc",
			p:        "*?c",
			expected: true,
		},
		{
			s:        "CABCB",
			p:        "C?BC*",
			expected: true,
		},
		{
			s:        "AAAA",
			p:        "BAAC*?",
			expected: false,
		},
		{
			s:        "ACAB",
			p:        "A?*AB",
			expected: true,
		},
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "*",
			expected: true,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "b",
			p:        "?*",
			expected: true,
		},
		{
			s:        "ab",
			p:        "?*",
			expected: true,
		},
		{
			s:        "ab",
			p:        "??",
			expected: true,
		},
		{
			s:        "ab",
			p:        "**",
			expected: true,
		},
		{
			s:        "ab",
			p:        "*",
			expected: true,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: false,
		},
		{
			s:        "AC",
			p:        "?*C",
			expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isMatchWildCard(item.s, item.p)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	var list = []item9{
		{
			text1:    "AAAT",
			text2:    "ATACATAC",
			expected: 4,
		},
		{
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = longestCommonSubsequence(item.text1, item.text2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestLongestCommonSubstring(t *testing.T) {
	var list = []item9{
		{
			text1:    "AAAT",
			text2:    "ATACATAC",
			expected: 2,
		},
		{
			text1:    "bacde",
			text2:    "ace",
			expected: 2,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = longestCommonSubstring(item.text1, item.text2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestFindLength(t *testing.T) {
	var list = []item1{
		{
			nums1:    []int{1, 2, 3, 2, 1},
			nums2:    []int{3, 2, 1, 4, 7},
			expected: 3,
			//示例2：输入：nums1=[0,0,0,0,0],nums2=[0,0,0,0,0]输出：5
		},
		{
			nums1:    []int{0, 0, 0, 0, 0},
			nums2:    []int{0, 0, 0, 0, 0},
			expected: 5,
		},
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2},
			expected: 2,
		},
	}
	var res int
	for _, item := range list {
		res = findLength(item.nums1, item.nums2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinDistance(t *testing.T) {
	var list = []item9{
		{
			text1:    "TAAGC",
			text2:    "GAACT",
			expected: 3,
		},
		{
			text1:    "horse",
			text2:    "ros",
			expected: 3,
		},
		{
			text1:    "intention",
			text2:    "execution",
			expected: 5,
		},
	}
	var res int
	for _, item := range list {
		res = minDistance(item.text1, item.text2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestCoinChange(t *testing.T) {
	var list = []item10{
		{
			nums:     []int{2},
			target:   3,
			expected: -1,
		},
		{
			nums:     []int{1},
			target:   0,
			expected: 0,
		},
		{
			nums:     []int{1, 2, 5},
			target:   11,
			expected: 3,
		},
		{
			nums:     []int{5, 2, 1},
			target:   11,
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3, 4},
			target:   6,
			expected: 2,
		},
		{
			nums:     []int{5, 7, 2, 1, 6},
			target:   6,
			expected: 1,
		},
		{
			nums:     []int{186, 419, 83, 408},
			target:   6249,
			expected: 20,
		},
	}
	var res int
	for _, item := range list {
		res = coinChange(item.nums, item.target)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestChange(t *testing.T) {
	var list = []item10{
		{
			nums:     []int{2},
			target:   3,
			expected: 0,
		},
		{
			nums:     []int{1},
			target:   0,
			expected: 1,
		},
		{
			nums:     []int{1},
			target:   5,
			expected: 1,
		},
		{
			nums:     []int{5},
			target:   10,
			expected: 1,
		},
		{
			nums:     []int{1, 2, 5},
			target:   5,
			expected: 4,
		},
		{
			nums:     []int{5, 2, 1},
			target:   5,
			expected: 4,
		},
		{
			nums:     []int{7},
			target:   0,
			expected: 1,
		},
		{
			nums:     []int{10},
			target:   10,
			expected: 1,
		},
		{
			nums:     []int{7},
			target:   0,
			expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = change1(item.target, item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = change2(item.target, item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestCanPartition(t *testing.T) {
	var list = []item11{
		{
			nums:     []int{1, 0, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 2},
			expected: true,
		},
		{
			nums:     []int{1, 3, 5, 3},
			expected: true,
		},
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			nums:     []int{1, 6},
			expected: false,
		},
		{
			nums:     []int{2},
			expected: false,
		},
		{
			nums:     []int{1, 2, 2},
			expected: false,
		},
		{
			nums:     []int{1, 2, 5},
			expected: false,
		},
		//{
		//	nums: []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		//		100, 100, 100, 100, 100, 100, 99, 97},
		//	expected: false,
		//},
	}

	var res bool
	for _, item := range list {
		res = canPartition1(item.nums)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		res = canPartition2(item.nums)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		res = canPartition3(item.nums)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestFindTargetSumWays(t *testing.T) {
	var list = []item10{
		{
			nums:     []int{1, 1, 1, 1, 1},
			target:   3,
			expected: 5,
		},
		{
			nums:     []int{1},
			target:   1,
			expected: 1,
		},
		{
			nums:     []int{1},
			target:   2,
			expected: 0,
		},
		{
			nums:     []int{1, 0},
			target:   1,
			expected: 2,
		},
		{
			nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
			target:   1,
			expected: 256,
		},
		{
			nums:     []int{1, 2, 1},
			target:   0,
			expected: 2,
		},
	}

	var res int
	for _, item := range list {
		res = findTargetSumWays11(item.nums, item.target)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = findTargetSumWays12(item.nums, item.target)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = findTargetSumWays13(item.nums, item.target)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestFindMaxForm(t *testing.T) {
	var list = []item13{
		{
			str: []string{"10", "0001", "111001", "1", "0"},
			item4: item4{
				num1:     5,
				num2:     3,
				expected: 4,
			},
		},
		{
			str: []string{"10", "0", "1"},
			item4: item4{
				num1:     1,
				num2:     1,
				expected: 2,
			},
		},
	}

	var res int
	for _, item := range list {
		res = findMaxForm1(item.str, item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsSubsequence(t *testing.T) {
	var list = []item8{
		{
			s:        "abc",
			p:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			p:        "ahbgdc",
			expected: false,
		},
		{
			s:        "",
			p:        "ahbgdc",
			expected: true,
		},
		{
			s:        "aaaaaa",
			p:        "bbaaaa",
			expected: false,
		},
		{
			s:        "bb",
			p:        "ahbgdc",
			expected: false,
		},
	}
	var res bool
	for _, item := range list {
		res = isSubsequence1(item.s, item.p)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		res = isSubsequence3(item.s, item.p)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		res = isSubsequence4(item.s, item.p)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsSubstring(t *testing.T) {
	var list = []item8{
		{
			s:        "abc",
			p:        "ahbgdc",
			expected: false,
		},
		{
			s:        "hbg",
			p:        "ahbgdc",
			expected: true,
		},
		{
			s:        "hbg",
			p:        "hbg",
			expected: true,
		},
		{
			s:        "axc",
			p:        "ahbgdc",
			expected: false,
		},
		{
			s:        "",
			p:        "ahbgdc",
			expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isSubstring1(item.s, item.p)
		t.Logf("%t, res-expected: %t:%t, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestTwoEggDrop(t *testing.T) {
	var list = []item2{
		{
			num:      1,
			expected: 1,
		},
		{
			num:      2,
			expected: 2,
		},
		{
			num:      3,
			expected: 2,
		},
		{
			num:      4,
			expected: 3,
		},
		{
			num:      5,
			expected: 3,
		},
		{
			num:      100,
			expected: 14,
		},
	}
	var res int
	for _, item := range list {
		res = twoEggDrop1(item.num)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = twoEggDrop2(item.num)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = twoEggDrop3(item.num)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = twoEggDrop4(item.num)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestSuperEggDrop(t *testing.T) {
	var list = []item4{
		{
			num1:     1,
			num2:     2,
			expected: 2,
		},
		{
			num1:     3,
			num2:     7,
			expected: 3,
		},
		{
			num1:     2,
			num2:     6,
			expected: 3,
		},
		{
			num1:     3,
			num2:     14,
			expected: 4,
		},
		{
			num1:     7,
			num2:     10000,
			expected: 15,
		},
	}
	var res int
	for _, item := range list {
		res = superEggDrop1(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = superEggDrop2(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = superEggDrop3(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = superEggDrop4(item.num1, item.num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestLastStoneWeight1(t *testing.T) {
	var list = []item3{
		{
			nums:     []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = lastStoneWeight1(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit1(t *testing.T) {
	var list = []item3{
		{
			nums:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			nums:     []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = maxProfit11(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = maxProfit12(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = maxProfit13(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit2(t *testing.T) {
	var list = []item3{
		{
			nums:     []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			nums:     []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = maxProfit21(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = maxProfit22(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		//res = maxProfit13(item.nums)
		//t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit3(t *testing.T) {
	var list = []item3{
		{
			nums:     []int{3, 3, 5, 0, 0, 3, 1, 4},
			expected: 6,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			nums:     []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			nums:     []int{1},
			expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = maxProfit31(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = maxProfit32(item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit4(t *testing.T) {
	var list = []item10{
		{
			nums:     []int{2, 4, 1},
			target:   2,
			expected: 2,
		},
		{
			nums:     []int{3, 2, 6, 5, 0, 3},
			target:   2,
			expected: 7,
		},
		{
			nums:     []int{2, 4, 1, 8, 2, 5, 9, 11},
			target:   3,
			expected: 18,
		},
	}

	var res int
	for _, item := range list {
		res = maxProfit41(item.target, item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = maxProfit42(item.target, item.nums)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}
