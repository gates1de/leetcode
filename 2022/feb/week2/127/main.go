package main
import (
	"fmt"
	"reflect"
)

type Node struct {
	visited bool
	depth int
	val string
	nextNodeList []*Node
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(endWord, wordList) ||
		reflect.DeepEqual(wordList, []string{beginWord, endWord}) {
		return 0
	}
	if diff(beginWord, endWord) == 1 {
		return 2
	}

	graph := makeGraph(beginWord, endWord, wordList)
	result := bfs(beginWord, endWord, graph)
	return result
}

func bfs(beginWord string, endWord string, node *Node) int {
	queue := []*Node{}
	queue = append(queue, node)
	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode.visited {
			continue
		}
		currentNode.visited = true
		if currentNode.val == endWord {
			return currentNode.depth + 1
		}
		for _, n := range currentNode.nextNodeList {
			if n.depth != 0 {
				continue
			}
			n.depth = currentNode.depth + 1
			queue = append(queue, n)
		}
	}
	return 0
}

func appendNextNodeList(targetNode *Node, nodeList []*Node, visited map[string]bool) {
	nextNodeList := []*Node{}
	targetNodeVal := targetNode.val
	for _, node := range nodeList {
		if diff(targetNodeVal, node.val) != 1 {
			continue
		}
		nextNodeList = append(nextNodeList, node)
	}

	for _, nextNode := range nextNodeList {
		if visited[nextNode.val] {
			continue
		}
		visited[nextNode.val] = true
		appendNextNodeList(nextNode, nodeList, visited)
	}
	targetNode.nextNodeList = nextNodeList
}

func makeGraph(beginWord string, endWord string, wordList []string) *Node {
	nodeList := make([]*Node, 0, len(wordList))
	beginNode := &Node{visited: false, depth: 0, val: beginWord}
	nodeList = append(nodeList, beginNode)
	for _, word := range wordList {
		node := &Node{visited: false, depth: 0, val: word}
		nodeList = append(nodeList, node)
	}

	appendNextNodeList(beginNode, nodeList, map[string]bool{})
	return beginNode
}

func copySlice(target []string) []string {
	result := make([]string, len(target))
	copy(result, target)
	return result
}

func contains(word string, wordList []string) bool {
	for _, w := range wordList {
		if word == w {
			return true
		}
	}
	return false
}

func diff(word1 string, word2 string) int {
	count := int(0)
	for i, _ := range word1 {
		if word1[i] != word2[i] {
			count++
		}
	}
	return count
}

func main() {
	// result: 5
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log","cog"}

	// result: 0
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log"}

	// result: 0
	// beginWord := "hotdog"
	// endWord := "hatbot"
	// wordList := []string{""}

	// result: 2
	beginWord := "h"
	endWord := "g"
	wordList := []string{"h","a","g"}

	// result: 
	// beginWord := ""
	// endWord := ""
	// wordList := []string{}

	result := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("result = %v\n", result)
}

