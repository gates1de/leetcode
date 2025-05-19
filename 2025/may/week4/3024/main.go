package main
import (
	"fmt"
	"sort"
)

func triangleType(nums []int) string {
	sort.Ints(nums)

	if nums[0] + nums[1] <= nums[2] {
		return "none"
	} else if nums[0] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[1] == nums[2] {
		return "isosceles"
	}

	return "scalene"
}

func main() {
	// result: "equilateral"
	// nums := []int{3,3,3}

	// result: "scalene"
	nums := []int{3,4,5}

	// result: ""
	// nums := []int{}

	result := triangleType(nums)
	fmt.Printf("result = %v\n", result)
}
