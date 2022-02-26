package insert

import (
	"fmt"
	"testing"

	"github.com/wafaking/goalgorithm/sort/utils"
)

func TestInsertSort(t *testing.T) {
	for i := 2; i < 30; i++ {
		list := utils.GetArrayOfSize(i)

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
		list := utils.GetArrayOfSize(n)
		InsertSort(list)
	}
}

func BenchmarkInsertSort100(b *testing.B)    { benchmarkInsertSort(100, b) }
func BenchmarkInsertSort1000(b *testing.B)   { benchmarkInsertSort(1000, b) }
func BenchmarkInsertSort10000(b *testing.B)  { benchmarkInsertSort(10000, b) }
func BenchmarkInsertSort100000(b *testing.B) { benchmarkInsertSort(100000, b) }
