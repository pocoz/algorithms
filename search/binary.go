package search

func Binary(item int, arr []int) (interface{}, int) {
	var (
		low  = 0
		high = len(arr) - 1
		i    = 0
	)

	for {
		i++
		if low <= high {
			mid := (low + high) / 2
			guess := arr[mid]
			if guess == item {
				return guess, i
			}
			if guess < item {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			return nil, i
		}
	}
}
