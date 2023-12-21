package queue_stack

//设计自助结算系统(LCR184)
//请设计一个自助结账系统，该系统需要通过一个队列来模拟顾客通过购物车的结算过程，需要实现的功能有：
//	get_max()：获取结算商品中的最高价格，如果队列为空，则返回-1
//	add(value)：将价格为value的商品加入待结算商品队列的尾部
//	remove()：移除第一个待结算的商品价格，如果队列为空，则返回-1
//注意，为保证该系统运转高效性，以上函数的均摊时间复杂度均为O(1)
//示例1：输入:["Checkout","add","add","get_max","remove","get_max"]
//	[[],[4],[7],[],[],[]]
//	输出:[null,null,null,7,4,7]
//示例2：输入:["Checkout","remove","get_max"]
//	[[],[],[]]
//	输出:[null,-1,-1]

type Checkout struct {
	queue      []int
	helpQueue1 []int
	helpQueue2 []int
}

func MaxInQueueConstructor1() Checkout {
	return Checkout{
		queue:      make([]int, 0),
		helpQueue1: make([]int, 0),
		helpQueue2: make([]int, 0),
	}
}

func (this *Checkout) Get_max() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.helpQueue1[0]
}

func (this *Checkout) Add(value int) {
	this.queue = append(this.queue, value)
	if len(this.queue) == 1 {
		this.helpQueue1 = append(this.helpQueue1, value)
		return
	}
	if this.helpQueue1[0] < value {
		this.helpQueue1 = append([]int{}, value)
		return
	}
	for len(this.helpQueue1) > 0 {
		if value <= this.helpQueue1[0] {
			this.helpQueue2 = append(this.helpQueue2, this.helpQueue1[0])
		}
		this.helpQueue1 = this.helpQueue1[1:]
	}
	this.helpQueue1 = this.helpQueue2
	this.helpQueue1 = append(this.helpQueue1, value)
	this.helpQueue2 = nil
}

func (this *Checkout) Remove() int {
	if len(this.queue) == 0 {
		return -1
	}
	value := this.queue[0]
	this.queue = this.queue[1:]
	if value == this.helpQueue1[0] {
		this.helpQueue1 = this.helpQueue1[1:]
	}
	return value
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
