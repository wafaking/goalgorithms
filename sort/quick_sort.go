package sort

// 使用数组
func quickSort1(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var (
		l, r = 0, len(nums) - 1
		// 此处可以选用随机元素代替中间元素
		mid = (r-l)/2 + l
		// 分别用于存储小于和等于pivot元素
		left, right = make([]int, 0), make([]int, 0)
	)

	for i := range nums {
		if i == mid { // 忽略当前元素
			continue
		} else if nums[i] > nums[mid] {
			right = append(right, nums[i])
		} else {
			left = append(left, nums[i])
		}
	}

	res1 := quickSort1(left)
	res2 := quickSort1(right)
	return append(append(res1, nums[mid]), res2...)
}

// 选取末尾元素作为pivot
func quickSort2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var (
		l, r = 0, len(nums) - 1
		// 每次都取最右侧元素为比较元素
		pivot = nums[r]
	)

	for i := range nums {
		// 将小于Pivot的元素与左侧元素交换，其中l指向待交换的位置，即大于pivot的元素位置
		if nums[i] < pivot {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
	// 交换完成后，l位置元素即为第一个大于等于pivot的元素
	// 将l与pivot元素交换，l位置即成为大于和小于Pivot元素的分界点
	nums[l], nums[r] = nums[r], nums[l]

	quickSort2(nums[:l])
	quickSort2(nums[l+1:])
	return nums
}

// 使用指针而不拷贝数组
func quickSort3(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var f func(l, r int)
	f = func(l, r int) {
		if l >= r {
			return
		}
		var (
			pivot = nums[r]
			idx   = l
		)
		// mid := (r-1)/2 + l
		for i := l; i < r; i++ {
			if nums[i] < pivot {
				nums[i], nums[idx] = nums[idx], nums[i]
				idx++
			}
		}
		nums[r], nums[idx] = nums[idx], nums[r]
		f(l, idx-1)
		f(idx+1, r)
	}
	f(0, len(nums)-1)
	return nums
}

// 分区找出分界点
func quickSort4(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var (
		partition func(l, r int) int
		sort      func(l, r int)
	)
	partition = func(l, r int) int {
		// pivot指定第一个元素
		pivot := nums[l]
		for l < r {
			// 先从右侧比较元素，找出小于pivot的元素
			for l < r && pivot <= nums[r] {
				r--
			}

			// 将小于pivot的元素放到l处
			// 此时相当于r处的元素被移走了
			nums[l] = nums[r]

			// 交替，从左侧开始比较元素，
			// 找到比pivot大的元素放到r处
			for l < r && pivot >= nums[l] {
				l++
			}
			// 将大于pivot的元素放到r片
			// 此时相当于l处的元素被移走了
			nums[r] = nums[l]

			// 下一次循环又从右侧开始
		}
		// 循环结束，相当于l位置的元素被移空
		// 将pivot元素放到l处，l处元素将nums分成小于和大小pivot元素的两部分
		nums[l] = pivot
		return l
	}

	sort = func(l, r int) {
		if l >= r {
			return
		}

		pivot := partition(l, r)
		// pivot位置已排序，可以不包含
		sort(l, pivot-1)
		sort(pivot+1, r)
	}

	sort(0, len(nums)-1)
	return nums
}
