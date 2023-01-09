package array

import "log"

var (
	nums   []int
	d2Nums [][]int // 二维数组
)

type item struct {
	nums     []int
	target   int
	expected []int
}

type item2 struct {
	nums     []int
	expected int
}

type item3 struct {
	nums     []int
	target   int
	expected int
}

type item4 struct {
	target   int
	expected int
}

func InitNums(sli []int) {
	nums = sli
}

func Init2DNums(d2Sli [][]int) {
	d2Nums = d2Nums
}

func PrintNums() {
	log.Printf("*******nums: %#v*************\n", nums)
}
func Println2DNums() {
	log.Printf("*******nums: %#v*************\n", d2Nums)
}
