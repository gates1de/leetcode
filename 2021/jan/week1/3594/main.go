package main
import (
	"fmt"
)

func findKthPositive(arr []int, k int) int {
	maxArrLength := 1000
	loopMax := 2000
	missingArr := []int{}
	arrCopy := copySlice(arr)

	for i := 1; i <= loopMax; i++ {
		if len(arrCopy) > 0 {
			arrVal := arrCopy[0]
			if i != arrVal {
				missingArr = append(missingArr, i)
			} else {
				arrCopy = arrCopy[1:]
				continue
			}
		} else {
			missingArr = append(missingArr, i)
		}

		if len(missingArr) == k || len(missingArr) >= maxArrLength {
			return missingArr[k - 1]
		}
	}
	return -1
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func main() {
	// list := []int{2, 3, 4, 7, 11}
	list := []int{1, 2, 3, 4}
	k := 2
	result := findKthPositive(list, k)
	fmt.Printf("result = %v\n", result)
}

