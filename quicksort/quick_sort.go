package main

import "fmt"

func main() {
	list := []int{1, 3, 4, 5, 6, 100, 43, 3}
	fmt.Println("before quick sort:", list)
	fmt.Println("after quick sort:", QuickSort(list))
}

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
	} else {
		return list
	}
}
