package main
import (
	"fmt"
	"container/heap"
	"sort"
)

type Task struct {
	EnqueueTime int
	ProcessingTime int
	Index int
} 

type Heap []*Task

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool {
	if h[i].ProcessingTime == h[j].ProcessingTime {
		return h[i].Index < h[j].Index
	}
	return h[i].ProcessingTime < h[j].ProcessingTime
}

func (h Heap) Swap(i, j int) { 
    h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(*Task))
}
func (h *Heap) Pop() interface{} { 
    n := h.Len()
    old := *h
    x := old[n - 1]
    *h = old[:n - 1]
    return x
}

func getOrder(tasks [][]int) []int {
	h := &Heap{}
	heap.Init(h)

	indexedTasks := make([]*Task, 0, len(tasks))
	for i, task := range tasks {
		indexedTasks = append(indexedTasks, &Task{
			EnqueueTime: task[0],
			ProcessingTime: task[1],
			Index: i,
		})
	}
	sort.Slice(indexedTasks, func(a, b int) bool {
		return indexedTasks[a].EnqueueTime < indexedTasks[b].EnqueueTime
	})

	i := int(0)
	currentTime := indexedTasks[0].EnqueueTime
	result := make([]int, 0, len(tasks))
    for h.Len() > 0 || i < len(tasks) {
        for i < len(tasks) && indexedTasks[i].EnqueueTime <= currentTime {
            heap.Push(h, indexedTasks[i])
            i++
        } 

        if h.Len() == 0 && i < len(indexedTasks) {
            currentTime = indexedTasks[i].EnqueueTime
            continue
        }

        task := heap.Pop(h).(*Task)
        currentTime += task.ProcessingTime
        result = append(result, task.Index)
    }

	return result
}

func main() {
	// result: [0,2,3,1]
	// tasks := [][]int{{1,2},{2,4},{3,2},{4,1}}

	// result: [4,3,2,0,1]
	tasks := [][]int{{7,10},{7,12},{7,5},{7,4},{7,2}}

	// result: 
	// tasks := [][]int{}

	result := getOrder(tasks)
	fmt.Printf("result = %v\n", result)
}

