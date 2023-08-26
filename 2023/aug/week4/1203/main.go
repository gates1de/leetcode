package main
import (
	"fmt"
)

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {	    
	groupId := m
	for i, id := range group {
		if id == -1 {
			group[i] = groupId
			groupId++
		}
	}

	itemGraph := make(map[int][]int)
	itemIndegree := make([]int, n)
	for i := 0; i < n; i++ {
		itemGraph[i] = make([]int, 0)
	}

	groupGraph := make(map[int][]int)
	groupIndegree := make([]int, groupId)
	for i := 0; i < groupId; i++ {
		groupGraph[i] = make([]int, 0)
	}

	for curr := 0; curr < n; curr++ {
		for _, prev := range beforeItems[curr] {
			itemGraph[prev] = append(itemGraph[prev], curr)
			itemIndegree[curr]++

			if group[curr] != group[prev] {
				groupGraph[group[prev]] = append(groupGraph[group[prev]], group[curr])
				groupIndegree[group[curr]]++
			}
		}
	}

	result := make([]int, 0)
	itemOrder := topologicalSort(itemGraph, itemIndegree);
	groupOrder := topologicalSort(groupGraph, groupIndegree);

	if len(itemOrder) == 0 || len(groupOrder) == 0 {
		return result
	}

	orderedGroups := make(map[int][]int)
	for _, item := range itemOrder {
		if orderedGroups[group[item]] == nil {
			orderedGroups[group[item]] = make([]int, 0)
		}
		orderedGroups[group[item]] = append(orderedGroups[group[item]], item)
	}

	for _, groupIndex := range groupOrder {
		if orderedGroups[groupIndex] == nil {
			continue
		}
		for _, val := range orderedGroups[groupIndex] {
			result = append(result, val)
		}
	}

	return result
}

func topologicalSort(graph map[int][]int, indegree []int) []int {
	visited := make([]int, 0)
	stack := make([]int, 0)

	for key, _ := range graph {
		if indegree[key] == 0 {
			stack = append(stack, key)
		}
	}

	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		visited = append(visited, curr)

		for _, prev := range graph[curr] {
			indegree[prev]--
			if indegree[prev] == 0 {
				stack = append(stack, prev)
			}
		}
	}

	if len(visited) == len(graph) {
		return visited
	}

	return make([]int, 0)
}

func main() {
	// result: [6,3,4,1,5,2,0,7]
	// n := int(8)
	// m := int(2)
	// group := []int{-1,-1,1,0,0,1,0,-1}
	// beforeItems := [][]int{{},{6},{5},{6},{3,6},{},{},{}}

	// result: []
	n := int(8)
	m := int(2)
	group := []int{-1,-1,1,0,0,1,0,-1}
	beforeItems := [][]int{{},{6},{5},{6},{3},{},{4},{}}

	// result: 
	// n := int(0)
	// m := int(0)
	// group := []int{}
	// beforeItems := [][]int{}

	result := sortItems(n, m, group, beforeItems)
	fmt.Printf("result = %v\n", result)
}

