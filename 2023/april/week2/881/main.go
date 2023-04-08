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
	// people := []int{1,2}
	// limit := int(3)

	// result: 3
	// people := []int{3,2,2,1}
	// limit := int(3)

	// result: 4
	people := []int{3,5,3,4}
	limit := int(5)

	// result: 
	// people := []int{}
	// limit := int(0)

	result := numRescueBoats(people, limit)
	fmt.Printf("result = %v\n", result)
}

