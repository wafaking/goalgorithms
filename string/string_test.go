package str

import (
	"fmt"
	"log"
	"testing"
)

func TestSwapTwoNum(t *testing.T) {
	var a, b = 10, 20
	SwapTwoNum(a, b)
}

func TestReverseVowels(t *testing.T) {
	var s1 = "leetcode"
	// var s2 = "hello"
	res := reverseVowels(s1)
	fmt.Println("res: ", res)
}

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	res := replaceSpace(s)
	log.Println("res: ", res)
}

func TestMyAtoi(t *testing.T) {
	s := map[string]int{
		"42":              42,
		"   -42":          -42,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    -2147483648,
		"21474836460":     2147483647,
		"2147483648":      2147483647,
	}
	var res int
	for k, v := range s {
		res = myAtoi(k)
		if res != v {
			t.Fatalf("k: %s, res: %d, expect: %d\n", k, res, v)
		}
	}
}
func TestPermutation(t *testing.T) {
	res := permutation4("abc")
	log.Println("res: ", res)
}

func TestCheckInclusion(t *testing.T) {
	//res := checkInclusion1("ab", "eidbaooo")
	res := checkInclusion("ab", "eidbocoaoo")
	log.Println("res: ", res)
}

func TestRemoveKdigits(t *testing.T) {
	sli := map[string]int{
		//"1432219": 3,
		//"10200":   1,
		//"10":      2,
		"10001": 4,
	}
	for str, k := range sli {
		res := removeKdigits(str, k)
		log.Printf("res---- str:%s, k: %d, res: %s\n", str, k, res)
	}
}

func TestChineseToNumber(t *testing.T) {
	var list = []string{
		// "四千五百一十三",
		// "四万",
		// "四万二千五百一十三",
		// "四万零五百一十三",
		// "四万二千五百一十三",
		"七千零七十九万",
		"一亿七千七十九万七千一百九十七",
		"一千亿七千七十九万七千一百九十七",
		"一万亿七千七十九万七千一百九十七",
	}
	for _, v := range list {
		//res := chineseToNumber(v)
		res := ZHToInt64(v)
		log.Printf("v:%s, res: %d \n", v, res)
	}
}
