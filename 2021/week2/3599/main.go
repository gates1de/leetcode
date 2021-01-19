package main
import (
	"fmt"
)

func createSortedArray(instructions []int) int {
	sortedList := make([]int, 0, len(instructions))
	cost := int(0)
	for _, val := range instructions {
		// fmt.Printf("val = %v, sortedList = %v\n", val, sortedList)
		sortedListLength := len(sortedList)

		if sortedListLength == 0 {
			sortedList = append(sortedList, val)
			continue
		}

		if sortedList[sortedListLength - 1] <= val {
			sortedList = append(sortedList, val)
			continue
		}

		for j, v := range sortedList {
			if v < val {
				continue
			}

			newList := make([]int, 0, sortedListLength + 1)
			newList = append(copySlice(sortedList[:j]), val)
			lessCost := j
			greaterCost := 0
			if j + 1 <= sortedListLength {
				newList = append(newList, sortedList[j:]...)
				sameValCount := count(sortedList[j:], val)
				greaterCost = sortedListLength - j - sameValCount
			}
			cost += min(lessCost, greaterCost)
			sortedList = newList
			// fmt.Printf("val %v: cost = min(%v, %v), total cost = %v\n", val, lessCost, greaterCost, cost)
			break
		}
	}

	return cost
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func count(list []int, target int) int {
	result := int(0)
	for _, v := range list {
		if target == v {
			result++
		}
	}
	return result
}

func copySlice(target []int) []int {
	result := make([]int, len(target))
	copy(result, target)
	return result
}

func main() {
	// instructions := []int{1, 5, 6, 2}
	// instructions := []int{1, 2, 3, 6, 5, 4}
	instructions := []int{1, 3, 3, 3, 2, 4, 2, 1, 2}
	result := createSortedArray(instructions)
	fmt.Printf("result = %v\n", result)
}

