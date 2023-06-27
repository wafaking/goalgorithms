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
			Expected: 1,
		},
		{
			Sli1:     []int{1, 2},
			Sli2:     []int{1, 2, 3},
			Expected: 2,
		},
		{
			Sli1:     []int{10, 9, 8, 7},
			Sli2:     []int{5, 6, 7, 8},
			Expected: 2,
		},
	}

	for _, item := range list {
		res := findContentChildren(item.Sli1, item.Sli2)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestClimbStairs(t *testing.T) {
	var list = []item2{
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
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		res = climbStairs2(item.Num)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	var list = []item3{
		{
			Sli:      []int{10, 15, 20},
			Expected: 15,
		},
		{
			Sli:      []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			Expected: 6,
		},
	}

	for _, item := range list {
		res := minCostClimbingStairs1(item.Sli)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		res = minCostClimbingStairs2(item.Sli)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePaths(t *testing.T) {
	var list = []item4{
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

	for _, item := range list {
		res := uniquePaths1(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = uniquePaths2(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		res = uniquePaths3(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		res = uniquePaths4(item.Num1, item.Num2)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	var list = []item5{
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
			Grid:     [][]int{{0}, {0}},
			Expected: 1,
		},
		{
			Grid:     [][]int{{1}, {0}},
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

	for _, item := range list {
		res := uniquePathsWithObstacles1(item.Grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = uniquePathsWithObstacles2(item.Grid)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		res = uniquePathsWithObstacles3(item.Grid)
		t.Logf("%t, res-expected: %d:%d", res == item.Expected, res, item.Expected)
		t.Log("--------------split----------------split--------------")
	}
}

func TestIntegerBreak(t *testing.T) {
	var list = []item2{
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
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob(t *testing.T) {
	var list = []item3{
		//{Sli: []int{}, Expected: 0},
		//{Sli: []int{1}, Expected: 1},
		//{Sli: []int{3, 1}, Expected: 3},
		//{Sli: []int{1, 2, 4, 3, 0, 9}, Expected: 14},
		//{Sli: []int{1, 2, 3, 1}, Expected: 4},
		{Sli: []int{2, 7, 9, 3, 1}, Expected: 12},
		{Sli: []int{2, 1, 1, 2}, Expected: 4},
	}
	for _, item := range list {
		res := rob11(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob12(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = rob13(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob2(t *testing.T) {
	var list = []item3{
		{Sli: []int{}, Expected: 0},
		{Sli: []int{1}, Expected: 1},
		{Sli: []int{1, 2}, Expected: 2},
		{Sli: []int{3, 1}, Expected: 3},
		{Sli: []int{2, 3, 2}, Expected: 3},
		{Sli: []int{1, 2, 3}, Expected: 3},
		{Sli: []int{1, 2, 1, 1}, Expected: 3},
		{Sli: []int{1, 2, 3, 1}, Expected: 4},
		{Sli: []int{2, 7, 9, 3, 1}, Expected: 11},
		{Sli: []int{2, 1, 1, 2}, Expected: 3},
		//例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestRob3(t *testing.T) {
	var list = []item3{
		{Sli: []int{}, Expected: 0},
		{Sli: []int{1}, Expected: 1},
		{Sli: []int{1, 2}, Expected: 2},
		{Sli: []int{3, 1}, Expected: 3},
		{Sli: []int{2, 3, 2}, Expected: 3},
		{Sli: []int{1, 2, 3}, Expected: 3},
		{Sli: []int{1, 2, 1, 1}, Expected: 3},
		{Sli: []int{1, 2, 3, 1}, Expected: 4},
		{Sli: []int{2, 7, 9, 3, 1}, Expected: 11},
		{Sli: []int{2, 1, 1, 2}, Expected: 3},
		//例1：输入：nums=[2,3,2], 输出：3,解释：你不能先偷窃1号房屋（金额=2），然后偷窃3号房屋（金额=2）,因为他们是相邻的。
		//示例2：输入：nums=[1,2,3,1], 输出：4,解释：你可以先偷窃1号房屋（金额=1），然后偷窃3号房屋(金额=3)偷窃到的最高金额=1+3=4。
		//示例3：输入：nums=[1,2,3],输出：3;

	}
	for _, item := range list {
		res := rob21(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMaxValue(t *testing.T) {
	var list = []item5{
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
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = maxValue11(item.Grid)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}

}

func TestLengthOfLIS11(t *testing.T) {
	var list = []item3{
		{
			Sli:      []int{10, 9, 2, 5, 3, 7, 101, 18},
			Expected: 4,
		},
		{
			Sli:      []int{0, 1, 0, 3, 2, 3},
			Expected: 4,
		},
		{
			Sli:      []int{7, 7, 7, 7, 7, 7, 7},
			Expected: 1,
		},
	}
	for _, item := range list {
		res := lengthOfLIS11(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS12(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS13(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS14(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
		res = lengthOfLIS15(item.Sli)
		t.Logf("%t, res-expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")

	}
}
