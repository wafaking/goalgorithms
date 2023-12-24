package queue_stack

import (
	"fmt"
	"goalgorithms/common"
	"testing"
)

func TestFit(t *testing.T) {

}

func TestPMMD(t *testing.T) {

	// ( 4 - 3 ) * 5 - 2 + ( 6 - 2 ) / 2 * 3 - 3 + 10 = 16
	// [4 3 - 5 * 2 - 6 2 - 2 / 3 * + 1 - 5 +]
	// var exp = []string{"(", "4", "-", "3", ")", "*", "5", "-", "2", "+", "(", "6", "-", "2", ")", "/", "2", "*", "3", "-", "1", "+", "5"}
	// 9+(3-1)*3+10/2 =20
	// [9 3 1 - 3 * + 10 2 / +]
	// var exp = []string{"9", "+", "(", "3", "-", "1", ")", "*", "3", "+", "10", "/", "2"}

	var exp = "( 4 - 3 ) * 15 - 2 + ( 6 - 2 ) / 2 * 3 - 10 + 10"

	rear := midConvertRear(exp)
	fmt.Println(rear)
	result := rearCalculation(rear)
	fmt.Println(result)
}

func TestFibonacci(t *testing.T) {
	var list = []common.Item4{
		{Num: 2, Expected: 1},
		{Num: 3, Expected: 2},
		{Num: 4, Expected: 3},
		{Num: 4, Expected: 3},
		{Num: 6, Expected: 8},
		{Num: 8, Expected: 21},
	}
	var res int
	for _, item := range list {
		res = fibonacci1(item.Num)
		t.Logf("res: %t, res: %d, item:%+v", res == item.Expected, res, item)
		res = fibonacci2(item.Num)
		t.Logf("res: %t, res: %d, item:%+v", res == item.Expected, res, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestBigFibonacci(t *testing.T) {
	var list = []common.Item37{
		{Num: 100, Expected: "354224848179261915075"},
		{Num: 101, Expected: "573147844013817084101"},
		{Num: 102, Expected: "927372692193078999176"},
		{Num: 200, Expected: "280571172992510140037611932413038677189525"},
	}
	var res string
	for _, item := range list {
		res = bigFibonacci1(item.Num)
		t.Logf("%t, res: %s, item:%+v", res == item.Expected, res, item)
		res = bigFibonacci2(item.Num)
		t.Logf("%t, res: %s, item:%+v", res == item.Expected, res, item)
		t.Log("------------------------SPLIT------------------------")
	}
}

func TestNumWays(t *testing.T) {
	for i := 0; i < 100; i++ {
		res1 := numWays1(i)
		res2 := numWays2(i)
		t.Logf("n:%d, res:  %d, %d, %t", i, res1, res2, res1 == res2)
	}
}

func TestMaxInQueue1(t *testing.T) {
	var (
		checkout      Checkout
		res, expected = make([]int, 0), make([]int, 0)
	)

	checkout = MaxInQueueConstructor1()
	expected = []int{7, 4, 7}
	checkout.Add(4)
	checkout.Add(7)
	res = append(res, checkout.Get_max())
	res = append(res, checkout.Remove())
	res = append(res, checkout.Get_max())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

	checkout = MaxInQueueConstructor1()
	res = make([]int, 0)
	expected = []int{-1, -1}
	res = append(res, checkout.Remove())
	res = append(res, checkout.Get_max())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

	checkout = MaxInQueueConstructor1()
	res = make([]int, 0)
	expected = []int{-1, -1, -1, 94, 16, 89}
	res = append(res, checkout.Get_max())
	res = append(res, checkout.Remove())
	res = append(res, checkout.Remove())
	checkout.Add(94)
	checkout.Add(16)
	checkout.Add(89)
	res = append(res, checkout.Remove())
	checkout.Add(22)
	res = append(res, checkout.Remove())
	res = append(res, checkout.Get_max())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

	checkout = MaxInQueueConstructor1()
	res = make([]int, 0)
	expected = []int{-1, -1, -1, -1, -1, 15, 15}
	res = append(res, checkout.Remove())
	res = append(res, checkout.Remove())
	res = append(res, checkout.Remove())
	res = append(res, checkout.Remove())
	res = append(res, checkout.Remove())
	checkout.Add(15)
	res = append(res, checkout.Get_max())
	checkout.Add(9)
	res = append(res, checkout.Get_max())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)
}

func TestMaxInQueue2(t *testing.T) {
	var (
		checkout      MaxQueue
		res, expected = make([]int, 0), make([]int, 0)
	)

	checkout = MaxInQueueConstructor2()
	expected = []int{7, 4, 7}
	checkout.Push_back(4)
	checkout.Push_back(7)
	res = append(res, checkout.Max_value())
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Max_value())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

	checkout = MaxInQueueConstructor2()
	res = make([]int, 0)
	expected = []int{-1, -1}
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Max_value())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

	checkout = MaxInQueueConstructor2()
	res = make([]int, 0)
	expected = []int{-1, -1, -1, 94, 16, 89}
	res = append(res, checkout.Max_value())
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Pop_front())
	checkout.Push_back(94)
	checkout.Push_back(16)
	checkout.Push_back(89)
	res = append(res, checkout.Pop_front())
	checkout.Push_back(22)
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Max_value())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

	checkout = MaxInQueueConstructor2()
	res = make([]int, 0)
	expected = []int{-1, -1, -1, -1, -1, 15, 15}
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Pop_front())
	res = append(res, checkout.Pop_front())
	checkout.Push_back(15)
	res = append(res, checkout.Max_value())
	checkout.Push_back(9)
	res = append(res, checkout.Max_value())
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)

}

func TestTwoStack2Queue(t *testing.T) {
	var (
		myQueue       MyQueue
		res, expected = make([]int, 0), make([]int, 0)
	)
	myQueue = TwoStack2QueueConstructor()
	expected = []int{1, 1, 0}
	myQueue.Push(1)
	myQueue.Push(2)
	res = append(res, myQueue.Peek())
	res = append(res, myQueue.Pop())
	res = append(res, common.BoolToInt(myQueue.Empty()))
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)
}

func TestTwoQueue2Stack(t *testing.T) {
	var (
		myStack       MyStack
		res, expected = make([]int, 0), make([]int, 0)
	)
	myStack = TwoQueue2StackConstructor()
	expected = []int{2, 2, 0}
	myStack.Push(1)
	myStack.Push(2)
	res = append(res, myStack.Top())
	res = append(res, myStack.Pop())
	res = append(res, common.BoolToInt(myStack.Empty()))
	t.Logf("%v, res:%+v, expected:%+v", common.DiffTwoIntSlice(res, expected), res, expected)
}

func TestRemoveKDigits(t *testing.T) {
	var list = []common.Item33{
		{Str: "10001", N: 1, Expected: "1"},
		{Str: "100010", N: 1, Expected: "10"},
		{Str: "1432219", N: 3, Expected: "1219"},
		{Str: "10200", N: 1, Expected: "200"},
		{Str: "10", N: 1, Expected: "0"},
		{Str: "10", N: 2, Expected: "0"},
		{Str: "10", N: 3, Expected: "0"},
		{Str: "10001", N: 4, Expected: "0"},
		{Str: "10001", N: 2, Expected: "0"},
		{Str: "1102912", N: 2, Expected: "2912"},
		{Str: "1102912", N: 3, Expected: "212"},
		{Str: "11111", N: 3, Expected: "11"},
		{Str: "12345", N: 2, Expected: "123"},
	}
	var res string
	for _, item := range list {
		res = removeKdigits1(item.Str, item.N)
		t.Logf("%v, res:%s, item:%+v", res == item.Expected, res, item)
		res = removeKdigits2(item.Str, item.N)
		t.Logf("%v, res:%s, item:%+v", res == item.Expected, res, item)
		t.Log("----------------------SPLIT-------------------")
	}
}
