package main
import (
	"fmt"
)

func validMountainArray(arr []int) bool {
	if len(arr) <= 1 || arr[0] >= arr[1] {
		return false
	}

	pre := arr[0]
	isDec := false
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		if pre == num {
			return false
		} else if pre > num {
			isDec = true
			pre = num
		} else {
			if isDec {
				return false
			}
		}
		pre = num
	}

	return isDec
}

func main() {
	// result: false
	// arr := []int{2, 1}

	// result: false
	// arr := []int{3, 5, 5}

	// result: true
	// arr := []int{0, 3, 2, 1}

	// result: true
	// arr := []int{0, 2, 3, 4, 5, 2, 1, 0}

	// result: false
	// arr := []int{0, 2, 3, 3, 5, 2, 1, 0}

	// result: false
	// arr := []int{0, 1, 2}

	// result: false
	arr := []int{1}

	// result: 
	// arr := []int{}

	result := validMountainArray(arr)
	fmt.Printf("result = %v\n", result)
}

