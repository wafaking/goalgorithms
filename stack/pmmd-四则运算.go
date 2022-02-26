package stack

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 后缀表达式求值
func rearCalculation(expression []string) (result int) {
	var stack []int
	for _, v := range expression {
		switch v {
		default:
			// 数值转换
			vInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("unknown err: ", err)
				return
			}
			stack = append(stack, vInt)
			continue
		case "+":
			result = stack[len(stack)-2] + stack[len(stack)-1]
		case "-":
			result = stack[len(stack)-2] - stack[len(stack)-1]
		case "*":
			result = stack[len(stack)-2] * stack[len(stack)-1]
		case "/":
			result = stack[len(stack)-2] / stack[len(stack)-1]
		}
		stack = stack[:len(stack)-2]  // 弹出栈顶两元素
		stack = append(stack, result) // 将栈顶两元素运算结果添加进栈
	}
	return
}

// 中缀表达式转后缀表达式
func midConvertRear(expression string) (rear []string) {
	var stack []string
	for i := 0; i < len(expression); i++ {
		v := strings.TrimSpace(string(expression[i]))
		switch {

		case v >= "0" && v <= "9":
			var temp string
			for ; i < len(expression); i++ { // 处理大于9的数值
				va := strings.TrimSpace(string(expression[i]))
				if va < "0" || va > "9" {
					break
				}
				temp = fmt.Sprintf("%s%s", temp, va)
			}
			rear = append(rear, temp)

		case v == "(":
			stack = append(stack, v)

		case v == ")":
			// for k, v := range stack { // 将从“(”开始的符号都弹出栈,伪栈(应该从栈顶弹出)
			// 	if v == "(" {
			// 		rear = append(rear, stack[k+1:]...)
			// 		stack = stack[:k]
			// 	}
			// }
			for k := len(stack) - 1; k >= 0; k-- {
				if stack[k] != "(" { // 不是"("就一直弹出
					rear = append(rear, stack[k])
					stack = stack[:k]
					continue
				}
				// 否则,将"("弹出，结束
				stack = stack[:k]
				break

			}
		case v == "+" || v == "-":
			if len(stack) == 0 { // 空栈直接添加
				stack = append(stack, v)
				continue
			}

			for i := len(stack) - 1; i >= 0; i-- { // 从栈顶处理
				topE := stack[i]
				if topE == "(" { // 直接添加
					break
				}
				// 将>=其优先级的符号都弹出，再将其本身添加符号栈
				rear = append(rear, topE)
				stack = stack[:i]
				continue
			}
			stack = append(stack, v)

		case v == "*" || v == "/":
			if len(stack) == 0 { // 空栈直接添加
				stack = append(stack, v)
				continue
			}

			for i := len(stack) - 1; i >= 0; i-- { // 从栈顶处理
				topE := stack[i]
				if topE == "(" || topE == "+" || topE == "-" { // 比其优先级低，就直接添加到符号栈
					break
				}

				// 将>=其优先级的符号都弹出,再将其本身添加进去
				rear = append(rear, topE)
				stack = stack[:i]
				continue
			}
			stack = append(stack, v)
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		rear = append(rear, stack[i]) // 注：此处一定要从栈顶弹出
	}

	return
}
