package str

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lastStrLength() int {
	r := bufio.NewReader(os.Stdin)
	l, _, err := r.ReadLine()
	if err != nil {
		fmt.Printf("err:%s", err)
		return 0
	}
	res := strings.Split(string(l), " ")
	return len(res[len(res)-1])
}
