package main

import "fmt"

// 归并排序
func mergeSort(list []int) []int {
	n := len(list)
	if n < 2 {
		return list
	}

	mid := n/2

	llist := mergeSort(list[:mid])
	rlist := mergeSort(list[mid:])

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

func main() {
	list := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	fmt.Println(mergeSort(list))
}
