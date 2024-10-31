package main

func BinarySearch(arr []int, elem int) int {
	var left, right, mid int = 0, len(arr), 0
	for left < right {
		mid = (left + right) / 2
		if arr[mid] < elem {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
