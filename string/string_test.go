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
	var list = []string{
		"abc",
		"aac",
		//"中1c",
	}
	for _, item := range list {
		res := permutation1(item)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = permutation2(item)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		res = permutation3(item)
		t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		//res = permutation4(item)
		//t.Logf("res: %v\n", res)
		//t.Log("----------SPLIT------------")
		//res = permutation5(item)
		//t.Logf("res: %v\n", res)
		//t.Log("----------SPLIT------------")

	}
}

func TestLetterCasePermutation(t *testing.T) {
	var list = []string{
		"ab",
		"abc",
		"a1b2",
		//"1a2b3c",
		//"3z4",
		//"123",
		//"中1c",
	}
	for _, item := range list {
		res := letterCasePermutation1(item)
		t.Logf("res: %v\n", res)
		res = letterCasePermutation2(item)
		t.Logf("res: %v\n", res)
		res = letterCasePermutation3(item)
		t.Logf("res: %v\n", res)
		res = letterCasePermutation4(item)
		t.Logf("res: %v\n", res)
		res = letterCasePermutation5(item)
		t.Logf("res: %v\n", res)
		//res = letterCasePermutation6(item)
		//t.Logf("res: %v\n", res)
		t.Log("----------SPLIT------------")
		//res = permutation3(item)
		//t.Logf("res: %v\n", res)
		//t.Log("----------SPLIT------------")
		//res = permutation4(item)
		//t.Logf("res: %v\n", res)
		//t.Log("----------SPLIT------------")
		//res = permutation5(item)
		//t.Logf("res: %v\n", res)
		//t.Log("----------SPLIT------------")

	}
}

func TestCheckInclusion(t *testing.T) {
	//res := checkInclusion1("ab", "eidbaooo")
	res := checkInclusion("ab", "eidbocoaoo")
	log.Println("res: ", res)
}

func TestRemoveKdigits(t *testing.T) {
	var list = []item2{
		{Num: "10001", N: 1, Expected: "1"},
		{Num: "100010", N: 1, Expected: "10"},
		{Num: "1432219", N: 3, Expected: "1219"},
		{Num: "10200", N: 1, Expected: "200"},
		{Num: "10", N: 1, Expected: "0"},
		{Num: "10", N: 2, Expected: "0"},
		{Num: "10", N: 3, Expected: "0"},
		{Num: "10001", N: 4, Expected: "0"},
		{Num: "10001", N: 2, Expected: "0"},
		{Num: "1102912", N: 2, Expected: "2912"},
		{Num: "1102912", N: 3, Expected: "212"},
		{Num: "11111", N: 3, Expected: "11"},
		{Num: "12345", N: 2, Expected: "123"},
	}
	for _, item := range list {
		res := removeKdigits1(item.Num, item.N)
		if res != item.Expected {
			t.Fatalf("removeKdigits1 FAILED, res: %s, expected:%s, item:%#v", res, item.Expected, item)
		}
		res = removeKdigits2(item.Num, item.N)
		if res != item.Expected {
			t.Fatalf("removeKdigits2 FAILED, res: %s, expected:%s, item:%#v", res, item.Expected, item)
		}
	}
}

func TestChineseToNumber(t *testing.T) {
	var listMap = map[string]int64{
		//"四千五百一十三":          4513,
		//"四万":               40000,
		//"四万二千五百一十三":        42513,
		//"四万零五百一十三":         40513,
		//"七千零七十九万":          70790000,
		//"一亿七千七十九万七千一百九十七":  170797197,
		//"一亿七千零七十九万七千一百九十七": 170797197,
		//"一千亿七千七十九万七千一百九十七": 100070797197,
		//"一万亿七千七十九万七千一百九十七": 1000070807197,
		"11000万": 1000,
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
	res := AddStrings2(num1, num2)
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
		"a":           "a",
		"ab":          "a",
		"abc":         "a",
		"aba":         "aba",
		"cbbd":        "bb",
		"babad":       "bab",
		"aacabdkacaa": "aca",
	}

	for str, expected := range m {
		//res := longestPalindrome1(str)
		res := longestPalindrome2(str)
		t.Logf(" Right: %t, res：%s, expected:%s, str: %s", res == expected, res, expected, str)
	}

}
func TestMyPow(t *testing.T) {
	list := []item{
		{
			Num:      2.00000,
			N:        10,
			Expected: 1024.00000,
		},
		{
			Num:      2.10000,
			N:        3,
			Expected: 9.26100,
		},
		{
			Num:      2.00000,
			N:        -2,
			Expected: 0.25000,
		},
	}

	for _, item := range list {
		res := myPow11(item.Num, item.N)
		t.Logf("%f(res): %f(expected),num: %f, n=%d", res, item.Expected, item.Num, item.N)
		t.Log("-----------SPLIT----------")
		res = myPow12(item.Num, item.N)
		t.Logf("%f(res): %f(expected),num: %f, n=%d", res, item.Expected, item.Num, item.N)
		t.Log("-----------SPLIT----------")
		res = myPow13(item.Num, item.N)
		t.Logf("%f(res): %f(expected),num: %f, n=%d", res, item.Expected, item.Num, item.N)
		t.Log("-----------SPLIT----------")
	}
}

func TestIsPalindromeInt1(t *testing.T) {
	var list = []int{
		-121,
		0,
		1,
		10,
		11,
		1221,
		1224,
		1001,
		12321,
		12341,
		10021,
		100021,
	}
	for _, num := range list {
		res := isPalindrome1(num)
		t.Logf("res: %t, num=%d", res, num)
		//t.Log("-----------SPLIT----------")
		res = isPalindrome21(num)
		t.Logf("res: %t, num=%d", res, num)
		res = isPalindrome22(num)
		t.Logf("res: %t, num=%d", res, num)
		res = isPalindrome3(num)
		t.Logf("res: %t, num=%d", res, num)
		t.Log("-----------SPLIT----------")
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var list = []struct {
		Str      string
		Expected int
	}{
		{Str: "abcabcbb", Expected: 3},
		{Str: "bbbbb", Expected: 1},
		{Str: "pwwkew", Expected: 3},
	}
	for _, item := range list {
		res := lengthOfLongestSubstring2(item.Str)
		if res != item.Expected {
			log.Fatalf("lengthOfLongestSubstring2 failed, res: %d, item:%#v", res, item)
		}

		res = lengthOfLongestSubstring3(item.Str)
		if res != item.Expected {
			log.Fatalf("lengthOfLongestSubstring3 failed, res: %d, item:%#v", res, item)
		}
	}
}
