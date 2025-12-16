package main

import (
	"fmt"
)

type Result struct {
	DP0  []int
	DP1  []int
	Size int
}

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n)
	for i := range g {
			g[i] = make([]int, 0)
	}
	for _, e := range hierarchy {
			g[e[0] - 1] = append(g[e[0] - 1], e[1] - 1)
	}
	return dfs(0, present, future, g, budget).DP0[budget]
}

func dfs(u int, present []int, future []int, g [][]int, budget int) Result {
	cost := present[u]
	dCost := present[u] / 2
	dp0 := make([]int, budget + 1)
	dp1 := make([]int, budget + 1)
	subProfit0 := make([]int, budget + 1)
	subProfit1 := make([]int, budget + 1)

	size := cost
	for _, v := range g[u] {
		childResult := dfs(v, present, future, g, budget)
		size += childResult.Size
		for i := budget; i >= 0; i-- {
			for sub := 0; sub <= min(childResult.Size, i); sub++ {
				if i - sub >= 0 {
					subProfit0[i] = max(subProfit0[i], subProfit0[i - sub] + childResult.DP0[sub])
					subProfit1[i] = max(subProfit1[i], subProfit1[i - sub] + childResult.DP1[sub])
				}
			}
		}
	}

	for i := 0; i <= budget; i++ {
		dp0[i] = subProfit0[i]
		dp1[i] = subProfit0[i]
		if i >= dCost {
			dp1[i] = max(subProfit0[i], subProfit1[i - dCost] + future[u] - dCost)
		}
		if i >= cost {
			dp0[i] = max(subProfit0[i], subProfit1[i - cost] + future[u] - cost)
		}
	}

	return Result{DP0: dp0, DP1: dp1, Size: size}
}

func main() {
	// result: 5
	// n := int(2)
	// present := []int{1,2}
	// future := []int{4,3}
	// hierarchy := [][]int{{1,2}}
	// budget := int(3)

	// result: 4
	// n := int(2)
	// present := []int{3,4}
	// future := []int{5,8}
	// hierarchy := [][]int{{1,2}}
	// budget := int(4)

	// result: 10
	n := int(3)
	present := []int{4,6,8}
	future := []int{7,9,11}
	hierarchy := [][]int{{1,2},{1,3}}
	budget := int(10)

	// result: 
	// n := int(0)
	// present := []int{}
	// future := []int{}
	// hierarchy := [][]int{}
	// budget := int(0)

	result := maxProfit(n, present, future, hierarchy, budget)
	fmt.Printf("result = %v\n", result)
}

