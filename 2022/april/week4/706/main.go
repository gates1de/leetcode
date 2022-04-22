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
	fmt.Println(this.Map)
}

func (this *MyHashMap) Get(key int) int {
	if _, ok := this.Map[key]; !ok {
		return -1
	}
	return this.Map[key]
}

func (this *MyHashMap) Remove(key int)  {
    delete(this.Map, key)
	fmt.Println(this.Map)
}

func main() {
	// result: [null, null, null, 1, -1, null, 1, null, -1]
	operations := []map[string]int{
		{"operation": 0, "key": 1, "value": 1},
		{"operation": 0, "key": 2, "value": 2},
		{"operation": 1, "key": 1},
		{"operation": 1, "key": 3},
		{"operation": 0, "key": 2, "value": 1},
		{"operation": 1, "key": 2},
		{"operation": 2, "key": 2},
		{"operation": 1, "key": 2},
	}

	// result: 
	// operations := []map[string]int{
	// 	{"operation": 0, "key": 0},
	// }

	obj := Constructor()

	for _, ope := range operations {
		if ope["operation"] == 0 {
			obj.Put(ope["key"], ope["value"])
		} else if ope["operation"] == 1 {
			fmt.Printf("obj.Get(%v) = %v\n", ope["key"], obj.Get(ope["key"]))
		} else if ope["operation"] == 2 {
			obj.Remove(ope["key"])
		}
	}
}

