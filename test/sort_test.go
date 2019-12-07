package test

import (
	"testing"

	"reflect"

	"github.com/ThreeBearsDan/Algorithm/sort"
)

func Testsortinplace(name string, f func([]int), t *testing.T) {
	tests := []struct{ seq, want []int }{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{3, 2, 1, 10, 9, 6, 5, 7, 4, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{2, 2, 1, 5, 4, 2, 2, 6, 2, 9, 10}, []int{1, 2, 2, 2, 2, 2, 4, 5, 6, 9, 10}},
	}

	for _, v := range tests {
		f(v.seq)
		if !reflect.DeepEqual(v.seq, v.want) {
			t.Errorf("sort name: %s, want: %v, got: %v", name, v.want, v.seq)
		}
	}
}

func Testsortnonplace(name string, f func([]int) []int, t *testing.T) {
	tests := []struct{ seq, want []int }{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{3, 2, 1, 10, 9, 6, 5, 7, 4, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{2, 2, 1, 5, 4, 2, 2, 6, 2, 9, 10}, []int{1, 2, 2, 2, 2, 2, 4, 5, 6, 9, 10}},
	}

	for _, v := range tests {
		if !reflect.DeepEqual(f(v.seq), v.want) {
			t.Errorf("sort name: %s, want: %v, got: %v", name, v.want, f(v.seq))
		}
	}
}

func TestBubbleSort(t *testing.T) {
	Testsortinplace("bubble sort", sort.BubbleSort, t)
}

func TestHeapSort(t *testing.T) {
	Testsortinplace("heap sort", sort.HeapSort, t)
}

func TestInsertionSort(t *testing.T) {
	Testsortinplace("insertion sort", sort.InsertionSort, t)
}

func TestMergeSort(t *testing.T) {
	Testsortnonplace("merge sort", sort.MergeSort, t)
}

func TestQuickSort(t *testing.T) {
	Testsortnonplace("quick sort with extra space", sort.QuickSort, t)
}

func TestQuickSortInplace(t *testing.T) {
	Testsortinplace("quick sort in place", sort.QuickSortInplace, t)
}

func TestShellSort(t *testing.T) {
	Testsortinplace("shell sort", sort.ShellSort, t)
}

func TestSelectSort(t *testing.T) {
	Testsortinplace("select sort", sort.SelectSort, t)
}
