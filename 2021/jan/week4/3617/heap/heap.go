package heap

type Heap [][]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a][2] < h[b][2] }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }
func (h Heap) Peek() []int        { return h[0] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

