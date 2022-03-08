package str

// ZHToInt64 中文数字转阿拉伯数字
func ZHToInt64(chinese string) int64 {
	var (
		res     int64 // 最终结果
		tempRes int64 // 上一次计数
		tempNum int64
	)
	for i, v := range chinese {
		num, ok := mapZHToInt[string(v)]
		if ok { // 数值
			if num > 0 {
				tempNum = num
			}
			if i == len(chinese)-3 { // 最后一位，结束
				return res + tempRes + tempNum
			}

		} else { // 单位
			unit, ok := mapZHUnitToInt[string(v)]
			if !ok {
				return -1
			}

			if string(v) == "万" || string(v) == "亿" {
				res += (tempRes + tempNum) * unit
				tempRes = 0
			} else {
				tempRes += tempNum * unit
			}
			tempNum = 0
			unit = 0

			// todo
			if i == len(chinese)-3 { // 最后一位，结束
				if string(v) == "万" || string(v) == "亿" {
					return res
				} else {
					return res + tempRes
				}
			}
		}
	}

	return res
}

var mapZHToInt = map[string]int64{
	"零": 0,
	"一": 1,
	"二": 2,
	"三": 3,
	"四": 4,
	"五": 5,
	"六": 6,
	"七": 7,
	"八": 8,
	"九": 9,
}

var mapZHUnitToInt = map[string]int64{
	"十": 10,
	"百": 100,
	"千": 1000,
	"万": 10000,
	"亿": 100000000,
}
