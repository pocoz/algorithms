package main

import (
	"fmt"
	"time"

	"github.com/pocoz/algorithms/search"
)

var begin time.Time

func main() {
	// wide search
	begin = time.Now()

	graph := make(map[string][]string, 0)
	graph["me"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{"bob"}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	queue := graph["me"]
	fmt.Println(search.Wide(queue, graph))
	fmt.Println(time.Since(begin))

	//// sort quick
	//begin = time.Now()
	//arr := []int{34, 6, 13, 22, 2, 54, 67, 333, 188}
	//fmt.Println(sorts.Quick(arr))
	//fmt.Println(time.Since(begin))
	//
	//// sort slow
	//arr = []int{34, 6, 13, 22, 2, 54, 67, 333, 188}
	//begin = time.Now()
	//fmt.Println(sorts.Slow(arr))
	//fmt.Println(time.Since(begin))
	//
	//// binary search
	//arr = []int{1, 4, 6, 13, 22, 44, 54, 67, 73, 188}
	//item := 34
	//res, i := search.Binary(item, arr)
	//if res != nil {
	//	fmt.Println("find: ", res)
	//	fmt.Println("iter: ", i)
	//} else {
	//	fmt.Println("element does not exist")
	//	fmt.Println("iter: ", i)
	//}
}
