package main
import (
	"fmt"
	heap "container/heap"
	"math"
	"sort"
)

type Worker struct {
    Quality int
    Wage int
    WagePerQuality float64  
}

type Workers []Worker

func(w Workers) Len() int {
    return len(w)
}

func(w Workers) Less(i, j int) bool {
    return w[i].WagePerQuality < w[j].WagePerQuality
}

func (w Workers) Swap(i, j int) {
    w[i], w[j] = w[j], w[i]
}

type WorkerHeap []int

func (w WorkerHeap) Len() int           { return len(w) }
func (w WorkerHeap) Less(i, j int) bool { return w[i] < w[j] }
func (w WorkerHeap) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w *WorkerHeap) Push(x interface{}) { *w = append(*w, x.(int)) }
func (w *WorkerHeap) Pop() interface{} {
    x := (*w)[len(*w)-1]
    *w = (*w)[:len(*w)-1]
    return x
}

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
    n := len(quality)
    workers := make(Workers, n)
    
	for i, _ := range workers {
        workers[i] = Worker{ 
            Quality: quality[i], 
            Wage: wage[i], 
            WagePerQuality: float64(wage[i]) / float64(quality[i]),
        }
    }

    sort.Sort(workers)

    result := math.MaxFloat64
    sumQuality := int(0)
    wh := &WorkerHeap{}
    
	for _, worker := range workers{
        heap.Push(wh, -worker.Quality)
        sumQuality += worker.Quality
        if wh.Len() > K {
            sumQuality += heap.Pop(wh).(int)
        }
        if wh.Len() == K {
            result = min(result, float64(sumQuality) * worker.WagePerQuality)
        }
    }

    return result
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func main() {
	// result: 105.00000
	// quality := []int{10,20,5}
	// wage := []int{70,50,30}
	// k := int(2)

	// result: 30.66667
	quality := []int{3,1,10,10,1}
	wage := []int{4,8,2,2,7}
	k := int(3)

	// result: 
	// quality := []int{}
	// wage := []int{}
	// k := int(0)

	result := mincostToHireWorkers(quality, wage, k)
	fmt.Printf("result = %v\n", result)
}

