package str

import (
	"goalgorithms/common"
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
		res := common.Reverse(k)
		log.Printf("isRight: %t, res:%s, expected:%s", res == expected, res, expected)
	}
}

func TestSwapTwoNum(t *testing.T) {
	var a, b = 10, 20
	SwapTwoNum(a, b)
}

func TestReverseVowels(t *testing.T) {
	var list = []common.Item36{
		{
			Str:      "leetcode",
			Expected: "leotcede",
		},
		{
			Str:      "hello",
			Expected: "holle",
		},
		{
			Str:      "hhhhh",
			Expected: "hhhhh",
		},
	}
	var res string
	for _, item := range list {
		res = reverseVowels1(item.Str)
		t.Logf("res: %t, res: %+vs, item:%+v", res == item.Expected, res, item)
		res = reverseVowels2(item.Str)
		t.Logf("res: %t, res: %+vs, item:%+v", res == item.Expected, res, item)
		res = reverseVowels3(item.Str)
		t.Logf("res: %t, res: %+vs, item:%+v", res == item.Expected, res, item)
		t.Log("-------------------------SPLIT----------------")
	}
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
	}
}

func TestCheckInclusion(t *testing.T) {
	res := checkInclusion("ab", "eidbocoaoo")
	log.Println("res: ", res)
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
	var list = []common.Item38{
		{
			Str1:     "333",
			Str2:     "222",
			Expected: "555",
		},
		{
			Str1:     "33",
			Str2:     "222",
			Expected: "255",
		},
		{
			Str1:     "399",
			Str2:     "21",
			Expected: "420",
		},
		{
			Str1:     "59",
			Str2:     "93",
			Expected: "152",
		},
	}
	var res string
	for _, item := range list {
		res = addStrings1(item.Str1, item.Str2)
		t.Logf("%t, res: %s, item:%+v", res == item.Expected, res, item)
		res = addStrings2(item.Str1, item.Str2)
		t.Logf("%t, res: %s, item:%+v", res == item.Expected, res, item)
		t.Log("------------------------SPLIT------------------------")
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
	list := []common.Item34{
		{
			Num:      2.00000,
			N:        8,
			Expected: 256.00000,
		},
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

	var res float64
	for _, item := range list {
		res = myPow11(item.Num, item.N)
		t.Logf("%f(res): %f(expected),num: %f, n=%d", res, item.Expected, item.Num, item.N)
		res = myPow12(item.Num, item.N)
		t.Logf("%f(res): %f(expected),num: %f, n=%d", res, item.Expected, item.Num, item.N)
		res = myPow13(item.Num, item.N)
		t.Logf("%f(res): %f(expected),num: %f, n=%d", res, item.Expected, item.Num, item.N)
		t.Log("----------------------SPLIT---------------------")
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

		res = lengthOfLongestSubstring21(item.Str)
		if res != item.Expected {
			log.Fatalf("lengthOfLongestSubstring3 failed, res: %d, item:%#v", res, item)
		}
	}
}

func TestKMPSearch(t *testing.T) {
	var list = []common.Item16{
		{
			Text1:    "ABC ABCDAB ABCDABCDABDE",
			Text2:    "ABCDABD",
			Expected: 15,
		},
		{
			Text1:    "AAAAAAAAAAAAAAAAAAS",
			Text2:    "AAAAAAAAB",
			Expected: -1,
		},
	}
	for _, item := range list {
		res := strStrV2(item.Text1, item.Text2)
		t.Logf("result: %t, res-expect:%d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
	}
}

func TestGenerateNext(t *testing.T) {
	var list = []common.Item35{
		{
			S:        "ABCDABD",
			Expected: []int{-1, 0, 0, 0, 0, 1, 2},
		},
	}
	for _, item := range list {
		res := generateNext(item.S)
		t.Logf("str: %s, res: %+v", item.S, res)
		t.Log("------------------------")
	}
}

func TestIsHappy(t *testing.T) {
	var list = []common.Item31{
		{
			Num:      19,
			Expected: true,
		},
		{
			Num:      2,
			Expected: false,
		},
		{
			Num:      1,
			Expected: true,
		},
		{
			Num:      11,
			Expected: false,
		},
	}

	var res bool
	for _, item := range list {
		res = isHappy1(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		res = isHappy2(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		res = isHappy3(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestIsUgly(t *testing.T) {
	var list = []common.Item31{
		{
			Num:      0,
			Expected: false,
		},
		{
			Num:      210,
			Expected: false,
		},
		{
			Num:      6,
			Expected: true,
		},
		{
			Num:      1,
			Expected: true,
		},
		{
			Num:      14,
			Expected: false,
		},
	}

	var res bool
	for _, item := range list {
		res = isUgly1(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestNthUglyNumber(t *testing.T) {
	var list = []common.Item4{
		{
			Num:      1,
			Expected: 1,
		},
		{
			Num:      4,
			Expected: 4,
		},
		{
			Num:      5,
			Expected: 5,
		},
		{
			Num:      6,
			Expected: 6,
		},
		{
			Num:      10,
			Expected: 12,
		},
	}

	var res int
	for _, item := range list {
		res = nthUglyNumber1(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		res = nthUglyNumber2(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestNumSquares(t *testing.T) {
	var list = []common.Item4{
		{
			Num:      12,
			Expected: 3,
		},
		{
			Num:      7,
			Expected: 4,
		},
		{
			Num:      13,
			Expected: 2,
		},
		{
			Num:      1,
			Expected: 1,
		},
		{
			Num:      16,
			Expected: 1,
		},
		{
			Num:      25,
			Expected: 1,
		},
	}

	var res int
	for _, item := range list {
		res = numSquares1(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		res = numSquares2(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestCountPrimes(t *testing.T) {
	var list = []common.Item4{
		//{
		//	Num:      1,
		//	Expected: 0,
		//},
		{
			Num:      10,
			Expected: 4,
		},
		{
			Num:      20,
			Expected: 8,
		},
		{
			Num:      31,
			Expected: 10,
		},
		{
			Num:      45,
			Expected: 14,
		},
		{
			Num:      499979,
			Expected: 41537,
		},
	}

	var res int
	for _, item := range list {
		res = countPrimes1(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		res = countPrimes2(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		res = countPrimes3(item.Num)
		t.Logf("%t, res:%+v, item:%+v", item.Expected == res, res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestFindGCD(t *testing.T) {
	var list = []common.Item2{
		{[]int{2, 5, 6, 9, 10}, 2},
		{[]int{7, 5, 6, 8, 3}, 1},
		{[]int{3, 3}, 3},
	}
	var res int
	for _, item := range list {
		res = findGCD1(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findGCD2(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findGCD3(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		res = findGCD4(item.Nums)
		t.Logf("res: %t, res-Expected: %d:%d, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestReverseString(t *testing.T) {
	var list = []common.Item32{
		{
			Bytes:    []byte{'h', 'e', 'l', 'l', 'o'},
			Expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			Bytes:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			Expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, item := range list {
		var temp = make([]byte, len(item.Bytes))
		copy(temp, item.Bytes)
		reverseString1(temp)
		t.Logf("res: %t, res: %+vs, item:%+v", string(temp) == string(item.Expected), temp, item)
		copy(temp, item.Bytes)
		reverseString2(temp)
		t.Logf("res: %t, res: %+vs, item:%+v", string(temp) == string(item.Expected), temp, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestReverseStr(t *testing.T) {
	var list = []common.Item33{
		{
			Str:      "abcdefg",
			N:        2,
			Expected: "bacdfeg",
		},
		{
			Str:      "abcd",
			N:        2,
			Expected: "bacd",
		},
	}

	var res string
	for _, item := range list {
		res = reverseStr1(item.Str, item.N)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = reverseStr2(item.Str, item.N)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = reverseStr3(item.Str, item.N)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = reverseStr4(item.Str, item.N)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestAddDigits(t *testing.T) {
	var list = []common.Item4{
		{
			Num:      38,
			Expected: 2,
		},
		{
			Num:      0,
			Expected: 0,
		},
		{
			Num:      999,
			Expected: 9,
		},
	}

	var res int
	for _, item := range list {
		res = addDigits1(item.Num)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = addDigits2(item.Num)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = addDigits3(item.Num)
		t.Logf("res: %t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("------------------------SPLIT------------------------")
	}
}
