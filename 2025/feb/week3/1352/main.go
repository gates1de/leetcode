package main
import (
	"fmt"
)

type ProductOfNumbers struct {
	Nums []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{Nums: make([]int, 0)}
}

func (this *ProductOfNumbers) Add(num int)  {
	this.Nums = append(this.Nums, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.Nums) < k {
		return -1
	}

	result := int(1)
	for i := 0; i < k; i++ {
		result *= this.Nums[len(this.Nums) - i - 1]
	}

	return result
}

func main() {
	// result: [null,null,null,null,null,null,20,40,0,null,32]
	operations := [][]int{{0,3},{0,0},{0,2},{0,5},{0,4},{1,2},{1,3},{1,4},{0,8},{1,2}}

	obj := Constructor()

	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Add(ope[1])
		} else if ope[0] == 1 {
			fmt.Printf("obj.GetProduct(%d) = %d\n", ope[1], obj.GetProduct(ope[1]))
		}
	}
}
