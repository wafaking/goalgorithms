package queue

// 用两个栈实现队列(sword2 09)
//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail(队列尾部插入整数)
//和deleteHead(队列头部删除整数)。(若队列中没有元素，deleteHead 操作返回 -1 )

//示例 1：
//输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]

//示例 2：
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]

type CQueue struct {
	stack     []int // 主
	helpStack []int // 辅助
}

func Constructor2() CQueue {
	var cQueue = CQueue{
		stack:     make([]int, 0, 100),
		helpStack: make([]int, 0, 100),
	}
	return cQueue
}

func (this *CQueue) AppendTail(value int) {
	this.stack = append(this.stack, value)
}

func (this *CQueue) DeleteHead() int {
	// 先取辅助栈里元素
	if len(this.helpStack) > 0 {
		val := this.helpStack[0]
		this.helpStack = this.helpStack[1:]
		return val
	}
	// 取主栈里元素
	if len(this.stack) == 0 {
		return -1
	}

	// 将主栈中的元素都添加到辅助栈中
	this.helpStack = append(this.helpStack, this.stack...)
	this.stack = this.stack[:0]
	return this.DeleteHead() // 重新删除
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
