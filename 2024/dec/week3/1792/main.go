package main
import (
	"fmt"
	heap "container/heap"
)

type Heap [][]float64

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a][0] > h[b][0]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]float64))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := &Heap{}
	heap.Init(pq)

	for _, class := range classes {
		passes := class[0]
		totalStudents := class[1]
		gain := calculateGain(passes, totalStudents)
		heap.Push(pq, []float64{gain, float64(passes), float64(totalStudents)})
	}

	for extraStudents > 0 {
		current := heap.Pop(pq).([]float64)
		passes := int(current[1])
		totalStudents := int(current[2])

		heap.Push(pq, []float64{calculateGain(passes + 1, totalStudents + 1), float64(passes + 1), float64(totalStudents + 1)})
		extraStudents--
	}

	totalPassRatio := float64(0)
	for pq.Len() > 0 {
		current := heap.Pop(pq).([]float64)
		totalPassRatio += current[1] / current[2]
	}

	return totalPassRatio / float64(len(classes))
}

func calculateGain(passes int, totalStudents int) float64 {
	return float64(passes + 1) / float64(totalStudents + 1) - float64(passes) / float64(totalStudents)
}

func main() {
	// result: 0.78333
	classes := [][]int{{1,2},{3,5},{2,2}}
	extraStudents := int(2)

	// result: 0.53485
	// classes := [][]int{{2,4},{3,9},{4,5},{2,10}}
	// extraStudents := int(4)

	// result: 
	// classes := [][]int{}
	// extraStudents := int(0)

	result := maxAverageRatio(classes, extraStudents)
	fmt.Printf("result = %v\n", result)
}

