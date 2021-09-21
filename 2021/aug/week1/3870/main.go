package main
import (
	"fmt"
)

func stoneGame(piles []int) bool {
    memo := map[[2]int]int{}
    return score(piles, 0, len(piles) - 1, memo) > 0
}

func score(piles []int, leftIndex int, rightIndex int, memo map[[2]int]int) int {
    if leftIndex == rightIndex {
        return piles[leftIndex]
    }

    if v, ok := memo[[2]int{leftIndex, rightIndex}]; ok {
        return v
    }

    a := piles[leftIndex] - score(piles, leftIndex + 1, rightIndex, memo)
    b := piles[rightIndex] - score(piles, leftIndex, rightIndex - 1, memo)
    result := max(a, b)
    memo[[2]int{leftIndex, rightIndex}] = result
    return result
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}


// Time Limit Exceeded

func ngSolution(piles []int) bool {
	return helper(0, 0, 1, piles)
}

func helper(alex int, lee int, turn int, piles []int) bool {
	// fmt.Printf("alex = %v, lee = %v, turn = %v, piles = %v\n", alex, lee, turn, piles)
	if len(piles) == 0 {
		return alex > lee
	}

	if turn % 2 == 1 {
		return helper(alex + piles[0], lee, turn + 1, piles[1:]) ||
			helper(alex + piles[len(piles) - 1], lee, turn + 1, piles[:len(piles) - 1])
	}

	return helper(alex, lee + piles[len(piles) - 1], turn + 1, piles[1:]) ||
		helper(alex, lee + piles[0], turn + 1, piles[:len(piles) - 1])
}

func main() {
	// result: true
	// piles := []int{5, 3, 4, 5}

	// result: 
	piles := []int{70,92,7,148,182,97,140,195,144,148,39,30,173,71,192,90,103,186,10,198,199,83,44,84,157,173,60,109,144,109,117,5,6,84,140,98,53,60,195,119,66,7,116,11,142,112,35,44,50,137,61,169,130,15,92,153,26,36,158,131,134,24,33,143,94,148,114,199,140,170,154,176,179,155,115,129,151,144,84,167,86,119,4,81,128,16,15,164,194,133,152,158,47,125,93,47,198,32,68,105,25,5,84,171,101,14,139,161,113,54,28,161,158,137,139,170,183,147,56,45,147,181,168,152,34,54,47,134,153,98,187,91,136,185,183,82,104,172,160,9,167,91,69,122,50,1,93,133,77,107,144,45,20,129,68,200,21,59,181,62,5,8,175,175,105,194,161,92,111,123,104,54,107,155,70,4,53,8,82,17,129,73,110,121,119,76,37,75,96,94,17,192,82,77,79,54,54,104,171,162}

	// result: 
	// piles := []int{}

	result := stoneGame(piles)
	fmt.Printf("result = %v\n", result)
}

