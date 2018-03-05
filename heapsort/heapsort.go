package main

import "fmt"

// 堆排序
func heapSort(list []int) {
	// 构建最大堆
	l := len(list)
	m := l / 2
	for i := m; i > -1; i-- {
		heap(list, i, l - 1)
	}

	// 交换根尾节点并删除根节点
	for i := l - 1; i > 0; i-- {
		list[i], list[0] = list[0], list[i]
		heap(list, 0, i-1)
	}
    
}

func heap(list []int, i, end int) {
	// 左节点
	ln := 2*i+1
	if ln > end {
		return
	}
	n := ln

	// 右节点
	rn := 2*i+2
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

func main() {
	list := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	heapSort(list)
	fmt.Println(list)
}