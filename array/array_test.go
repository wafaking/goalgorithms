package array

import (
	"goalgorithms/common"
	"sort"
	"testing"
)

func TestMain(t *testing.M) {
	//sli := []int{5, 3, 7, 2, 0, 4, 1, 2}
	//sli := []int{4, 7, 5, 9, 2, 3}
	//sli := []int{2, 3, 1, 0, 2, 5, 3}
	//sli := []int{2, 7, 11, 15}
	//sli := []int{3, 2, 4}
	//sli := []int{3, 0, -2, -1, -1, -1, 1, 2}
	//sli := []int{3, 2, 5, 4, 7, 2, 6, 5}
	//sli := []int{1, 2, 2, 3, 4, 5, 6, 7}
	//sli := []int{1, 2, 3, 4, 4}

	//InitNums(sli)
	//PrintNums()

	//Init2DNums()
	//Println2DNums()
	t.Run()
}

func TestTwoSum(t *testing.T) {
	var samples = []common.Item1{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, v := range samples {
		res := twoSum11(v.Nums, v.Target)
		if len(res) != len(v.Expected) {
			t.Failed()
		}
		for i := 0; i < len(res); i++ {
			if res[i] != v.Expected[i] {
				t.Fatal()
			}
		}
	}
	t.Log("------------SPLIT-------------")
	for _, v := range samples {
		res := twoSum12(v.Nums, v.Target)
		if len(res) != len(v.Expected) {
			t.Failed()
		}
		for i := 0; i < len(res); i++ {
			if res[i] != v.Expected[i] {
				t.Fatal()
			}
		}
	}
}

func TestTwoSum21(t *testing.T) {
	var list = []common.Item1{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	var res = make([]int, 0, 2)
	for _, item := range list {
		res = twoSum21(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		res = twoSum22(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("------------------SPLIT-----------------------")
	}
}

func TestFindRepeatNumber(t *testing.T) {
	var list = []common.Item1{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 0, nil},
		{[]int{3, 4, 2, 1, 1, 0}, 0, nil},
	}
	for _, v := range list {
		res := findRepeatNumber1(v.Nums)
		t.Log("res: ", res)
	}
	t.Log("-----------SPLIT---------")
	for _, v := range list {
		res := findRepeatNumber2(v.Nums)
		t.Log("res: ", res)
	}
	t.Log("-----------SPLIT---------")
	for _, v := range list {
		res := findRepeatNumber3(v.Nums)
		t.Log("res: ", res)
	}
	t.Log("-----------SPLIT---------")
	for _, v := range list {
		res := findRepeatNumber4(v.Nums)
		t.Log("res: ", res)
	}
}

func TestGetDuplications(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{3, 0, -2, -1, -1, -1, 1, 2},
			Expected: -2,
		},
	}
	for _, item := range list {
		res := getDuplications(item.Nums, len(item.Nums))
		t.Log("res: ", res)
	}
}

func TestThreeSum(t *testing.T) {
	var samples = [][]int{
		//{0, 1, 1},
		//{0, 0, 0},
		//{-1, 0, 1, 2, -1, -4},
		{-4, -1, -1, -1, -1, 0, 1, 2, 2},
	}
	for _, nums := range samples {
		res := threeSum1(nums)
		t.Log("res: ", res)
	}
	t.Log("----------SPLIT------------")
	for _, nums := range samples {
		res := threeSum2(nums)
		t.Log("res: ", res)
	}
}

func TestThreeSumClosest(t *testing.T) {
	var list = []common.Item3{
		{Nums: []int{-1, 2, 1, -4}, Target: 1, Expected: 2},
		{Nums: []int{0, 0, 0}, Target: 1, Expected: 0},
		{Nums: []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, Target: -1, Expected: -1},
	}
	for _, item := range list {
		res := threeSumClosest1(item.Nums, item.Target)
		t.Logf("%t, res: %d, Expected:%d", res == item.Expected, res, item.Expected)
		res = threeSumClosest2(item.Nums, item.Target)
		t.Logf("%t, res: %d, Expected:%d", res == item.Expected, res, item.Expected)
		t.Log("----------SPLIT------------")
	}
}

func TestFourSumCount(t *testing.T) {
	var list = [][4][]int{
		{
			{1, 2},
			{-2, -1},
			{-1, 2},
			{0, 2},
		},
		{
			{0},
			{0},
			{0},
			{0},
		},
	}
	for _, nums := range list {
		res := fourSumCount11(nums[0], nums[1], nums[2], nums[3])
		t.Log("res: ", res)
		t.Log("----------SPLIT------------")
		res = fourSumCount12(nums[0], nums[1], nums[2], nums[3])
		t.Log("res: ", res)
		t.Log("----------SPLIT------------")
	}
}

func TestFourSum(t *testing.T) {
	var samples = []common.Item1{
		//{[]int{2, 2, 2, 2, 2}, 8, nil},
		//{[]int{1, 0, -1, 0, -2, 2}, 0, nil},
		{[]int{-4, -1, -1, -1, -1, 0, 1, 2, 2, 3}, 0, nil},
	}
	// 示例1： 输入：nums=[1,0,-1,0,-2,2], target=0, 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	// 示例2： 输入：nums = [2,2,2,2,2], target=8, 输出：[[2,2,2,2]]

	for _, item := range samples {
		res := fourSum(item.Nums, item.Target)
		t.Log("res: ", res)
	}
	t.Log("----------SPLIT------------")
}

func TestSearchMatrix1(t *testing.T) {
	var list = []common.Item27{
		{
			Grid: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			Target:   5,
			Expected: true,
		},
		{
			Grid: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			Target:   20,
			Expected: false,
		},
	}
	var res bool
	for _, item := range list {
		res = searchMatrix21(item.Grid, item.Target)
		t.Logf("%t, res:%t, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestSearchMatrix2(t *testing.T) {
	var list = []common.Item27{
		{
			Grid: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			Target:   5,
			Expected: true,
		},
		{
			Grid: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			Target:   20,
			Expected: false,
		},
		{
			Grid: [][]int{
				{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
			},
			Target:   13,
			Expected: false,
		},
		{
			Grid: [][]int{
				{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
			},
			Target:   3,
			Expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = searchMatrix11(item.Grid, item.Target)
		t.Logf("%t, res:%t, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestMerge(t *testing.T) {
	var list = []common.Item43{
		{

			Nums1:    []int{1, 2, 3, 0, 0, 0},
			Nums2:    []int{2, 5, 6},
			N1:       3,
			N2:       3,
			Expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			Nums1:    []int{0},
			Nums2:    []int{1},
			N1:       0,
			N2:       1,
			Expected: []int{1},
		},
		{
			Nums1:    []int{1},
			Nums2:    []int{},
			N1:       1,
			N2:       0,
			Expected: []int{1},
		},
		{
			Nums1:    []int{4, 5, 6, 0, 0, 0},
			Nums2:    []int{1, 2, 3},
			N1:       3,
			N2:       3,
			Expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, item := range list {
		temp := make([]int, len(item.Nums1))
		copy(temp, item.Nums1)
		merge1(temp, item.N1, item.Nums2, item.N2)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(temp, item.Expected), temp, item)
		copy(temp, item.Nums1)
		merge2(temp, item.N1, item.Nums2, item.N2)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(temp, item.Expected), temp, item)
		copy(temp, item.Nums1)
		merge3(temp, item.N1, item.Nums2, item.N2)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(temp, item.Expected), temp, item)
		t.Log("------------------------SPLIT------------------------")
	}

}

func TestTotalDestinationPath(t *testing.T) {
	res := TotalDestinationPath(1, 1)
	t.Log("res: ", res)
}

func TestMinArray(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{},
			Expected: -1,
		},
		{
			Nums:     []int{2},
			Expected: 2,
		},
		{
			Nums:     []int{3, 2},
			Expected: 2,
		},
		{
			Nums:     []int{3, 4, 5, 1, 2},
			Expected: 1,
		},
		{
			Nums:     []int{2, 2, 2, 0, 1},
			Expected: 0,
		},
		{
			Nums:     []int{3, 4, 5, 6, 1, 2},
			Expected: 1,
		},
		{
			Nums:     []int{2, 2, 2, 2, 2, 2, 1, 2, 2},
			Expected: 1,
		},
	}
	var res int
	for _, item := range list {
		res = minArray(item.Nums)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}

func TestMovingCount(t *testing.T) {
	var list = []common.Item45{
		{
			Num1:     2,
			Num2:     3,
			Target:   1,
			Expected: 3,
		},
		{
			Num1:     3,
			Num2:     1,
			Target:   0,
			Expected: 1,
		},
		{
			Num1:     3,
			Num2:     2,
			Target:   17,
			Expected: 6,
		},
		{
			Num1:     16,
			Num2:     7,
			Target:   17,
			Expected: 112,
		},
		{
			Num1:     13,
			Num2:     15,
			Target:   20,
			Expected: 195,
		},
	}

	var res int
	for _, item := range list {
		res = movingCount1(item.Num1, item.Num2, item.Target) // Expected: 15
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = movingCount2(item.Num1, item.Num2, item.Target) // Expected: 15
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = movingCount3(item.Num1, item.Num2, item.Target) // Expected: 15
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}

func TestGetWordSum(t *testing.T) {
	res := getWordSum(10)
	t.Logf("res: %d", res)
}

func TestMaxArea(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expected: 49,
		},
		{
			Nums:     []int{1, 1},
			Expected: 1,
		},
		{
			Nums:     []int{8, 1, 6, 2, 5, 8, 4, 3, 7},
			Expected: 56,
		},
		{
			Nums:     []int{1, 6, 2, 5, 8, 4, 8, 3, 7},
			Expected: 42,
		},
	}
	var res int
	for _, item := range list {
		res = maxArea1(item.Nums)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		res = maxArea2(item.Nums)
		common.PrintDiffTwoInt(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}

func TestPermute(t *testing.T) {
	var numsList = [][]int{
		//{},
		//{1},
		//{0, 1},
		{1, 2, 3},
	}
	for _, nums := range numsList {
		res := permute11(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute12(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute21(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute22(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		res = permute3(nums)
		t.Logf("res: %v\n", res)
		t.Log("-------------SPLIT----------")
		//res = permute32(nums)
		//t.Logf("res: %v\n", res)
		//t.Log("-------------SPLIT----------")
	}
}

func TestPermuteUnique(t *testing.T) {
	var numsList = [][]int{
		//{},
		//{1},
		//{0, 1},
		//{1, 2, 3},
		{1, 1, 2},
		//{1, 2, 1},
		//{1, 1, 1},
		{1, 2, 1, 1},
		{-1, 2, -1, 2, 1, -1, 2, 1},
		{0, 1, 0, 0, 9},
	}
	for _, nums := range numsList {
		res := permuteUnique11(nums)
		t.Logf("res: %v\n", res)
	}
	t.Log("----------SPLIT------------")
	for _, nums := range numsList {
		//res := permuteUnique2(nums)
		res := permuteUnique2(nums)
		t.Logf("res: %v\n", res)
	}
}

func TestCombine(t *testing.T) {
	var list = []common.Item30{
		{
			Num:      4,
			Target:   2,
			Expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			Num:      1,
			Target:   1,
			Expected: [][]int{{1}},
		},
		{
			Num:      4,
			Target:   3,
			Expected: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}},
		},
	}
	var res = make([][]int, 0)
	for _, item := range list {
		res = combine1(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combine2(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combine3(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combine4(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combine5(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combine6(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------------SPLIT---------------------------")
	}
}

func TestCombinationSum1(t *testing.T) {
	var list = []common.Item26{
		{[]int{2}, 1, [][]int{}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
	}

	var res = make([][]int, 0)
	for _, item := range list {
		res = combinationSum11(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum12(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------------SPLIT---------------------------")
	}
}

func TestCombinationSum2(t *testing.T) {
	var list = []common.Item26{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		// 示例1: 输入: candidates=[10,1,2,7,6,1,5], target=8, 输出: {{1,1,6},{1,2,5},{1,7},{2,6}}
		// 示例2: 输入: candidates=[2,5,2,1,2], target=5, 输出: [[1,2,2],[5]]
	}
	var res = make([][]int, 0)
	for _, item := range list {
		res = combinationSum21(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum22(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum23(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum24(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------------SPLIT---------------------------")
	}
}

func TestCombinationSum3(t *testing.T) {
	var list = []common.Item30{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{4, 1, [][]int{}},
	}

	var res = make([][]int, 0)
	for _, item := range list {
		res = combinationSum31(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum32(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum33(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		res = combinationSum34(item.Num, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------------SPLIT---------------------------")
	}
}

func TestCombinationSum4(t *testing.T) {
	var list = []common.Item3{
		{[]int{1, 2, 3}, 7, 44},
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
	}

	var res int
	for _, item := range list {
		res = combinationSum41(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", res == item.Expected, res, item.Expected, item)
		res = combinationSum42(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", res == item.Expected, res, item.Expected, item)
		res = combinationSum43(item.Nums, item.Target)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------------SPLIT---------------------------")
	}
}

func TestFindDuplicate(t *testing.T) {
	var list = []common.Item2{
		{[]int{1, 2, 2, 3}, 2},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 2, 3, 3, 3, 5, 4}, 3},
	}
	var res int
	for _, item := range list {
		res = findDuplicate1(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findDuplicate21(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findDuplicate22(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findDuplicate3(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findDuplicate4(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestSingleNumber(t *testing.T) {
	var list = []common.Item2{
		{[]int{1, 2, 2}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for _, item := range list {
		res := singleNumber1(item.Nums)
		t.Logf("res: %t\n", res == item.Expected)
	}
	t.Log("----------SPLIT------------")
}

func TestSingleNumber2(t *testing.T) {
	var list = []common.Item2{
		{[]int{6, 3, 6, 6}, 3},
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		{[]int{-2, -2, 1, -2, 1, 4, 1}, 4},
		{[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}, -4},
		{[]int{-2, -2, -2, -4}, -4},
		{[]int{1, 1, 1, -4}, -4},
	}
	var res int
	for _, item := range list {
		res = singleNumber21(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = singleNumber22(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = singleNumber23(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = singleNumber24(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = singleNumber25(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestSingleNumber3(t *testing.T) {
	var list = []common.Item20{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{2, 2, 3, 3, 4, 5}, []int{4, 5}},
		{[]int{-2, -2, -3, -4, 1, 1}, []int{-3, -4}},
		{[]int{-1, 0}, []int{-1, 0}},
		{[]int{0, 1}, []int{0, 1}},
	}
	var res = make([]int, 0, 2)
	for _, item := range list {
		sort.Ints(item.Expected)
		res = singleNumber31(item.Nums)
		sort.Ints(res)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		res = singleNumber32(item.Nums)
		sort.Ints(res)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestSubsets1(t *testing.T) {
	var list = []common.Item1{
		{[]int{1, 2, 3}, 0, nil},
		{[]int{0}, 0, nil},
		// 示例1：输入：nums=[1,2,3], 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
		// 示例2：输入：nums = [0], 输出：[[],[0]]
	}
	for _, item := range list {
		res := subsets1(item.Nums)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = subsets2(item.Nums)
		t.Logf("res: %v\n", res)
	}
}

func TestKthLargest(t *testing.T) {
	var nums = []int{4, 5, 8, 2}
	var list = []common.Item4{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]], 输出：[null, 4, 5, 5, 8, 8]
	// [4, 5, 8, 2]的第3大元素：4
	// [4, 5, 8, 2, 3]---> 4
	// [4, 5, 8, 2, 3, 5]---> 5
	// [4, 5, 8, 2, 3, 5, 10]---> 5
	// [4, 5, 8, 2, 3, 5, 10, 9]---> 8
	// [4, 5, 8, 2, 3, 5, 10, 9, 4]---> 8
	var kth = Constructor(3, nums)
	for _, item := range list {
		res := kth.Add(item.Num)
		t.Logf("%t, Expected: %d, res:%d", res == item.Expected, item.Expected, res)
	}
}

func TestFindKthLargest(t *testing.T) {
	var list = []common.Item3{
		{Nums: []int{6, 8, 10, 7, 4, 5, 9, 2, 3, 1}, Target: 3, Expected: 8},
		{Nums: []int{3, 2, 1, 5, 6, 4}, Target: 2, Expected: 5},
		{Nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, Target: 4, Expected: 4},
	}

	for _, item := range list {
		var nums = item.Nums
		res := findKthLargest1(item.Nums, item.Target)
		if res != item.Expected {
			t.Fatalf("res: %d, origin:%+v", res, item)
		}

		t.Log("-------SPLIT-----------SPLIT------SPLIT----------")

		nums = item.Nums
		res = findKthLargest2(nums, item.Target)
		if res != item.Expected {
			t.Fatalf("res: %d, origin:%+v", res, item)
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	var list = []common.Item5{
		{
			Num: 1, Expected: [][]string{{"Q"}},
		},
		{
			Num: 4, Expected: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
	}
	for _, item := range list {
		res := solveNQueens1(item.Num)
		t.Logf("res: %v\n, expect:%v", res, item.Expected)
		t.Log("----------SPLIT------------")
	}
}

func TestSortColors(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{1, 0},
			Expected: []int{0, 1},
		},
		{
			Nums:     []int{2, 0},
			Expected: []int{0, 2},
		},
		{
			Nums:     []int{2, 0, 1},
			Expected: []int{0, 1, 2},
		},
		{
			Nums:     []int{2, 0, 2, 1, 1, 0},
			Expected: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, item := range list {
		var res = make([]int, len(item.Nums))
		copy(res, item.Nums)
		sortColors1(res)
		t.Logf("%v, res:%+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		copy(res, item.Nums)
		sortColors2(res)
		t.Logf("%v, res:%+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		copy(res, item.Nums)
		sortColors3(res)
		t.Logf("%v, res:%+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------")
	}
}

func TestSortArray(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{1, 3, 2},
			Expected: []int{1, 2, 3},
		},
		{
			Nums:     []int{2, 1, 4, 3, 7, 8, 5, 6},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			Nums:     []int{2, 1, 4, 3, 7, 8, 5, 9, 6},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			Nums:     []int{2, 0, 2, 1, 1, 0},
			Expected: []int{0, 0, 1, 1, 2, 2},
		},
	}
	var res []int
	for _, item := range list {
		res = make([]int, len(item.Nums))
		copy(res, item.Nums)
		res = sortArray1(res)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)

		copy(res, item.Nums)
		res = sortArray2(res)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)

		copy(res, item.Nums)
		res = sortArray3(res)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------")
	}
}

func TestConstructorLRU(t *testing.T) {

	var (
		lRUCache LRUCache
		res      int
	)

	/*
		lRUCache = ConstructorLRU(2)
		lRUCache.Put(1, 1)
		lRUCache.Put(2, 2)
		res = lRUCache.Get(1)
		t.Logf("%t, res:expect==%d:%d", res == 1, res, 1)
		lRUCache.Put(3, 3)

		res = lRUCache.Get(2)
		t.Logf("%t, res:expect==%d:%d", res == -1, res, -1)

		lRUCache.Put(4, 4)

		res = lRUCache.Get(1)
		t.Logf("%t, res:expect==%d:%d", res == -1, res, -1)

		res = lRUCache.Get(3)
		t.Logf("%t, res:expect==%d:%d", res == 3, res, 3)

		res = lRUCache.Get(4)
		t.Logf("%t, res:expect==%d:%d", res == 4, res, 4)
	*/

	t.Log("--------------------------SPLIT-----------------------------")
	//["LRUCache","put","put","put","put", "get","get","get","get", "put", "get","get","get","get","get"]
	//[[3],[1,1],[2,2],[3,3],[4,4], [4],[3],[2],[1], [5,5], [1],[2],[3],[4],[5]]
	//
	//添加到测试用例
	//输出
	//[null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
	//预期结果
	//[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]

	lRUCache = ConstructorLRU(3)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	lRUCache.Put(3, 3)
	lRUCache.Put(4, 4)

	//var res int
	res = lRUCache.Get(4)
	t.Logf("%t, res:expect==%d:%d", res == 4, res, 4)
	res = lRUCache.Get(3)
	t.Logf("%t, res:expect==%d:%d", res == 3, res, 3)
	res = lRUCache.Get(2)
	t.Logf("%t, res:expect==%d:%d", res == 2, res, 2)
	res = lRUCache.Get(1)
	t.Logf("%t, res:expect==%d:%d", res == -1, res, -1)

	lRUCache.Put(5, 5)

	res = lRUCache.Get(1)
	t.Logf("%t, res:expect==%d:%d", res == -1, res, -1)
	res = lRUCache.Get(2)
	t.Logf("%t, res:expect==%d:%d", res == 2, res, 2)
	res = lRUCache.Get(3)
	t.Logf("%t, res:expect==%d:%d", res == 3, res, 3)
	res = lRUCache.Get(4)
	t.Logf("%t, res:expect==%d:%d", res == -1, res, -1)
	res = lRUCache.Get(5)
	t.Logf("%t, res:expect==%d:%d", res == 5, res, 5)
}

func TestRemoveElement(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{3, 2, 2, 3},
			Target:   3,
			Expected: 2,
		},
		{
			Nums:     []int{},
			Target:   1,
			Expected: 0,
		},
		{
			Nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			Target:   2,
			Expected: 5,
		},
		{
			Nums:     []int{5, 2, 1},
			Target:   11,
			Expected: 3,
		},
	}
	var res int
	for _, item := range list {
		tempArray := make([]int, len(item.Nums))
		copy(tempArray, item.Nums)
		res = removeElement1(tempArray, item.Target)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		copy(tempArray, item.Nums)
		res = removeElement2(tempArray, item.Target)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestSortedSquares(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{-6, -4, -3, -1, -1, 0, 3, 3, 10},
			Expected: []int{0, 1, 1, 9, 9, 9, 16, 36, 100},
		},
		{
			Nums:     []int{-4, -1, 0, 3, 10},
			Expected: []int{0, 1, 9, 16, 100},
		},
		{
			Nums:     []int{-7, -3, 2, 3, 11},
			Expected: []int{4, 9, 9, 49, 121},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{-1},
			Expected: []int{1},
		},
	}
	var res []int
	for _, item := range list {
		tempArray := make([]int, len(item.Nums))
		copy(tempArray, item.Nums)
		res = sortedSquares1(tempArray)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		copy(tempArray, item.Nums)
		res = sortedSquares2(tempArray)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		copy(tempArray, item.Nums)
		res = sortedSquares3(tempArray)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		copy(tempArray, item.Nums)
		res = sortedSquares4(tempArray)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMinSubArrayLen(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{2, 3, 1, 2, 4, 3},
			Target:   7,
			Expected: 2,
		},
		{
			Nums:     []int{1, 4, 4},
			Target:   4,
			Expected: 1,
		},
		{
			Nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			Target:   11,
			Expected: 0,
		},
		{
			Nums:     []int{2, 3, 1, 2, 4, 3},
			Target:   6,
			Expected: 2,
		},
		{
			Nums:     []int{2, 3, 1, 2, 4, 3},
			Target:   2,
			Expected: 1,
		},
		{
			Nums:     []int{2, 3, 1, 2, 4, 3},
			Target:   3,
			Expected: 1,
		},
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Target:   11,
			Expected: 3,
		},
		{
			Nums:     []int{},
			Target:   0,
			Expected: 0,
		},
	}
	var res int
	for _, item := range list {
		res = minSubArrayLen1(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = minSubArrayLen2(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = minSubArrayLen3(item.Target, item.Nums)
		t.Logf("%t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestSpiralOrder(t *testing.T) {
	var list = []common.Item28{
		{
			Grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			Expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			Grid:     [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			Expected: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			Grid:     [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			Expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	var res = make([]int, 0)
	for _, item := range list {
		res = spiralOrder1(item.Grid)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		res = spiralOrder2(item.Grid)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		res = spiralOrder3(item.Grid)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestGenerateMatrix(t *testing.T) {
	var list = []common.Item29{
		{
			Num:      3,
			Expected: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			Num:      1,
			Expected: [][]int{{1}},
		},
	}
	var res = make([][]int, 0)
	for _, item := range list {
		res = generateMatrix1(item.Num)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestMissingNumber(t *testing.T) {
	var list = []common.Item2{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
		{[]int{1, 2}, 0},
	}
	var res int
	for _, item := range list {
		res = missingNumber1(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("------------------------SPLIT------------------------")

	}
}

func TestMaxSlidingWindow(t *testing.T) {
	var list = []common.Item21{
		{
			Nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			Target:   3,
			Expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			Nums:     []int{1, 3, -4, -1, -5, -3, 5, 3, 6, 7},
			Target:   3,
			Expected: []int{3, 3, -1, -1, 5, 5, 6, 7},
		},
		{
			Nums:     []int{1},
			Target:   1,
			Expected: []int{1},
		},
	}

	var (
		res = make([]int, 0)
	)
	for _, item := range list {
		res = maxSlidingWindow1(item.Nums, item.Target)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		res = maxSlidingWindow2(item.Nums, item.Target)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestIntersection(t *testing.T) {
	var list = []common.Item19{
		{
			Nums1:    []int{1, 2, 2, 1},
			Nums2:    []int{2, 2},
			Expected: []int{2},
		},
		{
			Nums1:    []int{4, 9, 5},
			Nums2:    []int{9, 4, 9, 8, 4},
			Expected: []int{9, 4},
		},
	}

	var res = make([]int, 0)
	for _, item := range list {
		sort.Ints(item.Expected)
		res = intersection1(item.Nums1, item.Nums2)
		sort.Ints(res)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestIntersect(t *testing.T) {
	var list = []common.Item19{
		{
			Nums1:    []int{1, 2, 2, 1},
			Nums2:    []int{2, 2},
			Expected: []int{2, 2},
		},
		{
			Nums1:    []int{4, 9, 5},
			Nums2:    []int{9, 4, 9, 8, 4},
			Expected: []int{4, 9},
		},
	}

	var res = make([]int, 0)
	for _, item := range list {
		sort.Ints(item.Expected)
		res = intersect1(item.Nums1, item.Nums2)
		sort.Ints(res)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		res = intersect2(item.Nums1, item.Nums2)
		sort.Ints(res)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		res = intersect3(item.Nums1, item.Nums2)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestExistInMatrix(t *testing.T) {
	var list = []common.Item44{
		{
			Grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			Word:     "ABCCED",
			Expected: true,
		},
		{
			Grid: [][]byte{
				{'a', 'b'},
				{'d', 'c'},
			},
			Word:     "abcd",
			Expected: true,
		},
		{
			Grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			Word:     "ABCCEF",
			Expected: false,
		},
	}

	var res bool
	for _, item := range list {
		res = exist(item.Grid, item.Word)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}
