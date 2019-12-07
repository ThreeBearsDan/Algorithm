package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tables := []struct {
		data []int
		want []int
	}{
		{
			[]int{1, 3, 2, 10, 100, 99, 8},
			[]int{1, 2, 3, 8, 10, 99, 100},
		},
		{
			[]int{65, 24, 13, 2, 10, 19},
			[]int{2, 10, 13, 19, 24, 65},
		},
		{
			[]int{2, 2, 2, 3, 4, 2, 2, 2, 2},
			[]int{2, 2, 2, 2, 2, 2, 2, 3, 4},
		},
	}

	for _, v := range tables {
		got := QuickSort(v.data)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("got [%v] want [%v]", got, v.want)
		}

		data := v.data
		QuickSortInplace(data)
		if !reflect.DeepEqual(data, v.want) {
			t.Errorf("got [%v] want [%v]", got, v.want)
		}
	}
}
