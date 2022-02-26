package link

import (
	"fmt"
	"log"
)

var head *ListNode

// ListNode 结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func SetListNode(sli []int) {
	head = makeListNode(sli)
}

func makeListNode(sli []int) *ListNode {

	dummy := &ListNode{-1, nil}
	cur := dummy
	for _, v := range sli {
		node := &ListNode{v, nil}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

func PrintHead() {
	var valSli []int
	//fmt.Println("head11: ", head)
	cur := head
	for cur != nil {
		valSli = append(valSli, cur.Val)
		cur = cur.Next
	}
	fmt.Println("***********************************")
	log.Printf("*******NodeList: %#v ********\n", valSli)
	fmt.Println("***********************************")
}

// 链表添加尾部元素
func addTailNode(val int) {
	if head == nil {
		head = &ListNode{Val: val, Next: nil}
		return
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
}

func removeNode(val int) {
	if head == nil {
		return
	}
	if head.Val == val {
		head = head.Next
		return
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}
	return
}
