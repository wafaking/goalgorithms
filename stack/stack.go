package stack

type Stacker interface {
	InitStack()
	DestroyStack()
	ClearStack()
	StackEmpty() bool // 是否为空
	GetTop() interface{}
	Push(data interface{}) interface{}
	Pop() interface{}
	StackLength() int
}

// ListStack 顺序栈
type ListStack struct {
	Top  int           // 栈顶指针
	Data []interface{} // 数据列表
}

func (s *ListStack) InitStack() {
	s = new(ListStack)
	s.Data = make([]interface{}, 0)
	return
}

func (s *ListStack) StackLength() int {
	if s == nil {
		panic("invalid stack")
	}
	return len(s.Data)
}

func (s *ListStack) GetTop() interface{} {
	if s == nil || s.Top == 0 {
		return nil
	}
	return s.Data[s.Top]
}

func (s *ListStack) Push(data interface{}) interface{} {
	if s == nil {
		s.InitStack()
	}
	s.Data = append(s.Data, data)
	s.Top++
	return data
}

func (s *ListStack) Pop() interface{} {
	if s == nil || s.Top == 0 {
		return nil
	}
	v := s.Data[s.Top]
	s.Data = s.Data[:s.Top]
	return v
}

func (s *ListStack) DestroyStack() {
	s = nil
	return
}

func (s *ListStack) ClearStack() {
	if s == nil || s.Top == 0 {
		return
	}
	s.Top = 0
	s.Data = []interface{}{}
	return
}

// 是否为空
func (s *ListStack) StackEmpty() bool {
	if s == nil {
		return true
	}
	return s.Top == 0
}
