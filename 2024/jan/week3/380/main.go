package main
import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	List []int
	Map map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{List: make([]int, 0, 10000), Map: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Map[val]; ok {
		return false
	}

	this.List = append(this.List, val)
	this.Map[val] = len(this.List) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Map[val]; !ok {
		return false
	}

	index := this.Map[val]
	this.List[index] = this.List[len(this.List) - 1]
	this.Map[this.List[index]] = index
	this.List = this.List[:len(this.List) - 1]
	delete(this.Map, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.List[rand.Intn(len(this.List))]
}

func main() {
	// result: [null, true, false, true, 2, true, false, 2]
	operations := [][]int{{0,1},{1,2},{0,2},{2,0},{1,1},{0,2},{2,0}}

	// result: 
	// operations := [][]int{}

	obj := Constructor()
	for _, ope := range operations {
		if ope[0] == 0 {
			fmt.Printf("obj.Insert(%v) = %v\n", ope[1], obj.Insert(ope[1]))
		} else if ope[0] == 1 {
			fmt.Printf("obj.Remove(%v) = %v\n", ope[1], obj.Remove(ope[1]))
		} else if ope[0] == 2 {
			fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
		}
	}
}

