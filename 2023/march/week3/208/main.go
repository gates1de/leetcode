package main
import (
	"fmt"
)

type Trie struct {
	Word string
	Nodes map[rune]*Trie
}

func Constructor() Trie {
	return Trie{Word: "", Nodes: make(map[rune]*Trie, 26)}
}

func (this *Trie) Insert(word string)  {
    for _, char := range word {
        t := Trie{}
        if this.Nodes == nil {
            this.Nodes = make(map[rune]*Trie)
        }
        if this.Nodes[char] == nil {
            this.Nodes[char] = &t
        }
        this = this.Nodes[char]
    }
    this.Word = word
}

func (this *Trie) Search(word string) bool {
    for _, char := range word {
        if this.Nodes[char] == nil {
            return false
        }
        this = this.Nodes[char]
    }
    return this.Word == word
}

func (this *Trie) StartsWith(prefix string) bool {
    for _, char := range prefix {
        if this.Nodes[char] == nil {
            return false
        }
        this = this.Nodes[char]
    }
    return true
}

func main() {
	// result: [null, null, true, false, true, null, true]
	operations := [][]string{
		{"0", "apple"},
		{"1", "apple"},
		{"1", "app"},
		{"2", "app"},
		{"0", "app"},
		{"1", "app"},
	}

	// result: 
	// operations := [][]string{
	// 	{"0", "insert"},
	// 	{"1", "search"},
	// 	{"2", "startWith"},
	// }

	obj := Constructor()

	for _, val := range operations {
		operation := val[0]
		word := val[1]
		if operation == "0" {
			obj.Insert(word)
		} else if operation == "1" {
			fmt.Printf("obj.Search(%v) = %v\n", word, obj.Search(word))
		} else if operation == "2" {
			fmt.Printf("obj.SearchWith(%v) = %v\n", word, obj.StartsWith(word))
		}
	}
}

