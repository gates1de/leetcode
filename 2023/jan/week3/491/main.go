package main
import (
	"bytes"
	"fmt"
	"strconv"
)

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0, 10000)
	m := make(map[string]bool, 10000)
	backtrack(make([]int, 0, 15), 0, nums, m, &result)
	return result
}

func backtrack(sequence []int, index int, nums []int, m map[string]bool, result *[][]int) {
	n := len(nums)
	if index == n {
		hashKey := generateHashKey(sequence)
		if len(sequence) >= 2 && !m[hashKey] {
			*result = append(*result, sequence)
			m[hashKey] = true
		}
		return
	}

	num := nums[index]
	nextSequence := make([]int, len(sequence), 15)
	copy(nextSequence, sequence)

	if len(nextSequence) == 0 {
		nextSequence = append(nextSequence, num)
		backtrack(nextSequence, index + 1, nums, m, result)
	} else {
		last := nextSequence[len(nextSequence) - 1]
		if last <= num {
			nextSequence = append(nextSequence, num)
			backtrack(nextSequence, index + 1, nums, m, result)
		}
	}

	backtrack(sequence, index + 1, nums, m, result);
}

func generateHashKey(sequence []int) string {
	var result bytes.Buffer
    for i := 0; i < len(sequence); i++ {
        result.WriteString(strconv.Itoa(sequence[i]))
        result.WriteString("+")
    }
    return result.String()
}

// Wrong Answer
func ngSolution(nums []int) [][]int {
	result := make([][]int, 0, 10000)
	m := make(map[string]bool, 10000)
	for i, _ := range nums {
		helper(make([]int, 0, 15), i, nums, m, &result)
	}
	return result
}

func helper(current []int, index int, nums []int, m map[string]bool, result *[][]int) {
	if index >= len(nums) {
		concatStr := ""
		for _, num := range current {
			concatStr += fmt.Sprintf("%v", num)
		}
		if len(current) >= 2 && !m[concatStr] {
			*result = append(*result, current)
			m[concatStr] = true
		}
		return
	}

	num := nums[index]
	if len(current) == 0 {
		current = append(current, num)
		helper(current, index + 1, nums, m, result)
		return
	}

	last := current[len(current) - 1]
	if last <= num {
		newCurrent := make([]int, len(current), 15)
		copy(newCurrent, current)
		newCurrent = append(newCurrent, num)
		helper(newCurrent, index + 1, nums, m, result)
	}
	helper(current, index + 1, nums, m, result)
}

func main() {
	// result: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
	// nums := []int{4,6,7,7}

	// result: [[4,4]]
	nums := []int{4,4,3,2,1}

	// result: 
	// nums := []int{}

	// result: 
	// nums := []int{}

	result := findSubsequences(nums)
	fmt.Printf("result = %v\n", result)
}

