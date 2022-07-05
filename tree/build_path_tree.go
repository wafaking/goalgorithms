package tree

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type PathNode struct {
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Value    string      `json:"value"`
	Children []*PathNode `json:"children"`
}

func NewPathNode(name string, value string, path string) *PathNode {
	return &PathNode{
		Name:     name,
		Value:    value,
		Path:     path,
		Children: []*PathNode{},
	}
}

func (n *PathNode) Insert(path string, value string) {
	n.insert(path, value, path)
}

//根据路径插入值
func (n *PathNode) insert(path string, value string, orginPath string) {
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
		newPathNode := &PathNode{
			Name:     strings.TrimSpace(ps[0]),
			Children: []*PathNode{},
		}
		if length == 1 {
			newPathNode.Value = value
			newPathNode.Path = strings.TrimSpace(orginPath)
		} else if length > 1 {
			newPathNode.insert(path[len(ps[0]):], value, orginPath)
			newPathNode.Path = orginPath[:strings.LastIndex(orginPath, ps[0])+len(ps[0])]
		}
		n.Children = append(n.Children, newPathNode)
	}
	return
}

func BuildPathTree() {
	node := NewPathNode("", "", "")
	node.Insert("hello/world", `{"env":"dev"}`)
	node.Insert("hello/world ", `{"env":"dev"}`)
	//node.Insert("hello1/world", `{"env":"dev"}`)
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
