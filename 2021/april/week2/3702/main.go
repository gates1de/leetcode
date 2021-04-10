package main
import (
	"container/heap"
	"fmt"
)

type Heap []string

func (h Heap) Len() int           { return len(h) }

func (h Heap) Less(a, b int) bool { return h[a] > h[b] }

func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type AlianHeap struct {
	Heap
	OrderMap map[rune]int
}

func (h AlianHeap) Less(i, j int) bool {
	for index, iRune := range h.Heap[i] {
		if index >= len(h.Heap[j]) {
			return false
		}
		jRune := rune(h.Heap[j][index])
		// fmt.Printf("i = %v, j = %v, i order = %v, j order = %v\n", string(iRune), string(jRune), h.OrderMap[iRune], h.OrderMap[jRune])
		if h.OrderMap[iRune] < h.OrderMap[jRune] {
			return true
		} else if h.OrderMap[iRune] > h.OrderMap[jRune] {
			return false
		}
	}
	return true
}

func isAlienSorted(words []string, order string) bool {
	m := map[rune]int{}
	for i, r := range order {
		m[r] = i
	}
	h := &AlianHeap{OrderMap: m}
	heap.Init(h)
	for _, word := range words {
		heap.Push(h, word)
	}
	// fmt.Printf("h = %v\n", h.Heap)
	return deepEqual(h.Heap, words)
}

func deepEqual(a, b []string) bool {
	for i, aWord := range a {
		bWord := b[i]
		if aWord != bWord {
			return false
		}
	}
	return true
}

func main() {
	// result: true
	// words := []string{"hello", "leetcode"}
	// order := "hlabcdefgijkmnopqrstuvwxyz"

	// result: false
	// words := []string{"word", "world", "row"}
	// order := "worldabcefghijkmnpqstuvxyz"

	// result: false
	// words := []string{"apple", "app"}
	// order := "abcdefghijklmnopqrstuvwxyz"

	// result: 
	words := []string{"iekm", "tpnhnbe"}
	order := "loxbzapnmstkhijfcuqdewyvrg"

	// result: 
	// words := []string{}
	// order := ""

	result := isAlienSorted(words, order)
	fmt.Printf("result = %v\n", result)
}

