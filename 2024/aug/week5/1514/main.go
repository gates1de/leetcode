package main

import (
	"fmt"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	maxProb := make([]float64, n)
	maxProb[start] = 1.0
	for i := 0; i < n - 1; i++ {
		hasUpdate := false
		for j := 0; j < len(edges); j++ {
			u := edges[j][0]
			v := edges[j][1]

			pathProb := succProb[j]
			if maxProb[u]*pathProb > maxProb[v] {
				maxProb[v] = maxProb[u] * pathProb
				hasUpdate = true
			}
			if maxProb[v]*pathProb > maxProb[u] {
				maxProb[u] = maxProb[v] * pathProb
				hasUpdate = true
			}
		}

		if !hasUpdate {
			break
		}
	}

	return maxProb[end]
}

func main() {
	// result: 0.2500
	// n := int(3)
	// edges := [][]int{{0,1},{1,2},{0,2}}
	// succProb := []float64{0.5,0.5,0.2}
	// start := int(0)
	// end := int(2)

	// result: 0.3000
	// n := int(3)
	// edges := [][]int{{0,1},{1,2},{0,2}}
	// succProb := []float64{0.5,0.5,0.3}
	// start := int(0)
	// end := int(2)

	// result: 0.0000
	n := int(3)
	edges := [][]int{{0, 1}}
	succProb := []float64{0.5}
	start := int(0)
	end := int(2)

	// result:
	// n := int(0)
	// edges := [][]int{}
	// succProb := []float64{}
	// start := int(0)
	// end := int(0)

	result := maxProbability(n, edges, succProb, start, end)
	fmt.Printf("result = %v\n", result)
}
