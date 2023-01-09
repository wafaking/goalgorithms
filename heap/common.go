package heap

// Heap 用[]int,实现小顶堆
type Heap []int

// len get h's length(注：可以不用*Heap，因heap是引用，只要长度不会变，操作的同一地址)
func (h Heap) len() int {
	return len(h)
}

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	return
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

// up to up adjust the index i's position
func (h *Heap) up(i int) {
	for {
		var parent = (i - 1) / 2 // 注：parent编号与数组中的位置不同
		if (i > h.len()-1) || (parent < 0) || h.less(parent, i) {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

// down to down adjust the index i's position
func (h *Heap) down(i int) {
	for {
		var subLeft = 2*i + 1  // 左子节点
		var subRight = 2*i + 2 // 右子节点

		// 判断是不是左、右子节点：
		// 如果2*k>n，则此节点无左子节点(k为编号)
		// 如果2*k+1>n，则此节点无右子节点
		if subLeft >= (h.len() - 1) {
			break
		}

		// 与子节点中较小的交换(每个节点的值都小于其左、右子节点的值)
		var minSub = subLeft
		if (subRight <= h.len()-1) && h.less(subRight, subLeft) {
			// subRight存在是前提
			minSub = subRight
		}
		if h.less(minSub, i) {
			h.swap(i, minSub)
		}
		i = minSub
	}
}

//从最小堆中删除节点，分为以下三个步骤
//把最末端的节点和要删除节点的位置进行交换。
//删除末端节点
//原来的末端节点需要与新位置上的父节点做比较，如果小于要做 up（看上面的方法），如果大于父节点，则再和子节点做比较，即 down 操作，直到该节点下降到小于最小子节点为止。
// remove 移除指定位置值
func (h *Heap) remove(i int) int {
	// 把节点i与最末端的节点交换
	h.swap(i, h.len()-1)
	// 删除末端节点
	var remVal = (*h)[h.len()-1]
	*h = (*h)[:h.len()-1]
	if h.len() == 0 {
		return remVal
	}

	// 将节点i与父节点比较，如果小于则up,否则down
	var parent = (i - 1) / 2
	if h.less(i, parent) {
		h.up(i)
	} else {
		h.down(i)
	}

	return remVal
}

func (h *Heap) Pop() int {
	return h.remove(0)
}

func (h *Heap) Push(val int) {
	*h = append(*h, val)
	h.up(len(*h) - 1)
}

func BuildHeap(array []int) Heap {
	var h = Heap(array)
	//建立堆的过程就是完全二叉树从下到上调整堆的过程，
	//从i=len(arr)/2开始依次向上调整，i=len(arr)/2是堆中末尾节点的父节点， i=0是根节点。
	for i := len(h)/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

// HeapSort 堆排序
func HeapSort(array []int) []int {
	heap := BuildHeap(array)
	var res = make([]int, 0, heap.len())
	for len(heap) > 0 {
		res = append(res, heap.Pop())
	}
	return res
}
