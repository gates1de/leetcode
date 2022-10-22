package main
import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	indexes := map[int]int{}
	for i, num := range nums {
		if index, ok := indexes[num]; ok && i - index <= k {
			return true
		}
		indexes[num] = i
	}
	return false
}

func main() {
	// result: true
	// nums := []int{1,2,3,1}
	// k := int(3)

	// result: true
	// nums := []int{1,0,1,1}
	// k := int(1)

	// result: false
	nums := []int{1,2,3,1,2,3}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := containsNearbyDuplicate(nums, k)
	fmt.Printf("result = %v\n", result)
}

