package main

func MergeSort(right, left []int) []int {
	var leftCount, rightCount int = 0, 0
	var i, j = 0, 0
	result := make([]int, 0)

	for i < len(right) && j < len(left) {
		if left[i] < right[j] {
			result = append(result, left[i])
			leftCount++
			rightCount = 0
			i++
			if i == 7 {
				pos := BinarySearch(left[i:], right[j])
				tmp := i
				for index := 0; index < pos; index++ {
					result = append(result, left[tmp+index])
				}
			}
		} else {
			result = append(result, right[j])
			rightCount++
			leftCount = 0
			j++
			if j == 7 {
				pos := BinarySearch(right[j:], left[i])
				tmp := i
				for index := 0; index < pos; index++ {
					result = append(result, right[tmp+index])
				}
			}
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
