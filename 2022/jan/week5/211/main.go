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

// Wrong Answer
// type WordDictionary struct {
// 	Root *TrieNode
// }
// 
// type TrieNode struct {
// 	Char rune
// 	Next [27]*TrieNode
// }
// 
// func Constructor() WordDictionary {
// 	return WordDictionary{Root: &TrieNode{Char: rune('.')}}
// }
// 
// 
// func (this *WordDictionary) AddWord(word string) {
// 	root := this.Root
// 	for _, char := range word {
// 		if root.Next[char - 'a'] == nil {
// 			root.Next[char - 'a'] = &TrieNode{Char: char}
// 		}
// 		root = root.Next[char - 'a']
// 	}
// 	root.Next[26] = &TrieNode{Char: rune('X')}
// }
// 
// 
// func (this *WordDictionary) Search(word string) bool {
// 	nodes := []*TrieNode{this.Root}
// 
// 	TOP:
// 	for i, char := range word {
// 		isLastChar := i == len(word) - 1
// 		if len(nodes) == 0 {
// 			return false
// 		}
// 
// 		if char == '.' {
// 			newNodes := []*TrieNode{}
// 			for _, node := range nodes {
// 				for _, n := range node.Next {
// 					if n == nil {
// 						continue
// 					}
// 					if isLastChar && n.Next[26] != nil {
// 						return true
// 					}
// 					newNodes = append(newNodes, n)
// 				}
// 			}
// 			nodes = newNodes
// 			continue
// 		}
// 
// 		for len(nodes) > 0 {
// 			node := nodes[0]
// 			nodes = nodes[1:]
// 
// 			if node.Next[char - 'a'] != nil {
// 				if isLastChar && node.Next[char - 'a'].Next[26] != nil {
// 					return true
// 				}
// 
// 				nodes = []*TrieNode{node.Next[char - 'a']}
// 				continue TOP
// 			}
// 		}
// 
// 		return false
// 	}
// 
// 	// fmt.Println(string(nodes[0].Char))
// 	return len(nodes) == 0
// }
// 
// func printNodes(root *TrieNode) {
// 	nodes := []*TrieNode{root}
// 	for len(nodes) > 0 {
// 		node := nodes[0]
// 		nodes = nodes[1:]
// 		for _, n := range node.Next {
// 			if n == nil {
// 				continue
// 			}
// 			fmt.Printf("node.Char = %v\n", string(n.Char))
// 			nodes = append(nodes, n)
// 		}
// 		fmt.Println("--------------")
// 	}
// }

func main() {
	obj := Constructor()

	// result: [null,null,null,null,false,true,true,true]
	// obj.AddWord("bad")
	// obj.AddWord("dad")
	// obj.AddWord("mad")
	// fmt.Printf("obj.Search(\"pad\") = %v\n", obj.Search("pad"))
	// fmt.Printf("obj.Search(\"bad\") = %v\n", obj.Search("bad"))
	// fmt.Printf("obj.Search(\".ad\") = %v\n", obj.Search(".ad"))
	// fmt.Printf("obj.Search(\"b.\") = %v\n", obj.Search("b.."))

	// result: [null,null,null,null,null,null,true,true,false,true,true,false,true,false,true,true,true,true,true,true,true,true]
	// obj.AddWord("b")
	// obj.AddWord("ba")
	// obj.AddWord("bad")
	// obj.AddWord("dad")
	// obj.AddWord("mad")
	// fmt.Printf("obj.Search(\".\") = %v\n", obj.Search("."))
	// fmt.Printf("obj.Search(\"b\") = %v\n", obj.Search("b"))
	// fmt.Printf("obj.Search(\"b\") = %v\n", obj.Search("a"))
	// fmt.Printf("obj.Search(\"..\") = %v\n", obj.Search(".."))
	// fmt.Printf("obj.Search(\"b.\") = %v\n", obj.Search("b."))
	// fmt.Printf("obj.Search(\"a.\") = %v\n", obj.Search("a."))
	// fmt.Printf("obj.Search(\".a\") = %v\n", obj.Search(".a"))
	// fmt.Printf("obj.Search(\".d\") = %v\n", obj.Search(".d"))
	// fmt.Printf("obj.Search(\"...\") = %v\n", obj.Search("..."))
	// fmt.Printf("obj.Search(\"b..\") = %v\n", obj.Search("b.."))
	// fmt.Printf("obj.Search(\".a.\") = %v\n", obj.Search(".a."))
	// fmt.Printf("obj.Search(\"..d\") = %v\n", obj.Search("..d"))
	// fmt.Printf("obj.Search(\"ba.\") = %v\n", obj.Search("ba."))
	// fmt.Printf("obj.Search(\"b.d\") = %v\n", obj.Search("b.d"))
	// fmt.Printf("obj.Search(\".ad\") = %v\n", obj.Search(".ad"))
	// fmt.Printf("obj.Search(\"bad\") = %v\n", obj.Search("bad"))

	// result: [null,null,null,null,null,null,null,null,true,false,true,true,true,false,true,true,false,false]
	obj.AddWord("ran")
	obj.AddWord("rune")
	obj.AddWord("runner")
	obj.AddWord("runs")
	obj.AddWord("add")
	obj.AddWord("adds")
	obj.AddWord("adder")
	obj.AddWord("addee")
	fmt.Printf("obj.Search(\"r.n\") = %v\n", obj.Search("r.n"))
	fmt.Printf("obj.Search(\"ru.n.e\") = %v\n", obj.Search("ru.n.e"))
	fmt.Printf("obj.Search(\"add\") = %v\n", obj.Search("add"))
	fmt.Printf("obj.Search(\"add.\") = %v\n", obj.Search("add."))
	fmt.Printf("obj.Search(\"adde.\") = %v\n", obj.Search("adde."))
	fmt.Printf("obj.Search(\".an.\") = %v\n", obj.Search(".an."))
	fmt.Printf("obj.Search(\"...s\") = %v\n", obj.Search("...s"))
	fmt.Printf("obj.Search(\"....e.\") = %v\n", obj.Search("....e."))
	fmt.Printf("obj.Search(\".......\") = %v\n", obj.Search("......."))
	fmt.Printf("obj.Search(\"..n.r\") = %v\n", obj.Search("..n.r"))
	// obj.AddWord("")
	// fmt.Printf("obj.Search(\"\") = %v\n", obj.Search(""))
}

