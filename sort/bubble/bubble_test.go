package bubble

import (
	"fmt"
	"testing"

	"github.com/wafaking/goalgorithm/sort/utils"
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
		list := utils.GetArrayOfSize(i)

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
	for i := 2; i < 30; i++ {
		list := utils.GetArrayOfSize(i)

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
	for i := 2; i < 30; i++ {
		list := utils.GetArrayOfSize(i)

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
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		BubbleSortStandardOptimize(list)
	}
}
func BenchmarkBubbleSortStandardOptimize100(b *testing.B) { benchmarkBubbleSortStandardOptimize(100, b) }
func BenchmarkBubbleSortStandardOptimize1000(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(1000, b)
}
func BenchmarkBubbleSortStandardOptimize10000(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(10000, b)
}
func BenchmarkBubbleSortStandardOptimize100000(b *testing.B) {
	benchmarkBubbleSortStandardOptimize(100000, b)
}
