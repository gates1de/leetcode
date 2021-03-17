package main
import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		headNum := nums[0]
		tailNum := nums[len(nums) - 1]
		// fmt.Printf("%v: head = %v, tail = %v\n", i, headNum, tailNum)
		if headNum > target {
			break
		}
		if tailNum < target {
			continue
		}
		if headNum <= target && target <= tailNum {
			for _, num := range nums {
				if num == target {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := int(5)

	// matrix := [][]int{
	// 	{2, 8, 14, 22, 30},
	// 	{4, 10, 16, 24, 38},
	// 	{6, 12, 18, 32, 44},
	// 	{20, 26, 28, 34, 48},
	// 	{36, 42, 46, 52, 60},
	// }
	// target := int(21)

	result := searchMatrix(matrix, target)
	fmt.Printf("result = %v\n", result)
}

