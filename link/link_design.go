package link

// 设计链表(leetcode-707)
// 你可以选择使用单链表或者双链表，设计并实现自己的链表。
// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
// 实现 MyLinkedList类：
// 	MyLinkedList()初始化 MyLinkedList 对象。
// 	int get(int index)获取链表中下标为index的节点的值。如果下标无效，则返回-1。
// 	void addAtHead(int val) 将一个值为val的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
// 	void addAtTail(int val) 将一个值为val的节点追加到链表中作为链表的最后一个元素。
// 	void addAtIndex(int index, int val)将一个值为val的节点插入到链表中下标为index的节点之前。
// 		如果index等于链表的长度，那么该节点会被追加到链表的末尾。
// 		如果 index比长度更大，该节点将不会插入到链表中。
// 	void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为index的节点。
// 示例：
// 输入: ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
// 	[[], [1], [3], [1, 2], [1], [1], [1]]
// 输出: [null, null, null, null, 2, null, 3]
// 解释
// 	MyLinkedList myLinkedList = new MyLinkedList();
// 	myLinkedList.addAtHead(1);
// 	myLinkedList.addAtTail(3);
// 	myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
// 	myLinkedList.get(1);              // 返回 2
// 	myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
// 	myLinkedList.get(1);              // 返回 3

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  -1,
		Next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	for cur, counter := this, 0; cur.Next != nil; cur, counter = cur.Next, counter+1 {
		if index == counter {
			return cur.Next.Val
		}
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	next := this.Next
	this.Next = &MyLinkedList{
		Val:  val,
		Next: next,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	var cur = this
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &MyLinkedList{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	for cur, counter := this, 0; cur != nil; counter++ {
		if counter == index {
			next := cur.Next
			node := &MyLinkedList{Val: val, Next: next}
			cur.Next = node
			break
		}
		cur = cur.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	for cur, counter := this, 0; cur.Next != nil; cur, counter = cur.Next, counter+1 {
		if counter == index {
			next := cur.Next.Next
			cur.Next = next
			break
		}
	}
}

func (this *MyLinkedList) Print() []int {
	var res = make([]int, 0)
	for cur := this; cur.Next != nil; cur = cur.Next {
		res = append(res, cur.Next.Val)
	}
	return res
}
