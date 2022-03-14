package main

import (
	"fmt"
	"log"
)

func main() {
	var stu = &Stu{
		Name: "wafa",
		Age:  10,
	}
	stu2 := stu
	fmt.Printf("%p, %p\n", stu, stu2)
	var num = 10
	num1 := &num
	fmt.Printf("%p, %p\n", &num, num1)

	var s = "abcd"
	s = s[2:]
	log.Println(s)

}

func reverse(str string) string {
	bt := []byte(str)
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		bt[start], bt[end] = bt[end], bt[start]
	}
	return string(bt)
}

type Stu struct {
	Name string
	Age  int
}

// 0, 1, 1, 2, 3, 5, 8
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
