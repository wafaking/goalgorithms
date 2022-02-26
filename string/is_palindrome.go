package str

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	var build strings.Builder

	for _, v := range s {
		if ('a' <= v && v <= 'z') ||
			(v >= '0' && v <= '9') {
			build.WriteRune(v)
		} else if v >= 'A' && v <= 'Z' {
			v += 'a' - 'A'
			build.WriteRune(v)
		}
	}

	s = build.String()
	var start, end = 0, len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isPalindrome1(s string) bool {
	var start, end = 0, len(s) - 1
	s = strings.ToLower(s)
	fmt.Println(s, start, end)

	for start < end {
		if !(('a' <= s[start] && s[start] <= 'z') ||
			(s[start] >= '0' && s[start] <= '9')) {
			fmt.Println("st: ", fmt.Sprintf("%v", s[start]))
			start++
			continue
		}

		if !(('a' <= s[end] && s[end] <= 'z') ||
			(s[end] >= '0' && s[end] <= '9')) {
			end--
			fmt.Println("end: ", s[end])
			continue
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
