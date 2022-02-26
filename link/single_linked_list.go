package link

/*
	线性表--链式存储结构的实现与基本操作
*/

type SqSingleListOPeration interface {
	//InitList() *SingleLink              // 初始化
	IsEmpty() bool                      // 判断是否为空,true/false
	Length() int                        // 长度
	ClearList()                         // 清空
	Add(data string)                    // 头部插入
	Append(data string)                 // 尾部插入
	Insert(index int, data string) bool // 插入
	GetElem(i int) *Node                // 获取元素
	LocateElem(Node) int                // 查找元素位置
	Remove(int) (bool, *Node)           // 删除
	CreateSingleList(length int)        // 创建一个指定长度的随机单链表
}

/*
// Node 单结构
type Node struct {
	Next *Node
	Data string
}

type SingleLink struct {
	Size int
	Data *Node
}

func (sq *SingleLink) InitList() *SingleLink {
	return new(SingleLink)
}

func (sq *SingleLink) Length() (length int) {
	if sq == nil {
		sq = sq.InitList()
	}
	return sq.Size
}

func (sq *SingleLink) IsEmpty() bool {
	return sq.Length() == 0
}

func (sq *SingleLink) Add(data string) {
	sq.IsEmpty()

	node := &Node{Data: data}
	node.Next = sq.Data
	sq.Data = node
	sq.Size++
	return
}

func (sq *SingleLink) Append(data string) {

	node := &Node{Data: data}
	if sq.IsEmpty() {
		sq.Data = node
		sq.Size++
		return
	}

	cur := sq.Data
	for cur.Next != nil {
		cur = cur.Next
		continue
	}
	cur.Next = node
	sq.Size++
	return
}

func (sq *SingleLink) Insert(pos int, data string) {
	if sq.IsEmpty() {
		sq.Add(data)
		return
	}

	if pos <= 0 { // 头部插入
		sq.Add(data)
		return
	} else if pos > sq.Size { // 尾部插入
		sq.Append(data)
		return
	}

	var curNode = sq.Data
	for i := 0; i < pos-1; i++ { // 如：pos=3，需要定位到index=2的位置；
		curNode = curNode.Next
	}

	node := &Node{Data: data}
	node.Next = curNode
	curNode.Next = node
	sq.Size++

	return
}

func (sq *SingleLink) GetElem(pos int) (has bool, data string) {
	if sq.Size <= pos || pos <= 0 {
		return
	}

	node := sq.Data
	for i := 0; i < sq.Size; i++ {
		if pos == (i + 1) {
			return true, node.Data
		}
		node = node.Next
		continue
	}
	return
}

func (sq *SingleLink) Remove(i int) (bool, *Node) {
	var (
		hasRomoved bool
		rNode      *Node
	)
	if sq.Size <= i || i <= 0 {
		return hasRomoved, rNode
	}

	var inode = sq.Data

	for num := 0; num < i; num++ {
		num++
		curNode := inode
		if num != i {
			inode = inode.Next
			continue
		}
		curNode.Next = inode.Next
		rNode = inode
		hasRomoved = true
		sq.Size--
	}
	return hasRomoved, rNode
}

func (sq *SingleLink) ClearList() {
	if sq.Size == 0 {
		return
	}
	sq.Size = 0
	sq.Data = nil
	return
}

func (sq *SingleLink) LocateElem(node Node) int {
	var index = -1
	if sq.Size == 0 {
		return index
	}

	var inode = sq.Data
	for num := 0; num < sq.Size; num++ {
		num++
		if node != *inode {
			inode = inode.Next
			continue
		}
		index = num
	}
	return index
}


*/
