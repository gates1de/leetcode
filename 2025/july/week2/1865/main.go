package main

import (
	"fmt"
)

type FindSumPairs struct {
	Nums1 []int
	Nums2 []int
	Freq map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	freq := make(map[int]int)
	for _, num := range nums2 {
		freq[num]++
	}

	return FindSumPairs{
		Nums1: nums1,
		Nums2: nums2,
		Freq: freq,
	}
}

func (this *FindSumPairs) Add(index int, val int)  {
	oldVal := this.Nums2[index]
	this.Freq[oldVal]--
	this.Nums2[index] += val
	this.Freq[this.Nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	result := int(0)
	for _, num := range this.Nums1 {
		rest := tot - num
		result += this.Freq[rest]
	}
	return result
}

func main() {
	// result: [null, 8, null, 2, 1, null, null, 11]
	nums1 := []int{1,1,2,2,2,3}
	nums2 := []int{1,4,5,2,5,4}
	operations := [][]int{{0,7},{1,3,2},{0,8},{0,4},{1,0,1},{1,1,1},{0,7}}

	// result: []
	// nums1 := []int{}
	// nums2 := []int{}
	// operations := [][]int{}

	obj := Constructor(nums1, nums2)
	for _, ope := range operations {
		command := ope[0]

		if command == 0 {
			tot := ope[1]
			fmt.Printf("obj.Count(%v) = %v\n", tot, obj.Count(tot))
		} else if command == 1 {
			obj.Add(ope[1], ope[2])
		}
	}
}

