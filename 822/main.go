package main

import (
	"fmt"
)

func validArrangement(pairs [][]int) [][]int {
	graph := make(map[int][]int)
	degrees := make(map[int]int)
	for _, pair := range pairs {
		graph[pair[0]] = []int{}
		graph[pair[1]] = []int{}
		degrees[pair[0]] = 0
		degrees[pair[1]] = 0
	}
	for _, pair := range pairs {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		degrees[pair[0]]++
		degrees[pair[1]]--
	}
	from := pairs[0][0]
	for cur, degree := range degrees {
		if degree == 1 {
			from = cur
			break
		}
	}
	record := [][]int{}
	dfs(from, graph, &record)
	n := len(record)
	ans := make([][]int, n)
	for i, j := n-1, 0; j < n; i, j = i-1, j+1 {
		ans[i] = record[j]
	}
	return ans
}

func dfs(from int, graph map[int][]int, record *[][]int) {

	next := graph[from]

	for len(next) > 0 {

		to := next[0]

		next = next[1:]

		dfs(to, graph, record)

		*record = append(*record, []int{from, to})

	}

}

func main() {

	pairs := [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}

	result := validArrangement(pairs)

	fmt.Println(result)

}
