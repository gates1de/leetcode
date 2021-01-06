package main
import (
	"fmt"
)

func countArrangement(n int) int {
	result := 0
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i - 1] = i
	}

    init := make([][]int, 0, n)
	perm := generatePermute(nums, []int{}, init)

TOP:
	for _, p := range perm {
		// fmt.Printf("perm = %v\n", p)
		for i := 1; i <= len(p); i++ {
			p_val := p[i - 1]
			if (i % p_val) != 0 && (p_val % i) != 0 {
				continue TOP
			}
		}
		// fmt.Printf("beatiful arrangement = %v\n", p)
		result++
	}
	return result
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func generatePermute(nums []int, perm []int, current [][]int) [][]int {
	n := len(nums)

	if n == 0 {
		// fmt.Printf("added perm = %v\n", perm)
		current = append(current, perm)
		return current
	}

	for i := 0; i < n; i++ {
		new_nums := make([]int, 0, n)
		new_nums = append(new_nums, nums[:i]...)
		new_nums = append(new_nums, nums[i + 1:]...)
		current = generatePermute(new_nums, append(perm, nums[i]), current)
	}
	return current
}

func main() {
	n := 6
	r := countArrangement(n)
	fmt.Printf("result = %v\n", r)
}

