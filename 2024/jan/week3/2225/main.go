package main
import (
	"fmt"
)

func findWinners(matches [][]int) [][]int {
	losses := make([]int, 100001)
	for _, match := range matches {
		win := match[0]
		loss := match[1]

		if losses[win] == 0 {
			losses[win] = -1
		}

		if losses[loss] == -1 {
			losses[loss] = 1
		} else {
			losses[loss]++
		}
	}

	zeroLoss := make([]int, 0, 100001)
	oneLoss := make([]int, 0, 100001)

	result := make([][]int, 2)

	for i, loss := range losses {
		if loss == -1 {
			zeroLoss = append(zeroLoss, i)
		} else if loss == 1 {
			oneLoss = append(oneLoss, i)
		}
	}

	result[0] = zeroLoss
	result[1] = oneLoss
	return result
}

func main() {
	// result: [[1,2,10],[4,5,7,8]]
	// matches := [][]int{{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9}}

	// result: [[1,2,5,6],[]]
	matches := [][]int{{2,3},{1,3},{5,4},{6,4}}

	// result: 
	// matches := [][]int{}

	result := findWinners(matches)
	fmt.Printf("result = %v\n", result)
}

