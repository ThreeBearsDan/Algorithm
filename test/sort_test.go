package test

import (
	"testing"

	"github.com/ThreeBearsDan/Algorithm/sort"
	"github.com/stretchr/testify/assert"
)

var sortlist = []int{1, 2, 3, 3, 3, 4, 6, 7, 10, 12}

func TestBubbleSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	sort.BubbleSort(unsortlist)
	if ok := assert.Equal(t, sortlist, unsortlist); ok != true {
		t.Error("test bubble sort failed")
	}
}

func TestHeapSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	sort.HeapSort(unsortlist)
	if ok := assert.Equal(t, sortlist, unsortlist); ok != true {
		t.Error("test heap sort failed")
	}
}

func TestInsertionSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	sort.InsertionSort(unsortlist)
	if ok := assert.Equal(t, sortlist, unsortlist); ok != true {
		t.Error("test insertion sort failed")
	}
}

func TestMergeSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	if ok := assert.Equal(t, sortlist, sort.MergeSort(unsortlist)); ok != true {
		t.Error("test merge sort failed")
	}
}

func TestQuickSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	if ok := assert.Equal(t, sortlist, sort.QuickSort(unsortlist)); ok != true {
		t.Error("test quick sort failed")
	}
}

func TestShellSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	sort.ShellSort(unsortlist)
	if ok := assert.Equal(t, sortlist, unsortlist); ok != true {
		t.Error("test shell sort failed")
	}
}

func TestSelectSort(t *testing.T) {
	unsortlist := []int{2, 3, 4, 1, 10, 6, 7, 3, 3, 12}
	sort.SelectSort(unsortlist)
	if ok := assert.Equal(t, sortlist, unsortlist); ok != true {
		t.Error("test select sort failed")
	}
}
