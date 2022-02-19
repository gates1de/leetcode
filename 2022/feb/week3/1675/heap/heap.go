package heap

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a] < h[b] }
func (h Heap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap struct {
	Heap
}

func (h MaxHeap) Less(i, j int) bool { return h.Heap[i] > h.Heap[j] }

