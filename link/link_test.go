package link

import (
	"log"
	"testing"
)

func TestMain(t *testing.M) {

	//ssli := [][]int{
	//	{1, 2, 2, 3, 3, 4, 5},
	//	{8, 4, 0, 6, 5, 6, 5, 7, 9},
	//	{1},
	//	{},
	//}
	//for _, sli := range ssli {
	//	SetListNode(sli)
	//	PrintHead()
	t.Run()
	//}
}

func TestPrintListNode(t *testing.T) {
	PrintHead()
}

func TestDeleteDuplicatesAll(t *testing.T) {
	deleteDuplicatesAll(head)
	PrintHead()
}

func TestReversePrint(t *testing.T) {
	var (
		list = [][]int{
			{1, 2, 3},
			{1, 2, 3, 4},
		}
	)
	for _, sli := range list {
		res := reversePrint1(makeListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
		res = reversePrint2(makeListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
		res = reversePrint3(makeListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
		res = reversePrint4(makeListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
	}
	//reversePrint4(head)
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
	var list = []array{
		//{list1: []int{1}, list2: nil, expected: []int{1}},
		//{list1: []int{1, 2}, list2: nil, expected: []int{2}},
		{list1: []int{1, 2, 3, 4, 5}, list2: nil, expected: []int{3, 4, 5}},
		// 1->3, 2->5
		// 1->3, 2->5, 3->nil
		{list1: []int{1, 2, 3, 4, 5, 6}, list2: nil, expected: []int{4, 5, 6}},
	}
	for _, item := range list {
		res := middleNode1(makeListNode(item.list1))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.expected)
	}
	t.Log("------------SPLIT-----------")
	for _, item := range list {
		res := middleNode2(makeListNode(item.list1))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.expected)
	}
}

func TestDeleteNode(t *testing.T) {
	// 此测试方法不准确
	node := &ListNode{Val: 1}
	deleteNode(node)
	PrintHead()
}

func TestMergeTwoLists(t *testing.T) {
	var list = []array{
		{list1: []int{}, list2: []int{}, expected: []int{}},
		{list1: []int{1}, list2: []int{3}, expected: []int{1, 3}},
		{list1: []int{1}, list2: []int{1, 3, 4}, expected: []int{1, 1, 3, 4}},
	}
	for _, item := range list {
		res := mergeTwoLists1(makeListNode(item.list1), makeListNode(item.list2))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.expected)
	}
	t.Log("------------SPLIT-----------")
	for _, item := range list {
		res := mergeTwoLists2(makeListNode(item.list1), makeListNode(item.list2))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.expected)
	}
	t.Log("------------SPLIT-----------")
	for _, item := range list {
		res := mergeTwoLists3(makeListNode(item.list1), makeListNode(item.list2))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.expected)
	}
	t.Log("------------SPLIT-----------")
}

func TestAddTwoNumbers1(t *testing.T) {
	var list = []array{
		{list1: []int{2, 4, 3}, list2: []int{5, 6, 4}, expected: []int{7, 0, 8}},
		{list1: []int{0}, list2: []int{0}, expected: []int{0}},
		{list1: []int{9, 9, 9, 9, 9, 9, 9}, list2: []int{9, 9, 9, 9}, expected: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, item := range list {
		res := addTwoNumbers11(makeListNode(item.list1), makeListNode(item.list2))
		t.Logf("res: %v, expect: %v", PrintList2(res), item.expected)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	//s1 := []int{1, 2, 4}
	//s1 := []int{7, 2, 4, 2}
	s1 := []int{7, 2, 4, 2}
	//s2 := []int{1, 3, 4}
	s2 := []int{5, 6, 4}
	list1 := makeListNode(s1)
	PrintList(list1)
	list2 := makeListNode(s2)
	PrintList(list2)

	//res := addTwoNumbers1(list1, list2)
	//res := addTwoNumbers2(list1, list2)
	res := addTwoNumbers23(list1, list2)
	PrintList(res)
}

func TestMergeInBetween(t *testing.T) {
	var array = array{
		list1:    []int{0, 1, 2, 3, 4, 5, 6},
		list2:    []int{1000000, 1000001, 1000002},
		expected: []int{0, 1, 1000000, 1000001, 1000002, 5, 6},
	}
	var list = [][2]int{
		//{2, 4},
		//{4, 2},
		{-2, -1},
		//{0, 6},
	}

	for _, item := range list {
		list1 := makeListNode(array.list1)
		list2 := makeListNode(array.list2)
		res := mergeInBetween1(list1, item[0], item[1], list2)
		t.Logf("res: %v, expected: %+v", PrintList2(res), array.expected)
		t.Logf("---------------SPLIT-------------")

		list1 = makeListNode(array.list1)
		list2 = makeListNode(array.list2)
		res = mergeInBetween2(list1, item[0], item[1], list2)
		t.Logf("res: %v, expected: %+v", PrintList2(res), array.expected)

		t.Logf("---------------SPLIT-------------")
	}
}
