package main
import (
	"fmt"
)

func singleNonDuplicate(nums []int) int {
	n := len(nums)
    if n == 0 {
        return -1
    } else if n == 1 {
        return nums[0]
    }

    mid := (n - 1) / 2
    if nums[mid] != nums[mid + 1] {
        mid++
    }

    if len(nums[:mid]) % 2 == 0 {
        return singleNonDuplicate(nums[mid:])
    }
    return singleNonDuplicate(nums[:mid])
}

func main() {
	// result: 2
	// nums := []int{1,1,2,3,3,4,4,8,8}

	// result: 10
	nums := []int{3,3,7,7,10,11,11}

	// result: 
	// nums := []int{}

	result := singleNonDuplicate(nums)
	fmt.Printf("result = %v\n", result)
}

