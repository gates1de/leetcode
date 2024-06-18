package main
import (
	"fmt"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	maxAbility := int(0)
	for _, ability := range worker {
		maxAbility = max(maxAbility, ability)
	}

	jobs := make([]int, maxAbility + 1)
	for i := 0; i < len(difficulty); i++ {
		if difficulty[i] <= maxAbility {
			jobs[difficulty[i]] = max(jobs[difficulty[i]], profit[i])
		}
	}

	for i := 1; i <= maxAbility; i++ {
		jobs[i] = max(jobs[i], jobs[i - 1])
	}

	result := int(0)
	for _, ability := range worker {
		result += jobs[ability]
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 100
	// difficulty := []int{2,4,6,8,10}
	// profit := []int{10,20,30,40,50}
	// worker := []int{4,5,6,7}

	// result: 0
	difficulty := []int{85,47,57}
	profit := []int{24,66,99}
	worker := []int{40,25,25}

	// result: 
	// difficulty := []int{}
	// profit := []int{}
	// worker := []int{}

	result := maxProfitAssignment(difficulty, profit, worker)
	fmt.Printf("result = %v\n", result)
}

