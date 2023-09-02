package array

//LFU缓存(leetcode-460)
//请你为最不经常使用（LFU）缓存算法设计并实现数据结构。
//实现LFUCache类：
//	LFUCache(int capacity)-用数据结构的容量capacity初始化对象
//	int get(int key)-如果键key存在于缓存中，则获取键的值，否则返回-1。
//	void put(int key,int value)-如果键key已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量capacity时，
//	则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除最近最久未使用的键。
//	为了确定最不常使用的键，可以为缓存中的每个键维护一个使用计数器。使用计数最小的键是最久未使用的键。
//	当一个键首次插入到缓存中时，它的使用计数器被设置为1(由于put操作)。对缓存中的键执行get或put操作，使用计数器的值将会递增。
//函数get和put必须以O(1)的平均时间复杂度运行。

//示例：输入：["LFUCache","put","put","get","put","get","get","put","get","get","get"]
//	[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
//	输出： [null,null,null,1,null,-1,3,null,-1,3,4]
//解释：
//	//cnt(x)=键x的使用计数
//	//cache=[]将显示最后一次使用的顺序（最左边的元素是最近的）
//	LFUCache lfu=newLFUCache(2);
//	lfu.put(1,1);//cache=[1,_],cnt(1)=1
//	lfu.put(2,2);//cache=[2,1],cnt(2)=1,cnt(1)=1
//	lfu.get(1);//返回1
//	//cache=[1,2],cnt(2)=1,cnt(1)=2
//	lfu.put(3,3);//去除键2，因为cnt(2)=1，使用计数最小
//	//cache=[3,1],cnt(3)=1,cnt(1)=2
//	lfu.get(2);//返回-1（未找到）
//	lfu.get(3);//返回3
//	//cache=[3,1],cnt(3)=2,cnt(1)=2
//	lfu.put(4,4);//去除键1，1和3的cnt相同，但1最久未使用
//	//cache=[4,3],cnt(4)=1,cnt(3)=2
//	lfu.get(1);//返回-1（未找到）
//	lfu.get(3);//返回3
//	//cache=[3,4],cnt(4)=1,cnt(3)=3
//	lfu.get(4);//返回4
//	//cache=[3,4],cnt(4)=2,cnt(3)=3

// TODO
type LFUCache struct {
	Capacity int //容量
	Size     int //元素数量
	Counter  int //计数器
	M        map[int]*LinkedNode
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{Capacity: capacity}
}

func (this *LFUCache) Get(key int) int {
	return 0
}

func (this *LFUCache) Put(key int, value int) {

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
