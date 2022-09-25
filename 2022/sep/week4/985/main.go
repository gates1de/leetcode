package main
import (
	"fmt"
)

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	evenSum := int(0)
	for _, num := range nums {
		if num % 2 != 0 {
			continue
		}
		evenSum += num
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		add := query[0]
		index := query[1]
		num := nums[index]
		newNum := num + add

		if num % 2 == 0 {
			evenSum -= num
		}
		if newNum % 2 == 0 {
			evenSum += newNum
		}

		nums[index] = newNum
		result[i] = evenSum
	}

	return result
}

func main() {
	// result: [8,6,2,4]
	// nums := []int{1,2,3,4}
	// queries := [][]int{{1,0},{-3,1},{-4,0},{2,3}}

	// result: [0]
	nums := []int{1}
	queries := [][]int{{4,0}}

	// result: 
	// nums := []int{}
	// queries := [][]int{}

	result := sumEvenAfterQueries(nums, queries)
	fmt.Printf("result = %v\n", result)
}

