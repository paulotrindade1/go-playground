package main

import "fmt"

func SearhIslands(graph [][]int) int {
	islands := 0
	visited := map[int]map[int]bool{}

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			if visited[i] == nil {
				visited[i] = map[int]bool{}
			}

			if ok := visited[i][j]; !ok && graph[i][j] == 1 {
				dfs(graph, visited, i, j)
				islands++
			}
		}
	}

	return islands
}

func dfs(graph [][]int, visited map[int]map[int]bool, i, j int) {
	if visited[i] == nil {
		visited[i] = map[int]bool{}
	}

	if ok := visited[i][j]; ok {
		return
	}

	if i < 0 || j < 0 || i >= len(graph) || j >= len(graph[0]) || graph[i][j] == 0 {
		return
	}

	visited[i][j] = true

	dfs(graph, visited, i+1, j)
	dfs(graph, visited, i-1, j)
	dfs(graph, visited, i, j+1)
	dfs(graph, visited, i, j-1)
}

func main() {
	graph := [][]int{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
	} // 2 islands

	fmt.Println(SearhIslands(graph))
}
