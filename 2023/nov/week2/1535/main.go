package main
import (
	"fmt"
)

func getWinner(arr []int, k int) int {
	maxNum := arr[0]
	for _, num := range arr[1:] {
		maxNum = max(maxNum, num)
	}

	current := arr[0]
	streak := int(0)

	for _, opponent := range arr[1:] {
		if current > opponent {
			streak++
		} else {
			current = opponent
			streak = 1
		}

		if streak == k || current == maxNum {
			return current
		}
	}

	return -1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// arr := []int{2,1,3,5,4,6,7}
	// k := int(2)

	// result: 3
	arr := []int{3,2,1}
	k := int(10)

	// result: 
	// arr := []int{}
	// k := int(0)

	result := getWinner(arr, k)
	fmt.Printf("result = %v\n", result)
}

