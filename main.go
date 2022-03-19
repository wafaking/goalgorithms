package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println(rand.Intn(3))
	base := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	log.Println(len(base))
	GetRandomStr2(3)
}

const (
	base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	//l    = len(base)
)

func GetRandomStr2(length int) {

	//var words = []byte(base)
	var words = strings.Fields("we are family")
	log.Println(words)
	//var l = len(words)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	log.Println(words)

	return
}

func GetRandomStr(length int) string {
	if length <= 0 {
		return ""
	}

	var res string
	var l = len(base)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		res += string(base[rand.Intn(l)])
	}
	return res
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
