package main

import "fmt"

// 插入排序
func insertionSort(list []int) {
    for i := 1; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}

func main() {
	list := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	insertionSort(list)
	fmt.Println(list)
}