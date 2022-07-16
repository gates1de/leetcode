package main
import (
	"fmt"
	"sort"
)


func makesquare(matchsticks []int) bool {
    length := int(0)
    sort.Slice(matchsticks, func (a, b int) bool { return matchsticks[a] > matchsticks[b] })
    for _, num := range matchsticks {
        length += num
    }

    if length % 4 != 0 {
        return false
    }

    oneLen := length / 4
    square := [4]int{0, 0, 0, 0}

    return helper(oneLen, square, matchsticks)
}

func helper(oneLen int, square [4]int, sticks []int) bool {
    if len(sticks) == 0 {
        for _, side := range square {
            if side != oneLen {
                return false
            }
        }
        return true
    }

    for i, side := range square {
        if side + sticks[0] > oneLen {
            continue
        }
        newSquare := square
        newSquare[i] += sticks[0]
        if helper(oneLen, newSquare, sticks[1:]) {
            return true
        }
    }
    return false
}

func main() {
	// result: true
	// matchsticks := []int{1,1,2,2,2}

	// result: false
	// matchsticks := []int{3,3,3,3,4}

	// result: true
	// matchsticks := []int{1,2,4,7,2,8,8}

	// result: true
	// matchsticks := []int{5,5,5,5,4,4,4,4,3,3,3,3}

	// result: false
	 matchsticks := []int{2,2,2,2,2,6}

	// result: 
	// matchsticks := []int{}

	result := makesquare(matchsticks)
	fmt.Printf("result = %v\n", result)
}

