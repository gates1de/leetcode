package main
import (
	"fmt"
	"math"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	// fmt.Printf("matrix = [\n")
	// for _, cols := range matrix {
	// 	fmt.Printf("%v\n", cols)
	// }
	// fmt.Printf("]\n")
	s := make([][]int, len(matrix) + 1)
	for i, _ := range s {
		s[i] = make([]int, len(matrix[0]) + 1)
	}
	return helper(s, matrix, k)
}

func helper(s [][]int, matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			s[i][j] = matrix[i - 1][j - 1]
			s[i][j] += s[i][j - 1]
			for _, cols := range matrix[:i - 1] {
				s[i][j] += cols[j - 1]
			}
			if s[i][j] == k {
				return k
			}
		}
	}

	// fmt.Printf("s = [\n")
	// for _, cols := range s {
	// 	fmt.Printf("%v\n", cols)
	// }
	// fmt.Printf("]\n")

	copyS := make([][]int, len(s))
	copy(copyS, s)
	result := math.MinInt32
	for i, cols := range matrix {
		for j, _ := range cols {
			if i == 0 && j == 0 {
				continue
			}
			for ii := i; ii < m; ii++ {
				for jj := j; jj < n; jj++ {
					// partSum := s[ii][jj] - s[i][jj] - s[ii][j] + s[i][j]
					partSum := s[ii + 1][jj + 1] - s[i][jj + 1] - s[ii + 1][j] + s[i][j]
					// fmt.Printf("[%v, %v] ~ [%v, %v] = %v\n", i, j, ii, jj, partSum)
					if partSum == k {
						return k
					} else if partSum < k {
						result = max(result, partSum)
					}
				}
			}
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
	// result: 2
	// matrix := [][]int{{1, 0, 1}, {0, -2, 3}}
	// k := int(2)

	// result: 3
	// matrix := [][]int{{2, 2, -1}}
	// k := int(3)

	// result: 
	matrix := [][]int{{1, 0, -3}, {2, 5, 4}, {3, 6, 9}}
	k := int(7)

	// result: 
	// matrix := [][]int{}
	// k := int(0)

	result := maxSumSubmatrix(matrix, k)
	fmt.Printf("result = %v\n", result)
}

