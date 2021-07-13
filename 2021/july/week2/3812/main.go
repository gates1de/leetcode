package main
import (
	"fmt"
	// "math"
	"math/rand"
	"time"
)

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	return binarySearch(0, len(nums) - 1, nums)
}

func binarySearch(left int, right int, nums []int) int {
	mid := left + (right - left) / 2
	if isOK(mid, nums) {
		return mid
	}
	if left >= right {
		return -1
	}
	result := binarySearch(left, mid - 1, nums)
	if result == -1 {
		result = binarySearch(mid + 1, right, nums)
	}
	return result
}

func isOK(i int, nums []int) bool {
	if i < 0 || len(nums) - 1 < i {
		return false
	}
	if i == 0 {
		return nums[i] > nums[i + 1]
	} else if i == len(nums) - 1 {
		return nums[i] > nums[i - 1]
	}
	return nums[i] > nums[i - 1] && nums[i] > nums[i + 1]
}

func generateNums(length int, min int, max int) []int {
	result := make([]int, length)
    rand.Seed(time.Now().UnixNano())
	for i, _ := range result {
        result[i] = rand.Intn(max - min) + min
    }
	return result
}

func printNums(nums []int) {
	fmt.Printf("[")
	comma := ""
	for _, num := range nums {
		fmt.Printf("%v%v", comma, num)
		comma = ","
	}
	fmt.Printf("]\n")
}

func main() {
	// result: 2
	// nums := []int{1, 2, 3, 1}

	// result: 1 or 5
	// nums := []int{1, 2, 1, 3, 5, 6, 4}

	// result: 3
	nums := []int{1, 2, 3, 4}

	// result: 
	// nums := []int{}

	// nums := generateNums(1000, int(-math.Pow(2.0, 31)), int(math.Pow(2.0, 31)) + 1)

	printNums(nums)
	result := findPeakElement(nums)
	fmt.Printf("result = %v\n", result)
}

