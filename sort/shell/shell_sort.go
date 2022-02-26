package shell

func ShellSort(sli []int) {
	length := len(sli)
	if length < 2 {
		return
	}
	step := length / 2

	for step >= 1 {
		insertSort(sli, step)
		step = step / 2
	}
}

func insertSort(sli []int, step int) {

	length := len(sli)

	for i := 1; i < length; i += step {
		for j := i; j > 0 && sli[j] < sli[j-1]; j -= step {
			sli[j], sli[j-1] = sli[j-1], sli[j]
		}
	}
}

func insertSort1(tree []int, step int) {
	for i := step; i < len(tree); i++ {
		for j := i; j >= step; j -= step {
			if tree[j] >= tree[j-step] {
				break
			}
			tree[j], tree[j-step] = tree[j-step], tree[j]
		}
	}
}

func ShellSort2(arr []int) {
	increment := len(arr) / 2
	for increment > 0 {
		for i := increment; i < len(arr); i++ {
			j := i
			temp := arr[i]

			for j >= increment && arr[j-increment] > temp {
				arr[j] = arr[j-increment]
				j = j - increment
			}
			arr[j] = temp
		}
		if increment == 2 {
			increment = 1
		} else {
			increment = int(increment * 5 / 11)
		}
	}
}
