package main
import (
	"fmt"
)

type NumArray struct {
	Nums []int
}


func Constructor(nums []int) NumArray {
	return NumArray{Nums: nums}
}


func (this *NumArray) Update(index int, val int)  {
	if index >= len(this.Nums) {
		return
	}
	this.Nums[index] = val
}


func (this *NumArray) SumRange(left int, right int) int {
	result := int(0)
	for _, num := range this.Nums[left:right + 1] {
		result += num
	}
	return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 */

func main() {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)
	fmt.Printf("obj.SumRange(0, 2) = %v\n", obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Printf("obj.SumRange(0, 2) = %v\n", obj.SumRange(0, 2))
}

