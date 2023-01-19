package tree

import (
	"strings"
)

func Diff3(fileO, fileA, fileB []string) []string {
	mapA := getDiffMap(fileO, fileA)
	mapB := getDiffMap(fileO, fileB)
	return GenerateTrunc(mapA, mapB, fileO, fileA, fileB)
}

// GenerateTrunc generate trunc
func GenerateTrunc(mapA, mapB map[int]int, fileO, fileA, fileB []string) []string {
	var (
		ans                    []string
		matchO, matchA, matchB = -1, -1, -1 // 记录上一次匹配的位置
		travers                func(start int)
		emitTrunc              func(startA, startB, endA, endB int)
	)

	emitTrunc = func(startA, startB, endA, endB int) {
		if startA < endA || startB < endB {
			ans = append(ans, "<<<<<<< fileA")
			for i := startA; i < endA; i++ {
				ans = append(ans, fileA[i])
			}
			ans = append(ans, "=======")
			for i := startB; i < endB; i++ {
				ans = append(ans, fileB[i])
			}
			ans = append(ans, ">>>>>>> fileB")
		}
	}

	travers = func(start int) {
		var o, a, b = findNextMatch(mapA, mapB, start, len(fileO))
		if o == -1 { // 无匹配项,添加并结束
			emitTrunc(matchA+1, matchB+1, len(fileA), len(fileB))
			return
		}

		// 有匹配项
		if a-1 > matchA || b-1 > matchB {
			// 此行匹配项与上一行匹配项之间有差异
			emitTrunc(matchA+1, matchB+1, a, b)
		}
		ans = append(ans, fileO[o])
		// 更新match索引
		matchO, matchA, matchB = o, a, b
		travers(matchO + 1)
	}

	// 从第一行开始
	travers(0)
	return ans
}

// findNextMatch if matches, return o, A, B
func findNextMatch(mapA, mapB map[int]int, start, total int) (int, int, int) {
	for o := start; o < total; o++ {
		idxA, okA := mapA[o]
		idxB, okB := mapB[o]
		if okA && okB {
			return o, idxA, idxB
		}
	}
	return -1, -1, -1
}

func getDiffMap(base, head []string) map[int]int {
	var mapA = make(map[int]int, 0)
	outPut := GenerateDiff(base, head)
	var headIdx, baseIdx = 0, 0
	for _, item := range outPut {
		item = strings.TrimSpace(item)
		if strings.HasPrefix(item, "+") {
			headIdx++
		} else if strings.HasPrefix(item, "-") {
			baseIdx++
		} else {
			mapA[baseIdx] = headIdx
			baseIdx++
			headIdx++
		}
	}
	return mapA
}
