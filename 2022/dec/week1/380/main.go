package main
import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
    Nums []int
    Indices map[int]int
}

func Constructor() RandomizedSet {
    return RandomizedSet{Nums: []int{}, Indices: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.Indices[val]; ok {
        return false
    }

    this.Nums = append(this.Nums, val)
    this.Indices[val] = len(this.Nums) - 1

    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.Indices[val]; !ok {
        return false
    }

    if len(this.Nums) < 2 {
        this.Nums = this.Nums[0:0]
    } else {
        i := this.Indices[val]
        swap := this.Nums[len(this.Nums) - 1]
        this.Indices[swap] = i
        this.Nums[i] = swap
        this.Nums = this.Nums[0:len(this.Nums) - 1]
    }

    delete(this.Indices, val)
    return true
}

func (this *RandomizedSet) GetRandom() int {
    return this.Nums[rand.Int() % len(this.Indices)]
}

// Wrong Answer
// type RandomizedSet struct {
// 	Counter map[int]int
// 	Nums []int
// 	Cursor int
// }
// 
// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		Counter: map[int]int{},
// 		Nums: make([]int, 0, 200000),
// 		Cursor: 0,
// 	}
// }
// 
// func (this *RandomizedSet) Insert(val int) bool {
// 	if this.Counter[val] > 0 {
// 		return false
// 	}
// 
// 	this.Counter[val]++
// 	this.Nums = append(this.Nums, val)
// 	return true
// }
// 
// func (this *RandomizedSet) Remove(val int) bool {
// 	if this.Counter[val] == 0 {
// 		return false
// 	}
// 
// 	this.Counter[val]--
// 	return true
// }
// 
// func (this *RandomizedSet) GetRandom() int {
// 	n := len(this.Nums)
// 	if n == 0 {
// 		return 0
// 	}
// 
// 	this.Cursor++
// 	if this.Cursor >= n {
// 		this.Cursor = 0
// 	}
// 
// 	result := this.Nums[this.Cursor]
// 	for this.Counter[result] == 0 {
// 		this.Cursor++
// 		if this.Cursor >= n {
// 			this.Cursor = 0
// 		}
// 		result = this.Nums[this.Cursor]
// 	}
// 
// 	return result
// }

func main() {
	// funcs[i][0] == 0 -> insert, funcs[i][0] == 1 -> remove, funcs[i][0] == 2 -> get random

	// result: [null, true, false, true, 2, true, false, 2]
	funcs := [][]int{
		{0, 1},
		{1, 2},
		{0, 2},
		{2, 0},
		{1, 1},
		{0, 2},
		{2, 0},
	}

	// result: []
	// funcs := [][]int{}

	obj := Constructor()
	for _, function := range funcs {
		val := function[1]
		switch function[0] {
		case 0:
			fmt.Printf("obj.Insert(%v) = %v\n", val, obj.Insert(val))
		case 1:
			fmt.Printf("obj.Remove(%v) = %v\n", val, obj.Remove(val))
		case 2:
			fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
		}
	}
}

