package array

import "log"

var nums []int

func InitNums(sli []int) {
	nums = sli
}

func PrintNums() {
	log.Printf("*******nums: %#v*************\n", nums)
}
