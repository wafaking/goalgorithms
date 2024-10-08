package link

import (
	"goalgorithms/common"
	"log"
	"testing"
)

// func TestMain(t *testing.M) {
//	t.Run()
// }

func TestPrintListNode(t *testing.T) {
	PrintHead()
}

func TestDeleteDuplicates(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{1, 2, 3, 3, 4, 4, 5},
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Nums:     []int{1, 1, 1, 2, 3},
			Expected: []int{1, 2, 3},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1, 1, 2, 3, 3},
			Expected: []int{1, 2, 3},
		},
		{
			Nums:     []int{1, 1},
			Expected: []int{1},
		},
	}

	var (
		head, resNode *common.ListNode
		res           = make([]int, 0)
	)

	for _, item := range list {
		head = BuildListNode(item.Nums)
		resNode = deleteDuplicates11(head)
		res = PrintList2(resNode)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		resNode = deleteDuplicates12(head)
		res = PrintList2(resNode)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
	}
	PrintHead()
}

func TestDeleteDuplicatesAll(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{1, 2, 3, 3, 4, 4, 5},
			Expected: []int{1, 2, 5},
		},
		{
			Nums:     []int{1, 1, 1, 2, 3},
			Expected: []int{2, 3},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1, 1},
			Expected: []int{},
		},
	}

	var (
		head, resNode *common.ListNode
		res           = make([]int, 0)
	)

	for _, item := range list {
		head = BuildListNode(item.Nums)
		resNode = deleteDuplicates21(head)
		res = PrintList2(resNode)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		head = BuildListNode(item.Nums)
		resNode = deleteDuplicates22(head)
		res = PrintList2(resNode)
		common.PrintSplit(t)
	}
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
		res := reversePrint1(BuildListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
		res = reversePrint2(BuildListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
		res = reversePrint3(BuildListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
		res = reversePrint4(BuildListNode(sli))
		t.Logf("res: %+v", res)
		t.Log("---------------SPLIT-------------SPLIT-------")
	}
	// reversePrint4(head)
	// log.Println("res: ", res)

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

	var (
		list = []common.Item20{
			{
				Nums:     []int{1, 2, 3},
				Expected: []int{3, 2, 1},
			},
			{
				Nums:     []int{1, 2, 3, 4},
				Expected: []int{4, 3, 2, 1},
			},
			{
				Nums:     []int{1, 2, 3, 4, 5},
				Expected: []int{5, 4, 3, 2, 1},
			},
			{
				Nums:     []int{},
				Expected: []int{},
			},
			{
				Nums:     []int{1},
				Expected: []int{1},
			},
		}
	)
	var res = make([]int, 0)
	for _, item := range list {
		res = PrintList2(reverseList1(BuildListNode(item.Nums)))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = PrintList2(reverseList2(BuildListNode(item.Nums)))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = PrintList2(reverseList3(BuildListNode(item.Nums)))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = PrintList2(reverseList4(BuildListNode(item.Nums)))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("-------------------------SPLIT-------------------------")

	}
}

func TestReverseBetween(t *testing.T) {

	var (
		list = []common.Item39{
			{
				Nums:     []int{},
				N1:       1,
				N2:       3,
				Expected: []int{},
			},
			{
				Nums:     []int{1, 2, 3, 4, 5},
				N1:       2,
				N2:       4,
				Expected: []int{1, 4, 3, 2, 5},
			},
			{
				Nums:     []int{5},
				N1:       1,
				N2:       1,
				Expected: []int{5},
			},
			{
				Nums:     []int{1, 2, 3, 4, 5},
				N1:       1,
				N2:       7,
				Expected: []int{5, 4, 3, 2, 1},
			},
			{
				Nums:     []int{1, 2, 3, 4, 5},
				N1:       1,
				N2:       5,
				Expected: []int{5, 4, 3, 2, 1},
			},
			{
				Nums:     []int{1, 2, 3},
				N1:       4,
				N2:       1,
				Expected: []int{1, 2, 3},
			},
		}
	)
	var res = make([]int, 0)
	for _, item := range list {
		res = PrintList2(reverseBetween1(BuildListNode(item.Nums), item.N1, item.N2))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = PrintList2(reverseBetween2(BuildListNode(item.Nums), item.N1, item.N2))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = PrintList2(reverseBetween3(BuildListNode(item.Nums), item.N1, item.N2))
		t.Logf("%v, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("-------------------------SPLIT-------------------------")
	}
}

func TestMiddleNode(t *testing.T) {
	var list = []common.Item19{
		// {Nums1: []int{1}, Nums2: nil, Expected: []int{1}},
		// {Nums1: []int{1, 2}, Nums2: nil, Expected: []int{2}},
		{Nums1: []int{1, 2, 3, 4, 5}, Nums2: nil, Expected: []int{3, 4, 5}},
		// 1->3, 2->5
		// 1->3, 2->5, 3->nil
		{Nums1: []int{1, 2, 3, 4, 5, 6}, Nums2: nil, Expected: []int{4, 5, 6}},
	}
	for _, item := range list {
		res := middleNode1(BuildListNode(item.Nums1))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.Expected)
	}
	t.Log("------------SPLIT-----------")
	for _, item := range list {
		res := middleNode2(BuildListNode(item.Nums1))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.Expected)
	}
}

func TestDeleteNode(t *testing.T) {
	// 此测试方法不准确
	node := &common.ListNode{Val: 1}
	deleteNode(node)
	PrintHead()
}

func TestMergeTwoLists(t *testing.T) {
	var list = []common.Item19{
		{Nums1: []int{}, Nums2: []int{}, Expected: []int{}},
		{Nums1: []int{1}, Nums2: []int{3}, Expected: []int{1, 3}},
		{Nums1: []int{1}, Nums2: []int{1, 3, 4}, Expected: []int{1, 1, 3, 4}},
	}
	for _, item := range list {
		res := mergeTwoLists1(BuildListNode(item.Nums1), BuildListNode(item.Nums2))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.Expected)
	}
	t.Log("------------SPLIT-----------")
	for _, item := range list {
		res := mergeTwoLists2(BuildListNode(item.Nums1), BuildListNode(item.Nums2))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.Expected)
	}
	t.Log("------------SPLIT-----------")
	for _, item := range list {
		res := mergeTwoLists3(BuildListNode(item.Nums1), BuildListNode(item.Nums2))
		t.Logf("res: %v, expected:%v", PrintList2(res), item.Expected)
	}
	t.Log("------------SPLIT-----------")
}

func TestAddTwoNumbers1(t *testing.T) {
	var list = []common.Item19{
		{Nums1: []int{2, 4, 3}, Nums2: []int{5, 6, 4}, Expected: []int{7, 0, 8}},
		{Nums1: []int{0}, Nums2: []int{0}, Expected: []int{0}},
		{Nums1: []int{9, 9, 9, 9, 9, 9, 9}, Nums2: []int{9, 9, 9, 9}, Expected: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, item := range list {
		res := addTwoNumbers11(BuildListNode(item.Nums1), BuildListNode(item.Nums2))
		t.Logf("res: %v, expect: %v", PrintList2(res), item.Expected)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	// s1 := []int{1, 2, 4}
	// s1 := []int{7, 2, 4, 2}
	s1 := []int{7, 2, 4, 2}
	// s2 := []int{1, 3, 4}
	s2 := []int{5, 6, 4}
	list1 := BuildListNode(s1)
	PrintList(list1)
	list2 := BuildListNode(s2)
	PrintList(list2)

	// res := addTwoNumbers1(list1, list2)
	// res := addTwoNumbers2(list1, list2)
	res := addTwoNumbers23(list1, list2)
	PrintList(res)
}

func TestMergeInBetween(t *testing.T) {
	var array = common.Item19{
		Nums1:    []int{0, 1, 2, 3, 4, 5, 6},
		Nums2:    []int{1000000, 1000001, 1000002},
		Expected: []int{0, 1, 1000000, 1000001, 1000002, 5, 6},
	}
	var list = [][2]int{
		// {2, 4},
		// {4, 2},
		{-2, -1},
		// {0, 6},
	}

	for _, item := range list {
		list1 := BuildListNode(array.Nums1)
		list2 := BuildListNode(array.Nums2)
		res := mergeInBetween1(list1, item[0], item[1], list2)
		t.Logf("res: %v, expected: %+v", PrintList2(res), array.Expected)
		t.Logf("---------------SPLIT-------------")

		list1 = BuildListNode(array.Nums1)
		list2 = BuildListNode(array.Nums2)
		res = mergeInBetween2(list1, item[0], item[1], list2)
		t.Logf("res: %v, expected: %+v", PrintList2(res), array.Expected)

		t.Logf("---------------SPLIT-------------")
	}
}

func TestGetDecimalValue(t *testing.T) {
	var array = []common.Item19{
		{
			Nums1:    []int{1, 0, 1},
			Expected: []int{5},
		},
		{
			Nums1:    []int{0, 0},
			Expected: []int{0},
		},
		{
			Nums1:    []int{1},
			Expected: []int{1},
		},
		{
			Nums1:    []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			Expected: []int{18880},
		},
	}
	for _, item := range array {
		res := getDecimalValue1(BuildListNode(item.Nums1))
		t.Logf("res: %t, %d, list:%v", res == item.Expected[0], res, item.Nums1)
		res = getDecimalValue2(BuildListNode(item.Nums1))
		t.Logf("res: %t, %d, list:%v", res == item.Expected[0], res, item.Nums1)
		res = getDecimalValue3(BuildListNode(item.Nums1))
		t.Logf("res: %t, %d, list:%v", res == item.Expected[0], res, item.Nums1)
	}
}

func TestRemoveElements(t *testing.T) {
	var list = []common.Item1{
		{
			Nums:     []int{1, 2, 6, 3, 4, 5, 6},
			Target:   6,
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Nums:     []int{},
			Target:   1,
			Expected: []int{},
		},
		{
			Nums:     []int{7, 7, 7, 7},
			Target:   7,
			Expected: []int{},
		},
	}
	var (
		res  = make([]int, 0)
		head = new(common.ListNode)
	)
	for _, item := range list {
		head = removeElements1(BuildListNode(item.Nums), item.Target)
		res = PrintList2(head)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		head = removeElements2(BuildListNode(item.Nums), item.Target)
		res = PrintList2(head)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		t.Log("--------------split----------------split--------------")
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	var list = []common.Item1{
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Target:   2,
			Expected: []int{1, 2, 3, 5},
		},
		{
			Nums:     []int{1},
			Target:   2,
			Expected: []int{1},
		},
		{
			Nums:     []int{},
			Target:   2,
			Expected: []int{},
		},
		{
			Nums:     []int{1},
			Target:   1,
			Expected: []int{},
		},
		{
			Nums:     []int{1, 2},
			Target:   1,
			Expected: []int{1},
		},
	}
	var (
		res  = make([]int, 0)
		head = new(common.ListNode)
	)
	for _, item := range list {
		head = removeNthFromEnd1(BuildListNode(item.Nums), item.Target)
		res = PrintList2(head)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		head = removeNthFromEnd2(BuildListNode(item.Nums), item.Target)
		res = PrintList2(head)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		head = removeNthFromEnd3(BuildListNode(item.Nums), item.Target)
		res = PrintList2(head)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		head = removeNthFromEnd4(BuildListNode(item.Nums), item.Target)
		res = PrintList2(head)
		t.Logf("%v, res-expected:%+v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item.Expected, item)
		t.Log("--------------split----------------split--------------")
	}
}

func TestSwapPairs(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: []int{2, 1, 4, 3, 5},
		},
		{
			Nums:     []int{1, 2, 3, 4},
			Expected: []int{2, 1, 4, 3},
		},
		{
			Nums:     []int{1},
			Expected: []int{1},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1, 2},
			Expected: []int{2, 1},
		},
	}
	var (
		res  = make([]int, 0)
		head = new(common.ListNode)
	)
	for _, item := range list {
		head = swapPairs1(BuildListNode(item.Nums))
		res = PrintList2(head)
		t.Logf("%t, res:%+v item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		common.PrintSplit(t)
	}
}

func TestHasCycle(t *testing.T) {
	var list = []common.Item8{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: false,
		},
		{
			Nums:     []int{1, 2, 2, 3, 4, 4, 3},
			Expected: true,
		},
		{
			Nums:     []int{1, 2, 2, -1, 3, -1, 3},
			Expected: false,
		},
		{
			Nums:     []int{},
			Expected: true,
		},
		{
			Nums:     []int{1},
			Expected: true,
		},
	}

	var res bool
	for _, item := range list {
		root := BuildListNode(item.Nums)
		res = hasCycle1(root)
		common.PrintDiffTwoBool(res, item.Expected, item, t)

		root = BuildListNode(item.Nums)
		res = hasCycle2(root)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}

func TestIsPalindrome(t *testing.T) {
	var list = []common.Item8{
		{
			Nums:     []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Expected: false,
		},
		{
			Nums:     []int{1, 2, 2, 1},
			Expected: true,
		},
		{
			Nums:     []int{8, 1, 6, 2, 2, 6, 1, 8},
			Expected: true,
		},
		{
			Nums:     []int{},
			Expected: true,
		},
		{
			Nums:     []int{1},
			Expected: true,
		},
	}
	var res bool
	for _, item := range list {
		res = isPalindrome1(BuildListNode(item.Nums))
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = isPalindrome2(BuildListNode(item.Nums))
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = isPalindrome3(BuildListNode(item.Nums))
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = isPalindrome4(BuildListNode(item.Nums))
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		common.PrintSplit(t)
	}
}

func TestSortList(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{0, 3, 0, 2, 3},
			Expected: []int{0, 0, 2, 3, 3},
		},
		{
			Nums:     []int{1, 3, 4, 2, 3, 4, 5},
			Expected: []int{1, 2, 3, 3, 4, 4, 5},
		},
		{
			Nums:     []int{4, 2, 1, 3},
			Expected: []int{1, 2, 3, 4},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
	}

	var (
		head, resNode *common.ListNode
		res           = make([]int, 0)
	)

	for _, item := range list {
		head = BuildListNode(item.Nums)
		resNode = sortList1(head)
		res = PrintList2(resNode)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		head = BuildListNode(item.Nums)
		resNode = sortList2(head)
		res = PrintList2(resNode)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		head = BuildListNode(item.Nums)
		resNode = sortList3(head)
		res = PrintList2(resNode)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
	PrintHead()
}

func TestLinkDesign(t *testing.T) {
	var res = make([]int, 0)
	myLink := Constructor()
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{}, nil, t)

	myLink.AddAtHead(2)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{2}, nil, t)

	// ["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
	// 	[[],[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]
	// 	输出
	// 	[null,null,null,null,null,null,null,null,null,5,null,null]
	// 	预期结果
	// 	[null,null,null,null,null,null,null,null,null,2,null,null]

	myLink.DeleteAtIndex(1)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{2}, nil, t)

	myLink.AddAtHead(2)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{2, 2}, nil, t)

	myLink.AddAtHead(7)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{7, 2, 2}, nil, t)

	myLink.AddAtHead(3)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{3, 7, 2, 2}, nil, t)

	myLink.AddAtHead(2)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{2, 3, 7, 2, 2}, nil, t)

	myLink.AddAtHead(5)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{5, 2, 3, 7, 2, 2}, nil, t)

	myLink.AddAtTail(5)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{5, 2, 3, 7, 2, 2, 5}, nil, t)

	resInt := myLink.Get(5)
	common.PrintDiffTwoInt(resInt, 2, nil, t)

	myLink.DeleteAtIndex(6)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{5, 2, 3, 7, 2, 2}, nil, t)

	myLink.DeleteAtIndex(4)
	res = myLink.Print()
	common.PrintDiffTwoIntSlice(res, []int{5, 2, 3, 7, 2}, nil, t)
}
