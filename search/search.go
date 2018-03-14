package search

// 折半查找非递归算法
func BinarySearch(s []int, key int) (pos int) {
	low, high := 0, len(s)-1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if s[mid] == key {
			return mid
		}

		if s[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 折半查找递归算法
func RecursionBinarySearch(s []int, key, low, high int) (pos int) {
	mid := (low + high) / 2
	if s[mid] == key {
		return mid
	}

	if low > high {
		return -1
	}

	if s[mid] > key {
		return RecursionBinarySearch(s, key, low, mid-1 )
	} else {
		return RecursionBinarySearch(s, key, mid+1, high)
	}
}