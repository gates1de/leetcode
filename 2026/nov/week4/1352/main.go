package main

import (
	"fmt"
)

type ProductOfNumbers struct {
	Products []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{Products: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.Products = this.Products[:1]
		return
	}

	last := this.Products[len(this.Products)-1]
	this.Products = append(this.Products, last*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.Products) {
		return 0
	}

	last := len(this.Products) - 1
	return this.Products[last] / this.Products[last-k]
}

func main() {
	// result: [null,null,null,null,null,null,20,40,0,null,32]
	operations := [][]int{{0, 3}, {0, 0}, {0, 2}, {0, 5}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {0, 8}, {1, 2}}

	obj := Constructor()

	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Add(ope[1])
		} else if ope[0] == 1 {
			fmt.Printf("obj.GetProduct(%d) = %d\n", ope[1], obj.GetProduct(ope[1]))
		}
	}
}
