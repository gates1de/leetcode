package main
import (
	"fmt"
	"math"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
    n := len(costs)
    sort.Slice(costs, func(i, j int) bool {
        return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
    })

    sum := int(0)
    for l, r := 0, n - 1; l < r; l, r = l + 1, r - 1 {
        sum += costs[l][0] + costs[r][1]
    }

    return sum
}

// Slow & Use more memory
func mySolution(costs [][]int) int {
	result := math.MaxInt32
	m := map[[2]int]int{}
	helper(0, 0, 0, len(costs), costs, m, &result)
	return result
}

func helper(cost int, aCount int, bCount int, n int, costs [][]int, m map[[2]int]int, result *int) {
	if len(costs) == 0 {
		if cost < *result {
			*result = cost
		}
		return
	}

	if m[[2]int{aCount, bCount}] > 0 && m[[2]int{aCount, bCount}] <= cost {
		return
	}

	m[[2]int{aCount, bCount}] = cost

	if aCount < n / 2 {
		helper(cost + costs[0][0], aCount + 1, bCount, n, costs[1:], m, result)
	}
	if bCount < n / 2 {
		helper(cost + costs[0][1], aCount, bCount + 1, n, costs[1:], m, result)
	}
}

func main() {
	// result: 110
	// costs := [][]int{{10,20},{30,200},{400,50},{30,20}}

	// result: 1859
	// costs := [][]int{{259,770},{448,54},{926,667},{184,139},{840,118},{577,469}}

	// result: 3086
	// costs := [][]int{{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42}}

	// result: 2
	// costs := [][]int{{1,1},{1,2}}

	// result: 37740
	costs := [][]int{{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512},{515,563},{451,713},{537,709},{343,819},{855,779},{457,60},{650,359},{631,42},{333,100},{1000,512}}

	// result: 
	// costs := [][]int{}

	result := twoCitySchedCost(costs)
	fmt.Printf("result = %v\n", result)
}

