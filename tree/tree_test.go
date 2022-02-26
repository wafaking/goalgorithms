package tree

import (
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

func TestPermute(t *testing.T) {
	log.Println(permute([]int{1, 2, 3}))
}

func TestPermuteUnique(t *testing.T) {
	var list = []int{1, 2, 1, 3}
	log.Println(PermuteUnique(list))
}

func TestPreOrder(t *testing.T) {
	res1 := PreOrder2(root)
	log.Println("res: ", res1)
	log.Println("--------")
	res := PreOrder3(root)
	log.Println("res: ", res)

}
func TestInOrder(t *testing.T) {
	InOrder1(root)
	res := InOrder2(root)
	log.Println("res: ", res)
	res = InOrder3(root)
	log.Println("res: ", res)
}
func TestPostOrder(t *testing.T) {
	PostOrder1(root)
	res := PostOrder2(root)
	log.Println("res: ", res)
	res = PostOrder3(root)
	log.Println("res: ", res)
}

func TestLevelOrder(t *testing.T) {
	res := LevelOrder1(root)
	log.Println("res: ", res)
}

func TestBuildTree(t *testing.T) {
	preOrder := PreOrder2(root)
	log.Println("preOrder: ", preOrder)
	inOrder := InOrder2(root)
	log.Println(" inOrder: ", inOrder)

	tree := buildTree11(preOrder, inOrder)
	//log.Println("res: ", res)
	log.Println("--------- ")

	preOrder2 := PreOrder2(tree)
	log.Println("tree preOrder: ", preOrder2)
	inOrder2 := InOrder2(tree)
	log.Println("tree inOrder: ", inOrder2)
}

func TestPostBuildTree(t *testing.T) {
	preOrder := PreOrder2(root)
	log.Println("preOrder: ", preOrder)
	inOrder := InOrder2(root)
	log.Println(" inOrder: ", inOrder)

	tree := buildTree11(preOrder, inOrder)
	//log.Println("res: ", res)
	log.Println("--------- ")

	preOrder2 := PreOrder2(tree)
	log.Println("tree preOrder: ", preOrder2)
	inOrder2 := InOrder2(tree)
	log.Println("tree inOrder: ", inOrder2)
}
