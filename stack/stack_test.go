package stack

import (
	"fmt"
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
	for i := 0; i < 10; i++ {
		res1 := fibonacci1(i)
		res2 := fibonacci2(i)
		t.Logf("res:  %d, %d", res1, res2)
	}
}

func TestNumWays(t *testing.T) {
	for i := 0; i < 100; i++ {
		res1 := numWays1(i)
		res2 := numWays2(i)
		t.Logf("n:%d, res:  %d, %d, %t", i, res1, res2, res1 == res2)
	}
}
