package main
import (
	"fmt"
	"container/heap"
)

type Job struct {
    StartTime int
    EndTime int
    Profit int
}

type JobHeap []Job
func (h *JobHeap) Less(i, j int) bool { return (*h)[i].EndTime < (*h)[j].EndTime }
func (h *JobHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *JobHeap) Len() int           { return len(*h) }
func (h *JobHeap) Pop() (v interface{}) {
    *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return v
}
func (h *JobHeap) Peek() Job { return (*h)[0] }
func (h *JobHeap) Push(v interface{}) { *h = append(*h, v.(Job)) }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
    h := &JobHeap{}
    heap.Init(h)
    for i, s := range startTime {
        e := endTime[i]
        p := profit[i]
        heap.Push(h, Job{StartTime: s, EndTime: e, Profit: p})
    }

    jobs := make([]Job, len(startTime))
    i := int(0)
    for h.Len() > 0 {
        job := heap.Pop(h).(Job)
        jobs[i] = job
        i++
    }

    return helper(jobs)
}

func helper(jobs []Job) int {
    n := len(jobs)
    dp := make([]int, n)
    dp[0] = jobs[0].Profit

    for i := 1; i < n; i++ {
        profit := jobs[i].Profit
        l := search(jobs, i)
        if l != -1 {
            profit += dp[l]
        }
        dp[i] = max(profit, dp[i - 1])
    }

    return dp[n - 1]
}

func search(jobs []Job, index int) int {
    start := int(0)
    end := index - 1
    for start <= end {
        mid := (start + end) / 2
        if jobs[mid].EndTime <= jobs[index].StartTime {
            if jobs[mid + 1].EndTime <= jobs[index].StartTime {
                start = mid + 1
            } else {
                return mid
            }
        } else {
            end = mid - 1
        }
    }

    return -1
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func main() {
	// result: 120
	// startTime := []int{1, 2, 3, 3}
	// endTime := []int{3, 4, 5, 6}
	// profit := []int{50, 10, 40, 70}

	// result: 150
	// startTime := []int{1, 2, 3, 4, 6}
	// endTime := []int{3, 5, 10, 6, 9}
	// profit := []int{20, 20, 100, 70, 60}

	// result: 6
	startTime := []int{1, 1, 1}
	endTime := []int{2, 3, 4}
	profit := []int{5, 6, 4}

	// result: 
	// startTime := []int{}
	// endTime := []int{}
	// profit := []int{}

	result := jobScheduling(startTime, endTime, profit)
	fmt.Printf("result = %v\n", result)
}

