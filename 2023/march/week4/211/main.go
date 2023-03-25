package main
import (
	"fmt"
)


type WordDictionary struct {
    isWord   bool
    Next [26]*WordDictionary
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
    node := this
    for _, char := range word {
        if node.Next[char - 'a'] == nil {
            node.Next[char - 'a'] = &WordDictionary{}
        }
        node = node.Next[char - 'a']
    }
    node.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
    node := this
    for i, char := range word {
        if char == '.' {
            for _, node := range node.Next {
                if node == nil {
                    continue
                }
                if node.Search(word[i+1:]) {
                    return true
                }
            }
            return false
        }
        if node.Next[char - 'a'] == nil {
            return false
        }
        node = node.Next[char - 'a']
    }
    return node.isWord
}

func main() {
	// result: [null,null,null,null,false,true,true,true]
	operations := [][]string{
		{"0", "bad"},
		{"0", "dad"},
		{"0", "mad"},
		{"1", "pad"},
		{"1", "bad"},
		{"1", ".ad"},
		{"1", "b.."},
	}

	// result: 
	// operations := [][]string{
	// 	{"0", "add"},
	// 	{"1", "search"},
	// }

	obj := Constructor()

	for _, val := range operations {
		operation := val[0]
		word := val[1]
		if operation == "0" {
			obj.AddWord(word)
		} else if operation == "1" {
			fmt.Printf("obj.Search(%v) = %v\n", word, obj.Search(word))
		}
	}
}

