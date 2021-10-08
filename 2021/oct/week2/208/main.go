package main
import (
	"fmt"
)

type Trie struct {
	Word string
	Nodes map[rune]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string)  {
	for _, c := range word {
		// fmt.Printf("c = %v, this.Nodes[c] = %v\n", c, this.Nodes[c])
		t := Trie{}
		if this.Nodes == nil {
			this.Nodes = map[rune]*Trie{}
		}
		if this.Nodes[c] == nil {
			this.Nodes[c] = &t
		}
		this = this.Nodes[c]
	}
	this.Word = word
}

func (this *Trie) Search(word string) bool {
	// fmt.Printf("this = %v\n", this)
	for _, c := range word {
		// fmt.Printf("c = %v, this.Nodes[c] = %v\n", c, this.Nodes[c])
		if this.Nodes[c] == nil {
			return false
		}
		this = this.Nodes[c]
	}

	return this.Word == word
}


func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.Nodes[c] == nil {
			return false
		}
		this = this.Nodes[c]
	}

	return true
}

func main() {
	obj := Constructor()

	// obj.Insert("apple")
	// fmt.Printf("obj.Search(\"apple\") = %v\n", obj.Search("apple")) // true
	// fmt.Printf("obj.Search(\"app\") = %v\n", obj.Search("app")) // false
	// fmt.Printf("obj.StartsWith(\"app\") = %v\n", obj.StartsWith("app")) // true
	// obj.Insert("app")
	// fmt.Printf("obj.Search(\"app\") = %v\n", obj.Search("app")) // true

	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	fmt.Printf("obj.Search(\"apps\") = %v\n", obj.Search("apps")) // false
	fmt.Printf("obj.Search(\"app\") = %v\n", obj.Search("app")) // true
	fmt.Printf("obj.Search(\"\") = %v\n", obj.Search("ad")) // false
	fmt.Printf("obj.Search(\"\") = %v\n", obj.Search("applepie")) // false
	fmt.Printf("obj.Search(\"\") = %v\n", obj.Search("rest")) // false

	// obj.Insert("")
	// fmt.Printf("obj.Search(\"\") = %v\n", obj.Search(""))
	// fmt.Printf("obj.Search(\"\") = %v\n", obj.StartsWith(""))
}

