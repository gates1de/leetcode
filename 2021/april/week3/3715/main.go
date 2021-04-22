package main
import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }
    if len(triangle) == 1 {
        return triangle[0][0]
    }
    for i := len(triangle) - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            min := triangle[i + 1][j]
            if triangle[i + 1][j + 1] < min {
                min = triangle[i + 1][j + 1]
            }
            triangle[i][j] += min
        }
    }
    return triangle[0][0]
}

// Out of memory
func ngSolution(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	sumList := []int{}
	dfs(triangle, 1, 0, triangle[0][0], &sumList)
	return min(sumList)
}

func dfs(triangle [][]int, i int, j int, currentSum int, sumList *[]int) {
	num := triangle[i][j]
	nextNum := triangle[i][j + 1]

	if i == len(triangle) - 1 {
		if num <= nextNum {
			*sumList = append(*sumList, currentSum + num)
		} else {
			*sumList = append(*sumList, currentSum + nextNum)
		}
		return
	}

	// fmt.Printf("currentSum = %v, num = %v, nextNum = %v\n", currentSum, num, nextNum)
	dfs(triangle, i + 1, j, currentSum + num, sumList)
	dfs(triangle, i + 1, j + 1, currentSum + nextNum, sumList)
}

func min(nums []int) int {
    result := nums[0]
    for _, num := range nums {
        if result > num {
            result = num
        }
    }
    return result
}

func main() {
	// result: 11
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	// result: -10
	// triangle := [][]int{{-10}}

	// result: -1
	// triangle := [][]int{{-1}, {2, 3}, {1, -1, -3}}

	// result: 
	// triangle := [][]int{}

	result := minimumTotal(triangle)
	fmt.Printf("result = %v\n", result)
}

