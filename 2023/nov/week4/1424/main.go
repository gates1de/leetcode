package main
import (
	"fmt"
)

func findDiagonalOrder(nums [][]int) []int {
	n := len(nums)
	queue := make([][]int, 0, n)
	queue = append(queue, []int{0, 0})
	temp := make([]int, 0)

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		row := pair[0]
		column := pair[1]

		temp = append(temp, nums[row][column])

		if column == 0 && row + 1 < n {
			queue = append(queue, []int{row + 1, column})
		}
		if column + 1 < len(nums[row]) {
			queue = append(queue, []int{row, column + 1})
		}
	}

	result := make([]int, len(temp))
	for i, num := range temp {
		result[i] = num
	}

	return result
}

func main() {
	// result: {1,4,2,7,5,3,8,6,9}
	// nums := [][]int{{1,2,3},{4,5,6},{7,8,9}}

	// result: {1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16}
	nums := [][]int{{1,2,3,4,5},{6,7},{8},{9,10,11},{12,13,14,15,16}}

	// result: 
	// nums := [][]int{}

	result := findDiagonalOrder(nums)
	fmt.Printf("result = %v\n", result)
}

