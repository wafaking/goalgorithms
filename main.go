package main

import "log"

func main() {
	var a = 'A'
	var b = 'a'
	v := b - a
	log.Printf("a: %d, %b, b: %d, %b, %c", a, a, b, b, a^(v))
	//1000001 65
	// 100000
	//1100001 97
}

type Student struct {
}
