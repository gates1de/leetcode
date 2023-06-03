package main
import (
	"fmt"
)

type MyHashSet struct {
    Map map[int]bool
}

func Constructor() MyHashSet {
    return MyHashSet{Map: map[int]bool{}}
}

func (this *MyHashSet) Add(key int)  {
    this.Map[key] = true
}

func (this *MyHashSet) Remove(key int)  {
    delete(this.Map, key)
}

func (this *MyHashSet) Contains(key int) bool {
    return this.Map[key]
}

func main() {
	// result: [null, null, null, true, false, null, true, null, false]
	operations := []map[string]int{
		{"operation": 0, "key": 1},
		{"operation": 0, "key": 2},
		{"operation": 2, "key": 1},
		{"operation": 2, "key": 3},
		{"operation": 0, "key": 2},
		{"operation": 2, "key": 2},
		{"operation": 1, "key": 2},
		{"operation": 2, "key": 2},
	}

	// result: 
	// operations := []map[string]int{
	// 	{"operation": 0, "key": 0},
	// }

	obj := Constructor()

	for _, ope := range operations {
		if ope["operation"] == 0 {
			obj.Add(ope["key"])
		} else if ope["operation"] == 1 {
			obj.Remove(ope["key"])
		} else if ope["operation"] == 2 {
			fmt.Printf("obj.Contains(%v) = %v\n", ope["key"], obj.Contains(ope["key"]))
		}
	}
}

