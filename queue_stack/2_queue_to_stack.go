package queue_stack

//用队列实现栈(leetcode225)
//使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop和empty）。
//实现MyStack类：
//void push(int x) 将元素 x 压入栈顶。
//int pop() 移除并返回栈顶元素。
//int top() 返回栈顶元素。
//boolean empty() 如果栈是空的，返回 true ；否则，返回false。
// 注意：
// 你只能使用队列的基本操作 —— 也就是push to back、peek/pop from front、size 和is empty这些操作。
// 你所使用的语言也许不支持队列。你可以使用 array （列表）或者 deque（双端队列）来模拟一个队列, 只要是标准的队列操作即可。
//示例：输入：["MyStack","push","push","top","pop","empty"]
//	[[],[1],[2],[],[],[]]
//输出：[null,null,null,2,2,false]解释：
//	MyStack myStack=new MyStack();
//	myStack.push(1);
//	myStack.push(2);
//	myStack.top();//返回2
//	myStack.pop();//返回2
//	myStack.empty();//返回False

type MyStack struct {
	queue     []int
	helpQueue []int
}

// 1,2,3,4

func TwoQueue2StackConstructor() MyStack {
	q := MyStack{
		queue:     make([]int, 0),
		helpQueue: make([]int, 0),
	}
	return q
}

// 添加元素为O(n)
// 每次先往辅助队列中添加，再将主队列中的元素添加到辅助队列中，最后互换主辅
func (this *MyStack) Push(x int) {
	// 1.新元素先入辅助队列
	this.helpQueue = append(this.helpQueue, x)
	// 2. 将主队列元素依次入辅
	if len(this.queue) > 0 {
		this.helpQueue = append(this.helpQueue, this.queue...)
		this.queue = this.queue[:0]
	}
	// 3. 将主辅互换
	this.queue, this.helpQueue = this.helpQueue, this.queue
}

func (this *MyStack) Pop() int {
	if len(this.queue) > 0 {
		val := this.queue[0]
		this.queue = this.queue[1:]
		return val
	}
	return -1
}

func (this *MyStack) Top() int {
	if len(this.queue) > 0 {
		return this.queue[0]
	}
	return -1
}

// 如果栈是空的，返回true；否则返回false;
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// -------法二--------

func (this *MyStack) Push2(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop2() int {
	if this.Empty() {
		return -1
	}

	for len(this.queue) != 1 {
		this.helpQueue = append(this.helpQueue, this.queue[0])
		this.queue = this.queue[1:]
	}
	ans := this.queue[0]
	this.queue = append([]int{}, this.helpQueue...)
	this.helpQueue = nil
	return ans
}

func (this *MyStack) Top2() int {
	if this.Empty() {
		return -1
	}
	ans := this.Pop()
	this.queue = append(this.queue, ans)
	return ans
}
