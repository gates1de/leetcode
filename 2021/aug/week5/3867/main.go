package main
import (
	"fmt"
	"sort"
)

func minPatches(nums []int, n int) int {
	result := int(0)
	helper(-1, nums, n, &result)
	return result
}

func helper(addTarget int, nums []int, n int, result *int) {
	m := map[int]bool{}
	if addTarget == -1 {
		combinations([]int{}, nums, m, 0)
	} else {
		*result++
		index := sort.SearchInts(nums, addTarget)
		if index >= len(nums) {
			nums = append(nums, addTarget)
		} else {
			nums = append(nums[:index + 1], nums[index:]...)
			nums[index] = addTarget
		}
		combinations([]int{}, nums, m, 0)
	}
	// fmt.Printf("m = %v\n", m)

	addTarget = -1
	for i := 1; i <= n; i++ {
		if !m[i] {
			addTarget = i
			break
		}
	}

	// fmt.Printf("addTarget = %v\n", addTarget)
	if addTarget == -1 {
		return
	}
	helper(addTarget, nums, n, result)
}

func combinations(current []int, nums []int, m map[int]bool, index int) {
	// fmt.Printf("current = %v\n", current)
	s := sum(current)
	m[s] = true
	for i := index; i < len(nums); i++ {
		newCurrent := append(current, nums[i])
		combinations(newCurrent, nums, m, i + 1)
	}
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func contains(target int, nums []int) bool {
   high := len(nums) - 1
   low := 0
   var mid int
   for low <= high {
      mid = (high + low) / 2
      if nums[mid] == target {
		  return true
      } else if nums[mid] > target {
         high = mid
      } else {
         low = mid + 1
      }
   }
   return false
}

func main() {
	// result: 1
	// nums := []int{1, 3}
	// n := int(6)

	// result: 2
	// nums := []int{1, 5, 10}
	// n := int(20)

	// result: 0
	// nums := []int{1, 2, 2}
	// n := int(5)

	// result: 
	nums := []int{1, 3, 9, 10000}
	n := int(100000000)

	// result: 
	// nums := []int{}
	// n := int(0)

	result := minPatches(nums, n)
	fmt.Printf("result = %v\n", result)
}

