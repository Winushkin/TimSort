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

func reverse(arr []int) {
	var left, right int = 0, len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func mergeRuns(runs Stack) []int {
	var runsAmount int = runs.Len()
	stack := Stack{}

	for runs.Len() > 0 {
		stack.PushBack(runs.PopFront())
	}

	if runsAmount == 1 {
		return runs.Peek()

	} else if runsAmount == 2 {
		var X []int = stack.PopFront()
		var Y []int = stack.PopFront()
		return MergeSort(X, Y)

	} else {
		for stack.Len() >= 3 {
			var X []int = stack.PopBack()
			var Y []int = stack.PopBack()
			var Z []int = stack.PopBack()

			if len(Y) <= len(X) {
				X = MergeSort(X, Y)
				runs.PushBack(Z)
				runs.PushBack(X)

			} else if len(Z) <= len(X)+len(Y) {
				Y = MergeSort(Y, Z)
				runs.PushBack(Y)
				runs.PushBack(X)

			} else {
				Y = MergeSort(X, Y)
				runs.PushBack(Y)
			}
		}

	}

	if stack.Len() == 2 {
		var X []int = runs.PopFront()
		var Y []int = runs.PopFront()
		return MergeSort(X, Y)
	}

	return runs.Peek()
}

func TimSort(arr []int) []int {
	var N int = len(arr)
	var runs Stack = Stack{}
	var index int = 0
	var RunFailFlag bool = false
	var run []int

	if N <= 64 {
		InsertionSort(arr)
		return arr
	}
	var minRun int = getMinRun(N)

	if len(arr) <= 1 {
		return arr
	}

	for index < len(arr) {
		index += 2
		run = arr[index-2 : index]

		if index >= len(arr) {
			if arr[index-2] > arr[index-1] {
				reverse(run)
			}
			runs.PushBack(run)
		}

		if arr[index-2] > arr[index-1] {
			for arr[index-1] > arr[index] {
				run = append(run, arr[index])
				index++
				if index >= len(arr) {
					break
				}
			}
		} else if arr[index-2] <= arr[index-1] {
			for arr[index-1] <= arr[index] {
				run = append(run, arr[index])
				index++
				if index >= len(arr) {
					break
				}
			}
			reverse(run)
		}
		for len(run) < minRun && index < len(arr) {
			RunFailFlag = true
			run = append(run, arr[index])
			index++
			if index >= len(arr) {
				break
			}
		}
		if RunFailFlag {
			InsertionSort(run)
		}

		runs.PushBack(run)
	}

	return mergeRuns(runs)
}
