package main

import (
	"log"
	"math/rand"
	"path"
	"strings"
	"time"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var stus = []*Student{
		{"wafa", 20},
		{"king", 30},
	}
	opt(stus)
	for _, v := range stus {
		log.Printf("v:%#v", v)
		log.Printf("%#v", v)
	}

	rand.Seed(time.Now().UnixNano())
	log.Println(rand.Intn(3))
	base := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	log.Println(len(base))
	GetRandomStr2(3)

	p := "12345"
	log.Println(p[:3] + "xxxx" + p[len(p)-4:])

	//log.Println(time.Now().Unix()-1648637864 > 86400)
	//treeNames, treePaths := getParentTreeFields("conf/dir")
	//log.Printf("%+v, %+v", treeNames, treePaths)
	//var pa = "base/.gitkeep"
	var pa = "/base/a.gitkeep"
	log.Println(path.Ext(pa))
	log.Println(path.Split(pa))
}

func opt(stus []*Student) {
	for k, v := range stus {
		if v.Name == "wafa" {
			stus[k].Age++
		} else if v.Name == "king" {
			v.Age++
		}
		log.Println(v)
	}
}

func getParentTreeFields(treePath string) (treeNames []string, treePaths []string) {
	if len(treePath) == 0 {
		return treeNames, treePaths
	}

	treeNames = strings.Split(treePath, "/")
	treePaths = make([]string, len(treeNames))
	for i := range treeNames {
		treePaths[i] = strings.Join(treeNames[:i+1], "/")
	}
	return treeNames, treePaths
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
