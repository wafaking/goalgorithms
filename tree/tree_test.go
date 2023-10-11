package tree

import (
	"encoding/json"
	"goalgorithms/common"
	"log"
	"testing"
)

/*
			 5
			/ \
		   4   8
		  /   / \
		 11  13  4
		/  \    / \
	   7    2  5   1
*/

// 原sli：4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
// pre 4 3 1 2 8 5 7 6 9 10
// in 1 2 3 4 5 6 7 8 9 10
// post 2 1 3 6 7 5 10 9 8 4
// layer 4 3 8 1 5 9 2 7 10 6

func TestMain(t *testing.M) {
	list := []int{
		4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
	}
	for _, v := range list {
		Insert(v)
	}

	t.Run()
}

func TestBuildBinaryTree(t *testing.T) {
	var list = [][]int{
		{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1},
		{4, 3, 8, 1, 5, 9, 2, 7, 10, 6},
		{3, 2, 3, -1, 3, -1, 1},
	}
	for _, nums := range list {
		tree := BuildBinaryTree(nums)
		res := levelOrder11(tree)
		t.Logf("org:%+v, res:%+v", nums, res)
	}
}

func TestPreorderTraversal(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			Nums:     []int{1, -1, 2, 3},
			Expected: []int{1, 2, 3},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1},
			Expected: []int{1},
		},
		{
			Nums:     []int{1, 2},
			Expected: []int{1, 2},
		},
		{
			Nums:     []int{1, -1, 2},
			Expected: []int{1, 2},
		},
	}

	var res []int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = preOrderTraversal1(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		res = preOrderTraversal2(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		res = preOrderTraversal3(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestPostorderTraversal(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			Nums:     []int{1, -1, 2, 3},
			Expected: []int{3, 2, 1},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1},
			Expected: []int{1},
		},
		{
			Nums:     []int{1, 2},
			Expected: []int{2, 1},
		},
		{
			Nums:     []int{1, -1, 2},
			Expected: []int{2, 1},
		},
	}

	var res []int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = postOrderTraversal1(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		res = postOrderTraversal2(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		res = postOrderTraversal3(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestInOrderTraversal(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			Nums:     []int{1, -1, 2, 3},
			Expected: []int{1, 3, 2},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1},
			Expected: []int{1},
		},
		{
			Nums:     []int{1, 2},
			Expected: []int{2, 1},
		},
		{
			Nums:     []int{1, -1, 2},
			Expected: []int{1, 2},
		},
	}

	var res []int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = inOrderTraversal1(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		res = inOrderTraversal2(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestLevelOrder21(t *testing.T) {
	var list = []common.Item7{
		{
			Nums:     []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1},
			Expected: [][]int{{5}, {4, 8}, {11, 13, 4}, {7, 2, 5, 1}},
		},
		{
			Nums:     []int{4, 3, 8, 1, 5, 9, 2, 7, 10, 6},
			Expected: [][]int{{4}, {3, 8}, {1, 5, 9, 2}, {7, 10, 6}},
		},
		{
			Nums:     []int{3, 2, 3, -1, 3, -1, 1},
			Expected: [][]int{{3}, {2, 3}, {3, 1}},
		},
	}

	//[0,2,4,1,null,3,-1,5,1,null,6,null,8]
	defaultNullTreeVal = -99999
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res := levelOrder21(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffDoubleSlice(item.Expected, res), res, item)
	}
}

//func TestPostOrder(t *testing.T) {
//	PostOrder1(root)
//	res := PostOrder2(root)
//	log.Println("res: ", res)
//	res = PostOrder3(root)
//	log.Println("res: ", res)
//}

func TestLevelOrder(t *testing.T) {
	res := levelOrder11(root)
	log.Println("res: ", res)
}

// Todo
func TestBuildTree(t *testing.T) {
	//preOrder := PreOrder2(root)
	//log.Println("preOrder: ", preOrder)
	//inOrder := InOrder2(root)
	//log.Println(" inOrder: ", inOrder)
	//
	//tree := buildTree11(preOrder, inOrder)
	////log.Println("res: ", res)
	//log.Println("--------- ")
	//
	//preOrder2 := PreOrder2(tree)
	//log.Println("tree preOrder: ", preOrder2)
	//inOrder2 := InOrder2(tree)
	//log.Println("tree inOrder: ", inOrder2)
}

// TODO
func TestPostBuildTree(t *testing.T) {
	//preOrder := PreOrder2(root)
	//log.Println("preOrder: ", preOrder)
	//inOrder := InOrder2(root)
	//log.Println(" inOrder: ", inOrder)
	//
	//tree := buildTree11(preOrder, inOrder)
	////log.Println("res: ", res)
	//log.Println("--------- ")
	//
	//preOrder2 := PreOrder2(tree)
	//log.Println("tree preOrder: ", preOrder2)
	//inOrder2 := InOrder2(tree)
	//log.Println("tree inOrder: ", inOrder2)
}

func TestBuildPathTree(t *testing.T) {
	var (
		fileList = []item{
			{Path: "a/1.txt", Content: "1"},
			{Path: "a/b/2.txt", Content: "2"},
			{Path: "a/b/3.txt", Content: "3"},
			{Path: "a/b/c/4.ppt", Content: "4"},

			// new add
			{Path: "a/c/ac.txt", Content: "1"},
			{Path: "m/n/mn.txt", Content: "1"},
			{Path: "m/m.txt", Content: "1"},
		}

		gitKeepList = []item{
			{Path: ".gitkeep", Content: "1"},
			{Path: "a/b/c/.gitkeep", Content: "3"},
			{Path: "a/b/.gitkeep", Content: "3"},
			{Path: "a/c/.gitkeep", Content: "1"},
			{Path: "m/n/.gitkeep", Content: "1"},
			{Path: "m/.gitkeep", Content: "1"},
		}
		// result:
		//		a/b --> del
		//		a/b/c --> del
		//		a/c --> add
		//		m	--> add
		//		m/n --> add
	)

	pathNode := BuildPathTree(fileList)

	SignDirDelta(gitKeepList, pathNode)

	res, err := json.Marshal(pathNode)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(res))
}
