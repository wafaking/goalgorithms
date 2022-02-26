package str

import "fmt"

func SwapTwoNum(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
