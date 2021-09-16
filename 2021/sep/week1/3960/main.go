package main
import (
	"fmt"
)

func arrayNesting(nums []int) int {
	res := 1
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		cur := 0
		val := nums[i]
		for !visited[val] {
			cur++
			visited[val] = true
			val = nums[val]
		}
		if res < cur {
			res = cur
		}
	}
	return res
}

// Too Slow & Use more memory
func mySolution(nums []int) int {
	m := map[int]int{}

	result := int(0)
	for i, _ := range nums {
		count := helper(i, nums, m)
		if result < count {
			result = count
		}
	}
	return result
}

func helper(start int, nums []int, m map[int]int) int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	count := int(0)
	i := start
	// fmt.Printf("start = %v, m = %v\n", start, m)
	for true {
		if copyNums[i] == -1 {
			break
		}
		count++
		if m[i] > count {
			break
		}
		m[i] = count
		copyNums[i] = -1
		i = nums[i]
	}

	return count
}

func main() {
	// result: 4
	// nums := []int{5, 4, 0, 3, 1, 6, 2}

	// result: 1
	// nums := []int{0, 1, 2}

	// result: 1
	nums := []int{0}

	// result: 
	// nums := []int{}

	// result: 
	// nums := []int{}

	result := arrayNesting(nums)
	fmt.Printf("result = %v\n", result)
}

