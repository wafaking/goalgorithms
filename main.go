package main

import (
	"log"
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
}
