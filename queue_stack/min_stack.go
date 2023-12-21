package queue_stack

//最小栈(leetcode-155/LCR-147)
//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min函数在该栈中，
//调用min、push及pop的时间复杂度都是O(1)。
//示例:
//	MinStack minStack = new MinStack();
//	minStack.push(-2);
//	minStack.push(0);
//	minStack.push(-3);
//	minStack.min();   --> 返回 -3.
//	minStack.pop();
//	minStack.top();      --> 返回 0.
//	minStack.min();   --> 返回 -2.

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

type MinStack struct {
	stack     []int // 存放元素
	helpStack []int // 辅助栈
}

func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		helpStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	length := len(this.helpStack)
	if length == 0 {
		this.helpStack = append(this.helpStack, val)
		return
	}
	min := this.helpStack[length-1]
	if min > val { // 辅助栈里的值要大于x，则添加x
		this.helpStack = append(this.helpStack, val)
	} else {
		this.helpStack = append(this.helpStack, min)
	}
}

func (this *MinStack) Pop() {

	length := len(this.stack)
	if length <= 0 {
		return
	}

	this.stack = this.stack[:length-1]
	this.helpStack = this.helpStack[:length-1]
}

func (this *MinStack) Top() int {
	length := len(this.stack)
	if length <= 0 {
		return -1
	}
	return this.stack[length-1]
}

func (this *MinStack) GetMin() int {
	length := len(this.helpStack)
	if length <= 0 {
		return 0
	}
	return this.helpStack[length-1]
}
