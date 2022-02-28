package str

import (
	"fmt"
	"strconv"
)

// 数字
var chnNumChar = [10]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

//权位
var chnUnitSection = [4]string{"", "万", "亿", "万亿"}

// 数字权位
var chnUnitChar = [4]string{"", "十", "百", "千"}

// ZHToInt64 中文数字转阿拉伯数字
func ZHToInt64(chinese string) int64 {
	var (
		res int64 // 最终结果
		//section int64 // 单位
		tempRes int64 // 上一次计数
		tempNum int64
	)
	//"二千五百一十三",
	//"四万",
	//"四万零五百一十三",
	//"四十二万二千五百一十三",

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

			// TODO 考虑万亿这种情况
			// todo
			// "七千零七十九万",

			//"二千五百一十三",
			//"四万二千五百一十三",
			// 如果是万，亿
			// TODO duplication
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

type chnNameValue struct {
	name    string
	value   int
	secUnit bool
}

//权位于结点的关系
var chnValuePair = []chnNameValue{{"十", 10, false},
	{"百", 100, false},
	{"千", 1000, false},
	{"万", 10000, true},
	{"亿", 100000000, true},
}

// 阿拉伯数字转汉字
func numberToChinese(num int64) (numStr string) {
	var unitPos = 0
	var needZero = false

	for num > 0 { //小于零特殊处理
		section := num % 10000 // 已万为小结处理
		if needZero {
			numStr = chnNumChar[0] + numStr
		}
		strIns := sectionToChinese(section)
		if section != 0 {
			strIns += chnUnitSection[unitPos]
		} else {
			strIns += chnUnitSection[0]
		}
		numStr = strIns + numStr
		/*千位是 0 需要在下一个 section 补零*/
		needZero = (section < 1000) && (section > 0)
		num = num / 10000
		unitPos++
	}
	return
}
func mapZh2Int(chnStr string) (num int) {
	switch chnStr {
	case "零":
		num = 0
	case "一":
		num = 1
	case "二":
		num = 2
	case "三":
		num = 3
	case "四":
		num = 4
	case "五":
		num = 5
	case "六":
		num = 6
	case "七":
		num = 7
	case "八":
		num = 8
	case "九":
		num = 9
	default:
		num = -1
	}
	return
}

func main() {
	for {
		var typeStr string
		var scanStr string

		fmt.Println("1 阿拉伯转中文数字 2 中文数字转阿拉伯数字")
		fmt.Println("请输入")

		//fmt.Scanf("%s", &a)
		fmt.Scan(&typeStr)

		fmt.Println("请输入要转换的内容")
		fmt.Scan(&scanStr)
		if typeStr == "1" {
			num, _ := strconv.ParseInt(scanStr, 10, 64)
			var chnStr = numberToChinese(num)
			fmt.Println(chnStr)
		} else {
			var numInt = chineseToNumber(scanStr)
			fmt.Println(numInt)
		}
	}
}

func sectionToChinese(section int64) (chnStr string) {
	var strIns string
	var unitPos = 0
	var zero = true
	for section > 0 {
		var v = section % 10
		if v == 0 {
			if !zero {
				zero = true /*需要补，zero 的作用是确保对连续的多个，只补一个中文零*/
				chnStr = chnNumChar[v] + chnStr
			}
		} else {
			zero = false                   //至少有一个数字不是
			strIns = chnNumChar[v]         //此位对应的中文数字
			strIns += chnUnitChar[unitPos] //此位对应的中文权位
			chnStr = strIns + chnStr
		}
		unitPos++ //移位
		section = section / 10
	}
	return
}

// 汉字转阿拉伯数字
func chineseToNumber(chnStr string) (rtnInt int) {
	var section = 0
	var number = 0
	for index, value := range chnStr {

		// 中文一：1
		var num = mapZh2Int(string(value))
		if num > 0 {
			number = num
			length := len(chnStr)
			if index == length-3 { // 中文数字长度为3,此处为最后一位
				section += number
				rtnInt += section
				break
			}
		} else { // 单位
			unit, secUnit := chineseToUnit(string(value))
			if secUnit {
				section = (section + number) * unit
				rtnInt += section
				section = 0

			} else {
				section += (number * unit)

			}
			number = 0
			if index == len(chnStr)-3 {
				rtnInt += section
				break
			}
		}
	}

	return
}

func chineseToUnit(chnStr string) (unit int, secUnit bool) {

	for i := 0; i < len(chnValuePair); i++ {
		if chnValuePair[i].name == chnStr {
			unit = chnValuePair[i].value
			secUnit = chnValuePair[i].secUnit
		}
	}
	return
}
