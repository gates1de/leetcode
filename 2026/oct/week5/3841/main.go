package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	if n == 0 {
		return []bool{}
	}

	adjacency := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjacency[u] = append(adjacency[u], v)
		adjacency[v] = append(adjacency[v], u)
	}

	parent := make([]int, n)
	depth := make([]int, n)
	size := make([]int, n)
	heavy := make([]int, n)
	for i := range heavy {
		heavy[i] = -1
	}
	var findHeavy func(int, int)
	findHeavy = func(node, previous int) {
		parent[node] = previous
		size[node] = 1
		largest := 0
		for _, next := range adjacency[node] {
			if next == previous {
				continue
			}
			depth[next] = depth[node] + 1
			findHeavy(next, node)
			size[node] += size[next]
			if size[next] > largest {
				largest = size[next]
				heavy[node] = next
			}
		}
	}
	findHeavy(0, -1)

	head := make([]int, n)
	position := make([]int, n)
	index := 0
	var decompose func(int, int)
	decompose = func(node, chainHead int) {
		head[node] = chainHead
		position[node] = index
		index++
		if heavy[node] != -1 {
			decompose(heavy[node], chainHead)
		}
		for _, next := range adjacency[node] {
			if next != parent[node] && next != heavy[node] {
				decompose(next, next)
			}
		}
	}
	decompose(0, 0)

	fenwick := make([]int, n+1)
	add := func(i, value int) {
		for i++; i <= n; i += i & -i {
			fenwick[i] ^= value
		}
	}
	prefix := func(i int) int {
		result := 0
		for ; i > 0; i -= i & -i {
			result ^= fenwick[i]
		}
		return result
	}
	rangeXor := func(left, right int) int {
		return prefix(right+1) ^ prefix(left)
	}

	characters := []byte(s)
	for node, character := range characters {
		add(position[node], 1<<(character-'a'))
	}

	pathXor := func(u, v int) int {
		result := 0
		for head[u] != head[v] {
			if depth[head[u]] < depth[head[v]] {
				u, v = v, u
			}
			result ^= rangeXor(position[head[u]], position[u])
			u = parent[head[u]]
		}
		if depth[u] > depth[v] {
			u, v = v, u
		}
		return result ^ rangeXor(position[u], position[v])
	}

	answer := make([]bool, 0)
	for _, query := range queries {
		parts := strings.Fields(query)
		if parts[0] == "update" {
			node, _ := strconv.Atoi(parts[1])
			next := parts[2][0]
			add(position[node], 1<<(characters[node]-'a')^1<<(next-'a'))
			characters[node] = next
			continue
		}
		u, _ := strconv.Atoi(parts[1])
		v, _ := strconv.Atoi(parts[2])
		answer = append(answer, bits.OnesCount(uint(pathXor(u, v))) <= 1)
	}

	return answer
}

func main() {
	// result: [true, false]
	// n := 3
	// edges := [][]int{{0, 1}, {1, 2}}
	// s := "aac"
	// queries := []string{"query 0 2", "update 1 b", "query 0 2"}

	// result: [false, false, true]
	n := 4
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}}
	s := "abca"
	queries := []string{"query 1 2", "update 0 b", "query 2 3", "update 3 a", "query 1 3"}

	// result: []
	// n := 0
	// edges := [][]int{}
	// s := ""
	// queries := []string{}

	result := palindromePath(n, edges, s, queries)
	fmt.Printf("result = %v\n", result)
}
