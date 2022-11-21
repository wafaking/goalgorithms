package bitmap

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	resMap := map[uint32]int{
		11:         3,
		128:        1,
		4294967293: 31,
		10:         2,
	}
	for n, expected := range resMap {
		res := hammingWeight1(n)
		t.Logf("res: %t, %d==%d(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT-----------")
	for n, expected := range resMap {
		res := hammingWeight2(n)
		t.Logf("res: %t, %d==%d(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT-----------")
	for n, expected := range resMap {
		res := hammingWeight3(n)
		t.Logf("res: %t, %d==%d(expected), n=%d", res == expected, res, expected, n)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	resMap := map[int]bool{
		0:   false,
		1:   true,
		3:   false,
		4:   true,
		16:  true,
		11:  false,
		128: true,
	}
	for n, expected := range resMap {
		res := isPowerOfTwo1(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfTwo2(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfTwo3(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfTwo4(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
}

func TestIsPowerOfThree(t *testing.T) {
	resMap := map[int]bool{
		0:  false,
		1:  true,
		3:  true,
		9:  true,
		12: false,
		15: false,
		18: false,
		21: false,
		24: false,
		27: true,
	}
	for n, expected := range resMap {
		res := isPowerOfThree1(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfThree11(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfThree2(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfThree3(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
}

func TestIsPowerOfFour(t *testing.T) {
	resMap := map[int]bool{
		0:  false,
		1:  true,
		2:  false,
		4:  true,
		8:  false,
		12: false,
		16: true,
		20: false,
		24: false,
		28: false,
		32: false,
	}
	for n, expected := range resMap {
		res := isPowerOfFour1(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfFour2(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
	for n, expected := range resMap {
		res := isPowerOfFour3(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
	t.Log("-----------SPLIT----------")
}
