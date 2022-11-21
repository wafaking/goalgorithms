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

type item struct {
	list1, list2, expected *ListNode
}

type array struct {
	list1, list2, expected []int
}

type itemExt struct {
	array
	start, end int
}

func SetListNode(sli []int) {
	head = makeListNode(sli)
}

func makeListNode(sli []int) *ListNode {
	dummy := &ListNode{-1, nil}
	cur := dummy
	for _, v := range sli {
		cur.Next = &ListNode{v, nil}
		cur = cur.Next
	}
	return dummy.Next
}

func PrintHead() {
	var valSli []int
	cur := head
	for cur != nil {
		valSli = append(valSli, cur.Val)
		cur = cur.Next
	}
	fmt.Println("***********************************")
	log.Printf("*******NodeList: %#v ********\n", valSli)
	fmt.Println("***********************************")
}

func PrintList2(list *ListNode) []int {
	var (
		cur = list
		sli = []int{}
	)

	for cur != nil {
		sli = append(sli, cur.Val)
		cur = cur.Next
	}
	return sli
}

func PrintList(list *ListNode) {
	var (
		cur = list
		sli = []int{}
	)

	for cur != nil {
		sli = append(sli, cur.Val)
		cur = cur.Next
	}
	fmt.Println("array: ", sli)
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
