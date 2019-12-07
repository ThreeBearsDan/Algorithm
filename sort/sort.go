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
			} else {
				break
			}
		}
	}
}

// Merge sort
func MergeSort(s []int) []int {
	length := len(s)
	if length < 2 {
		return s
	}

	ls := MergeSort(s[:length/2])
	rs := MergeSort(s[length/2:])

	return merge(ls, rs)
}

func merge(a []int, b []int) []int {
	result := make([]int, 0, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}

// Quick sort by allocating more space
func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	ls := make([]int, 0)
	rs := make([]int, 0)

	for i := 1; i < len(s); i++ {
		if s[0] > s[i] {
			ls = append(ls, s[i])
		} else {
			rs = append(rs, s[i])
		}
	}

	return append(append(QuickSort(ls), s[0]), QuickSort(rs)...)
}

// Quick sort in place
func QuickSortInplace(s []int) {
	if len(s) < 2 {
		return
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] > s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
			i++
		} else {
			s[j], s[i+1] = s[i+1], s[j]
			j--
		}
	}

	QuickSortInplace(s[:i])
	QuickSortInplace(s[i+1:])
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

// Select sort
func SelectSort(s []int) {
	if len(s) < 2 {
		return
	}

	for i := 0; i < len(s); i++ {
		idx := i
		min := s[i]
		for j := i + 1; j < len(s); j++ {
			if min > s[j] {
				idx = j
				min = s[j]
			}
		}

		s[i], s[idx] = s[idx], s[i]
	}
}
