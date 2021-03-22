package main
import (
	"fmt"
)

type FreqStack struct {
	nums []int
	numsCounter map[int]int
	currentMaxCount int
	freqKeys []int
}


func Constructor() FreqStack {
	stack := FreqStack{
		nums: []int{},
		numsCounter: map[int]int{},
		currentMaxCount: 0,
		freqKeys: []int{},
	}
	return stack
}


func (this *FreqStack) Push(val int)  {
	this.nums = append(this.nums, val)
	this.numsCounter[val]++
	if this.currentMaxCount < this.numsCounter[val] {
		this.freqKeys = []int{val}
		this.currentMaxCount = this.numsCounter[val]
	} else if this.currentMaxCount == this.numsCounter[val] {
		this.freqKeys = append(this.freqKeys, val)
	}
}


func (this *FreqStack) Pop() int {
	freqKey := int(-1)
	removeIndex := int(-1)
TOP:
	for i := len(this.nums) - 1; i >= 0; i-- {
		val := this.nums[i]
		lastKeyIndex := findLast(val, this.freqKeys)
		if lastKeyIndex >= 0 {
			removeIndex = i
			freqKey = val
			break TOP
		}
	}
	if removeIndex >= 0 {
		this.nums = append(this.nums[:removeIndex], this.nums[removeIndex + 1:]...)
	}
	this.numsCounter[freqKey]--
	if len(this.freqKeys) == 1 {
		this.currentMaxCount--
	}
	this.freqKeys = []int{}
	for key, val := range this.numsCounter {
		if this.currentMaxCount <= val {
			this.freqKeys = append(this.freqKeys, key)
		}
	}

	// fmt.Printf("nums = %v, numsCounter = %v, freqKeys = %v\n", this.nums, this.numsCounter, this.freqKeys)
	return freqKey
}

func findLast(target int, nums []int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		val := nums[i]
		if val == target {
			return i
		}
	}
	return -1
}

func main() {
	obj := Constructor()
	fmt.Printf("init = %v\n", obj)
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)
	fmt.Printf("popped = %v\n", obj.Pop())
	fmt.Printf("popped = %v\n", obj.Pop())
	fmt.Printf("popped = %v\n", obj.Pop())
	fmt.Printf("popped = %v\n", obj.Pop())
}

