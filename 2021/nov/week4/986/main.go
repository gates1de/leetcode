package main
import (
	"fmt"
)

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := [][]int{}
	n1 := len(firstList)
	n2 := len(secondList)
	if n1 == 0 || n2 == 0 {
		return result
	}

	i := int(0)
	j := int(0)
	for i < n1 && j < n2 {
		first := firstList[i]
		second := secondList[j]

		start := max(first[0], second[0])
		end := min(first[1], second[1])
		if start <= end {
			result = append(result, []int{start, end})
		}

		if first[1] <= second[1] {
			i++
		}
		if second[1] <= first[1] {
			j++
		}
	}
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// result: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
	// firstList := [][]int{{0,2},{5,10},{13,23},{24,25}}
	// secondList := [][]int{{1,5},{8,12},{15,24},{25,26}}

	// result: []
	// firstList := [][]int{{1,3},{5,9}}
	// secondList := [][]int{}

	// result: []
	// firstList := [][]int{}
	// secondList := [][]int{{4,8},{10,12}}

	// result: [[3,7]]
	// firstList := [][]int{{1,7}}
	// secondList := [][]int{{3,10}}

	// result: [[16,16]]
	// firstList := [][]int{{14,16}}
	// secondList := [][]int{{7,13},{16,20}}

	// result: [[8,10],[12,15]]
	firstList := [][]int{{8,15}}
	secondList := [][]int{{2,6},{8,10},{12,20}}

	// result: 
	// firstList := [][]int{}
	// secondList := [][]int{}

	result := intervalIntersection(firstList, secondList)
	fmt.Printf("result = %v\n", result)
}

