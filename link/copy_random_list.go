package link

// 复杂链表的复制(剑指 Offer 35)
// 请实现copyRandomList函数，复制一个复杂链表。
//在复杂链表中，每个节点除了有一个next指针指向下一个节点，还有一个random指针指向链表中的任意节点或者null。

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 法一：遍历(先生成链表，再遍历节点，根据Random节点所在位置赋值给新的Random)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 1. 先生成新链表主节点
	cur := head
	newHead := &Node{Val: cur.Val}
	newCur := newHead

	for cur.Next != nil {
		node := &Node{Val: cur.Next.Val}
		newCur.Next = node

		cur = cur.Next
		newCur = newCur.Next
	}

	// 2. 找出原链表节点所指的Random节点的位置，再赋值给新链表的Random节点
	cur = head // 此处要从头开始，上次遍历已将cur移动到了最后位置
	newCur = newHead
	for cur != nil {
		if cur.Random != nil {
			// 1. 取出Random的位置
			position := GetRandomPosition(head, cur.Random)
			// 2. 根据位置获取新节点的值
			newCur.Random = SetRandomByPosition(newHead, position)
		}
		cur = cur.Next
		newCur = newCur.Next
	}

	return newHead
}

func GetRandomPosition(head, random *Node) int {
	cur := head
	position := 0
	for cur != random {
		position++
		cur = cur.Next
	}
	return position
}

func SetRandomByPosition(head *Node, position int) *Node {
	cur := head
	for position > 0 {
		cur = cur.Next
		position--
	}
	return cur
}

// 法二：使用map记录
// 注：使用老的A节点映射新的节点A1
func copyRandomList2(head *Node) *Node {
	// 1. 第一次遍历，生成新的链表，并将旧、新节点映射
	if head == nil {
		return head
	}

	mapLink := map[*Node]*Node{}
	cur := head
	NewHead := &Node{Val: head.Val} // 新链表头节点
	newCur := NewHead
	mapLink[cur] = newCur // 老、新关系映射

	for cur.Next != nil {
		newNode := &Node{Val: cur.Next.Val}
		newCur.Next = newNode
		mapLink[cur.Next] = newNode

		newCur = newCur.Next
		cur = cur.Next
	}

	// 2. 更新新节点的Random节点
	cur = head       // 此片要从头开始，上次遍历已将cur移动到了最后位置
	newCur = NewHead // 同上
	for cur != nil {
		if cur.Random != nil {
			newCur.Random = mapLink[cur.Random]
		}
		newCur = newCur.Next
		cur = cur.Next
	}

	return NewHead
}

// 法三：在每个节点后添加一个新节点，再复制随机节点，最后再拆分奇偶节点
func copyRandomList3(head *Node) *Node {
	if head == nil {
		return head
	}
	// 1. 复制节点
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = cur.Next.Next
	}

	//2. 复制Random
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 3. 拆分奇、偶节点
	newNode := head.Next // 新的链表
	oldCur := head       // cur和newCur都是用于遍历的当前节点，用newNode自己处理后，返回的是最后节点的val
	newCur := newNode

	for newCur.Next != nil {
		oldCur.Next = oldCur.Next.Next
		newCur.Next = newCur.Next.Next

		//改变当前节点
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	// 最后一次遍历，当newCur.Next=nil时，oldCur.Next仍旧指向newCur，
	// 因此需要让oldCur.Next指向nil
	oldCur.Next = nil

	return newNode
}
