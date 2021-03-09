package heap

// h[][0]: sum of 1, h[][1]: mat index
type Heap [][]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(a, b int) bool {
	if h[a][0] == h[b][0] {
		return h[a][1] < h[b][1]
	}
	return h[a][0] < h[b][0]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

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

