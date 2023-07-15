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
			Sli1:     []int{1, 2, 3},
			Sli2:     []int{1, 1},
			expected: 1,
		},
		{
			Sli1:     []int{1, 2},
			Sli2:     []int{1, 2, 3},
			expected: 2,
		},
		{
			Sli1:     []int{10, 9, 8, 7},
			Sli2:     []int{5, 6, 7, 8},
			expected: 2,
		},
	}

	for _, item := range list {
		res := findContentChildren(item.Sli1, item.Sli2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestClimbStairs(t *testing.T) {
	var list = []item2{
		{
			Num:      2,
			expected: 2,
		},
		{
			Num:      3,
			expected: 3,
		},
		{
			Num:      4,
			expected: 5,
		},
	}

	for _, item := range list {
		res := climbStairs1(item.Num)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = climbStairs2(item.Num)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	var list = []item3{
		{
			Sli:      []int{10, 15, 20},
			expected: 15,
		},
		{
			Sli:      []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, item := range list {
		res := minCostClimbingStairs1(item.Sli)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = minCostClimbingStairs2(item.Sli)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePaths(t *testing.T) {
	var list = []item4{
		{
			Num1:     3,
			Num2:     7,
			expected: 28,
		},
		{
			Num1:     3,
			Num2:     2,
			expected: 3,
		},
		{
			Num1:     7,
			Num2:     3,
			expected: 28,
		},
		{
			Num1:     3,
			Num2:     3,
			expected: 6,
		},
		{
			Num1:     1,
			Num2:     2,
			expected: 1,
		},
		{
			Num1:     2,
			Num2:     1,
			expected: 1,
		},
	}

	for _, item := range list {
		res := uniquePaths1(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = uniquePaths2(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = uniquePaths3(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = uniquePaths4(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	var list = []item5{
		{
			Grid:     [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expected: 2,
		},
		{
			Grid:     [][]int{{0, 1}, {0, 0}},
			expected: 1,
		},
		{
			Grid:     [][]int{{0, 0}},
			expected: 1,
		},
		{
			Grid:     [][]int{{1, 0}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0}, {0}},
			expected: 1,
		},
		{
			Grid:     [][]int{{1}, {0}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0, 1, 0}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0, 0}, {1, 1}, {0, 0}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0, 0}, {0, 0}, {0, 1}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			expected: 0,
		},
		{
			Grid:     [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}},
			expected: 10,
		},
		{
			Grid: [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
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
		res := uniquePathsWithObstacles1(item.Grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = uniquePathsWithObstacles2(item.Grid)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		res = uniquePathsWithObstacles3(item.Grid)
		t.Logf("%t, res-expected: %d:%d", res == item.expected, res, item.expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIntegerBreak(t *testing.T) {
	var list = []item2{
		{
			Num:      2,
			expected: 1,
		},
		{
			Num:      3,
			expected: 2,
		},
		{
			Num:      4,
			expected: 4,
		},
		{
			Num:      5,
			expected: 6,
		},
		{
			Num:      6,
			expected: 9,
		},
		{
			Num:      7,
			expected: 12,
		},
		{
			Num:      8,
			expected: 18,
		},
		{
			Num:      9,
			expected: 27,
		},
		{
			Num:      10,
			expected: 36,
		},
		{
			Num:      11,
			expected: 54,
		},
		{
			Num:      12,
			expected: 81,
		},
	}

	for _, item := range list {
		res := integerBreak1(item.Num)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob(t *testing.T) {
	var list = []item3{
		//{Sli: []int{}, expected: 0},
		//{Sli: []int{1}, expected: 1},
		//{Sli: []int{3, 1}, expected: 3},
		//{Sli: []int{1, 2, 4, 3, 0, 9}, expected: 14},
		//{Sli: []int{1, 2, 3, 1}, expected: 4},
		{Sli: []int{2, 7, 9, 3, 1}, expected: 12},
		{Sli: []int{2, 1, 1, 2}, expected: 4},
	}
	for _, item := range list {
		res := rob11(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob12(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob13(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob2(t *testing.T) {
	var list = []item3{
		{Sli: []int{}, expected: 0},
		{Sli: []int{1}, expected: 1},
		{Sli: []int{1, 2}, expected: 2},
		{Sli: []int{3, 1}, expected: 3},
		{Sli: []int{2, 3, 2}, expected: 3},
		{Sli: []int{1, 2, 3}, expected: 3},
		{Sli: []int{1, 2, 1, 1}, expected: 3},
		{Sli: []int{1, 2, 3, 1}, expected: 4},
		{Sli: []int{2, 7, 9, 3, 1}, expected: 11},
		{Sli: []int{2, 1, 1, 2}, expected: 3},
		//例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob3(t *testing.T) {
	var list = []item3{
		{Sli: []int{}, expected: 0},
		{Sli: []int{1}, expected: 1},
		{Sli: []int{1, 2}, expected: 2},
		{Sli: []int{3, 1}, expected: 3},
		{Sli: []int{2, 3, 2}, expected: 3},
		{Sli: []int{1, 2, 3}, expected: 3},
		{Sli: []int{1, 2, 1, 1}, expected: 3},
		{Sli: []int{1, 2, 3, 1}, expected: 4},
		{Sli: []int{2, 7, 9, 3, 1}, expected: 11},
		{Sli: []int{2, 1, 1, 2}, expected: 3},
		//例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxValue(t *testing.T) {
	var list = []item5{
		{
			Grid: [][]int{
				{4, 2, 1},
			},
			expected: 7,
		},
		{
			Grid: [][]int{
				{1},
				{1},
				{4},
			},
			expected: 6,
		},
		{
			Grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 12,
		},
	}
	for _, item := range list {
		res := maxValue12(item.Grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = maxValue11(item.Grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}

}

func TestLengthOfLIS11(t *testing.T) {
	var list = []item3{
		{
			Sli:      []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			Sli:      []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			Sli:      []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
	}
	for _, item := range list {
		res := lengthOfLIS11(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS12(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS13(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS14(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS15(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")

	}
}

func TestWeightBag1(t *testing.T) {
	var list = []item6{
		{
			Weight:    []int{3, 2, 5, 5},
			Value:     []int{2, 8, 1, 3},
			BagWeight: 6,
			expected:  10,
		},
		{
			Weight:    []int{1, 4, 4, 3},
			Value:     []int{2, 1, 3, 1},
			BagWeight: 6,
			expected:  5,
		},
		{
			Weight:    []int{2, 3, 3, 1},
			Value:     []int{6, 8, 1, 9},
			BagWeight: 6,
			expected:  23,
		},
		{
			Weight:    []int{3, 3, 2, 2, 5, 5, 5, 5},
			Value:     []int{2, 2, 8, 8, 1, 1, 3, 3},
			BagWeight: 6,
			expected:  16,
		},
		{
			Weight:    []int{1, 1, 4, 4, 4, 4, 3, 3},
			Value:     []int{2, 2, 1, 1, 3, 3, 1, 1},
			BagWeight: 6,
			expected:  7,
		},
		{
			Weight:    []int{2, 2, 3, 3, 3, 3, 1, 1},
			Value:     []int{6, 6, 8, 8, 1, 1, 9, 9},
			BagWeight: 6,
			expected:  30,
		},
	}
	for _, item := range list {
		res := weightBag11(item.Weight, item.Value, item.BagWeight)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestWeightBag2(t *testing.T) {
	var list = []item7{
		{
			Weight:    []int{3, 2, 5, 5},
			Value:     []int{2, 8, 1, 3},
			Times:     []int{2, 2, 2, 2},
			BagWeight: 6,
			expected:  16,
		},
		{
			Weight:    []int{1, 4, 4, 3},
			Value:     []int{2, 1, 3, 1},
			Times:     []int{2, 2, 2, 2},
			BagWeight: 6,
			expected:  7,
		},
		{
			Weight:    []int{2, 3, 3, 1},
			Value:     []int{6, 8, 1, 9},
			Times:     []int{2, 2, 2, 2},
			BagWeight: 6,
			expected:  30,
		},
	}
	for _, item := range list {
		res := weightBag21(item.Weight, item.Value, item.Times, item.BagWeight)
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
			Sli1:     []int{1, 2, 3, 2, 1},
			Sli2:     []int{3, 2, 1, 4, 7},
			expected: 3,
			//示例2：输入：nums1=[0,0,0,0,0],nums2=[0,0,0,0,0]输出：5
		},
		{
			Sli1:     []int{0, 0, 0, 0, 0},
			Sli2:     []int{0, 0, 0, 0, 0},
			expected: 5,
		},
		{
			Sli1:     []int{1, 2, 3},
			Sli2:     []int{1, 2},
			expected: 2,
		},
	}
	var res int
	for _, item := range list {
		res = findLength(item.Sli1, item.Sli2)
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
			sli:      []int{2},
			target:   3,
			expected: -1,
		},
		{
			sli:      []int{1},
			target:   0,
			expected: 0,
		},
		{
			sli:      []int{1, 2, 5},
			target:   11,
			expected: 3,
		},
		{
			sli:      []int{5, 2, 1},
			target:   11,
			expected: 3,
		},
		{
			sli:      []int{1, 2, 3, 4},
			target:   6,
			expected: 2,
		},
		{
			sli:      []int{5, 7, 2, 1, 6},
			target:   6,
			expected: 1,
		},
		{
			sli:      []int{186, 419, 83, 408},
			target:   6249,
			expected: 20,
		},
	}
	var res int
	for _, item := range list {
		res = coinChange(item.sli, item.target)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestChange(t *testing.T) {
	var list = []item10{
		{
			sli:      []int{2},
			target:   3,
			expected: 0,
		},
		{
			sli:      []int{1},
			target:   0,
			expected: 1,
		},
		{
			sli:      []int{1},
			target:   5,
			expected: 1,
		},
		{
			sli:      []int{5},
			target:   10,
			expected: 1,
		},
		{
			sli:      []int{1, 2, 5},
			target:   5,
			expected: 4,
		},
		{
			sli:      []int{5, 2, 1},
			target:   5,
			expected: 4,
		},
		{
			sli:      []int{7},
			target:   0,
			expected: 1,
		},
		{
			sli:      []int{10},
			target:   10,
			expected: 1,
		},
		{
			sli:      []int{7},
			target:   0,
			expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = change1(item.target, item.sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		res = change2(item.target, item.sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.expected, res, item.expected, item)
		t.Log("--------------split----------------split--------------")
	}
}
