package selection

import (
	"fmt"
	"testing"

	"github.com/wafaking/goalgorithm/sort/utils"
)

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
