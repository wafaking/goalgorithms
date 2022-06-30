package sort

import (
	"fmt"
	"github.com/wafaking/goalgorithms/utils/sort"
	"testing"
)

var arraySli = [][]int{
	{10, 9, 7, 5, 6, 3},
	{1, 3, 46, 5, 2, 10, 2, 31, 18, 24, 30, 12, 9, 4},
	{1, 3, 2, 4, 5, 6, 7, 8, 9},
	{3, 9, 1, 4, 7},
	{432, 432432, 0, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4},
}

// -----TestBubbleSort---test---------------
func TestBubbleSort(t *testing.T) {
	for i := 2; i < 30; i++ {
		list := sort.GetArrayOfSize(i)

		BubbleSort(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()

			}
		}
	}

	for _, list := range arraySli {
		BubbleSort(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	list := sort.GetArrayOfSize(n)
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
	for i := 2; i < 30; i++ {
		list := sort.GetArrayOfSize(i)

		BubbleSortStandard(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}

	for _, list := range arraySli {
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
	list := sort.GetArrayOfSize(n)
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
	for i := 2; i < 30; i++ {
		list := sort.GetArrayOfSize(i)

		BubbleSortStandardOptimize(list)

		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+1] {
				fmt.Println(list)
				t.Error()
			}
		}
	}

	for _, list := range arraySli {
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
	list := sort.GetArrayOfSize(n)
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
	for i := 2; i < 30; i++ {
		list := sort.GetArrayOfSize(i)

		InsertSort(list)

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
		list := sort.GetArrayOfSize(n)
		InsertSort(list)
	}
}

func BenchmarkInsertSort100(b *testing.B)    { benchmarkInsertSort(100, b) }
func BenchmarkInsertSort1000(b *testing.B)   { benchmarkInsertSort(1000, b) }
func BenchmarkInsertSort10000(b *testing.B)  { benchmarkInsertSort(10000, b) }
func BenchmarkInsertSort100000(b *testing.B) { benchmarkInsertSort(100000, b) }

func TestSelectionSort(t *testing.T) {
	for i := 2; i < 30; i++ {
		list := sort.GetArrayOfSize(i)

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
		list := sort.GetArrayOfSize(num)
		SelectionSort(list)
	}
}

func BenchmarkSelectionSort100(b *testing.B)    { benchmarkSelectionSort(100, b) }
func BenchmarkSelectionSort1000(b *testing.B)   { benchmarkSelectionSort(1000, b) }
func BenchmarkSelectionSort10000(b *testing.B)  { benchmarkSelectionSort(10000, b) }
func BenchmarkSelectionSort100000(b *testing.B) { benchmarkSelectionSort(100000, b) }

func TestShellSort(t *testing.T) {
	for i := 2; i < 30; i++ {
		list := sort.GetArrayOfSize(i)

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
		list := sort.GetArrayOfSize(n)
		ShellSort(list)
	}
}

func BenchmarkShellSort100(b *testing.B)    { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }
