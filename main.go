package main

import (
	"log"
	"math"
	"runtime"
	//_ "net/http/pprof"
	//"github.com/wafaking/"
)

func main() {
	var p uint32 = 1
	q := p << 1
	m := q << 1
	n := m << 1
	log.Println(p, q)
	log.Println(1&11, q&11, m&11, n&11)
	log.Println(runtime.GOOS)

	//var str string = ""
	//var sli [][]int64
	//num := 170797197
	num := 1170797197
	//num := 7197
	// 170797197:    "一亿七千零七十九万七千一百九十七",
	// 1170797197:    "一十一亿七千零七十九万七千一百九十七",
	//w1 := num % 100000000
	w1 := num % 10000
	num /= 10000
	w2 := num % 10000
	log.Println("w1:", w1)
	log.Println("w2: ", w2)

	var num1 = 100070797197
	temp := math.MaxInt64
	log.Println(num1 > temp)

}

type Stu struct {
	Name string
	Age  int
}
