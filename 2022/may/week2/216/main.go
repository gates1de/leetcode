package main
import (
	"fmt"
)

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	for i := 1; i <= 9; i++ {
		visited := make([]bool, 9)
		visited[i - 1] = true
		helper(i, i, k - 1, n, visited, &result)
	}
	return result
}

func helper(sum int, min int, k int, n int, visited []bool, result *[][]int) {
	if k == 0 {
		if sum != n {
			return
		}
		nums := []int{}
		for i, isVisited := range visited {
			if isVisited {
				nums = append(nums, i + 1)
			}
		}
		*result = append(*result, nums)
		return
	}

	if k >= n - sum {
		return
	}

	loopMax := int(9)
	if loopMax > n - sum {
		loopMax = n - sum
	}
	for i := min + 1; i <= loopMax; i++ {
		if visited[i - 1] {
			continue
		}

		newVisited := make([]bool, 9)
		copy(newVisited, visited)
		newVisited[i - 1] = true
		if min < i {
			helper(sum + i, i, k - 1, n, newVisited, result)
		} else {
			helper(sum + i, min, k - 1, n, newVisited, result)
		}
	}
}

func main() {
	// result: [[1,2,4]]
	// k := int(3)
	// n := int(7)

	// result: [[1,2,6],[1,3,5],[2,3,4]]
	// k := int(3)
	// n := int(9)

	// result: []
	k := int(4)
	n := int(1)

	// result: 
	// k := int(0)
	// n := int(0)

	result := combinationSum3(k, n)
	fmt.Printf("result = %v\n", result)
}

