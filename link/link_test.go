package link

import (
	"log"
	"testing"
)

func TestMain(t *testing.M) {

	ssli := [][]int{
		{1, 2, 2, 3, 3, 4, 5},
		{8, 4, 0, 6, 5, 6, 5, 7, 9},
		{1},
		{},
	}
	for _, sli := range ssli {
		SetListNode(sli)
		PrintHead()
		t.Run()
	}
}

func TestPrintListNode(t *testing.T) {
	PrintHead()
}

func TestDeleteDuplicatesAll(t *testing.T) {
	deleteDuplicatesAll(head)
	PrintHead()
}

func TestReversePrint(t *testing.T) {
	//reversePrint(head)
	//res := reversePrint23(head)
	//log.Println("res: ", res)

	//log.Println("--------split-----------split------")

	reversePrint4(head)
	//log.Println("res: ", res)

}
func TestGetKthFromEnd(t *testing.T) {
	res := getKthFromEnd(head, 2)
	log.Printf("res: %#v", res)
}
func TestAddTailNode(t *testing.T) {
	addTailNode(100)
	PrintHead()
}
func TestRemoveNode(t *testing.T) {
	removeNode(1)
	PrintHead()
}

func TestDeleteValue(t *testing.T) {
	deleteValue(head, 1)
	PrintHead()
}
func TestReverseList(t *testing.T) {
	reverseList()
	PrintHead()
}

func TestMiddleNode(t *testing.T) {
	middleNode()
	PrintHead()
}

func TestDeleteNode(t *testing.T) {
	// 此测试方法不准确
	node := &ListNode{Val: 1}
	deleteNode(node)
	PrintHead()
}
