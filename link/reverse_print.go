package link

import (
	"fmt"
	"log"
)

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

//示例 1：
//输入：head = [1,3,2]
//输出：[2,3,1]

//法一： 顺序遍历，将元素添加至栈中，再将栈中元素输出；
func reversePrint1(head *ListNode) []int {
	dummy := head
	var sli []int
	for dummy != nil {
		sli = append(sli, dummy.Val)
		dummy = dummy.Next
	}

	if len(sli) <= 1 {
		return sli
	}
	n := len(sli)
	for i := 0; i < n/2+n%2; i++ { // 注意此处需要判别奇偶数
		sli[i], sli[n-1-i] = sli[n-1-i], sli[i]
	}
	//交换方法2
	//var start, end = 0, n - 1
	//for start < end {
	//	sli[start], sli[end] = sli[end], sli[start]
	//	start++
	//	end--
	//}
	return sli
}

//法二：使用defer打印(错误)
func reversePrint21(head *ListNode) []int {
	var (
		cur = head
		sli = []int{}
	)
	for cur != nil {
		defer func(val int) {
			sli = append(sli, val)
		}(cur.Val)
		cur = cur.Next
	}
	return sli
}

//法二：使用defer打印(正确)
func reversePrint22(head *ListNode) (sli []int) {
	var (
		cur = head
	)
	for cur != nil {
		defer func(val int) {
			sli = append(sli, val)
		}(cur.Val)
		cur = cur.Next
	}
	return sli
}

//法二：使用defer打印(正确)
func reversePrint23(head *ListNode) []int {
	return handleNode(head)
}

func handleNode(head *ListNode) (sli []int) {
	cur := head
	for cur != nil {
		defer func(val int) {
			sli = append(sli, val)
		}(cur.Val)
		cur = cur.Next
	}
	return sli
}

func reversePrint2(head *ListNode) (sli []int) {
	dummy := head
	for dummy != nil {
		val := dummy.Val
		defer func() {
			//sli = append(sli, dummy.Val) // 错误写法
			sli = append(sli, val)
		}()
		dummy = dummy.Next
	}
	return sli
}

//法三：将链表反转再输出

// 法四：使用递归
func reversePrint4(head *ListNode) {
	cur := head
	printListReverse4(cur)
	return
}

func printListReverse4(node *ListNode) {
	if node != nil {
		printListReverse4(node.Next)
		log.Println(node.Val)
	}
}

func reversePrint(head *ListNode) {
	dummy := head
	printListReverse(dummy)
}

func printListReverse(node *ListNode) {
	if node != nil {
		if node.Next != nil {
			printListReverse(node.Next)
		}
		fmt.Println("val: ", node.Val)
	}
}
