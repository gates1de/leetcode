package main
import (
	"fmt"
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	diagNum := m + n - 1
	diagMat := make([][]int, diagNum)
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			// fmt.Printf("i = %v, j = %v, diagMat[%v] append %v\n", i, j, m - i - 1 + j, mat[i][j])
			diagMat[m - i - 1 + j] = append(diagMat[m - i - 1 + j], mat[i][j])
		}
	}
	fmt.Printf("diagMat = %v\n", diagMat)
	sortedDiagMat := make([][]int, diagNum)
	for i, nums := range diagMat {
		sort.Slice(nums, func (i, j int) bool { return nums[i] > nums[j] })
		sortedDiagMat[i] = nums
	}
	fmt.Printf("sortedDiagMat = %v\n", sortedDiagMat)

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			nums := sortedDiagMat[m - i - 1 + j]
			mat[i][j], nums = pop(nums)
			sortedDiagMat[m - i - 1 + j] = nums
			// fmt.Printf("mat[%v][%v] = %v, nums = %v\n", i, j, mat[i][j], nums)
		}
	}
	return mat
}

func pop(nums []int) (int, []int) {
	if len(nums) == 0 {
		return -1, nums
	}
	result := nums[0]
	nums = nums[1:]
	return result, nums
}

func main() {
	// mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	mat := [][]int{{11,25,66,1,69,7}, {23,55,17,45,15,52}, {75,31,36,44,58,8}, {22,27,33,25,68,4}, {84,28,14,11,5,50}}
	result := diagonalSort(mat)
	fmt.Printf("result = %v\n", result)
}

