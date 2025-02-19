package sort

import (
	"fmt"
	"goalgorithms/common"
	"testing"

	"goalgorithms/sort/utils"
)

var commonListArray [][]int

func TestMain(t *testing.M) {
	var length = 5
	commonListArray = make([][]int, 0, length)
	for i := 0; i < length; i++ {
		list := utils.GetArrayOfSize(i)
		commonListArray = append(commonListArray, list)
	}
	t.Run()
}

// -----TestBubbleSort---test---------------
func TestBubbleSort(t *testing.T) {
	for _, list := range commonListArray {
		BubbleSort(list)
		t.Logf("array: %v", list)
		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		BubbleSort(list)
	}
}

func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }

// -----BubbleSortStandard---test---------------

func TestBubbleSortStandard(t *testing.T) {
	for _, list := range commonListArray {
		BubbleSortStandard(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkBubbleSortStandard(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		BubbleSortStandard(list)
	}
}
func BenchmarkBubbleSortStandard100(b *testing.B)    { benchmarkBubbleSortStandard(100, b) }
func BenchmarkBubbleSortStandard1000(b *testing.B)   { benchmarkBubbleSortStandard(1000, b) }
func BenchmarkBubbleSortStandard10000(b *testing.B)  { benchmarkBubbleSortStandard(10000, b) }
func BenchmarkBubbleSortStandard100000(b *testing.B) { benchmarkBubbleSortStandard(100000, b) }

// -----BubbleSortStandardOptimize---test---------------

// -----BubbleSortStandardOptimize---test---------------
func TestStandardBubbleSortOptimize(t *testing.T) {
	for _, list := range commonListArray {
		BubbleSortStandardOptimize(list)
		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkBubbleSortStandardOptimize(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		BubbleSortStandardOptimize(list)
	}
}
func BenchmarkBubbleSortStandardOptimize100(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(100, b)
}
func BenchmarkBubbleSortStandardOptimize1000(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(1000, b)
}
func BenchmarkBubbleSortStandardOptimize10000(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(10000, b)
}
func BenchmarkBubbleSortStandardOptimize100000(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(100000, b)
}

func TestInsertSort(t *testing.T) {
	for _, list := range commonListArray {
		t.Logf("start array: %v", list)
		InsertSort1(list)
		//InsertSort2(array)
		t.Logf("end array: %v", list)
		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkInsertSort(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := utils.GetArrayOfSize(n)
		InsertSort1(list)
	}
}

func BenchmarkInsertSort100(b *testing.B)    { benchmarkInsertSort(100, b) }
func BenchmarkInsertSort1000(b *testing.B)   { benchmarkInsertSort(1000, b) }
func BenchmarkInsertSort10000(b *testing.B)  { benchmarkInsertSort(10000, b) }
func BenchmarkInsertSort100000(b *testing.B) { benchmarkInsertSort(100000, b) }

func TestSelectionSort(t *testing.T) {
	for i := 2; i < 30; i++ {
		list := utils.GetArrayOfSize(i)

		SelectionSort(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkSelectionSort(num int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := utils.GetArrayOfSize(num)
		SelectionSort(list)
	}
}

func BenchmarkSelectionSort100(b *testing.B)    { benchmarkSelectionSort(100, b) }
func BenchmarkSelectionSort1000(b *testing.B)   { benchmarkSelectionSort(1000, b) }
func BenchmarkSelectionSort10000(b *testing.B)  { benchmarkSelectionSort(10000, b) }
func BenchmarkSelectionSort100000(b *testing.B) { benchmarkSelectionSort(100000, b) }

func TestShellSort(t *testing.T) {
	for i := 2; i < 30; i++ {
		list := utils.GetArrayOfSize(i)

		ShellSort(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkShellSort(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := utils.GetArrayOfSize(n)
		ShellSort(list)
	}
}

func BenchmarkShellSort100(b *testing.B)    { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }

func TestQuickSort2Ways(t *testing.T) {
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
	for _, item := range list {
		var nums = make([]int, len(item.Nums))
		copy(nums, item.Nums)
		QuickSort2Ways1(nums)
		common.PrintDiffTwoIntSlice(nums, item.Expected, item, t)
		copy(nums, item.Nums)
		QuickSort2Ways1(nums)
		common.PrintDiffTwoIntSlice(nums, item.Expected, item, t)

		common.PrintSplit(t)
	}
}

func TestQuickSort3Ways(t *testing.T) {
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
	for _, item := range list {
		var nums = make([]int, len(item.Nums))
		copy(nums, item.Nums)
		QuickSort3Ways(nums)
		t.Logf("res: %v, item:%+v", common.DiffTwoIntSlice(nums, item.Expected), item)
		t.Log("--------------------SPLIT--------------------")
	}
}

func TestQuickSort(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{3},
			Expected: []int{3},
		},
		{
			Nums:     []int{3, 1, 4, 1, 2, 7, 8, 5, 6},
			Expected: []int{1, 1, 2, 3, 4, 5, 6, 7, 8},
		},
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

	var res = make([]int, 0)
	for _, item := range list {
		tempNums := make([]int, len(item.Nums))

		copy(tempNums, item.Nums)
		res = quickSort1(tempNums)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		copy(tempNums, item.Nums)
		res = quickSort2(tempNums)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		copy(tempNums, item.Nums)
		res = quickSort3(tempNums)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		copy(tempNums, item.Nums)
		res = quickSort4(tempNums)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}

func TestMergeSort(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{3, 1, 4, 1, 2, 7, 8, 5, 6},
			Expected: []int{1, 1, 2, 3, 4, 5, 6, 7, 8},
		},
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

	var res = make([]int, 0)
	for _, item := range list {
		var temp = make([]int, len(item.Nums))
		copy(temp, item.Nums)
		res = mergeSort1(temp)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		copy(temp, item.Nums)
		res = mergeSort2(temp)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}
