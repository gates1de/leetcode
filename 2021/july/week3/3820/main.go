package main
import (
	"fmt"
	"math/rand"
	// "time"
)

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
    return Solution{arr: nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.arr
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    tmp := make([]int, len(this.arr))
    copy(tmp, this.arr)
    for i := 0; i < len(this.arr); i++ {
        r := rand.Intn(len(this.arr))
        tmp[i], tmp[r] = tmp[r], tmp[i]
    }
    return tmp
}

// MyAnswer: Slow
// type Solution struct {
// 	Original []int
// 	Nums []int
// }
//
// func Constructor(nums []int) Solution {
// 	copyNums := make([]int, len(nums))
// 	copy(copyNums, nums)
// 	return Solution{Original: nums, Nums: copyNums}
// }
// 
// func (this *Solution) Reset() []int {
// 	return this.Original
// }
// 
// func (this *Solution) Shuffle() []int {
//     n := len(this.Nums)
//     for i := n - 1; i >= 0; i-- {
// 		rand.Seed(time.Now().UnixNano())
//         j := rand.Intn(i + 1)
//         this.Nums[i], this.Nums[j] = this.Nums[j], this.Nums[i]
//     }
// 	return this.Nums
// }

func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	fmt.Printf("obj.Shuffle() = %v\n", obj.Shuffle())
	fmt.Printf("obj.Reset() = %v\n", obj.Reset())
	fmt.Printf("obj.Shuffle() = %v\n", obj.Shuffle())

	// nums := []int{}
	// obj := Constructor(nums)
	// fmt.Printf("obj.Shuffle() = %v\n", obj.Shuffle())
	// obj.Reset()
	// fmt.Printf("obj.Shuffle() = %v\n", obj.Shuffle())
}

