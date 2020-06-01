package search

import (
	"fmt"
	"reflect"
)

var dijkstraProcessed = make(map[string]bool, 0)

// Dijkstra dont work!!!
func Dijkstra(graph map[string]map[string]int, costs map[string]int, parents map[string]string) map[string]int {
	node := findLowestCostNode(costs)
	fmt.Println("node", node)
	if node != "" {
		fmt.Println("costs", costs)
		cost := costs[node]
		neighbors := graph[node]
		for k := range neighbors {
			newCost := cost + neighbors[k]
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = node
				dijkstraProcessed[node] = true
			}
		}
		fmt.Println("parents", parents)
	} else {
		return costs
	}
	return Dijkstra(graph, costs, parents)
}

func findLowestCostNode(costs map[string]int) string {
	if len(costs) == 0 {
		return ""
	}
	keys := reflect.ValueOf(costs).MapKeys()

	var lowestValue = costs[keys[0].String()]
	var lowestKey = keys[0].String()
	for k, v := range costs {
		_, ok := dijkstraProcessed[k]
		if v < lowestValue && !ok {
			lowestKey = k
		}
	}
	return lowestKey
}
