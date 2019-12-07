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

// Heap Sort
func HeapSort(s []int) {
	if len(s) < 2 {
		return
	}

	heap(s)
	HeapSort(s[:len(s)-1])
}

func heap(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		if s[(i-1)/2] < s[i] {
			s[(i-1)/2], s[i] = s[i], s[(i-1)/2]
		}
	}

	s[0], s[len(s)-1] = s[len(s)-1], s[0]
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
