package main
import (
	"fmt"
)

func countArrangement(n int) int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i - 1] = i
	}
	count := int(0)
	search(nums, 1, &count)
	return count
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

func search(nums []int, i int, count *int) {
	n := len(nums)

	if n == 0 {
		*count++
		return
	}

	for j := 0; j < n; j++ {
		val := nums[j]
		if (i % val) != 0 && (val % i) != 0 {
			continue
		}
		new_nums := remove(j, copySlice(nums))
		search(new_nums, i + 1, count)
	}
}

func main() {
	n := 11
	r := countArrangement(n)
	fmt.Printf("result = %v\n", r)
}

