package main

import (
	"fmt"
)

func main() {
	list := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	bubbleSort(list)
	fmt.Println(list)
}

func bubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}
