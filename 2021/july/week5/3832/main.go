package main
import (
	"fmt"
)

type MapSum struct {
	Memo map[string]int
	Maps map[string]int
}

func Constructor() MapSum {
	return MapSum{Memo: map[string]int{}, Maps: map[string]int{}}
}

func (this *MapSum) Insert(key string, val int)  {
	originalVal := val
	if this.Maps[key] > 0 {
		val -= this.Maps[key]
	}
	this.Maps[key] = originalVal
	current := ""
	for _, r := range key {
		current += string(r)
		this.Memo[current] += val
	}
}

func (this *MapSum) Sum(prefix string) int {
	return this.Memo[prefix]
}

func main() {
	obj := Constructor()

	// result: [null,null,3,null,5,null,null,9]
	// obj.Insert("apple", 3)
	// fmt.Printf("obj.Sum() = %v\n", obj.Sum("ap"))
	// obj.Insert("app", 2)
	// fmt.Printf("obj.Sum() = %v\n", obj.Sum("ap"))
	// obj.Insert("sub", 3)
	// obj.Insert("abc", 4)
	// fmt.Printf("obj.Sum() = %v\n", obj.Sum("a"))

	// result: 
	obj.Insert("apple", 3)
	fmt.Printf("obj.Sum() = %v\n", obj.Sum("ap"))
	obj.Insert("app", 2)
	fmt.Printf("obj.Sum() = %v\n", obj.Sum("ap"))
	obj.Insert("apple", 2)
	fmt.Printf("obj.Sum() = %v\n", obj.Sum("ap"))

	// obj.Insert("", )
	// fmt.Printf("obj.Sum() = %v\n", obj.Sum(""))
}

