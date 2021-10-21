package main
import (
	"fmt"
)

type RandomizedSet struct {
	Nums map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{Nums: map[int]bool{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if ok, _ := this.Nums[val]; ok {
		return false
	}
	this.Nums[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if ok, _ := this.Nums[val]; !ok {
		return false
	}
	delete(this.Nums, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	for key, _ := range this.Nums {
		return key
	}
	return 0
}

func main() {
	obj := Constructor()

	fmt.Printf("obj.Insert(1) = %v\n", obj.Insert(1))
	fmt.Printf("obj.Remove(2) = %v\n", obj.Remove(2))
	fmt.Printf("obj.Insert(2) = %v\n", obj.Insert(2))
	fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	fmt.Printf("obj.Remove(1) = %v\n", obj.Remove(1))
	fmt.Printf("obj.Insert(2) = %v\n", obj.Insert(2))
	fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())

	// fmt.Printf("obj.Insert() = %v\n", obj.Insert())
	// fmt.Printf("obj.Remove() = %v\n", obj.Remove())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
}

