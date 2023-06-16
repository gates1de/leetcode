package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numOfWays(nums []int) int {
	n := len(nums)
	table := make([][]int, n)
	for i, _ := range table {
		table[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		table[i][0] = 1
		table[i][i] = 1
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			table[i][j] = (table[i - 1][j - 1] + table[i - 1][j]) % modulo
		}
	}

	return (dfs(nums, table) - 1) % modulo
}

func dfs(nums []int, table [][]int) int {
	n := len(nums)
	if n < 3 {
		return 1
	}

	leftNodes := make([]int, 0, n)
	rightNodes := make([]int, 0, n)
	for i := 1; i < n; i++ {
		num := nums[i]
		if num < nums[0] {
			leftNodes = append(leftNodes, num)
		} else {
			rightNodes = append(rightNodes, num)
		}
	}

	leftWays := dfs(leftNodes, table) % modulo
	rightWays := dfs(rightNodes, table) % modulo

	return (((leftWays * rightWays) % modulo) * table[n - 1][len(leftNodes)]) % modulo
}

func main() {
	// result: 1
	// nums := []int{2,1,3}

	// result: 5
	// nums := []int{3,4,5,1,2}

	// result: 0
	nums := []int{1,2,3}

	// result: 
	// nums := []int{}

	result := numOfWays(nums)
	fmt.Printf("result = %v\n", result)
}

