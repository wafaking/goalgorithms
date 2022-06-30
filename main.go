package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	//urls := []string{}
	////初始化urls
	//
	//var (
	//	ch   = make(chan string, 100)
	//	wg   *sync.WaitGroup
	//	nums = 100
	//)
	//
	//wg.Add(2)
	//go producer(wg, ch, urls)
	//go consumer(wg, ch, nums)
	//
	//wg.Wait()
	//var path = "111/222/a.txt"
	//var path = "111/222/a"
	var path = "111/222/a.gitkeep"
	log.Printf("ext: %#v", filepath.Ext(path))
	log.Printf("base: %#v", filepath.Base(path))
	dir, name := filepath.Split(path)
	log.Printf("dir: %#v ,name: %#v", dir, name)
	log.Printf("%#v", filepath.SplitList(path))

	if filepath.Base(path) == ".gitkeep" {
		log.Printf("%v ----------", filepath.Dir(path))
	}

	//match()
	//fileMain()
}

type Student struct {
	Name string
	Age  int
	Sub  *Student
}

type FileInfo struct {
	Name string
	Path string
	Age  int
	Sub  *FileInfo
}

type MFileInfo struct {
	Name string
	Ojb  *MFileInfo
}

func match() {
	var nancy = FileInfo{
		Name: "111",
		Path: "111",
		Age:  1,
		Sub: &FileInfo{
			Name: "222",
			Path: "222",
			Age:  1,
			Sub: &FileInfo{
				Name: "readme.md",
				Path: "111/222/readme.md",
				Age:  0,
				Sub:  nil,
			},
		},
	}

	var wafa = Student{
		Name: "111/222/hehe.md",
		Age:  0,
		Sub:  nil,
	}
	//
	//var king = Student{
	//	Name: "king",
	//	Age:  21,
	//	Sub:  nil,
	//}
	log.Printf("nancy: %#v", nancy)
	var fileNameList = strings.Split(wafa.Name, "/")
	if nancy.Name == fileNameList[0] {

		//for _, sub := range nancy.Sub {
		//
		//}
	}

}

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

//层级节点
type Node struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Value    string  `json:"value"`
	Children []*Node `json:"children"`
}

func NewNode(name string, value string, path string) *Node {
	return &Node{
		Name:     name,
		Value:    value,
		Path:     path,
		Children: []*Node{},
	}
}

func (n *Node) Insert(path string, value string) {
	n.insert(path, value, path)
}

//根据路径插入值
func (n *Node) insert(path string, value string, orginPath string) {
	if string(path[0]) == "/" {
		path = path[1:]
	}
	ps := strings.Split(path, "/")
	length := len(ps)
	is_exist := false
	for _, p := range n.Children {
		if ps[0] == p.Name {
			is_exist = true
			if length == 1 {
				p.Value = value
				p.Path = orginPath
			}
			if length > 1 {
				p.insert(path[len(ps[0]):], value, orginPath)
				p.Path = orginPath[:strings.LastIndex(orginPath, ps[0])+len(ps[0])]
			}
		}
	}
	if !is_exist {
		newNode := &Node{
			Name:     ps[0],
			Children: []*Node{},
		}
		if length == 1 {
			newNode.Value = value
			newNode.Path = orginPath
		} else if length > 1 {
			newNode.insert(path[len(ps[0]):], value, orginPath)
			newNode.Path = orginPath[:strings.LastIndex(orginPath, ps[0])+len(ps[0])]
		}
		n.Children = append(n.Children, newNode)
	}
	return
}

func fileMain() {
	node := NewNode("", "", "")
	node.Insert("hello/world", `{"env":"dev"}`)
	node.Insert("hello1/world", `{"env":"dev"}`)
	//node.Insert("/hello1/world/haha", `{"env":"dev"}`)
	//node.Insert("/hello1/world1/haha", `{"env":"dev"}`)
	//node.Insert("/hello2/world", `{"env":"dev3"}`)
	//node.Insert("/hello3/world3/haha/hehe/aa", `{"env":"dev"}`)
	res, err := json.Marshal(node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
