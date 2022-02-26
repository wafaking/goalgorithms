package dynamics

import (
	"log"
	"testing"
)

func TestCuttingRope(t *testing.T) {
	var sampleMap = map[int]int{
		2:  1,
		3:  2,
		5:  6,
		6:  9,
		7:  12, //  3;4
		10: 36, // 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
	}
	for n, expected := range sampleMap {
		// res := cuttingRope1(n)
		// res := cuttingRope2(n)
		res := cuttingRope3(n)
		t.Logf("res: %t, n:%d, res:%d, expected:%d", res == expected, n, res, expected)
	}
}

func TestPermute(t *testing.T) {
	var numsList = [][]int{
		//{1},
		//{0, 1},
		//{1, 2, 3},
		{1, 2, 3, 4},
	}
	for _, nums := range numsList {
		res := permute2(nums)
		log.Printf("res: %v\n", res)

	}
}
