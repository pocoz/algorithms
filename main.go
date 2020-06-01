package main

import (
	"fmt"
	"github.com/pocoz/algorithms/core"
	"time"
)

var begin time.Time

func main() {
	fmt.Println(core.Find(2166,6099))

	//// dijkstra search
	//var (
	//	graph = map[string]map[string]int{
	//		"start": {"a": 6, "b": 2},
	//		"a":     {"fin": 1},
	//		"b":     {"a": 3, "fin": 5},
	//	}
	//	costs = map[string]int{
	//		"a":   6,
	//		"b":   2,
	//		"fin": math.MaxInt32,
	//	}
	//	parents = map[string]string{
	//		"a":   "start",
	//		"b":   "start",
	//		"fin": "",
	//	}
	//)
	//fmt.Println(search.Dijkstra(graph, costs, parents))

	//// wide search
	//begin = time.Now()
	//graph := make(map[string][]string, 0)
	//graph["me"] = []string{"alice", "bob", "claire"}
	//graph["bob"] = []string{"anuj", "peggy"}
	//graph["alice"] = []string{"peggy"}
	//graph["claire"] = []string{"thom", "jonny"}
	//graph["anuj"] = []string{"bob"}
	//graph["peggy"] = []string{}
	//graph["thom"] = []string{}
	//graph["jonny"] = []string{}
	//queue := graph["me"]
	//fmt.Println(search.Wide(queue, graph))
	//fmt.Println(time.Since(begin))

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
