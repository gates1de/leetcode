package main
import (
	"fmt"
	"math"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	result := math.MaxInt32
	helper(1, limit, people, limit, &result)
	return result
}

func helper(current int, remain int, people []int, limit int, result *int) {
	if len(people) == 0 {
		if current < *result {
			*result = current
		}
		return
	}

	if people[0] <= remain {
		index := sort.SearchInts(people, remain)
		// fmt.Println(people, index)

		if len(people) <= index {
			target := people[index - 1]
			if remain < limit {
				remain = 0
			} else {
				remain -= target
			}
			helper(current, remain, remove(people, index), limit, result)
		} else {
			if people[index] == remain {
				helper(current, 0, remove(people, index), limit, result)
			} else {
				if remain < limit {
					remain = 0
				} else {
					remain -= people[index - 1]
				}
				helper(current, remain, remove(people, index - 1), limit, result)
			}
		}
	} else {
		helper(current + 1, limit - people[0], people[1:], limit, result)
	}
}

func remove(arr []int, index int) []int {
	if len(arr) == index {
		return arr[:index - 1]
	}
	newArr := arr[:index]
	newArr = append(newArr, arr[index + 1:]...)
	return newArr
}

func main() {
	// result: 1
	// people := []int{1, 2}
	// limit := int(3)

	// result: 3
	// people := []int{3, 2, 2, 1}
	// limit := int(3)

	// result: 4
	// people := []int{3, 5, 3, 4}
	// limit := int(5)

	// result: 3
	// people := []int{21, 40, 16, 24, 30}
	// limit := int(50)

	// result: 11
	// people := []int{3,8,4,9,2,2,7,1,6,10,6,7,1,7,7,6,4,4,10,1}
	// limit := int(10)

	// result: 3
	people := []int{3, 2, 3, 2, 2}
	limit := int(6)

	// result: 
	// people := []int{}
	// limit := int(0)

	result := numRescueBoats(people, limit)
	fmt.Printf("result = %v\n", result)
}

