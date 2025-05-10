package main
import (
	"fmt"
	"math"
	heap "container/heap"
)

type State struct {
	X int
	Y int
	Distance int
}

type PriorityQueue []State

func (this PriorityQueue) Len() int {
	return len(this)
}

func (this PriorityQueue) Less(i, j int) bool {
	return this[i].Distance < this[j].Distance
}

func (this PriorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PriorityQueue) Push(x interface{}) {
	*this = append(*this, x.(State))
}

func (this *PriorityQueue) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n - 1]
	*this = old[:n - 1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])
	d := make([][]int, n)
	v := make([][]bool, n)

	for i := range d {
		d[i] = make([]int, m)
		v[i] = make([]bool, m)
		for j := range d[i] {
			d[i][j] = math.MaxInt32
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	d[0][0] = 0
	pq := &PriorityQueue{}
	heap.Push(pq, State{0, 0, 0})

	for pq.Len() > 0 {
		s := heap.Pop(pq).(State)

		if v[s.X][s.Y] {
			continue
		}

		v[s.X][s.Y] = true

		for _, dir := range dirs {
			nx, ny := s.X + dir[0], s.Y + dir[1]

			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}

			dist := max(d[s.X][s.Y], moveTime[nx][ny]) + 1
			if d[nx][ny] > dist {
				d[nx][ny] = dist
				heap.Push(pq, State{nx, ny, dist})
			}
		}
	}

	return d[n - 1][m - 1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// moveTime := [][]int{{0,4},{4,4}}

	// result: 3
	// moveTime := [][]int{{0,0,0},{0,0,0}}

	// result: 3
	moveTime := [][]int{{0,1},{1,2}}

	// result: 
	// moveTime := [][]int{}

	result := minTimeToReach(moveTime)
	fmt.Printf("result = %v\n", result)
}

