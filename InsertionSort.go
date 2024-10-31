package main

func InsertionSort(arr []int) {
	var x, j int

	for i := 1; i < len(arr); i++ {
		x = arr[i]
		for j = i; j > 0 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}
