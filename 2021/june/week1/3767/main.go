package main
import (
	"fmt"
	"container/heap"
)


func openLock(deadends []string, target string) int {
	deads := make(map[string]struct{})
	for _, dead := range deadends {
		deads[dead] = struct{}{}
	}
	visited := make(map[string]struct{})
	queue := []string{"0000"}
	var level int
	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			val := queue[0]
			queue = queue[1:]
			if _, ok := deads[val]; ok {
				continue
			}
			if val == target {
				return level
			}
			for j := 0; j < 4; j++ {
				next := getNextOrPrevValue(val[:j], val[j+1:], val[j], false)
				if _, ok := visited[next]; !ok {
					queue = append(queue, next)
					visited[next] = struct{}{}
				}
				prev := getNextOrPrevValue(val[:j], val[j+1:], val[j], true)
				if _, ok := visited[prev]; !ok {
					queue = append(queue, prev)
					visited[prev] = struct{}{}
				}
			}
		}
		level++
	}

	return -1
}

func getNextOrPrevValue(prefix, postfix string, char byte, prev bool) string {
	if prev {
		if char == '0' {
			char = '9'
		} else {
			char--
		}
	} else {
		if char == '9' {
			char = '0'
		} else {
			char++
		}
	}
	return prefix + string(char) + postfix
}


type Node struct {
	Val string
	Cost int
}
type MinHeap []Node

func (h *MinHeap) Less(i, j int) bool { return (*h)[i].Cost < (*h)[j].Cost }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Pop() (v interface{}) {
    *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return v
}
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(Node)) }


var prevs map[rune]string
var nexts map[rune]string

// Very slow & Use more memory
func mySolution(deadends []string, target string) int {
	if contains("0000", deadends) {
		return -1
	}

	prevs = map[rune]string{
		rune('0'): "9",
		rune('1'): "0",
		rune('2'): "1",
		rune('3'): "2",
		rune('4'): "3",
		rune('5'): "4",
		rune('6'): "5",
		rune('7'): "6",
		rune('8'): "7",
		rune('9'): "8",
	}
	nexts = map[rune]string{
		rune('0'): "1",
		rune('1'): "2",
		rune('2'): "3",
		rune('3'): "4",
		rune('4'): "5",
		rune('5'): "6",
		rune('6'): "7",
		rune('7'): "8",
		rune('8'): "9",
		rune('9'): "0",
	}

	visited := map[string]bool{}
	m := map[string]int{}
	for i := 0; i < 10000; i++ {
		m[fmt.Sprintf("%04d", i)] = 9999
	}
	dijkstra("0000", visited, m, deadends, target)
	if m[target] == 9999 {
		return -1
	}
	return m[target]
}

func dijkstra(current string, visited map[string]bool, m map[string]int, deadends []string, target string) {
	minHeap := new(MinHeap)
	heap.Push(minHeap, Node{Val: "0000", Cost: 0})
	count := int(0)
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		current := node.Val
		if visited[current] {
			continue
		}

		m[current] = min(m[current], count)
		visited[current] = true
		count = m[current] + 1
		for i, digit := range current {
			prev := string(current[:i]) + prevs[digit] + string(current[i + 1:])
			next := string(current[:i]) + nexts[digit] + string(current[i + 1:])
			if !contains(prev, deadends) && !visited[prev] && m[prev] > m[current] + 1 {
				m[prev] = m[current] + 1
				heap.Push(minHeap, Node{Val: prev, Cost: count})
			}
			if !contains(next, deadends) && !visited[next] && m[next] > m[current] + 1 {
				m[next] = m[current] + 1
				heap.Push(minHeap, Node{Val: next, Cost: count})
			}
		}
	}
}

func contains(target string, list []string) bool {
	for _, s := range list {
		if target == s {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	// target := "0202"

	// result: 1
	// deadends := []string{"8888"}
	// target := "0009"

	// result: -1
	// deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	// target := "8888"

	// result: -1
	deadends := []string{"0000"}
	target := "8888"

	// result: 
	// deadends := []string{}
	// target := ""

	result := openLock(deadends, target)
	fmt.Printf("result = %v\n", result)
}

