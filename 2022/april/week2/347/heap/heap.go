package heap

type Counter struct {
	Num int
	Count int
}

type Heap []Counter

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a].Count > h[b].Count }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Counter))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

