package main

func getMinRun(N int) int {
	var r int = 0
	for {
		if N <= 64 {
			break
		}
		r |= N & 1
		N >>= 1
	}
	return r + N
}

func mergeRuns() {

}

func TimSort(arr []int) {
	var N int = len(arr)

	if N <= 64 {
		InsertionSort(arr)
		return
	}
	var minRun int = getMinRun(N)

	var p *int = &arr[0]

}
