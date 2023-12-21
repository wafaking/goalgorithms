package queue_stack

//队列的最大值(sword-2-59/LCR184)
//请定义一个队列并实现函数max_value得到队列里的最大值，
//要求函数max_value、push_back和pop_front的均摊时间复杂度都是O(1)。
//若队列为空，pop_front和max_value需要返回-1
//示例1：输入:["MaxQueue","push_back","push_back","max_value","pop_front","max_value"],
//	[[],[1],[2],[],[],[]]
//输出:[null,null,null,2,1,2]
//示例2：输入:["MaxQueue","pop_front","max_value"][[],[],[]]输出:[null,-1,-1]

type MaxQueue struct {
	queue     []int
	helpQueue []int // 双端队列
}

func MaxInQueueConstructor2() MaxQueue {
	return MaxQueue{
		queue:     make([]int, 0), // 单向队列，先进先出(头queue[0])
		helpQueue: make([]int, 0), // 辅助双端队列，先进先出
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.helpQueue) == 0 {
		return -1
	}
	return this.helpQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)

	for len(this.helpQueue) > 0 {
		tail := this.helpQueue[len(this.helpQueue)-1]
		if tail > value { // 移除尾部元素
			break
		}
		this.helpQueue = this.helpQueue[:len(this.helpQueue)-1]
	}

	this.helpQueue = append(this.helpQueue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	val := this.queue[0] // 先进先出
	this.queue = this.queue[1:]

	// 只有弹出的元素大于辅助队列中的最大元素时，才将其弹出
	if val == this.helpQueue[0] {
		this.helpQueue = this.helpQueue[1:]
	}
	return val
}
