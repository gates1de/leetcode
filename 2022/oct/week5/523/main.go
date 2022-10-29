package main
import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	hashMap := map[int]int{0: 0}
	sum := int(0)
	for i, num := range nums {
        sum += num
		if val, ok := hashMap[sum % k]; ok {
			if val < i {
				return true
			}
		} else {
			hashMap[sum % k] = i + 1
		}
	}
	return false
}

func main() {
	// result: true
	// nums := []int{23,2,4,6,7}
	// k := int(6)

	// result: true
	// nums := []int{23,2,6,4,7}
	// k := int(6)

	// result: false
	// nums := []int{23,2,6,4,7}
	// k := int(13)

	// result: true
	nums := []int{2,4,3}
	k := int(6)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := checkSubarraySum(nums, k)
	fmt.Printf("result = %v\n", result)
}

