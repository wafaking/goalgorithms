package str

// 阿拉伯数字转中文
var digitalToZH = [10]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var bigUnitMap = [4]string{"", "万", "亿", "万亿"}
var unitMap = [4]string{"", "十", "百", "千"}

// 法一：阿拉伯数字转中文
func Number2ZH1(num int64) string {
	// 1. 先计算千以内的数字，如：2032，2003
	// 2. 再以万为单位计算，如：200012032
	return number2Zh(num)
}

func number2Zh(num int64) string {
	var (
		unit int
		res  string
	)

	if num == 0 {
		return digitalToZH[num]
	} else if num < 0 {
		return "unknown"
	}

	for num > 0 {
		tempNum := num % 10000
		if tempNum > 0 { //
			//40513:     "四万零五百一十三",
			tempTh := toZHWithinThousand(tempNum) // 小于1000的数计算
			if tempNum < 1000 && num/10000 > 0 {  //40513:     "四万零五百一十三",
				tempTh = "零" + tempTh
			}
			res = tempTh + bigUnitMap[unit] + res
		}
		unit++
		num /= 10000
	}
	return res

}

// 转换小于1000的数
func toZHWithinThousand(num int64) string {
	var (
		unit int
		res  string
	)

	for num > 0 {
		v := num % 10
		tempNum := digitalToZH[v]
		// 有单位
		if unit > 0 { // 此时零要展示
			tempUnit := unitMap[unit]
			if v == 0 { // 取值为0的情况
				if len(res) >= 3 && string(res[0:3]) != "零" { // 防止出现"零零"
					res = tempNum + res
				}
			} else {
				res = tempNum + tempUnit + res
			}
		} else { // 无单位
			// 第一位, 零不展示
			if v > 0 {
				res = tempNum + res
			}
		}
		unit++
		num /= 10
	}
	return res
}

// 法二：阿拉伯数字转汉字
func Number2ZH2(num int64) (numStr string) {
	var unitPos = 0
	var needZero = false

	for num > 0 { //小于零特殊处理
		section := num % 10000 // 已万为小结处理
		if needZero {
			numStr = digitalToZH[0] + numStr
		}
		strIns := sectionToChinese(section)
		if section != 0 {
			strIns += bigUnitMap[unitPos]
		} else {
			strIns += bigUnitMap[0]
		}
		numStr = strIns + numStr
		/*千位是 0 需要在下一个 section 补零*/
		needZero = (section < 1000) && (section > 0)
		num = num / 10000
		unitPos++
	}
	return
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
				chnStr = digitalToZH[v] + chnStr
			}
		} else {
			zero = false               //至少有一个数字不是
			strIns = digitalToZH[v]    //此位对应的中文数字
			strIns += unitMap[unitPos] //此位对应的中文权位
			chnStr = strIns + chnStr
		}
		unitPos++ //移位
		section = section / 10
	}
	return
}
