package test

import ( 
	"github.com/ThreeBearsDan/Algorithm/search"
	"testing"
)

var list = []int{1, 2, 5, 8, 10, 20, 100}

var testcase = map[int]int{
    1: 0,
	2: 1,
	5: 2,
	8: 3,
	10: 4,
	20: 5,
	100: 6,  
	3: -1,
	9: -1,
	18: -1,
	25: -1,
	200: -1,
}

func TestBinarySearch(t *testing.T) {
	for key, val := range testcase {
		if val != search.BinarySearch(list, key) {
			t.Errorf("search '%d' from '%v' error", key, list)
		}
	}
}

func TestRecursionBinaryTest(t *testing.T) {
	for key, val := range testcase {
		if val != search.RecursionBinarySearch(list, key, 0, len(list)-1 ) {
			t. Errorf("search '%d' from '%v' error", key, list)
		}
	}
}
