package tree

import (
	pathPkg "path/filepath"
	"sort"
	"strings"
)

// 根据文件路径构建目录树

type PathNode struct {
	Name     string      `json:"name"`
	FullPath string      `json:"path"`
	Value    string      `json:"value"`
	DiffType string      `json:"diff_type"`
	Children []*PathNode `json:"children"`
}

func NewPathNode(name string, value string, fullPath string) *PathNode {
	return &PathNode{
		Name:     name,
		Value:    value,
		FullPath: fullPath,
		Children: []*PathNode{},
	}
}

func (n *PathNode) Insert(fullPath string, value string) {
	n.insert(fullPath, value, fullPath)
}

//根据路径插入值
func (n *PathNode) insert(fullPath string, value string, originPath string) {
	if string(fullPath[0]) == "/" {
		fullPath = fullPath[1:]
	}
	ps := strings.Split(fullPath, "/")
	length := len(ps)
	isExist := false
	for _, p := range n.Children {
		if ps[0] == p.Name {
			isExist = true
			if length == 1 {
				p.Value = value
				p.FullPath = originPath
			}
			if length > 1 {
				//p.FullPath = originPath[:strings.Index(originPath, ps[0])+len(ps[0])]
				p.FullPath = originPath[:strings.Index(originPath, fullPath)+len(ps[0])]
				p.insert(fullPath[len(ps[0]):], value, originPath)
			}
		}
	}
	if !isExist {
		newPathNode := &PathNode{
			Name:     strings.TrimSpace(ps[0]),
			Children: []*PathNode{},
		}
		if length == 1 {
			newPathNode.Value = value
			newPathNode.FullPath = strings.TrimSpace(originPath)
		} else if length > 1 {
			//newPathNode.Path = originPath[:strings.LastIndex(originPath, ps[0])+len(ps[0])]
			newPathNode.FullPath = originPath[:strings.Index(originPath, fullPath)+len(ps[0])]
			newPathNode.insert(fullPath[len(ps[0]):], value, originPath)
		}
		n.Children = append(n.Children, newPathNode)
	}
	return
}

func (n *PathNode) Sign(fullPath string, value string) {
	fullPath = pathPkg.Dir(fullPath)
	n.sign(fullPath, value, fullPath)
}

//根据路径插入值
func (n *PathNode) sign(fullPath string, value string, originPath string) {
	if string(fullPath[0]) == "/" {
		fullPath = fullPath[1:]
	}
	ps := strings.Split(fullPath, "/")
	length := len(ps)

	//isExist := false
	for _, p := range n.Children {
		if ps[0] == p.Name {
			//isExist = true
			if length == 1 {
				//p.Value = value
				//p.Path = originPath
				p.DiffType = value
			}
			if length > 1 {
				//p.Path = originPath[:strings.Index(originPath, ps[0])+len(ps[0])]
				//p.FullPath = originPath[:strings.Index(originPath, fullPath)+len(ps[0])]
				p.sign(fullPath[len(ps[0]):], value, originPath)
			}
		}
	}
	//if !isExist {
	//	newPathNode := &PathNode{
	//		Name:     strings.TrimSpace(ps[0]),
	//		Children: []*PathNode{},
	//	}
	//	if length == 1 {
	//		//newPathNode.Value = value
	//		//newPathNode.FullPath = strings.TrimSpace(originPath)
	//		newPathNode.DiffType = value
	//	} else if length > 1 {
	//		//newPathNode.Path = originPath[:strings.LastIndex(originPath, ps[0])+len(ps[0])]
	//		newPathNode.FullPath = originPath[:strings.Index(originPath, fullPath)+len(ps[0])]
	//		newPathNode.sign(fullPath[len(ps[0]):], value, originPath)
	//	}
	//	n.Children = append(n.Children, newPathNode)
	//}
	return
}

type item struct {
	Path    string
	Content string
}

func BuildPathTree(fileList []item) *PathNode {
	node := NewPathNode("", "", "")
	//node.Insert("1/hello/world1/1/1/b.txt", "")
	//node.Insert("1/hello/world1/1/1/b.txt", "路")
	//node.Insert("1/hello/world ", `{"env":"dev"}`)
	//node.Insert("hello1/world", `{"env":"dev"}`)
	//node.Insert("/hello1/world/haha", `{"env":"dev"}`)
	//node.Insert("/hello1/world1/haha", `{"env":"dev"}`)
	//node.Insert("/hello2/world", `{"env":"dev3"}`)
	//node.Insert("/hello3/world3/haha/hehe/aa", `{"env":"dev"}`)
	for _, file := range fileList {
		node.Insert(file.Path, file.Content)
	}
	return node
}

func SignDirDelta(fileList []item, pathNode *PathNode) {
	if pathNode == nil {
		return
	}
	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].Path < fileList[j].Path
	})
	for _, v := range fileList {
		if v.Path == ".gitkeep" {
			continue
		}
		pathNode.Sign(v.Path, v.Content)
	}
}
