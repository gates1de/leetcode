package heap

func (h CompHeap) Len() int           { return len(h) }
func (h CompHeap) Less(a, b int) bool { return h[a].Count < h[b].Count }
func (h CompHeap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *CompHeap) Push(x interface{}) {
	*h = append(*h, x.(Comp))
}

func (h *CompHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Comp struct {
	Char byte
	Count int
}

type CompHeap []Comp

