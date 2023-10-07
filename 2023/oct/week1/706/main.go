package main
import (
	"fmt"
)

type MyHashMap struct {
    Map map[int]int
}

func Constructor() MyHashMap {
    return MyHashMap{Map: map[int]int{}}
}

func (this *MyHashMap) Put(key int, value int)  {
    this.Map[key] = value
}

func (this *MyHashMap) Get(key int) int {
    if _, ok := this.Map[key]; !ok {
        return -1
    }
    return this.Map[key]
}

func (this *MyHashMap) Remove(key int)  {
    delete(this.Map, key)
}

func main() {
	// result: [null, null, null, 1, -1, null, 1, null, -1]
	operations := [][]int{{0,1,1},{0,2,2},{1,1},{1,3},{0,2,1},{1,2},{2,2},{1,2}}

	// result: 
	// operations := [][]int{{0,1},{1,1},{2,1}}

	obj := Constructor()
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Put(ope[1], ope[2])
		} else if ope[0] == 1 {
			fmt.Printf("obj.Get(%v) = %v\n", ope[1], obj.Get(ope[1]))
		} else if ope[0] == 2 {
			obj.Remove(ope[1])
		}
	}
}

