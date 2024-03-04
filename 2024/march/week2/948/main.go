package main
import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    left := int(0)
    right := len(tokens) - 1
    currentPoint := int(0)
    result := int(0)

    for left <= right && (power >= tokens[left] || currentPoint > 0) {
        for left <= right && power >= tokens[left] {
            power -= tokens[left]
            left++
            currentPoint++
        }

        result = max(result, currentPoint)

        if left <= right && currentPoint > 0 {
            power += tokens[right]
            right--
            currentPoint--
        }
    }

    return result
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func main() {
	// result: 0
	// tokens := []int{100}
	// power := int(50)

	// result: 1
	// tokens := []int{200,100}
	// power := int(150)

	// result: 2
	tokens := []int{100,200,300,400}
	power := int(200)

	// result: 
	// tokens := []int{}
	// power := int(0)

	result := bagOfTokensScore(tokens, power)
	fmt.Printf("result = %v\n", result)
}

