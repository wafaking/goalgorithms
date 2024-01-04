package dynamics

import (
	"goalgorithms/common"
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
		t.Logf("res: %t, n:%d, res:%d, Expected:%d", res == expected, n, res, expected)
	}
}

func TestFindContentChildren(t *testing.T) {
	var list = []common.Item11{
		{
			Nums1:    []int{1, 2, 3},
			Nums2:    []int{1, 1},
			Expected: 1,
		},
		{
			Nums1:    []int{1, 2},
			Nums2:    []int{1, 2, 3},
			Expected: 2,
		},
		{
			Nums1:    []int{10, 9, 8, 7},
			Nums2:    []int{5, 6, 7, 8},
			Expected: 2,
		},
	}

	for _, item := range list {
		res := findContentChildren(item.Nums1, item.Nums2)
		t.Logf("%t, res-Expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestClimbStairs(t *testing.T) {
	var list = []common.Item4{
		{
			Num:      2,
			Expected: 2,
		},
		{
			Num:      3,
			Expected: 3,
		},
		{
			Num:      4,
			Expected: 5,
		},
	}

	for _, item := range list {
		res := climbStairs1(item.Num)
		t.Logf("%t, res-Expected: %d:%d", res == item.Expected, res, item.Expected)
		res = climbStairs2(item.Num)
		t.Logf("%t, res-Expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{10, 15, 20},
			Expected: 15,
		},
		{
			Nums:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			Expected: 6,
		},
	}

	for _, item := range list {
		res := minCostClimbingStairs1(item.Nums)
		t.Logf("%t, res-Expected: %d:%d", res == item.Expected, res, item.Expected)
		res = minCostClimbingStairs2(item.Nums)
		t.Logf("%t, res-Expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePaths(t *testing.T) {
	var list = []common.Item9{
		{
			Num1:     2,
			Num2:     2,
			Expected: 2,
		},
		{
			Num1:     3,
			Num2:     7,
			Expected: 28,
		},
		{
			Num1:     3,
			Num2:     2,
			Expected: 3,
		},
		{
			Num1:     7,
			Num2:     3,
			Expected: 28,
		},
		{
			Num1:     3,
			Num2:     3,
			Expected: 6,
		},
		{
			Num1:     1,
			Num2:     2,
			Expected: 1,
		},
		{
			Num1:     2,
			Num2:     1,
			Expected: 1,
		},
	}

	var res int
	for _, item := range list {
		res = uniquePaths1(item.Num1, item.Num2)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = uniquePaths2(item.Num1, item.Num2)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = uniquePaths3(item.Num1, item.Num2)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = uniquePaths4(item.Num1, item.Num2)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	var list = []common.Item12{
		{
			Grid:     [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			Expected: 2,
		},
		{
			Grid:     [][]int{{0, 1}, {0, 0}},
			Expected: 1,
		},
		{
			Grid:     [][]int{{0, 0}},
			Expected: 1,
		},
		{
			Grid:     [][]int{{1, 0}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{0, 1, 0}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{0, 0}, {1, 1}, {0, 0}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{0, 0}, {0, 0}, {0, 1}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0}},
			Expected: 10,
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
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = uniquePathsWithObstacles1(item.Grid)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		res = uniquePathsWithObstacles2(item.Grid)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		res = uniquePathsWithObstacles3(item.Grid)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinPathSum(t *testing.T) {
	var list = []common.Item12{
		{
			Grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			Expected: 7,
		},
		{
			Grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
			Expected: 12,
		},
		{
			Grid:     [][]int{{0, 0}},
			Expected: 0,
		},
		{
			Grid:     [][]int{{1, 0}},
			Expected: 1,
		},
		{
			Grid:     [][]int{{0, 1, 0}},
			Expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = minPathSum1(item.Grid)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestCalculateMinimumHP(t *testing.T) {
	var list = []common.Item12{
		{
			Grid:     [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}},
			Expected: 7,
		},
		{
			Grid:     [][]int{{-2, -1, 3}, {-5, -3, 1}, {10, 30, -5}},
			Expected: 5,
		},
		{
			Grid:     [][]int{{3, 4, 3}, {-2, 6, 1}, {10, 10, -5}},
			Expected: 1,
		},
		{
			Grid:     [][]int{{0}},
			Expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = calculateMinimumHP1(item.Grid)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		res = calculateMinimumHP2(item.Grid)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIntegerBreak(t *testing.T) {
	var list = []common.Item4{
		{
			Num:      2,
			Expected: 1,
		},
		{
			Num:      3,
			Expected: 2,
		},
		{
			Num:      4,
			Expected: 4,
		},
		{
			Num:      5,
			Expected: 6,
		},
		{
			Num:      6,
			Expected: 9,
		},
		{
			Num:      7,
			Expected: 12,
		},
		{
			Num:      8,
			Expected: 18,
		},
		{
			Num:      9,
			Expected: 27,
		},
		{
			Num:      10,
			Expected: 36,
		},
		{
			Num:      11,
			Expected: 54,
		},
		{
			Num:      12,
			Expected: 81,
		},
	}

	for _, item := range list {
		res := integerBreak1(item.Num)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob(t *testing.T) {
	var list = []common.Item2{
		//{Nums: []int{}, Expected: 0},
		//{Nums: []int{1}, Expected: 1},
		//{Nums: []int{3, 1}, Expected: 3},
		//{Nums: []int{1, 2, 4, 3, 0, 9}, Expected: 14},
		//{Nums: []int{1, 2, 3, 1}, Expected: 4},
		{Nums: []int{2, 7, 9, 3, 1}, Expected: 12},
		{Nums: []int{2, 1, 1, 2}, Expected: 4},
	}
	for _, item := range list {
		res := rob11(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob12(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob13(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob2(t *testing.T) {
	var list = []common.Item2{
		{Nums: []int{}, Expected: 0},
		{Nums: []int{1}, Expected: 1},
		{Nums: []int{1, 2}, Expected: 2},
		{Nums: []int{3, 1}, Expected: 3},
		{Nums: []int{2, 3, 2}, Expected: 3},
		{Nums: []int{1, 2, 3}, Expected: 3},
		{Nums: []int{1, 2, 1, 1}, Expected: 3},
		{Nums: []int{1, 2, 3, 1}, Expected: 4},
		{Nums: []int{2, 7, 9, 3, 1}, Expected: 11},
		{Nums: []int{2, 1, 1, 2}, Expected: 3},
		//例1：输入：Nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：Nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：Nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob3(t *testing.T) {
	var list = []common.Item2{
		{Nums: []int{}, Expected: 0},
		{Nums: []int{1}, Expected: 1},
		{Nums: []int{1, 2}, Expected: 2},
		{Nums: []int{3, 1}, Expected: 3},
		{Nums: []int{2, 3, 2}, Expected: 3},
		{Nums: []int{1, 2, 3}, Expected: 3},
		{Nums: []int{1, 2, 1, 1}, Expected: 3},
		{Nums: []int{1, 2, 3, 1}, Expected: 4},
		{Nums: []int{2, 7, 9, 3, 1}, Expected: 11},
		{Nums: []int{2, 1, 1, 2}, Expected: 3},
		//例1：输入：Nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：Nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：Nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxValue(t *testing.T) {
	var list = []common.Item12{
		{
			Grid: [][]int{
				{4, 2, 1},
			},
			Expected: 7,
		},
		{
			Grid: [][]int{
				{1},
				{1},
				{4},
			},
			Expected: 6,
		},
		{
			Grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			Expected: 12,
		},
	}
	for _, item := range list {
		res := maxValue12(item.Grid)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = maxValue11(item.Grid)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}

}

func TestLengthOfLIS11(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			Expected: 4,
		},
		{
			Nums:     []int{0, 1, 0, 3, 2, 3},
			Expected: 4,
		},
		{
			Nums:     []int{7, 7, 7, 7, 7, 7, 7},
			Expected: 1,
		},
	}
	for _, item := range list {
		res := lengthOfLIS11(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS12(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS13(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS14(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS15(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")

	}
}

func TestWeightBag1(t *testing.T) {
	var list = []common.Item13{
		{
			Weight:    []int{3, 2, 5, 5},
			Value:     []int{2, 8, 1, 3},
			BagWeight: 6,
			Expected:  10,
		},
		{
			Weight:    []int{1, 4, 4, 3},
			Value:     []int{2, 1, 3, 1},
			BagWeight: 6,
			Expected:  5,
		},
		{
			Weight:    []int{2, 3, 3, 1},
			Value:     []int{6, 8, 1, 9},
			BagWeight: 6,
			Expected:  23,
		},
		{
			Weight:    []int{3, 3, 2, 2, 5, 5, 5, 5},
			Value:     []int{2, 2, 8, 8, 1, 1, 3, 3},
			BagWeight: 6,
			Expected:  16,
		},
		{
			Weight:    []int{1, 1, 4, 4, 4, 4, 3, 3},
			Value:     []int{2, 2, 1, 1, 3, 3, 1, 1},
			BagWeight: 6,
			Expected:  7,
		},
		{
			Weight:    []int{2, 2, 3, 3, 3, 3, 1, 1},
			Value:     []int{6, 6, 8, 8, 1, 1, 9, 9},
			BagWeight: 6,
			Expected:  30,
		},
	}
	var res int
	for _, item := range list {
		res = weightBag11(item.Weight, item.Value, item.BagWeight)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = weightBag13(item.Weight, item.Value, item.BagWeight)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestWeightBag2(t *testing.T) {
	var list = []common.Item14{
		{
			Weight:    []int{3, 2, 5, 5},
			Value:     []int{2, 8, 1, 3},
			Times:     []int{2, 2, 2, 2},
			BagWeight: 6,
			Expected:  16,
		},
		{
			Weight:    []int{1, 4, 4, 3},
			Value:     []int{2, 1, 3, 1},
			Times:     []int{2, 2, 2, 2},
			BagWeight: 6,
			Expected:  7,
		},
		{
			Weight:    []int{2, 3, 3, 1},
			Value:     []int{6, 8, 1, 9},
			Times:     []int{2, 2, 2, 2},
			BagWeight: 6,
			Expected:  30,
		},
	}
	for _, item := range list {
		res := weightBag21(item.Weight, item.Value, item.Times, item.BagWeight)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsMatchRegx(t *testing.T) {
	var list = []common.Item15{
		{
			S:        "abc",
			P:        ".*c",
			Expected: true,
		},
		{
			S:        "abc",
			P:        "*.c",
			Expected: false,
		},
		{
			S:        "AAAA",
			P:        "BAAC*.",
			Expected: false,
		},
		{
			S:        "ACAB",
			P:        "A.B*AB",
			Expected: true,
		},
		{
			S:        "ABAA",
			P:        "A*CA.C",
			Expected: false,
		},
		{
			S:        "aa",
			P:        "a",
			Expected: false,
		},
		{
			S:        "aa",
			P:        "a*",
			Expected: true,
		},
		{
			S:        "ab",
			P:        ".*",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "..",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "**",
			Expected: false,
		},
		{
			S:        "aab",
			P:        "c*a*b",
			Expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isMatchRegx1(item.S, item.P)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsMatchWildCard(t *testing.T) {
	var list = []common.Item15{
		{
			S:        "abc",
			P:        "?*c",
			Expected: true,
		},
		{
			S:        "abc",
			P:        "*?c",
			Expected: true,
		},
		{
			S:        "CABCB",
			P:        "C?BC*",
			Expected: true,
		},
		{
			S:        "AAAA",
			P:        "BAAC*?",
			Expected: false,
		},
		{
			S:        "ACAB",
			P:        "A?*AB",
			Expected: true,
		},
		{
			S:        "aa",
			P:        "a",
			Expected: false,
		},
		{
			S:        "aa",
			P:        "*",
			Expected: true,
		},
		{
			S:        "aa",
			P:        "a*",
			Expected: true,
		},
		{
			S:        "b",
			P:        "?*",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "?*",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "??",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "**",
			Expected: true,
		},
		{
			S:        "ab",
			P:        "*",
			Expected: true,
		},
		{
			S:        "aab",
			P:        "c*a*b",
			Expected: false,
		},
		{
			S:        "AC",
			P:        "?*C",
			Expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isMatchWildCard(item.S, item.P)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	var list = []common.Item16{
		{
			Text1:    "AAAT",
			Text2:    "ATACATAC",
			Expected: 4,
		},
		{
			Text1:    "abcde",
			Text2:    "ace",
			Expected: 3,
		},
		{
			Text1:    "abc",
			Text2:    "abc",
			Expected: 3,
		},
		{
			Text1:    "abc",
			Text2:    "def",
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = longestCommonSubsequence(item.Text1, item.Text2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestLongestCommonSubstring(t *testing.T) {
	var list = []common.Item16{
		{
			Text1:    "AAAT",
			Text2:    "ATACATAC",
			Expected: 2,
		},
		{
			Text1:    "bacde",
			Text2:    "ace",
			Expected: 2,
		},
		{
			Text1:    "abc",
			Text2:    "abc",
			Expected: 3,
		},
		{
			Text1:    "abc",
			Text2:    "def",
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = longestCommonSubstring(item.Text1, item.Text2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestFindLength(t *testing.T) {
	var list = []common.Item11{
		{
			Nums1:    []int{1, 2, 3, 2, 1},
			Nums2:    []int{3, 2, 1, 4, 7},
			Expected: 3,
			//示例2：输入：nums1=[0,0,0,0,0],nums2=[0,0,0,0,0]输出：5
		},
		{
			Nums1:    []int{0, 0, 0, 0, 0},
			Nums2:    []int{0, 0, 0, 0, 0},
			Expected: 5,
		},
		{
			Nums1:    []int{1, 2, 3},
			Nums2:    []int{1, 2},
			Expected: 2,
		},
	}
	var res int
	for _, item := range list {
		res = findLength(item.Nums1, item.Nums2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinDistance(t *testing.T) {
	var list = []common.Item16{
		{
			Text1:    "TAAGC",
			Text2:    "GAACT",
			Expected: 3,
		},
		{
			Text1:    "horse",
			Text2:    "ros",
			Expected: 3,
		},
		{
			Text1:    "intention",
			Text2:    "execution",
			Expected: 5,
		},
	}
	var res int
	for _, item := range list {
		res = minDistance(item.Text1, item.Text2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestCoinChange(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{2},
			Target:   3,
			Expected: -1,
		},
		{
			Nums:     []int{1},
			Target:   0,
			Expected: 0,
		},
		{
			Nums:     []int{1, 2, 5},
			Target:   11,
			Expected: 3,
		},
		{
			Nums:     []int{5, 2, 1},
			Target:   11,
			Expected: 3,
		},
		{
			Nums:     []int{1, 2, 3, 4},
			Target:   6,
			Expected: 2,
		},
		{
			Nums:     []int{5, 7, 2, 1, 6},
			Target:   6,
			Expected: 1,
		},
		{
			Nums:     []int{186, 419, 83, 408},
			Target:   6249,
			Expected: 20,
		},
	}
	var res int
	for _, item := range list {
		res = coinChange(item.Nums, item.Target)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestChange(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{2},
			Target:   3,
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Target:   0,
			Expected: 1,
		},
		{
			Nums:     []int{1},
			Target:   5,
			Expected: 1,
		},
		{
			Nums:     []int{5},
			Target:   10,
			Expected: 1,
		},
		{
			Nums:     []int{1, 2, 5},
			Target:   5,
			Expected: 4,
		},
		{
			Nums:     []int{5, 2, 1},
			Target:   5,
			Expected: 4,
		},
		{
			Nums:     []int{7},
			Target:   0,
			Expected: 1,
		},
		{
			Nums:     []int{10},
			Target:   10,
			Expected: 1,
		},
		{
			Nums:     []int{7},
			Target:   0,
			Expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = change1(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = change2(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestCanPartition(t *testing.T) {
	var list = []common.Item8{
		{
			Nums:     []int{1, 0, 1},
			Expected: true,
		},
		{
			Nums:     []int{1, 2, 3, 2},
			Expected: true,
		},
		{
			Nums:     []int{1, 3, 5, 3},
			Expected: true,
		},
		{
			Nums:     []int{1, 5, 11, 5},
			Expected: true,
		},
		{
			Nums:     []int{1, 2, 3, 5},
			Expected: false,
		},
		{
			Nums:     []int{1, 6},
			Expected: false,
		},
		{
			Nums:     []int{2},
			Expected: false,
		},
		{
			Nums:     []int{1, 2, 2},
			Expected: false,
		},
		{
			Nums:     []int{1, 2, 5},
			Expected: false,
		},
		{
			Nums: []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
				100, 100, 100, 100, 100, 100, 99, 97},
			Expected: false,
		},
		{
			Nums: []int{
				66, 90, 7, 6, 32, 16, 2, 78, 69, 88, 85, 26, 3, 9, 58, 65, 30, 96, 11, 31, 99, 49, 63, 83, 79, 97, 20, 64, 81, 80, 25, 69, 9, 75, 23, 70, 26, 71, 25, 54, 1, 40, 41, 82, 32, 10, 26, 33, 50, 71, 5, 91, 59, 96, 9, 15, 46, 70, 26, 32, 49, 35, 80, 21, 34, 95, 51, 66, 17, 71, 28, 88, 46, 21, 31, 71, 42, 2, 98, 96, 40, 65, 92, 43, 68, 14, 98, 38, 13, 77, 14, 13, 60, 79, 52, 46, 9, 13, 25, 8,
			},
			Expected: true,
		},
	}

	var res bool
	for _, item := range list {
		res = canPartition1(item.Nums)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = canPartition2(item.Nums)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = canPartition3(item.Nums)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = canPartition4(item.Nums)
		common.PrintDiffTwoBool(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}

func TestFindTargetSumWays(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{1, 1, 1, 1, 1},
			Target:   3,
			Expected: 5,
		},
		{
			Nums:     []int{1},
			Target:   1,
			Expected: 1,
		},
		{
			Nums:     []int{1},
			Target:   2,
			Expected: 0,
		},
		{
			Nums:     []int{1, 0},
			Target:   1,
			Expected: 2,
		},
		{
			Nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
			Target:   1,
			Expected: 256,
		},
		{
			Nums:     []int{1, 2, 1},
			Target:   0,
			Expected: 2,
		},
	}

	var res int
	for _, item := range list {
		res = findTargetSumWays11(item.Nums, item.Target)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = findTargetSumWays12(item.Nums, item.Target)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = findTargetSumWays13(item.Nums, item.Target)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = findTargetSumWays14(item.Nums, item.Target)
		common.PrintDiffTwoInt(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}

func TestFindMaxForm(t *testing.T) {
	var list = []common.Item10{
		{
			Str: []string{"10", "0001", "111001", "1", "0"},
			Item9: common.Item9{
				Num1:     5,
				Num2:     3,
				Expected: 4,
			},
		},
		{
			Str: []string{"10", "0", "1"},
			Item9: common.Item9{
				Num1:     1,
				Num2:     1,
				Expected: 2,
			},
		},
	}

	var res int
	for _, item := range list {
		res = findMaxForm1(item.Str, item.Num1, item.Num2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsSubsequence(t *testing.T) {
	var list = []common.Item15{
		{
			S:        "abc",
			P:        "ahbgdc",
			Expected: true,
		},
		{
			S:        "axc",
			P:        "ahbgdc",
			Expected: false,
		},
		{
			S:        "",
			P:        "ahbgdc",
			Expected: true,
		},
		{
			S:        "aaaaaa",
			P:        "bbaaaa",
			Expected: false,
		},
		{
			S:        "bb",
			P:        "ahbgdc",
			Expected: false,
		},
	}
	var res bool
	for _, item := range list {
		res = isSubsequence1(item.S, item.P)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		res = isSubsequence3(item.S, item.P)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		res = isSubsequence4(item.S, item.P)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIsSubstring(t *testing.T) {
	var list = []common.Item15{
		{
			S:        "abc",
			P:        "ahbgdc",
			Expected: false,
		},
		{
			S:        "hbg",
			P:        "ahbgdc",
			Expected: true,
		},
		{
			S:        "hbg",
			P:        "hbg",
			Expected: true,
		},
		{
			S:        "axc",
			P:        "ahbgdc",
			Expected: false,
		},
		{
			S:        "",
			P:        "ahbgdc",
			Expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isSubstring1(item.S, item.P)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestTwoEggDrop(t *testing.T) {
	var list = []common.Item4{
		{
			Num:      1,
			Expected: 1,
		},
		{
			Num:      2,
			Expected: 2,
		},
		{
			Num:      3,
			Expected: 2,
		},
		{
			Num:      4,
			Expected: 3,
		},
		{
			Num:      5,
			Expected: 3,
		},
		{
			Num:      100,
			Expected: 14,
		},
	}
	var res int
	for _, item := range list {
		res = twoEggDrop1(item.Num)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = twoEggDrop2(item.Num)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = twoEggDrop3(item.Num)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = twoEggDrop4(item.Num)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestSuperEggDrop(t *testing.T) {
	var list = []common.Item9{
		{
			Num1:     1,
			Num2:     2,
			Expected: 2,
		},
		{
			Num1:     3,
			Num2:     7,
			Expected: 3,
		},
		{
			Num1:     2,
			Num2:     6,
			Expected: 3,
		},
		{
			Num1:     3,
			Num2:     14,
			Expected: 4,
		},
		{
			Num1:     7,
			Num2:     10000,
			Expected: 15,
		},
	}
	var res int
	for _, item := range list {
		res = superEggDrop1(item.Num1, item.Num2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = superEggDrop2(item.Num1, item.Num2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = superEggDrop3(item.Num1, item.Num2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = superEggDrop4(item.Num1, item.Num2)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestLastStoneWeight1(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{2, 7, 4, 1, 8, 1},
			Expected: 1,
		},
		{
			Nums:     []int{1},
			Expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = lastStoneWeight1(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit1(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{7, 1, 5, 3, 6, 4},
			Expected: 5,
		},
		{
			Nums:     []int{7, 6, 4, 3, 1},
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = maxProfit11(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit12(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit13(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit2(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{7, 1, 5, 3, 6, 4},
			Expected: 7,
		},
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: 4,
		},
		{
			Nums:     []int{7, 6, 4, 3, 1},
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = maxProfit21(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit22(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		//res = maxProfit13(item.Nums)
		//t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit3(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{3, 3, 5, 0, 0, 3, 1, 4},
			Expected: 6,
		},
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: 4,
		},
		{
			Nums:     []int{7, 6, 4, 3, 1},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = maxProfit31(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit32(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit4(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{2, 4, 1},
			Target:   2,
			Expected: 2,
		},
		{
			Nums:     []int{3, 2, 6, 5, 0, 3},
			Target:   2,
			Expected: 7,
		},
		{
			Nums:     []int{2, 4, 1, 8, 2, 5, 9, 11},
			Target:   3,
			Expected: 18,
		},
	}

	var res int
	for _, item := range list {
		res = maxProfit41(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit42(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit5(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 2, 3, 0, 2},
			Expected: 3,
		},
		{
			Nums:     []int{1},
			Expected: 0,
		},
	}

	var res int
	for _, item := range list {
		res = maxProfit51(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit52(item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxProfit6(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{1, 3, 2, 8, 4, 9},
			Target:   2,
			Expected: 8,
		},
		{
			Nums:     []int{1, 3, 7, 5, 10, 3},
			Target:   3,
			Expected: 6,
		},
	}

	var res int
	for _, item := range list {
		res = maxProfit61(item.Nums, item.Target)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = maxProfit62(item.Nums, item.Target)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}
