package list

import "log"

var (
	nums   []int
	d2Nums [][]int // 二维数组
)

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
