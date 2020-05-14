package search

var searched = make(map[string]bool)

func Wide(queue []string, graph map[string][]string) string {
	if len(queue) == 0 {
		return "not found"
	}
	for _, q := range queue {
		_, ok := searched[q]
		if ok {
			queue = queue[1:]
			continue
		}
		if isSeller(q) {
			return q
		}
		queue = queue[1:]
		queue = append(queue, graph[q]...)
		searched[q] = true
	}
	return Wide(queue, graph)
}

func isSeller(name string) bool {
	if name == "thom" {
		return true
	}
	return false
}
