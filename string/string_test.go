package str

import (
	"fmt"
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	var strRes = map[string]string{
		"abcd":  "dcba",
		"a":     "a",
		"":      "",
		"ab-cd": "dc-ba",
	}

	for k, expected := range strRes {
		res := reverse(k)
		log.Printf("isRight: %t, res:%s, expected:%s", res == expected, res, expected)
	}
}

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
	var listMap = map[string]int64{
		"四千五百一十三":          4513,
		"四万":               40000,
		"四万二千五百一十三":        42513,
		"四万零五百一十三":         40513,
		"七千零七十九万":          70790000,
		"一亿七千七十九万七千一百九十七":  170797197,
		"一亿七千零七十九万七千一百九十七": 170797197,
		"一千亿七千七十九万七千一百九十七": 100070797197,
		"一万亿七千七十九万七千一百九十七": 1000070807197,
	}
	for k, expected := range listMap {
		//res := chineseToNumber(v)
		res := ZHToInt64(k)
		log.Printf("isPass: %t,%s, v:%d, expectd:%d \n", res == expected, k, res, expected)
	}
}

func TestNumberToChinese(t *testing.T) {
	var listMap = map[int64]string{
		0:            "零",
		4000:         "四千",
		4510:         "四千五百一十",
		4513:         "四千五百一十三",
		4013:         "四千零一十三",
		4003:         "四千零三",
		40000:        "四万",
		42513:        "四万二千五百一十三",
		40513:        "四万零五百一十三",
		70790000:     "七千零七十九万",
		10797197:     "一千零七十九万七千一百九十七",
		100797197:    "一亿零七十九万七千一百九十七",
		170797197:    "一亿七千零七十九万七千一百九十七",
		100070797197: "一千亿七千零七十九万七千一百九十七",
	}
	for k, expected := range listMap {
		//log.Println("k: ", k)
		res := Number2ZH1(k)
		//res := Number2ZH2(k)
		log.Printf("isPass: %t, %d, v:%s, expectd:%s\n", res == expected, k, res, expected)
	}
}

func TestAddStrings(t *testing.T) {
	//var num1, num2 = "333", "222"
	//var num1, num2 = "33", "222"
	//var num1, num2 = "333", "adb"
	var num1, num2 = "399", "21"
	//var num1, num2 = "59", "93"
	res := addStrings2(num1, num2)
	t.Logf("res: %s, num1：%s, num2:%s", res, num1, num2)
}

func TestReverseStr(t *testing.T) {

	var strList = []string{
		"abcdefg",
		"abcdef",
		"abcd",
	}

	for _, v := range strList {
		//res := reverseStr1(v, 7)
		res := reverseStr2(v, 2)
		t.Logf("res:%s", res)
	}
}

func TestLongestPalindrome(t *testing.T) {
	var m = map[string]string{
		"a":     "a",
		"ab":    "a",
		"abc":   "a",
		"aba":   "aba",
		"cbbd":  "bb",
		"babad": "bab",
	}

	for str, expected := range m {
		//res := longestPalindrome1(str)
		res := longestPalindrome1(str)
		t.Logf(" Right: %t, res：%s, expected:%s, str: %s", res == expected, res, expected, str)
	}
}
