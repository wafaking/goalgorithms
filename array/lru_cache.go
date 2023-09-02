package array

// LRU缓存(leetcode-146)
//请你设计并实现一个满足LRU(最近最少使用)缓存约束的数据结构。
//实现LRUCache类：
//	LRUCache(int capacity)以正整数作为容量capacity初始化LRU缓存
//	int get(int key)如果关键字key存在于缓存中，则返回关键字的值，否则返回-1。
//	void put(int key,int value)如果关键字key已经存在，则变更其数据值value；
//	如果不存在，则向缓存中插入该组key-value。如果插入操作导致关键字数量超过capacity，则应该逐出最久未使用的关键字。
//函数get和put必须以O(1)的平均时间复杂度运行。

//示例：输入 ["LRUCache","put","put","get","put","get","put","get","get","get"]
//	[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
//	输出 [null,null,null,1,null,-1,null,-1,3,4]
//解释:
//	LRUCache lRUCache=newLRUCache(2);
//	lRUCache.put(1,1);//缓存是{1=1}
//	lRUCache.put(2,2);//缓存是{1=1,2=2}
//	lRUCache.get(1);//返回1
//	lRUCache.put(3,3);//该操作会使得关键字2作废，缓存是{1=1,3=3}
//	lRUCache.get(2);//返回-1(未找到)
//	lRUCache.put(4,4);//该操作会使得关键字1作废，缓存是{4=4,3=3}
//	lRUCache.get(1);//返回-1(未找到)
//	lRUCache.get(3);//返回3
//	lRUCache.get(4);//返回4

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//法一：双向链表(使用已有元素构建首尾)
//法二：双向链表(构造虚拟头尾元素，防止链表为空)

type LRUCache struct {
	Capacity, Len int
	Head          *LinkedNode
	Tail          *LinkedNode
	LinkedNodeMap map[int]*LinkedNode
}

type LinkedNode struct {
	Key, Value int
	Pre        *LinkedNode
	Next       *LinkedNode
}

func ConstructorLRU(capacity int) LRUCache {
	var l = LRUCache{
		Capacity:      capacity,
		LinkedNodeMap: make(map[int]*LinkedNode),
	}
	return l
}

func (this *LRUCache) Get(key int) int {
	if this.Capacity == 0 {
		return -1
	}

	item, ok := this.LinkedNodeMap[key]
	if !ok {
		return -1
	}

	this.first(item)
	return item.Value
}

func (this *LRUCache) first(item *LinkedNode) {
	if item.Pre == nil { // item is the first
		return
	}
	itemPre := item.Pre
	itemNext := item.Next

	if itemNext == nil { // item is the tail
		// 将item的pre与next相连
		itemPre.Next = itemNext

		// 将item放到头部
		head := this.Head
		head.Pre = item
		item.Pre = nil
		item.Next = head

		// 重新赋值head与tail
		this.Tail = itemPre
		this.Head = item

	} else { // item is not the tail

		// 将item的pre与next相连
		itemPre.Next = itemNext
		itemNext.Pre = itemPre

		// 将item放到头部
		head := this.Head
		head.Pre = item
		item.Pre = nil
		item.Next = head

		// 重新赋值head
		//this.Tail = itemPre
		this.Head = item
	}
}

func (this *LRUCache) Put(key int, value int) {

	obj, exist := this.LinkedNodeMap[key]
	if exist {
		obj.Value = value
		this.first(obj)
		return
	}

	// not exist
	obj = &LinkedNode{
		Key:   key,
		Value: value,
		Pre:   nil,
		Next:  this.Head,
	}
	this.Len++

	this.LinkedNodeMap[key] = obj

	// no head
	head := this.Head
	if head == nil {
		this.Head = obj
		this.Tail = obj
		return
	}

	head.Pre = obj
	this.Head = obj
	if this.Capacity < this.Len {
		oldTail := this.Tail
		delete(this.LinkedNodeMap, oldTail.Key)

		newTail := oldTail.Pre
		newTail.Next = nil
		this.Tail = newTail
		this.Len--
	}
}
