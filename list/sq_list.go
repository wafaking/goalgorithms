package list

/*
	线性表：顺序表的实现与基本操作
*/

// SqListOperation 基本操作
type SqListOperation interface {
	InitList() *Element           // 初始化
	ListEmpty() bool              // 判断是否为空,true/false
	ClearList()                   // 清空
	GetElem(i int) *Element       // 获取元素
	LocateElem(Element) int       // 查找元素位置
	ListInsert(int, Element) bool // 插入
	ListDelete(i int) bool        // 删除
	ListLength() int              // 长度
}

// SingleSq 单结构
type Element struct {
	Name string
	Age  int
}

type SqList struct {
	List   []Element
	Length int
}

func (sq *SqList) InitList() *SqList {
	return new(SqList)
}

func (sq *SqList) ListLength() (length int) {
	if sq != nil {
		return sq.Length
	}
	return len(sq.List)
}

func (sq *SqList) GetElem(i int) *Element {
	if sq.ListLength() < i || i < 0 {
		return nil
	}
	return &sq.List[i-1]
}

func (sq *SqList) ListEmpty() bool {
	return sq.ListLength() == 0
}

func (sq *SqList) ClearList() {
	if sq.ListLength() == 0 {
		return
	}
	sq.Length = 0
	sq.List = sq.List[:0]
}

func (sq *SqList) LocateElem(elem Element) int {
	if sq.ListLength() == 0 {
		return -1
	}
	for i := 0; i < sq.ListLength(); i++ {
		if sq.List[i] == elem {
			return i
		}
	}

	return -1
}

func (sq *SqList) ListInsert(i int, elem Element) bool {
	if sq == nil {
		return false
	}
	if sq == nil || sq.ListLength()+1 < i || i <= 0 {
		return false
	}
	sq.Length = sq.ListLength() + 1
	sq.List = append(append(sq.List[:i], elem), sq.List[i:]...)
	return true
}

func (sq *SqList) ListDelete(i int) bool {
	if sq.ListLength() < i || i < 0 {
		return false
	}
	sq.List = append(sq.List[0:i], sq.List[i+1:]...)
	sq.Length = sq.ListLength() - 1
	return true
}
