package main
import (
	"fmt"
)

func subarraysDivByK(nums []int, k int) int {
	prefixMod := int(0)
	result := int(0)
	modGroups := make([]int, k)
	modGroups[0] = 1

	for _, num := range nums {
		prefixMod = (prefixMod + num % k + k) % k
		result += modGroups[prefixMod]
		modGroups[prefixMod]++
	}
	return result
}

func main() {
	// result: 7
	// nums := []int{4,5,0,-2,-3,1}
	// k := int(5)

	// result: 0
	nums := []int{5}
	k := int(9)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := subarraysDivByK(nums, k)
	fmt.Printf("result = %v\n", result)
}

