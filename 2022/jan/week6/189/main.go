package main
import (
	"fmt"
)

func rotate(nums []int, k int)  {
	n := len(nums)
	if n <= 1 || k <= 0 || n == k {
		return
	}

	k %= n
	k = n - k
	result := make([]int, n)
	firstHalf := nums[k:]
	latterHalf := nums[:k]
	// fmt.Println(k, firstHalf, latterHalf)
	copy(result[:len(firstHalf)], firstHalf)
	copy(result[len(firstHalf):], latterHalf)
	copy(nums, result)
}

func main() {
	// result: [5,6,7,1,2,3,4]
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := int(3)

	// result: [3,99,-1,-100]
	// nums := []int{-1, -100, 3, 99}
	// k := int(2)

	// result: [1]
	// nums := []int{1}
	// k := int(0)

	// result: [4,3,6,1,8,7,6,2,1,8,9]
	 nums := []int{4,3,6,1,8,7,6,2,1,8,9,10,11}
	 k := int(11)

	// result: 
	// nums := []int{}
	// k := int(0)

	rotate(nums, k)
	fmt.Printf("result = %v\n", nums)
}

