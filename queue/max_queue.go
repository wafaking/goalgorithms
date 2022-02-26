package queue

//剑指 Offer 59 - II. 队列的最大值
//请定义一个队列并实现函数max_value得到队列里的最大值，
//要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//若队列为空，pop_front 和 max_value 需要返回 -1

//示例 1：
//输入:
//["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
//[[],[1],[2],[],[],[]]
//输出: [null,null,null,2,1,2]

//示例 2：
//输入:
//["MaxQueue","pop_front","max_value"]
//[[],[],[]]
//输出: [null,-1,-1]

type MaxQueue struct {
	queue     []int
	helpQueue []int
}

func Constructor() MaxQueue {
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

	// 把value添加到辅助队列的尾部，将小于value的元素弹出
	for len(this.helpQueue) > 0 && this.helpQueue[len(this.helpQueue)-1] < value {
		this.helpQueue = this.helpQueue[:len(this.helpQueue)-1] // 注：此处应将尾部元素弹出（双端队列）保证头部最大
	}
	this.helpQueue = append(this.helpQueue, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.queue) <= 0 {
		return n
	}
	n = this.queue[0] // 先进先出
	this.queue = this.queue[1:]
	if n == this.helpQueue[0] { // 只有弹出的元素大于辅助队列中的最大元素时，才将其弹出
		this.helpQueue = this.helpQueue[1:]
	}
	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

//type MaxQueue struct {
//	q1 []int
//	max []int
//}
//
//func Constructor() MaxQueue {
//	return MaxQueue{
//		make([]int,0),
//		make([]int,0),
//	}
//}
//
//func (this *MaxQueue) Max_value() int {
//	if len(this.max) == 0{
//		return -1
//	}
//	return this.max[0]
//}
//
//func (this *MaxQueue) Push_back(value int)  {
//	this.q1 = append(this.q1,value)
//	for len(this.max) != 0 && value > this.max[len(this.max)-1]{
//		this.max = this.max[:len(this.max)-1]
//	}
//	this.max = append(this.max,value)
//}
//
//func (this *MaxQueue) Pop_front() int {
//	n := -1
//	if len(this.q1) != 0{
//		n = this.q1[0]
//		this.q1 = this.q1[1:]
//		if this.max[0] == n{
//			this.max = this.max[1:]
//		}
//	}
//	return n
//}
