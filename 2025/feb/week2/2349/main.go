package main
import (
	"fmt"
	heap "container/heap"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Peek() int          { return h[0] }

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

type NumberContainers struct {
	NumberToIndices map[int]*Heap
	IndexToNumbers map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{NumberToIndices: make(map[int]*Heap), IndexToNumbers: make(map[int]int)}
}

func (this *NumberContainers) Change(index int, number int)  {
	this.IndexToNumbers[index] = number
	if this.NumberToIndices[number] == nil {
    h := &Heap{}
    heap.Init(h)
		this.NumberToIndices[number] = h
	}

	heap.Push(this.NumberToIndices[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if this.NumberToIndices[number] == nil {
		return -1
	}

	minHeap := this.NumberToIndices[number]

	for minHeap.Len() > 0 {
		index := minHeap.Peek()
		if this.IndexToNumbers[index] == number {
			return index
		}

		heap.Pop(minHeap)
	}

	return -1
}

func main() {
	// result: [null, -1, null, null, null, null, 1, null, 2]
	operations := [][]int{{0,0,0},{1,0,10},{2,2,10},{2,1,10},{2,3,10},{2,5,10},{1,0,10},{2,1,20},{1,0,10}}

	obj := Constructor()
	for _, ope := range operations {
		function := ope[0]
		index := ope[1]
		val := ope[2]

		if function == 1 {
			fmt.Printf("obj.Find(%d) = %d\n", val, obj.Find(val))
		} else if function == 2 {
			obj.Change(index, val)
		}
	}
}

