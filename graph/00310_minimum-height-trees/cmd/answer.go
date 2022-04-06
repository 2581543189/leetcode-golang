package main

import "fmt"

func main() {
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
}

func findMinHeightTrees(n int, edges [][]int) (nodes []int) {
	in, connect := make([]int, n), map[int][]int{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		in[a]++
		in[b]++
		connect[a] = append(connect[a], b)
		connect[b] = append(connect[b], a)
	}
	for i := 0; i < n; i++ {
		if in[i] < 2 {
			nodes = append(nodes, i)
		}
	}
	for n > 2 {
		s := len(nodes)
		n -= s
		for _, node := range nodes {
			for _, other := range connect[node] {
				in[other]--
				if in[other] == 1 {
					nodes = append(nodes, other)
				}
			}
		}
		nodes = nodes[s:]
	}
	return nodes
}
