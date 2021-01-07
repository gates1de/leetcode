package main
import (
	"fmt"
)

func countArrangement(n int) int {
	nums := make([]int, n)
	factorial := int(1)
	for i := 1; i <= n; i++ {
		nums[i - 1] = i
		factorial *= i
	}

    perm := make([][]int, 0, factorial)
	perm = generatePermute(nums, []int{}, perm)
	return len(perm)
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func remove(i int, nums []int) []int {
	nums = append(nums[:i], nums[i + 1:]...)
	result := copySlice(nums)
	return result
}

func generatePermute(nums []int, perm []int, current [][]int) [][]int {
	n := len(nums)

	if n == 0 {
		for i := 1; i <= len(perm); i++ {
			p_val := perm[i - 1]
			if (i % p_val) != 0 && (p_val % i) != 0 {
				return current
			}
		}
		return append(current, perm)
	}

	for i := 0; i < n; i++ {
		new_perm := copySlice(append(perm, nums[i]))
		new_nums := remove(i, copySlice(nums))
		current = generatePermute(new_nums, new_perm, current)
	}
	return current
}

func main() {
	n := 7
	r := countArrangement(n)
	fmt.Printf("result = %v\n", r)
}

