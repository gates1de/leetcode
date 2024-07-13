package main
import (
	"fmt"
)

func averageWaitingTime(customers [][]int) float64 {
	nextIdleTime := int(0)
	netWaitTime := int(0)

	for i, _ := range customers {
		nextIdleTime = max(customers[i][0], nextIdleTime) + customers[i][1]
		netWaitTime += nextIdleTime - customers[i][0]
	}

	return float64(netWaitTime) / float64(len(customers))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5.00000
	// customers := [][]int{{1,2},{2,5},{4,3}}

	// result: 3.25000
	customers := [][]int{{5,2},{5,4},{10,3},{20,1}}

	// result: 
	// customers := [][]int{}

	result := averageWaitingTime(customers)
	fmt.Printf("result = %v\n", result)
}

