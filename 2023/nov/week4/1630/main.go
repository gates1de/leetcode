package main
import (
	"fmt"
	"math"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))
	for i, left := range l {
		right := r[i]
		arr := make([]int, right - left + 1)
		for j, _ := range arr {
			arr[j] = nums[left + j]
		}

		result[i] = check(arr)
	}

	return result
}

func check(nums []int) bool {
	n := len(nums)
	minElement := math.MaxInt32
	maxElement := math.MinInt32
	set := make(map[int]bool)

	for _, num := range nums {
		minElement = min(minElement, num)
		maxElement = max(maxElement, num)
		set[num] = true
	}

	if (maxElement - minElement) % (n - 1) != 0 {
		return false
	}

	diff := (maxElement - minElement) / (n - 1)
	curr := minElement + diff

	for curr < maxElement {
		if exists, _ := set[curr]; !exists {
			return false
		}
		curr += diff
	}

	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [true,false,true]
	// nums := []int{4,6,5,9,3,7}
	// l := []int{0,0,2}
	// r := []int{2,3,5}


	// result: [false,true,false,false,true,true]
	nums := []int{-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10}
	l := []int{0,1,6,4,8,7}
	r := []int{4,4,9,7,9,10}

	// result: []
	// nums := []int{}
	// l := []int{}
	// r := []int{}

	result := checkArithmeticSubarrays(nums, l, r)
	fmt.Printf("result = %v\n", result)
}

