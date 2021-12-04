package main
import (
	"fmt"
)

type StreamChecker struct {
	root *TrieNode
	sb   []byte
}

type TrieNode struct {
	isWord bool
	next   [26]*TrieNode
}

func (sc *StreamChecker) createTrie(words []string) {
	for _, s := range words {
		node := sc.root
		ln := len(s)
		for i := ln - 1; i >= 0; i-- {
			c := s[i]
			if node.next[c-'a'] == nil {
				node.next[c-'a'] = &TrieNode{}
			}
			node = node.next[c-'a']
		}
		node.isWord = true
	}
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{}
	sc.root = &TrieNode{}
	sc.sb = []byte{}
	sc.createTrie(words)
	return sc
}

func (this *StreamChecker) Query(letter byte) bool {

	(*this).sb = append((*this).sb, letter)
	node := this.root
	for i := len((*this).sb) - 1; i >= 0 && node != nil; i-- {
		c := (*this).sb[i]
		node = node.next[c-'a']
		if node != nil && node.isWord {
			return true
		}
	}
	return false
}

// Accepted, but slow & use more memory
// type StreamChecker struct {
// 	Words []string
// 	Stream string
// }
// 
// func Constructor(words []string) StreamChecker {
// 	sc := StreamChecker{words, ""}
// 	return sc
// }
// 
// func (this *StreamChecker) Query(letter byte) bool {
// 	this.Stream += string(letter)
// 	for _, word := range this.Words {
// 		if len(word) > len(this.Stream) {
// 			continue
// 		}
// 		s := this.Stream[len(this.Stream) - len(word):]
// 		if s == word {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	// result: [null, false, false, false, true, false, true, false, false, false, false, false, true]
	words := []string{"cd", "f", "kl"}
	obj := Constructor(words)
	fmt.Printf("query = %v\n", obj.Query('a'))
	fmt.Printf("query = %v\n", obj.Query('b'))
	fmt.Printf("query = %v\n", obj.Query('c'))
	fmt.Printf("query = %v\n", obj.Query('d'))
	fmt.Printf("query = %v\n", obj.Query('e'))
	fmt.Printf("query = %v\n", obj.Query('f'))
	fmt.Printf("query = %v\n", obj.Query('g'))
	fmt.Printf("query = %v\n", obj.Query('h'))
	fmt.Printf("query = %v\n", obj.Query('i'))
	fmt.Printf("query = %v\n", obj.Query('j'))
	fmt.Printf("query = %v\n", obj.Query('k'))
	fmt.Printf("query = %v\n", obj.Query('l'))

	// result: 
	// words := []string{}
	// obj := Constructor(words)
	// fmt.Printf("query = %v\n", obj.Query(''))
}

