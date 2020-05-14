package sorts

func Quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	if len(arr) == 2 {
		return Slow(arr)
	}

	var (
		numPivot = len(arr) / 2
		pivot    = arr[numPivot]
		less     = make([]int, 0)
		greater  = make([]int, 0)
	)

	for i, a := range arr {
		if i == numPivot {
			continue
		}
		if a <= pivot {
			less = append(less, a)
		} else {
			greater = append(greater, a)
		}
	}

	return append(append(Quick(less), pivot), Quick(greater)...)
}
