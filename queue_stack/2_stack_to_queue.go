package queue_stack

//用栈实现队列(leetcode-232/sword-209)
//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//实现 MyQueue 类：
//	void push(int x) 将元素 x 推到队列的末尾
//	int pop() 从队列的开头移除并返回元素
//	int peek() 返回队列开头的元素
//	boolean empty() 如果队列为空，返回 true ；否则，返回 false
//说明：
//	你只能使用标准的栈操作——也就是只有pushtotop,peek/popfromtop,size,和isempty操作是合法的。
//	你所使用的语言也许不支持栈。你可以使用list或者deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//示例1：输入：["MyQueue","push","push","peek","pop","empty"]
//	[[],[1],[2],[],[],[]]
//输出：[null,null,null,1,1,false]
//解释：
//	MyQueue myQueue = new MyQueue();
//	myQueue.push(1); // queue is: [1]
//	myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//	myQueue.peek(); // return 1
//	myQueue.pop(); // return 1, queue is [2]
//	myQueue.empty(); // return false

type MyQueue struct {
	length    int   // 元素数量
	stack     []int // 栈顶为最后加入的元素
	helpStack []int // 栈顶为最先加入的元素
}

func TwoStack2QueueConstructor() MyQueue {
	return MyQueue{
		// 栈顶为stack[len(stack)-1]
		stack:     make([]int, 0),
		helpStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {

	// 优先将helpStack中的元素加入stack中
	for len(this.helpStack) > 0 {
		n := len(this.helpStack) - 1
		this.stack = append(this.stack, this.helpStack[n])
		this.helpStack = this.helpStack[:n]
	}
	this.stack = append(this.stack, x)
	this.length++
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	//优先从helpStack中pop,这样可以连续多次pop
	// 优先将stack中的元素添加到helpStack中
	var n = len(this.stack) - 1
	for len(this.stack) > 0 {
		this.helpStack = append(this.helpStack, this.stack[n])
		this.stack = this.stack[:n]
		n--
	}

	n = len(this.helpStack) - 1
	var ans = this.helpStack[n]
	this.helpStack = this.helpStack[:n]
	this.length--
	return ans
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	ans := this.Pop()
	this.helpStack = append(this.helpStack, ans)
	this.length++
	return ans
}

func (this *MyQueue) Empty() bool {
	return this.length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
