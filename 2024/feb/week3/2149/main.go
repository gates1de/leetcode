package main
import (
	"fmt"
)

func rearrangeArray(nums []int) []int {
	n := len(nums)
	negatives := make([]int, n / 2)
	positives := make([]int, n / 2)
	i := int(0)
	j := int(0)

	for _, num := range nums {
		if num < 0 {
			negatives[i] = num
			i++
		} else {
			positives[j] = num
			j++
		}
	}

	result := make([]int, n)
	i = 0
	j = 0
	for index, _ := range result {
		if index % 2 == 1 {
			result[index] = negatives[i]
			i++
		} else {
			result[index] = positives[j]
			j++
		}
	}

	return result
}

func main() {
	// result: [3,-2,1,-5,2,-4]
	nums := []int{3,1,-2,-5,2,-4}

	// result: [1,-1]
	// nums := []int{-1,1}

	// result: 
	// nums := []int{}

	result := rearrangeArray(nums)
	fmt.Printf("result = %v\n", result)
}

