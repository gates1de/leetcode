package main
import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
    for i := len(cost) - 3; i >= 0; i-- {
        oneStepCost := cost[i + 1]
        twoStepCost := cost[i + 2]
        if oneStepCost < twoStepCost {
            cost[i] += oneStepCost
        } else {
            cost[i] += twoStepCost
        }
    }
    
    if cost[0] < cost[1] {
        return cost[0]
    }
    return cost[1]
}

func main() {
	// result: 15
	// cost := []int{10,15,20}

	// result: 6
	// cost := []int{1,100,1,1,1,100,1,1,100,1}

	// result: 0
	cost := []int{1, 0}

	// result: 
	// cost := []int{}

	result := minCostClimbingStairs(cost)
	fmt.Printf("result = %v\n", result)
}

