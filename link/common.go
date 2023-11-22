package link

import (
	"fmt"
	"goalgorithms/common"
	"log"
)

var head *common.ListNode

func BuildListNode(sli []int) *common.ListNode {
	dummy := &common.ListNode{Val: -1}
	cur := dummy
	for _, v := range sli {
		cur.Next = &common.ListNode{Val: v}
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

func PrintList2(list *common.ListNode) []int {
	var (
		cur = list
		sli = make([]int, 0)
	)

	for cur != nil {
		sli = append(sli, cur.Val)
		cur = cur.Next
	}
	return sli
}

func PrintList(list *common.ListNode) {
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
		head = &common.ListNode{Val: val, Next: nil}
		return
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &common.ListNode{Val: val}
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
