package tree

// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

import (
	"fmt"
)

type operation uint

const (
	// INSERT add type
	INSERT operation = 1

	// DELETE  delete type
	DELETE = 2

	// MOVE change type
	MOVE = 3
)

func (op operation) String() string {
	switch op {
	case INSERT:
		return "INS"
	case DELETE:
		return "DEL"
	case MOVE:
		return "MOV"
	default:
		return "UNKNOWN"
	}
}

//var colors = map[operation]string{
//	INSERT: "\033[32m",
//	DELETE: "\033[31m",
//	MOVE:   "\033[39m",
//}

// GenerateDiff generate diff by source and destination
func GenerateDiff(src, dst []string) []string {

	var (
		script             = shortestEditScript(src, dst)
		srcIndex, dstIndex int
		res                []string
	)

	for _, op := range script {
		switch op {
		case INSERT:
			//fmt.Println(colors[op] + "+" + dst[dstIndex])
			res = append(res, fmt.Sprintf("+%s", dst[dstIndex]))
			dstIndex++

		case MOVE:
			//fmt.Println(colors[op] + " " + src[srcIndex])
			res = append(res, src[srcIndex])
			srcIndex++
			dstIndex++

		case DELETE:
			//fmt.Println(colors[op] + "-" + src[srcIndex])
			res = append(res, fmt.Sprintf("-%s", src[srcIndex]))
			srcIndex++
		}
	}

	return res
}

// shortestEditScript generate the shortest scripts
func shortestEditScript(src, dst []string) []operation {
	n := len(src)
	m := len(dst)
	max := n + m
	var trace []map[int]int
	var x, y int

loop:
	for d := 0; d <= max; d++ {
		// 最多只有 d+1 个 k
		v := make(map[int]int, d+2)
		trace = append(trace, v)

		// 需要注意处理对角线
		if d == 0 {
			t := 0
			for len(src) > t && len(dst) > t && src[t] == dst[t] {
				t++
			}
			v[0] = t
			//if t == len(src) && t == len(dst) {
			if len(dst) == len(src) && t == len(src) {
				break loop
			}
			continue
		}

		lastV := trace[d-1]

		for k := -d; k <= d; k += 2 {
			// 向下
			if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
				x = lastV[k+1]
			} else { // 向右
				x = lastV[k-1] + 1
			}

			y = x - k

			// 处理对角线
			for x < n && y < m && src[x] == dst[y] {
				x, y = x+1, y+1
			}

			v[k] = x

			if x == n && y == m {
				break loop
			}
		}
	}

	// for debug
	// printTrace(trace)

	// 反向回溯
	var script []operation

	x = n
	y = m
	var k, prevK, prevX, prevY int

	for d := len(trace) - 1; d > 0; d-- {
		k = x - y
		lastV := trace[d-1]

		if k == -d || (k != d && lastV[k-1] < lastV[k+1]) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		prevX = lastV[prevK]
		prevY = prevX - prevK

		for x > prevX && y > prevY {
			script = append(script, MOVE)
			x--
			y--
		}

		if x == prevX {
			script = append(script, INSERT)
		} else {
			script = append(script, DELETE)
		}

		x, y = prevX, prevY
	}

	if trace[0][0] != 0 {
		for i := 0; i < trace[0][0]; i++ {
			script = append(script, MOVE)
		}
	}

	return reverse(script)
}

// printTrace print trace
//func printTrace(trace []map[int]int) {
//	for d := 0; d < len(trace); d++ {
//		fmt.Printf("d = %d:\n", d)
//		v := trace[d]
//		for k := -d; k <= d; k += 2 {
//			x := v[k]
//			y := x - k
//			fmt.Printf("  k = %2d: (%d, %d)\n", k, x, y)
//		}
//	}
//}

// reverse function
func reverse(s []operation) []operation {
	result := make([]operation, len(s))

	for i, v := range s {
		result[len(s)-1-i] = v
	}

	return result
}
