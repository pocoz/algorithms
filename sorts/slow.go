package sorts

func findSmallest(arr []int) (int, int) {
	var (
		smallest = arr[0]
		index    = 0
	)

	for i, a := range arr {
		if a < smallest {
			smallest = a
			index = i
		}
	}
	return smallest, index
}

func removeFromArray(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr
}

func Slow(arr []int) []int {
	var sortArr = make([]int, 0)
	for range arr {
		smallest, index := findSmallest(arr)
		arr = removeFromArray(arr, index)
		sortArr = append(sortArr, smallest)
	}
	return sortArr
}
