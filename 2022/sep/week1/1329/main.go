package main
import (
	"fmt"
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	m := len(mat)
	n := len(mat[0])

	for i := m - 1; i >= 0; i-- {
		nums := []int{}
		for j := i; j < m; j++ {
			col := j
			row := j - i
			if row >= n {
				break
			}
			nums = append(nums, mat[col][row])
		}

		sort.Ints(nums)
		numsIndex := int(0)

		for j := i; j < m; j++ {
			col := j
			row := j - i
			if row >= n {
				break
			}
			mat[col][row] = nums[numsIndex]
			numsIndex++
		}
	}

	for i := 1; i < n; i++ {
		nums := []int{}
		for j := i; j < n; j++ {
			col := j - i
			row := j
			if col >= m {
				break
			}
			nums = append(nums, mat[col][row])
		}

		sort.Ints(nums)
		numsIndex := int(0)

		for j := i; j < n; j++ {
			col := j - i
			row := j
			if col >= m {
				break
			}
			mat[col][row] = nums[numsIndex]
			numsIndex++
		}
	}

	return mat
}

func main() {
	// result: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
	// mat := [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}

	// result: [[5,17,4,1,52,7],[11,11,25,45,8,69],[14,23,25,44,58,15],[22,27,31,36,50,66],[84,28,75,33,55,68]]
	// mat := [][]int{{11,25,66,1,69,7},{23,55,17,45,15,52},{75,31,36,44,58,8},{22,27,33,25,68,4},{84,28,14,11,5,50}}

	// result: [[2,1]]
	mat := [][]int{{2,1}}

	// result: [[1],[3],[2]]
	// mat := [][]int{{1},{3},{2}}

	// result: 
	// mat := [][]int{}

	result := diagonalSort(mat)
	fmt.Printf("result = %v\n", result)
}

