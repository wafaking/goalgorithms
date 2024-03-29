package tree

import (
	"encoding/json"
	"goalgorithms/common"
	"goalgorithms/link"
	"log"
	"sort"
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
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = preOrderTraversal2(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = preOrderTraversal3(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
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
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = postOrderTraversal2(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = postOrderTraversal3(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
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
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = inOrderTraversal2(root)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestLevelOrder1(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{5, 4, 8, 11, defaultNullTreeVal, 13, 4, 7, 2, defaultNullTreeVal, defaultNullTreeVal, 5, 1},
			Expected: []int{5, 4, 8, 11, defaultNullTreeVal, 13, 4, 7, 2, defaultNullTreeVal, defaultNullTreeVal, 5, 1},
		},
		{
			Nums:     []int{4, 3, 8, 1, 5, 9, 2, 7, 10, 6},
			Expected: []int{4, 3, 8, 1, 5, 9, 2, 7, 10, 6},
		},
		{
			Nums:     []int{3, 2, 3, defaultNullTreeVal, 3, defaultNullTreeVal, 1},
			Expected: []int{3, 2, 3, defaultNullTreeVal, 3, defaultNullTreeVal, 1},
		},
	}

	var res = make([]int, 0)
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestLevelOrder2(t *testing.T) {
	var list = []common.Item7{
		{
			Nums:     []int{5, 4, 8, 11, defaultNullTreeVal, 13, 4, 7, 2, defaultNullTreeVal, defaultNullTreeVal, 5, 1},
			Expected: [][]int{{5}, {4, 8}, {11, 13, 4}, {7, 2, 5, 1}},
		},
		{
			Nums:     []int{4, 3, 8, 1, 5, 9, 2, 7, 10, 6},
			Expected: [][]int{{4}, {3, 8}, {1, 5, 9, 2}, {7, 10, 6}},
		},
		{
			Nums:     []int{3, 2, 3, defaultNullTreeVal, 3, defaultNullTreeVal, 1},
			Expected: [][]int{{3}, {2, 3}, {3, 1}},
		},
	}

	var res = make([][]int, 0)
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res = levelOrder21(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		res = levelOrder22(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestLevelOrderBottom(t *testing.T) {
	var list = []common.Item7{
		{
			Nums:     []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1},
			Expected: [][]int{{7, 2, 5, 1}, {11, 13, 4}, {4, 8}, {5}},
		},
		{
			Nums:     []int{4, 3, 8, 1, 5, 9, 2, 7, 10, 6},
			Expected: [][]int{{7, 10, 6}, {1, 5, 9, 2}, {3, 8}, {4}},
		},
		{
			Nums:     []int{3, 2, 3, -1, 3, -1, 1},
			Expected: [][]int{{3, 1}, {2, 3}, {3}},
		},
	}

	//[0,2,4,1,null,3,-1,5,1,null,6,null,8]
	//defaultNullTreeVal =
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res := levelOrderBottom1(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
	}
}

func TestLevelOrder(t *testing.T) {
	list := []int{
		4, 3, 8, 1, 2, 5, 7, 9, 6, 10,
	}
	var root = new(common.TreeNode)
	for _, v := range list {
		Insert(root, v)
	}
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

func TestInvertTree(t *testing.T) {
	var list = []common.Item6{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: []int{1, 3, 2, 7, 6, 5, 4},
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
		{
			Nums:     []int{4, 2, 7, 1, 3, 6, 9},
			Expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			Nums:     []int{2, 1, 3},
			Expected: []int{2, 3, 1},
		},
	}

	var (
		res      *common.TreeNode
		listVals []int
	)
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = invertTree1(root)
		listVals = levelOrder11(res)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(listVals, item.Expected), listVals, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestIsSymmetric(t *testing.T) {
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
		root := BuildBinaryTree(item.Nums)
		res = isSymmetric2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = isSymmetric3(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestMaxDepth(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 2, 3},
			Expected: 2,
		},
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: 3,
		},
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: 3,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: 2,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 1,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = maxDepth1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = maxDepth2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = maxDepth3(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestMinDepth(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 2, 3},
			Expected: 2,
		},
		{
			Nums:     []int{2, defaultNullTreeVal, 3, defaultNullTreeVal, 4, defaultNullTreeVal, 5, defaultNullTreeVal, 6},
			Expected: 5,
		},
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: 2,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: 2,
		},
		{
			Nums:     []int{1, 2, 3, defaultNullTreeVal, 3},
			Expected: 2,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 1,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = minDepth1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = minDepth2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		//res = maxDepth3(root)
		//t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestCountNodes(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: 7,
		},
		{
			Nums:     []int{1, 2, 3, 4},
			Expected: 4,
		},
		{
			Nums:     []int{2, defaultNullTreeVal, 3, defaultNullTreeVal, 4, defaultNullTreeVal, 5, defaultNullTreeVal, 6},
			Expected: 5,
		},
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: 5,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: 2,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 1,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = countNodes1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = countNodes2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestIsBalanced(t *testing.T) {
	var list = []common.Item8{
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: true,
		},
		{
			Nums:     []int{1, 2, 2, 3, 3, defaultNullTreeVal, defaultNullTreeVal, 4, 4},
			Expected: false,
		},
		{
			Nums:     []int{1, 2, 2, -1, 3, -1, 3},
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
		root := BuildBinaryTree(item.Nums)
		res = isBalanced1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = isBalanced2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestBinaryTreePaths1(t *testing.T) {
	var list = []common.Item17{
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: []string{"1->2->4", "1->2->5", "1->3"},
		},
		{
			Nums:     []int{1, 2, 3, defaultNullTreeVal, 5},
			Expected: []string{"1->2->5", "1->3"},
		},
		{
			Nums:     []int{1, 2, 2, -1, 3, -1, 3},
			Expected: []string{"1->2->3", "1->2->3"},
		},
		{
			Nums:     []int{},
			Expected: []string{},
		},
		{
			Nums:     []int{1},
			Expected: []string{"1"},
		},
	}

	var res []string
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = binaryTreePaths1(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoStrSlice(res, item.Expected), res, item)
		res = binaryTreePaths2(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoStrSlice(res, item.Expected), res, item)
		res = binaryTreePaths3(root)
		sort.Strings(res)
		sort.Strings(item.Expected)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoStrSlice(res, item.Expected), res, item)
		res = binaryTreePaths4(root)
		sort.Strings(res)
		sort.Strings(item.Expected)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoStrSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestSumOfLeftLeaves(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: 10,
		},
		{
			Nums:     []int{1, 2, 3, 4},
			Expected: 4,
		},
		{
			Nums:     []int{2, defaultNullTreeVal, 3, defaultNullTreeVal, 4, defaultNullTreeVal, 5, defaultNullTreeVal, 6},
			Expected: 0,
		},
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: 24,
		},
		{
			Nums:     []int{1, 2, 3, defaultNullTreeVal, 4, 5, 6, 7, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 8},
			Expected: 20,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: 0,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 0,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = sumOfLeftLeaves1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = sumOfLeftLeaves2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = sumOfLeftLeaves3(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestHasPathSum(t *testing.T) {
	var list = []common.Item18{
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Target:   30,
			Expected: true,
		},
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Target:   12,
			Expected: true,
		},
		{
			Nums:     []int{5, 4, 8, 11, defaultNullTreeVal, 13, 4, 7, 2, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 1},
			Target:   22,
			Expected: true,
		},
		{
			Nums:     []int{1, -2, -3, 1, 3, -2, defaultNullTreeVal, -1},
			Target:   -4,
			Expected: true,
		},
		{
			Nums:     []int{1, 2, 3},
			Target:   6,
			Expected: false,
		},
		{
			Nums:     []int{},
			Target:   0,
			Expected: false,
		},
		{
			Nums:     []int{1},
			Target:   1,
			Expected: true,
		},
		{
			Nums:     []int{1, 2, 3, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 4},
			Target:   8,
			Expected: true,
		},
	}

	var res bool
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = hasPathSum1(root, item.Target)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = hasPathSum2(root, item.Target)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = hasPathSum3(root, item.Target)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

// TODO
func TestBuildTreeByPrIn(t *testing.T) {
	var list = []common.Item19{
		{
			Nums1:    []int{3, 9, 20, 15, 7},
			Nums2:    []int{9, 3, 15, 20, 7},
			Expected: []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
		},
		{
			Nums1:    []int{1, 2},
			Nums2:    []int{1, 2},
			Expected: []int{1, 2},
		},
		{
			Nums1:    []int{1, 2, 3},
			Nums2:    []int{3, 2, 1},
			Expected: []int{1, 2, defaultNullTreeVal, 3},
		},
	}

	var (
		root *common.TreeNode
		res  []int
	)
	for _, item := range list {
		root = buildTree11(item.Nums1, item.Nums2)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		root = buildTree12(item.Nums1, item.Nums2)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestBuildTreeByPostIn(t *testing.T) {
	var list = []common.Item19{
		{
			Nums1:    []int{9, 3, 15, 20, 7},
			Nums2:    []int{9, 15, 7, 20, 3},
			Expected: []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
		},
		{
			Nums1:    []int{1, 2},
			Nums2:    []int{1, 2},
			Expected: []int{2, 1, defaultNullTreeVal},
		},
		{
			Nums1:    []int{1, 2, 3},
			Nums2:    []int{3, 2, 1},
			Expected: []int{1, defaultNullTreeVal, 2, defaultNullTreeVal, 3},
		},
	}

	var (
		root *common.TreeNode
		res  []int
	)
	for _, item := range list {
		root = buildTree21(item.Nums1, item.Nums2)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		root = buildTree22(item.Nums1, item.Nums2)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestFindBottomLeftValue(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Expected: 4,
		},
		{
			Nums:     []int{2, 1, 3},
			Expected: 1,
		},
		{
			Nums:     []int{1, 2, 3, 4, defaultNullTreeVal, 5, 6, defaultNullTreeVal, defaultNullTreeVal, 7},
			Expected: 7,
		},
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: 15,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: 2,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 1,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = findBottomLeftValue1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = findBottomLeftValue2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = findBottomLeftValue3(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestLongestUnivaluePath(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{5, 4, 5, 1, 1, 5},
			Expected: 2,
		},
		{
			Nums:     []int{2, 1, 3},
			Expected: 0,
		},
		{
			Nums:     []int{1, 4, 5, 4, 4, 5},
			Expected: 2,
		},
		{
			Nums:     []int{4, 4, 5, 4, 4, 5},
			Expected: 2,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 1},
			Expected: 1,
		},
		{
			Nums:     []int{},
			Expected: 0,
		},
		{
			Nums:     []int{1},
			Expected: 0,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = longestUnivaluePath(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{3, 2, 1},
			Expected: []int{3, defaultNullTreeVal, 2, defaultNullTreeVal, 1},
		},
		{
			Nums:     []int{3, 2, 1, 6, 0, 5},
			Expected: []int{6, 3, 5, defaultNullTreeVal, 2, 0, defaultNullTreeVal, defaultNullTreeVal, 1},
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

	var (
		res  []int
		root *common.TreeNode
	)
	for _, item := range list {
		root = constructMaximumBinaryTree1(item.Nums)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		root = constructMaximumBinaryTree2(item.Nums)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestMergeTrees(t *testing.T) {
	var list = []common.Item19{
		{
			Nums1:    []int{1, 3, 2, 5},
			Nums2:    []int{2, 1, 3, defaultNullTreeVal, 4, defaultNullTreeVal, 7},
			Expected: []int{3, 4, 5, 5, 4, defaultNullTreeVal, 7},
		},
		{
			Nums1:    []int{1},
			Nums2:    []int{1, 2},
			Expected: []int{2, 2},
		},
		{
			Nums1:    []int{1, 2, defaultNullTreeVal, 3},
			Nums2:    []int{1, defaultNullTreeVal, 2, defaultNullTreeVal, 3},
			Expected: []int{2, 2, 2, 3, defaultNullTreeVal, defaultNullTreeVal, 3},
		},
		{
			Nums1:    []int{1},
			Nums2:    []int{},
			Expected: []int{1},
		},
		{
			Nums1:    []int{},
			Nums2:    []int{1},
			Expected: []int{1},
		},
		{
			Nums1:    []int{1, defaultNullTreeVal, 2, defaultNullTreeVal, 3},
			Nums2:    []int{1, 2},
			Expected: []int{2, 2, 2, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 3},
		},
	}

	var (
		root, root1, root2 *common.TreeNode
		res                []int
	)
	for _, item := range list {
		root1 = BuildBinaryTree(item.Nums1)
		root2 = BuildBinaryTree(item.Nums2)
		root = mergeTrees1(root1, root2)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		root1 = BuildBinaryTree(item.Nums1)
		root2 = BuildBinaryTree(item.Nums2)
		root = mergeTrees2(root1, root2)
		res = levelOrder11(root)
		t.Logf("%t, res: %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestSearchBST(t *testing.T) {
	var list = []common.Item21{
		{
			Nums:     []int{4, 2, 7, 1, 3},
			Target:   2,
			Expected: []int{2, 1, 3},
		},
		{
			Nums:     []int{4, 2, 7, 1, 3},
			Target:   5,
			Expected: []int{},
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2, defaultNullTreeVal, 3, defaultNullTreeVal, 4, defaultNullTreeVal, 5},
			Target:   4,
			Expected: []int{4, defaultNullTreeVal, 5},
		},
	}

	var (
		res        = make([]int, 0)
		root, node *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		node = searchBST1(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		//tree := BuildBinaryTree(item.Nums)
		node = searchBST2(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		node = searchBST3(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestIsValidBST(t *testing.T) {
	var list = []common.Item8{
		{
			Nums:     []int{2, 1, 3},
			Expected: true,
		},
		//示例2：输入：root=[5,1,4,null,null,3,6]输出：false,解释：
		{
			Nums:     []int{5, 1, 4, defaultNullTreeVal, defaultNullTreeVal, 3, 6},
			Expected: false,
		},
		{
			Nums:     []int{10, 8, 16, 3, 13, 7, 12, 2, 9},
			Expected: false,
		},
		{
			Nums:     []int{2, 2, 2},
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
		root := BuildBinaryTree(item.Nums)
		res = isValidBST1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = isValidBST2(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = isValidBST3(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = isValidBST4(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestGetMinimumDifference(t *testing.T) {
	var list = []common.Item2{
		{
			Nums:     []int{4, 2, 6, 1, 3},
			Expected: 1,
		},
		{
			Nums:     []int{1, 0, 48, defaultNullTreeVal, defaultNullTreeVal, 12, 49},
			Expected: 1,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: 1,
		},
		{
			Nums:     []int{236, 104, 701, defaultNullTreeVal, 227, defaultNullTreeVal, 911},
			Expected: 9,
		},
	}

	var res int
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		res = getMinimumDifference1(root)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestFindMode(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{1, defaultNullTreeVal, 2, 2},
			Expected: []int{2},
		},
		{
			Nums:     []int{0},
			Expected: []int{0},
		},
		{
			Nums:     []int{10, 8, 15, 8, 8, defaultNullTreeVal, 15},
			Expected: []int{8},
		},
		{
			Nums:     []int{10, 8, 15, 8, 8, 15, 15},
			Expected: []int{8, 15},
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 2},
			Expected: []int{1, 2},
		},
	}

	var res = make([]int, 0)
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res = findMode1(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		res = findMode2(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	var list = []common.Item22{
		{
			Item2: common.Item2{
				Nums:     []int{3, 5, 1, 6, 2, 0, 8, defaultNullTreeVal, defaultNullTreeVal, 7, 4},
				Expected: 3,
			},
			Nums1: []int{5, 6, 2, defaultNullTreeVal, defaultNullTreeVal, 7, 4},
			Nums2: []int{1, 0, 8},
		},
		{
			Item2: common.Item2{
				Nums:     []int{3, 5, 1, 6, 2, 0, 8, defaultNullTreeVal, defaultNullTreeVal, 7, 4},
				Expected: 5,
			},
			Nums1: []int{5, 6, 2, defaultNullTreeVal, defaultNullTreeVal, 7, 4},
			Nums2: []int{4},
		},
		{
			Item2: common.Item2{
				Nums:     []int{1, 2},
				Expected: 1,
			},
			Nums1: []int{1, 2},
			Nums2: []int{2},
		},
	}

	var (
		res, p, q = new(common.TreeNode), new(common.TreeNode), new(common.TreeNode)
	)
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		p = BuildBinaryTree(item.Nums1)
		q = BuildBinaryTree(item.Nums2)
		res = lowestCommonAncestor1(root, p, q)
		t.Logf("%t, res: %+v, item:%+v", res.Val == item.Expected, res.Val, item)
		res = lowestCommonAncestor2(root, p, q)
		t.Logf("%t, res: %+v, item:%+v", res.Val == item.Expected, res.Val, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestLowestCommonAncestorBST(t *testing.T) {

	var list = []common.Item22{
		{
			Item2: common.Item2{
				Nums:     []int{6, 2, 8, 0, 4, 7, 9, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
				Expected: 6,
			},
			Nums1: []int{2, 0, 4, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
			Nums2: []int{8, 7, 9},
		},
		{
			Item2: common.Item2{
				Nums:     []int{6, 2, 8, 0, 4, 7, 9, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
				Expected: 2,
			},
			Nums1: []int{2, 0, 4, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
			Nums2: []int{4, 3, 5},
		},
		{
			Item2: common.Item2{
				Nums:     []int{6, 2, 8, 0, 4, 7, 9, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
				Expected: 2,
			},
			Nums1: []int{0},
			Nums2: []int{4, 3, 5},
		},
		{
			Item2: common.Item2{
				Nums:     []int{6, 2, 8, 0, 4, 7, 9, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
				Expected: 6,
			},
			Nums1: []int{3},
			Nums2: []int{9},
		},
		{
			Item2: common.Item2{
				Nums:     []int{6, 2, 8, 0, 4, 7, 9, defaultNullTreeVal, defaultNullTreeVal, 3, 5},
				Expected: 4,
			},
			Nums1: []int{3},
			Nums2: []int{5},
		},
	}

	var (
		res, p, q = new(common.TreeNode), new(common.TreeNode), new(common.TreeNode)
	)
	for _, item := range list {
		root := BuildBinaryTree(item.Nums)
		p = BuildBinaryTree(item.Nums1)
		q = BuildBinaryTree(item.Nums2)
		res = lowestCommonAncestorBST1(root, p, q)
		t.Logf("%t, res: %+v, item:%+v", res.Val == item.Expected, res.Val, item)
		res = lowestCommonAncestorBST2(root, p, q)
		t.Logf("%t, res: %+v, item:%+v", res.Val == item.Expected, res.Val, item)
		res = lowestCommonAncestorBST3(root, p, q)
		t.Logf("%t, res: %+v, item:%+v", res.Val == item.Expected, res.Val, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestInsertIntoBST(t *testing.T) {
	var list = []common.Item21{
		{
			Nums:     []int{4, 2, 7, 1, 3},
			Target:   5,
			Expected: []int{4, 2, 7, 1, 3, 5},
		},
		{
			Nums:     []int{40, 20, 60, 10, 30, 50, 70},
			Target:   25,
			Expected: []int{40, 20, 60, 10, 30, 50, 70, defaultNullTreeVal, defaultNullTreeVal, 25},
		},
		{
			Nums:     []int{4, 2, 7, 1, 3, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal},
			Target:   5,
			Expected: []int{4, 2, 7, 1, 3, 5},
		},
		{
			Nums:     []int{3, 1, 5},
			Target:   2,
			Expected: []int{3, 1, 5, defaultNullTreeVal, 2},
		},
	}

	var (
		res        = make([]int, 0)
		root, node *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		node = insertIntoBST1(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		root = BuildBinaryTree(item.Nums)
		node = insertIntoBST2(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestDeleteNode(t *testing.T) {
	var list = []common.Item21{
		{
			Nums:     []int{5, 3, 6, 2, 4, defaultNullTreeVal, 7},
			Target:   3,
			Expected: []int{5, 4, 6, 2, defaultNullTreeVal, defaultNullTreeVal, 7},
		},
		{
			Nums:     []int{5, 3, 6, 2, 4, defaultNullTreeVal, 7},
			Target:   0,
			Expected: []int{5, 3, 6, 2, 4, defaultNullTreeVal, 7},
		},
		{
			Nums:     []int{},
			Target:   0,
			Expected: []int{},
		},
		{
			Nums:     []int{0},
			Target:   0,
			Expected: []int{},
		},
		{
			Nums:     []int{10, 6, 13, 3, 8, defaultNullTreeVal, defaultNullTreeVal, 1, 4, 7, 9},
			Target:   6,
			Expected: []int{10, 3, 13, 1, 8, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 7, 9, 4},
		},
	}

	var (
		res        = make([]int, 0)
		root, node *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		node = deleteNode1(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		root = BuildBinaryTree(item.Nums)
		node = deleteNode2(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		root = BuildBinaryTree(item.Nums)
		node = deleteNode3(root, item.Target)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestTrimBST(t *testing.T) {
	var list = []common.Item23{
		{
			Nums:     []int{1, 0, 2},
			Num1:     1,
			Num2:     2,
			Expected: []int{1, defaultNullTreeVal, 2},
		},
		{
			Nums:     []int{3, 0, 4, defaultNullTreeVal, 2, defaultNullTreeVal, defaultNullTreeVal, 1},
			Num1:     1,
			Num2:     3,
			Expected: []int{3, 2, defaultNullTreeVal, 1},
		},
		{
			Nums:     []int{10, 6, 13, 3, 8, defaultNullTreeVal, defaultNullTreeVal, 1, 4, 7, 9},
			Num1:     4,
			Num2:     7,
			Expected: []int{6, 4, 7},
		},
		{
			Nums:     []int{10, 6, 13, 3, 8, defaultNullTreeVal, defaultNullTreeVal, 1, 4, 7, 9},
			Num1:     3,
			Num2:     7,
			Expected: []int{6, 3, 7, defaultNullTreeVal, 4},
		},
	}

	var (
		res        = make([]int, 0)
		root, node *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		node = trimBST1(root, item.Num1, item.Num2)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		root = BuildBinaryTree(item.Nums)
		node = trimBST2(root, item.Num1, item.Num2)
		res = levelOrder11(node)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("----------------SPLIT------------------SPLIT----------")
	}
}

func TestSortedArrayToBST(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{-10, -3, 0, 5, 9},
			Expected: []int{0, -3, 9, -10, defaultNullTreeVal, 5},
		},
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: []int{3, 2, 4, 1, defaultNullTreeVal, defaultNullTreeVal, 5},
		},
		{
			Nums:     []int{1, 3},
			Expected: []int{3, 1},
		},
	}

	var (
		res  = make([]int, 0)
		tree *common.TreeNode
	)
	for _, item := range list {
		tree = sortedArrayToBST1(item.Nums)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		tree = sortedArrayToBST2(item.Nums)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestSortedListToBST(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{-10, -3, 0, 5, 9},
			Expected: []int{0, -3, 9, -10, defaultNullTreeVal, 5},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
		{
			Nums:     []int{1, 2, 3, 4, 5},
			Expected: []int{3, 2, 4, 1, defaultNullTreeVal, defaultNullTreeVal, 5},
		},
		{
			Nums:     []int{1, 3},
			Expected: []int{3, 1},
		},
	}

	var (
		res      = make([]int, 0)
		tree     *common.TreeNode
		listNode *common.ListNode
	)
	for _, item := range list {
		listNode = link.BuildListNode(item.Nums)
		tree = sortedListToBST1(listNode)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		listNode = link.BuildListNode(item.Nums)
		tree = sortedListToBST2(listNode)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		tree = sortedListToBST3(listNode)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestConvertBST(t *testing.T) {
	var list = []common.Item20{
		{
			Nums: []int{4, 1, 6, 0, 2, 5, 7, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal,
				3, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 8},
			Expected: []int{30, 36, 21, 36, 35, 26, 15, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal,
				33, defaultNullTreeVal, defaultNullTreeVal, defaultNullTreeVal, 8},
		},
		{
			Nums:     []int{0, defaultNullTreeVal, 1},
			Expected: []int{1, defaultNullTreeVal, 1},
		},
		{
			Nums:     []int{1, 0, 2},
			Expected: []int{3, 3, 2},
		},
		{
			Nums:     []int{3, 2, 4, 1},
			Expected: []int{7, 9, 4, 10},
		},
	}

	var (
		res  = make([]int, 0)
		tree *common.TreeNode
		root *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		tree = convertBST1(root)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		root = BuildBinaryTree(item.Nums)
		tree = convertBST2(root)
		res = levelOrder11(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestPreorderN(t *testing.T) {
	var list = []common.Item24{
		{
			Expected: []int{1, 3, 5, 6, 2, 4},
			NNode: &common.NTreeNode{
				Val: 1,
				Children: []*common.NTreeNode{
					{
						Val: 3,
						Children: []*common.NTreeNode{
							{
								Val:      5,
								Children: nil,
							},
							{
								Val:      6,
								Children: nil,
							},
						},
					},
					{
						Val:      2,
						Children: nil,
					},
					{
						Val:      4,
						Children: nil,
					},
				},
			},
		},
		{
			Expected: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
			NNode: &common.NTreeNode{
				Val: 1,
				Children: []*common.NTreeNode{
					{
						Val:      2,
						Children: nil,
					},
					{
						Val: 3,
						Children: []*common.NTreeNode{
							{
								Val:      6,
								Children: nil,
							},
							{
								Val: 7,
								Children: []*common.NTreeNode{
									{
										Val: 11,
										Children: []*common.NTreeNode{
											{
												Val:      14,
												Children: nil,
											},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*common.NTreeNode{
							{
								Val: 8,
								Children: []*common.NTreeNode{
									{
										Val:      12,
										Children: nil,
									},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*common.NTreeNode{
							{
								Val: 9,
								Children: []*common.NTreeNode{
									{
										Val:      13,
										Children: nil,
									},
								},
							},
							{
								Val:      10,
								Children: nil,
							},
						},
					},
				},
			},
		},
		{
			Expected: []int{},
			NNode:    &common.NTreeNode{},
		},
	}

	var res []int
	for _, item := range list {
		res = preorder1(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = preorder2(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = preorder3(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestLevelOrderN(t *testing.T) {
	var list = []common.Item25{
		{
			Expected: [][]int{{1}, {3, 2, 4}, {5, 6}},
			NNode: &common.NTreeNode{
				Val: 1,
				Children: []*common.NTreeNode{
					{
						Val: 3,
						Children: []*common.NTreeNode{
							{
								Val:      5,
								Children: nil,
							},
							{
								Val:      6,
								Children: nil,
							},
						},
					},
					{
						Val:      2,
						Children: nil,
					},
					{
						Val:      4,
						Children: nil,
					},
				},
			},
		},
		{
			Expected: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
			NNode: &common.NTreeNode{
				Val: 1,
				Children: []*common.NTreeNode{
					{
						Val:      2,
						Children: nil,
					},
					{
						Val: 3,
						Children: []*common.NTreeNode{
							{
								Val:      6,
								Children: nil,
							},
							{
								Val: 7,
								Children: []*common.NTreeNode{
									{
										Val: 11,
										Children: []*common.NTreeNode{
											{
												Val:      14,
												Children: nil,
											},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*common.NTreeNode{
							{
								Val: 8,
								Children: []*common.NTreeNode{
									{
										Val:      12,
										Children: nil,
									},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*common.NTreeNode{
							{
								Val: 9,
								Children: []*common.NTreeNode{
									{
										Val:      13,
										Children: nil,
									},
								},
							},
							{
								Val:      10,
								Children: nil,
							},
						},
					},
				},
			},
		},
		{
			Expected: [][]int{},
			NNode:    nil,
		},
	}

	var res [][]int
	for _, item := range list {
		res = levelOrder1(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item)
		res = levelOrder2(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item)
		res = levelOrder3(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoDoubleIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestPostorderN(t *testing.T) {
	var list = []common.Item24{
		{
			Expected: []int{5, 6, 3, 2, 4, 1},
			NNode: &common.NTreeNode{
				Val: 1,
				Children: []*common.NTreeNode{
					{
						Val: 3,
						Children: []*common.NTreeNode{
							{
								Val:      5,
								Children: nil,
							},
							{
								Val:      6,
								Children: nil,
							},
						},
					},
					{
						Val:      2,
						Children: nil,
					},
					{
						Val:      4,
						Children: nil,
					},
				},
			},
		},
		{
			Expected: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
			NNode: &common.NTreeNode{
				Val: 1,
				Children: []*common.NTreeNode{
					{
						Val:      2,
						Children: nil,
					},
					{
						Val: 3,
						Children: []*common.NTreeNode{
							{
								Val:      6,
								Children: nil,
							},
							{
								Val: 7,
								Children: []*common.NTreeNode{
									{
										Val: 11,
										Children: []*common.NTreeNode{
											{
												Val:      14,
												Children: nil,
											},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*common.NTreeNode{
							{
								Val: 8,
								Children: []*common.NTreeNode{
									{
										Val:      12,
										Children: nil,
									},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*common.NTreeNode{
							{
								Val: 9,
								Children: []*common.NTreeNode{
									{
										Val:      13,
										Children: nil,
									},
								},
							},
							{
								Val:      10,
								Children: nil,
							},
						},
					},
				},
			},
		},
		{
			Expected: []int{},
			NNode:    nil,
		},
	}

	var res []int
	for _, item := range list {
		res = postorder1(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = postorder2(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = postorder3(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		res = postorder4(item.NNode)
		t.Logf("res: %v, %+v, item:%+v", common.DiffTwoIntSlice(res, item.Expected), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestRightSideView(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{1, 2, 3, defaultNullTreeVal, 5, defaultNullTreeVal, 4},
			Expected: []int{1, 3, 4},
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 3},
			Expected: []int{1, 3},
		},
		{
			Nums:     []int{},
			Expected: []int{},
		},
	}

	var res = make([]int, 0)
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res = rightSideView1(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		res = rightSideView2(tree)
		//t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		res = rightSideView3(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestZigzagLevelOrder(t *testing.T) {
	var list = []common.Item7{
		{
			Nums:     []int{3, 9, 20, defaultNullTreeVal, defaultNullTreeVal, 15, 7},
			Expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			Nums:     []int{1},
			Expected: [][]int{{1}},
		},
		{
			Nums:     []int{},
			Expected: [][]int{},
		},
	}

	var res = make([][]int, 0)
	for _, item := range list {
		tree := BuildBinaryTree(item.Nums)
		res = zigzagLevelOrder1(tree)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		//res = zigzagLevelOrder2(tree)
		//t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		t.Log("--------------------------")
	}
}

func TestVerifyTreeOrder(t *testing.T) {
	var list = []common.Item8{
		{
			Nums:     []int{4, 9, 6, 9, 8},
			Expected: false,
		},
		{
			Nums:     []int{1, 2, 2, 3, 4, 4, 3},
			Expected: true,
		},
		{
			Nums:     []int{4, 6, 5, 9, 8},
			Expected: true,
		},
		{
			Nums:     []int{},
			Expected: true,
		},
	}

	var res bool
	for _, item := range list {
		res = verifyTreeOrder1(item.Nums)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		res = verifyTreeOrder2(item.Nums)
		t.Logf("%t, res: %+v, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestPathSum2(t *testing.T) {
	var list = []common.Item26{
		{
			Nums:     []int{5, 4, 8, 11, defaultNullTreeVal, 13, 4, 7, 2, defaultNullTreeVal, defaultNullTreeVal, 5, 1},
			Target:   22,
			Expected: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			Nums:     []int{1, 2, 3},
			Target:   55,
			Expected: [][]int{},
		},
		{
			Nums:     []int{1, 2},
			Target:   0,
			Expected: [][]int{},
		},
	}

	var (
		res  = make([][]int, 0)
		root *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		res = pathSum21(root, item.Target)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		res = pathSum22(root, item.Target)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		res = pathSum23(root, item.Target)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		res = pathSum24(root, item.Target)
		t.Logf("%t, res:%+v, item:%+v", common.DiffTwoDoubleIntSlice(item.Expected, res), res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestPathSum3(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{10, 5, -3, 3, 2, defaultNullTreeVal, 11, 3, -2, defaultNullTreeVal, 1},
			Target:   8,
			Expected: 3,
		},
		{
			Nums:     []int{5, 4, 8, 11, defaultNullTreeVal, 13, 4, 7, 2, defaultNullTreeVal, defaultNullTreeVal, 5, 1},
			Target:   22,
			Expected: 3,
		},
		{
			Nums:     []int{1},
			Target:   1,
			Expected: 1,
		},
		{
			Nums:     []int{},
			Target:   0,
			Expected: 0,
		},
	}

	var (
		res  int
		root *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		res = pathSum31(root, item.Target)
		t.Logf("%t, res:%d, item:%+v", res == item.Expected, res, item)
		res = pathSum32(root, item.Target)
		t.Logf("%t, res: %d, item:%+v", res == item.Expected, res, item)
		res = pathSum33(root, item.Target)
		t.Logf("%t, res:%d, item:%+v", res == item.Expected, res, item)
		t.Log("--------------------SPLIT--------------------------")
	}
}

func TestFindTarget(t *testing.T) {
	var list = []common.Item18{
		{
			Nums:     []int{2, 2, 2, 2, 3, 4, 5},
			Target:   8,
			Expected: true,
		},
		{
			Nums:     []int{4, 3, 2, 3, 5, 2, 1},
			Target:   5,
			Expected: true,
		},
		{
			Nums:     []int{4, 3, 2, 3, 5, 2, 1, 1, 4},
			Target:   11,
			Expected: false,
		},
		{
			Nums:     []int{2, 2},
			Target:   4,
			Expected: true,
		},
		{
			Nums:     []int{2, 2},
			Target:   5,
			Expected: false,
		},
		{
			Nums:     []int{5, 3, 6, 2, 4, defaultNullTreeVal, 7},
			Target:   9,
			Expected: true,
		},
		{
			Nums:     []int{5, 3, 6, 2, 4, defaultNullTreeVal, 7},
			Target:   28,
			Expected: false,
		},
		{
			Nums:     []int{2, 1, 3},
			Target:   3,
			Expected: true,
		},
	}

	var (
		res  bool
		root *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		res = findTarget1(root, item.Target)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = findTarget2(root, item.Target)
		common.PrintDiffTwoBool(res, item.Expected, item, t)
		res = findTarget3(root, item.Target)
		common.PrintDiffTwoBool(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}

func TestBstFromPreorder(t *testing.T) {
	var list = []common.Item20{
		{
			Nums:     []int{8, 5, 1, 7, 10, 12},
			Expected: []int{8, 5, 10, 1, 7, defaultNullTreeVal, 12},
		},
		{
			Nums:     []int{4, 2},
			Expected: []int{4, 2, defaultNullTreeVal},
		},
		{
			Nums:     []int{1, 3},
			Expected: []int{1, defaultNullTreeVal, 3},
		},
	}

	var (
		res  = make([]int, 0)
		root *common.TreeNode
	)
	for _, item := range list {
		root = bstFromPreorder1(item.Nums)
		res = levelOrder11(root)
		common.PrintDiffTwoIntSlice(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}

func TestFindTargetNode(t *testing.T) {
	var list = []common.Item3{
		{
			Nums:     []int{8, 5, 1, 7, 10, 12},
			Target:   3,
			Expected: 8,
		},
		{
			Nums:     []int{4, 2},
			Target:   2,
			Expected: 2,
		},
		{
			Nums:     []int{1, defaultNullTreeVal, 3},
			Target:   1,
			Expected: 3,
		},
		{
			Nums:     []int{1, 3},
			Target:   4,
			Expected: -1,
		},
		{
			Nums:     []int{7, 3, 9, 1, 5},
			Target:   2,
			Expected: 7,
		},
		{
			Nums:     []int{10, 5, 15, 2, 7, defaultNullTreeVal, 20, 1, defaultNullTreeVal, 6, 8},
			Target:   4,
			Expected: 8,
		},
	}

	var (
		res  int
		root *common.TreeNode
	)
	for _, item := range list {
		root = BuildBinaryTree(item.Nums)
		res = findTargetNode1(root, item.Target)
		common.PrintDiffTwoInt(res, item.Expected, item, t)

		common.PrintSplit(t)
	}
}
