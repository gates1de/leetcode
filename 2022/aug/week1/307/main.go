package main
import (
	"fmt"
)

type NumArray struct {
    BinaryIndexedTree []int
    Nums []int
}

func Constructor(nums []int) NumArray {
    res := NumArray{
		BinaryIndexedTree: make([]int, len(nums) + 1),
		Nums: nums,
    }
    
    for i := 1; i < len(nums) + 1; i++ {
        res.UpdateDiff(i, nums[i - 1])
    }
    
    return res    
}

func (this *NumArray) Update(i int, val int)  {
    diff := val - this.Nums[i]
    this.Nums[i] = val
    this.UpdateDiff(i + 1, diff)
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.Sum(j) - this.Sum(i - 1)
}

func (this *NumArray) Sum(i int) int {
    j := i + 1
    res := this.BinaryIndexedTree[j]
    for j - j & (-j) != 0 {
        j = j - j & (-j)
        res += this.BinaryIndexedTree[j]
    }
    return res
}

func (this *NumArray) UpdateDiff(i int, diff int) {
    this.BinaryIndexedTree[i] += diff
    for i + i & (-i) < len(this.BinaryIndexedTree) {
        i = i + i & (-i)
        this.BinaryIndexedTree[i] += diff
    }  
}

// Time Limit Exceeded
// type NumArray struct {
// 	Sums []int
// 	Nums []int
// 	Diffs map[int]int
// }
// 
// func Constructor(nums []int) NumArray {
// 	sums := make([]int, len(nums))
// 	sum := int(0)
// 	for i, num := range nums {
// 		sum += num
// 		sums[i] = sum
// 	}
// 	return NumArray{Sums: sums, Nums: nums, Diffs: map[int]int{}}
// }
// 
// func (this *NumArray) Update(index int, val int)  {
// 	if len(this.Nums) <= index {
// 		return
// 	}
// 	pre := this.Nums[index]
// 	this.Nums[index] = val
// 	this.Diffs[index] += val - pre
// }
// 
// func (this *NumArray) SumRange(left int, right int) int {
// 	if left < 0 || len(this.Sums) <= right {
// 		return 0
// 	}
// 
// 	sum := this.Sums[right]
// 	subtract := this.Sums[left] - this.Nums[left]
// 	for i, diff := range this.Diffs {
// 		if i <= left {
// 			subtract += diff
// 		}
// 		if i <= right {
// 			sum += diff
// 		}
// 	}
// 
// 	return sum - subtract
// }

func main() {
	/**
	 * query is
	 * index 0: 0 -> update, 1 -> sum range
	 * index 1, 2: index or values
	 */

	// result: [null, 9, null, 8]
	// initNums := []int{1, 3, 5}
	// queries := [][]int{
	// 	{1, 0, 2},
	// 	{0, 1, 2},
	// 	{1, 0, 2},
	// }

	// result: [null, 9, null, 8]
	initNums := []int{1, 3, 5, 7, 9}
	queries := [][]int{
		{1, 0, 4},
		{0, 2, 30},
		{1, 0, 1},
		{1, 0, 2},
		{0, 3, -10},
		{1, 2, 4},
		{1, 3, 4},
		{1, 4, 4},
		{0, 3, -100},
		{1, 2, 4},
		{1, 3, 4},
		{1, 4, 4},
	}

	// result: 
	// initNums := []int{}
	// queries := [][]int{}

	obj := Constructor(initNums)
	for _, query := range queries {
		if query[0] == 0 {
			obj.Update(query[1], query[2])
		} else if query[0] == 1 {
			fmt.Printf("obj.SumRange(%v, %v) = %v\n", query[1], query[2], obj.SumRange(query[1], query[2]))
		}
	}
}

