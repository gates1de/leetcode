package main
import (
	"fmt"
	heap "container/heap"
)

type Heap []byte

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a] > h[b]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(byte))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func repeatLimitedString(s string, repeatLimit int) string {
	freq := make(map[byte]int)
	for _, char := range s {
		freq[byte(char)]++
	}

	pq := &Heap{}
	heap.Init(pq)

	for char, _ := range freq {
		heap.Push(pq, char)
	}

	result := make([]byte, 0, len(s))

	for pq.Len() > 0 {
		current := heap.Pop(pq).(byte)
		count := freq[current]

		use := min(count, repeatLimit)
		for i := 0; i < use; i++ {
			result = append(result, current)
		}

		freq[current] = count - use

		if freq[current] > 0 && pq.Len() > 0 {
			next := heap.Pop(pq).(byte)
			result = append(result, next)
			freq[next]--

			if freq[next] > 0 {
				heap.Push(pq, next)
			}

			heap.Push(pq, current)
		}
	}

	return string(result)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: "zzcccac"
	// s := "cczazcc"
	// repeatLimit := int(3)

	// result: "bbabaa"
	s := "aababab"
	repeatLimit := int(2)

	// result: ""
	// s := ""
	// repeatLimit := int(0)

	result := repeatLimitedString(s, repeatLimit)
	fmt.Printf("result = %v\n", result)
}

