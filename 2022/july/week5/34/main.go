package main
import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	targets := []int{}
	binarySearch(0, n - 1, nums, target, &targets)
	sort.Ints(targets)

	if len(targets) == 0 {
		return []int{-1, -1}
	}
	return []int{targets[0], targets[len(targets) - 1]}
}

func binarySearch(i int, j int, nums []int, target int, targets *[]int) {
	n := len(nums)
	if i >= n || j >= n || i > j {
		return
	}

	if nums[i] == target {
		*targets = append(*targets, i)
	}
	if nums[j] == target {
		*targets = append(*targets, j)
	}
	if i == j || i + 1 == j {
		return
	}

	midIndex := (i + j) / 2
	mid := nums[midIndex]

	if mid == target {
		binarySearch(midIndex, j, nums, target, targets)
		binarySearch(i, midIndex, nums, target, targets)
	} else if mid < target {
		binarySearch(midIndex + 1, j, nums, target, targets)
	} else {
		binarySearch(i, midIndex - 1, nums, target, targets)
	}
}

func main() {
	// result: [3,4]
	// nums := []int{5,7,7,8,8,10}
	// target := int(8)

	// result: [-1,-1]
	// nums := []int{5,7,7,8,8,10}
	// target := int(6)

	// result: [-1,-1]
	// nums := []int{}
	// target := int(0)

	// result: [0,0]
	// nums := []int{5,7,7,8,8,10}
	// target := int(5)

	// result: [1,6]
	nums := []int{5,7,7,7,7,7,7,8,8,10}
	target := int(7)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := searchRange(nums, target)
	fmt.Printf("result = %v\n", result)
}

