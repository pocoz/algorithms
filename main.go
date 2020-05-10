package main

import (
	"fmt"
	"time"

	"github.com/pocoz/algorithms/binary_search"
	"github.com/pocoz/algorithms/sorts"
)

func main() {
	// sort quick
	begin := time.Now()
	arr := []int{34, 6, 13, 22, 2, 54, 67, 333, 188}
	fmt.Println(sorts.Quick(arr))
	fmt.Println(time.Since(begin))

	//sort slow
	arr = []int{34, 6, 13, 22, 2, 54, 67, 333, 188}
	begin = time.Now()
	fmt.Println(sorts.Slow(arr))
	fmt.Println(time.Since(begin))

	// binary search
	arr = []int{1,4,6,13,22,44,54,67,73,188}
	item := 34
	res, i := binary_search.Search(item, arr)
	if res != nil {
		fmt.Println("find: ", res)
		fmt.Println("iter: ", i)
	} else {
		fmt.Println("element does not exist")
		fmt.Println("iter: ", i)
	}
}
