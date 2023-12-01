package graph

import (
	"goalgorithms/common"
	"testing"
)

func TestCanFinish(t *testing.T) {
	var list = []common.Item27{
		{
			Grid:     [][]int{{1, 0}},
			Target:   2,
			Expected: true,
		},
		{
			Grid:     [][]int{{0, 1}, {1, 0}},
			Target:   2,
			Expected: false,
		},
		{
			Grid:     [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}},
			Target:   20,
			Expected: false,
		},
		{
			Grid:     [][]int{{1, 0}, {2, 0}, {2, 3}, {4, 1}, {5, 2}, {6, 3}, {7, 4}, {7, 5}, {7, 6}},
			Target:   8,
			Expected: true,
		},
	}

	var res bool
	for _, item := range list {
		res = canFinish1(item.Target, item.Grid)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		res = canFinish2(item.Target, item.Grid)
		t.Logf("%t, res-Expected: %t:%t, item:%+v", res == item.Expected, res, item.Expected, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}
