package sort

// 冒泡排序
func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}

// 堆排序
func HeapSort(list []int) {
	// 构建最大堆
	l := len(list)
	m := l / 2
	for i := m; i > -1; i-- {
		heap(list, i, l-1)
	}

	// 交换根尾节点并删除根节点
	for i := l - 1; i > 0; i-- {
		list[i], list[0] = list[0], list[i]
		heap(list, 0, i-1)
	}

}

func heap(list []int, i, end int) {
	// 左节点
	ln := 2*i + 1
	if ln > end {
		return
	}
	n := ln

	// 右节点
	rn := 2*i + 2
	if rn <= end && list[rn] > list[n] {
		n = rn
	}

	// 调整根节点
	if list[i] < list[n] {
		list[i], list[n] = list[n], list[i]
	}
	// 对所有节点进行调整
	heap(list, n, end)
}

// 插入排序
func InsertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

// 归并排序
func MergeSort(list []int) []int {
	n := len(list)
	if n < 2 {
		return list
	}

	mid := n / 2

	llist := MergeSort(list[:mid])
	rlist := MergeSort(list[mid:])

	return merge(llist, rlist)
}

func merge(a []int, b []int) []int {
	temp := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			temp = append(temp, a[i])
			i++
		} else {
			temp = append(temp, b[j])
			j++
		}
	}
	temp = append(temp, a[i:]...)
	temp = append(temp, b[j:]...)

	return temp
}

// 快速排序
func QuickSort(list []int) []int {
	if len(list) > 1 {
		rlist := make([]int, 0)
		llist := make([]int, 0)
		key := make([]int, 0)
		key = append(key, list[0])
		for i := 1; i < len(list); i++ {
			if list[i] < key[0] {
				llist = append(llist, list[i])
			} else if list[i] > key[0] {
				rlist = append(rlist, list[i])
			} else {
				key = append(key, list[i])
			}
		}
		return append(QuickSort(llist), append(key, QuickSort(rlist)...)...)
	}
	return list
}

// 希尔排序
func ShellSort(list []int) {
	n := len(list)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			for j := i; j > gap-1; j -= gap {
				if list[j] < list[j-gap] {
					list[j], list[j-gap] = list[j-gap], list[j]
				}
			}
		}
		gap = gap / 2
	}
}

// 选择排序
func SelectSort(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
