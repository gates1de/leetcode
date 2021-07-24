package main
import (
	"fmt"
	"reflect"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    var queue [][]string
    queue = append(queue, []string{beginWord})
    visited := make(map[string]bool)
    end := false
    for len(queue) > 0 && !end {
        l := len(queue)
        for i := 0; i < l; i++ {
            lastWord := queue[i][len(queue[i]) - 1]
            for j := 0; j < len(wordList); j++ {
                if canTransformed(lastWord, wordList[j]) {
                    if wordList[j] == endWord {
                        end = true
                    }
                    newPath := append([]string{}, queue[i]...)
                    newPath = append(newPath, wordList[j])
                    queue = append(queue, newPath)
                    visited[wordList[j]] = true
                }
            }
        }

        var newWordList []string
        for i := 0; i < len(wordList); i++ {
            if !visited[wordList[i]] {
                newWordList = append(newWordList, wordList[i])
            }
        }
        wordList = newWordList

        queue = queue[l:]
    }
    var result [][]string
    for _, v := range queue {
        if v[len(v) - 1] == endWord {
            result = append(result, v)
        }
    }
    return result
}

func canTransformed(a string, b string)  bool {
    diff := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            diff++
        }
    }
    if diff == 1 {
        return true
    } else {
        return false
    }
}

type Node struct {
    visited bool
    depth int
    val string
    nextNodeList []*Node
	path []string
}

// Wrong Answer
func ngSolution(beginWord string, endWord string, wordList []string) [][]string {
	result := [][]string{}
    if !contains(endWord, wordList) ||
        reflect.DeepEqual(wordList, []string{beginWord, endWord}) {
        return result
    }
    if diff(beginWord, endWord) == 1 {
		result = append(result, []string{beginWord, endWord})
        return result
    }

    graph := makeGraph(beginWord, endWord, wordList)
    // fmt.Printf("graph val = %v, nextNodeList = %v\n", graph.val, graph.nextNodeList[0])
    result = bfs(beginWord, endWord, graph)
    return result
}

func bfs(beginWord string, endWord string, node *Node) [][]string {
    queue := []*Node{}
    queue = append(queue, node)
	node.path = []string{beginWord}
	minDepth := int(1000)
	paths := [][]string{}
    for len(queue) != 0 {
        currentNode := queue[0]
        queue = queue[1:]
        // fmt.Printf("currentNode.val = %v, depth = %v\n", currentNode.val, currentNode.depth)
        currentNode.visited = true
        for _, n := range currentNode.nextNodeList {
            if n.depth != 0 && n.val != endWord {
                continue
            }
            n.depth = currentNode.depth + 1
			n.path = make([]string, len(currentNode.path) + 1)
			copy(n.path, currentNode.path)
			n.path[len(n.path) - 1] = n.val
            // fmt.Printf("pushed node = %v, n.path = %v\n", n.val, n.path)
			if n.val == endWord {
				depth := n.depth + 1
				if minDepth > depth {
					paths = [][]string{n.path}
					minDepth = depth
				} else if minDepth == depth {
					paths = append(paths, n.path)
				}
				continue
			}
            queue = append(queue, n)
        }
    }
	return paths
}

func appendNextNodeList(targetNode *Node, nodeList []*Node, visited map[string]bool) {
    nextNodeList := []*Node{}
    targetNodeVal := targetNode.val
    // fmt.Printf("targetNodeVal = %v\n", targetNodeVal)
    for _, node := range nodeList {
        if diff(targetNodeVal, node.val) != 1 {
            continue
        }
        // fmt.Printf("append node.val = %v\n", node.val)
        nextNodeList = append(nextNodeList, node)
    }

    for _, nextNode := range nextNodeList {
        if visited[nextNode.val] {
            // fmt.Printf("already visited = %v\n", nextNode.val)
            continue
        }
        visited[nextNode.val] = true
        appendNextNodeList(nextNode, nodeList, visited)
    }
	// fmt.Printf("%v: nextNodeList = %v\n", targetNodeVal, nextNodeList)
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
	// result: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log","cog"}

	// result: []
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log"}

	// result: [["hot","dot","dog"],["hot","hog","dog"]]
	beginWord := "hot"
	endWord := "dog"
	wordList := []string{"hot","cog","dog","tot","hog","hop","pot","dot"}

	// result: 
	// beginWord := ""
	// endWord := ""
	// wordList := []string{}

	result := findLadders(beginWord, endWord, wordList)
	fmt.Printf("result = %v\n", result)
}

