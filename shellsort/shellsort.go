package main

import "fmt"

// 希尔排序
func shellSort(list []int) {
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

func main() {
	list := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	shellSort(list)
	fmt.Println(list)
}
