package shell

import (
	"fmt"
	"testing"

	"github.com/wafaking/goalgorithm/sort/utils"
)

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
