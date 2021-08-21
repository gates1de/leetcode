package main
import (
	"fmt"
	"math/rand"
	"time"
)

type NumArray struct {
	RangeSums []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{RangeSums: make([]int, len(nums))}
	numArray.RangeSums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numArray.RangeSums[i] = nums[i] + numArray.RangeSums[i - 1]
	}

	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.RangeSums[right]
	}
	return this.RangeSums[right] - this.RangeSums[left - 1]
}

func generateNums(length int, min int, max int) []int {
    result := make([]int, length)
    rand.Seed(time.Now().UnixNano())
    for i, _ := range result {
        result[i] = rand.Intn(max - min) + min
    }
    return result
}

func main() {
	// nums := []int{-2, 0, 3, -5, 2, -1}
	// obj := Constructor(nums)
	// fmt.Printf("obj.SumRange(0, 2) = %v\n", obj.SumRange(0, 2))
	// fmt.Printf("obj.SumRange(2, 5) = %v\n", obj.SumRange(2, 5))
	// fmt.Printf("obj.SumRange(0, 5) = %v\n", obj.SumRange(0, 5))

	nums := generateNums(10000, -100000, 100000)
	obj := Constructor(nums)
	fmt.Printf("obj.SumRange(0, 9999) = %v\n", obj.SumRange(0, 9999))
	fmt.Printf("obj.SumRange(1000, 9998) = %v\n", obj.SumRange(1000, 9998))
	fmt.Printf("obj.SumRange(500, 500) = %v\n", obj.SumRange(500, 500))

	// nums := []int{}
	// obj := Constructor(nums)
	// fmt.Printf("obj.SumRange(, ) = %v\n", obj.SumRange(, ))
}

