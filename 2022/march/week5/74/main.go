package main
import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	for _, nums := range matrix {
		if target > nums[n - 1] {
			continue
		}

		for _, num := range nums {
			if num == target {
				return true
			} else if target < num {
				return false
			}
		}
		return false
	}

	return false
}

func main() {
	// result: true
	// matrix := [][]int{
	// 	{1,3,5,7},
	// 	{10,11,16,20},
	// 	{23,30,34,60},
	// }
	// target := int(3)

	// result: false
	// matrix := [][]int{
	// 	{1,3,5,7},
	// 	{10,11,16,20},
	// 	{23,30,34,60},
	// }
	// target := int(13)

	// result: true
	matrix := [][]int{{1}}
	target := int(1)

	// result: 
	// matrix := [][]int{}
	// target := int(0)

	result := searchMatrix(matrix, target)
	fmt.Printf("result = %v\n", result)
}

