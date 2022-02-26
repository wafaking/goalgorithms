package bitmap

import "testing"

func TestHammingWeight(t *testing.T) {

	resMap := map[uint32]int{
		11:         3,
		128:        1,
		4294967293: 31,
	}
	for n, expected := range resMap {
		res := hammingWeight1(n)
		// res := hammingWeight2(n)
		t.Logf("res: %t, %d==%d(expected), n=%d", res == expected, res, expected, n)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	resMap := map[int]bool{
		0:   false,
		1:   true,
		4:   true,
		16:  true,
		11:  false,
		128: true,
	}
	for n, expected := range resMap {
		// res := isPowerOfTwo(n)
		// res := isPowerOfTwo2(n)
		res := isPowerOfTwo3(n)
		t.Logf("res: %t, %t==%t(expected), n=%d", res == expected, res, expected, n)
	}
}
