package main
import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	counts := make([][2]int, len(mat))
	for i, nums := range mat {
		counts[i][0] = i
		counts[i][1] = oneCounter(nums)
	}

	sort.Slice(counts, func(a, b int) bool {
		if counts[a][1] == counts[b][1] {
			return counts[a][0] < counts[b][0]
		}
		return counts[a][1] < counts[b][1]
	})

	result := make([]int, k)
	for i, c := range counts {
		if i == k {
			break
		}
		result[i] = c[0]
	}

	return result
}

func oneCounter(nums []int) int {
	result := int(0)
	for _, num := range nums {
		if num == 0 {
			break
		}
		result++
	}
	return result
}

func main() {
	// result: [2,0,3]
	// mat := [][]int{
	// 	{1,1,0,0,0},
	// 	{1,1,1,1,0},
	// 	{1,0,0,0,0},
	// 	{1,1,0,0,0},
	// 	{1,1,1,1,1},
	// }
	// k := int(3)

	// result: [0,2]
	// mat := [][]int{
	// 	{1,0,0,0},
	// 	{1,1,1,1},
	// 	{1,0,0,0,},
	// 	{1,1,0,0,},
	// }
	// k := int(2)

	// result: [0]
	// mat := [][]int{{1,1},{1,1}}
	// k := int(1)

	// result: [4,6,0,1,2,3]
	mat := [][]int{{1,1,0},{1,1,0},{1,1,1},{1,1,1},{0,0,0},{1,1,1},{1,0,0}}
	k := int(6)

	// result: 
	// mat := [][]int{}
	// k := int(0)

	result := kWeakestRows(mat, k)
	fmt.Printf("result = %v\n", result)
}

