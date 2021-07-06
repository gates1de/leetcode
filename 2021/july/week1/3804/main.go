package main
import (
	"fmt"
	"sort"
)

func minSetSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	sort.Slice(arr, func (a, b int) bool { return arr[a] < arr[b] })

	pre := int(0)
	countArr := []int{}
	for _, num := range arr {
		if pre == num {
			countArr[len(countArr) - 1]++
			continue
		}
		pre = num
		countArr = append(countArr, 1)
	}
	sort.Slice(countArr, func (a, b int) bool { return countArr[a] > countArr[b] })
	halfSize := len(arr) / 2
	// fmt.Printf("countArr = %v, halfSize = %v\n", countArr, halfSize)
	sum := int(0)
	for i, count := range countArr {
		sum += count
		if sum >= halfSize {
			return i + 1
		}
	}
	return halfSize
}

func main() {
	// result: 2
	// arr := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}

	// result: 1
	// arr := []int{7, 7, 7, 7, 7, 7}

	// result: 1
	// arr := []int{1, 9}

	// result: 1
	// arr := []int{1000, 1000, 3, 7}

	// result: 5
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// result: 
	// arr := []int{}

	result := minSetSize(arr)
	fmt.Printf("result = %v\n", result)
}

