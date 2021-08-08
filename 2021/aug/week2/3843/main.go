package main
import (
	"fmt"
	"container/heap"
)

type Cell struct {
    Val int
    Row int
	Col int
	Rank int
}
type RankHeap []Cell

func (h *RankHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *RankHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *RankHeap) Len() int           { return len(*h) }
func (h *RankHeap) Pop() (v interface{}) {
    *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return v
}
func (h *RankHeap) Push(v interface{}) { *h = append(*h, v.(Cell)) }

func matrixRankTransform(matrix [][]int) [][]int {
	answer := make([][]int, len(matrix))
	rankHeap := &RankHeap{}
	heap.Init(rankHeap)
	m := map[int][]Cell{}

	for i, _ := range answer {
		answer[i] = make([]int, len(matrix[0]))
	}

	for i, row := range matrix {
		for j, num := range row {
			heap.Push(rankHeap, Cell{Val: num, Row: i, Col: j})
		}
	}

	for len(*rankHeap) > 0 {
		cell := heap.Pop(rankHeap).(Cell)
		rank := int(1)
		for i, otherRank := range answer[cell.Row] {
			if i == cell.Col {
				continue
			}

			otherVal := matrix[cell.Row][i]
			if cell.Val == otherVal {
				rank = max(rank, otherRank)
			}
			if cell.Val > otherVal {
				if otherRank > 0 {
					rank = max(rank, otherRank + 1)
				}
			}
		}
		for j, row := range answer {
			if j == cell.Row {
				continue
			}

			otherRank := row[cell.Col]
			otherVal := matrix[j][cell.Col]
			if cell.Val == otherVal {
				rank = max(rank, otherRank)
			}
			if cell.Val > otherVal {
				if otherRank > 0 {
					rank = max(rank, otherRank + 1)
				}
			}
		}
		cell.Rank = rank
		if m[cell.Val] == nil {
			m[cell.Val] = []Cell{cell}
		} else {
			for _, c := range m[cell.Val] {
				if c.Row != cell.Row || c.Col != c.Col {
					continue
				}
				if c.Rank < cell.Rank {
					c.Rank = cell.Rank
					answer[c.Row][c.Col] = cell.Rank
				} else if c.Rank > cell.Rank {
					cell.Rank = c.Rank
				}
			}
			m[cell.Val] = append(m[cell.Val], cell)
		}
		answer[cell.Row][cell.Col] = cell.Rank
		// fmt.Printf("cell = %v, answer = %v\n", cell, answer)
	}

	return answer
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [[1,2],[2,3]]
	// matrix := [][]int{{1, 2}, {3, 4}}

	// result: [[1,1],[1,1]]
	// matrix := [][]int{{7, 7}, {7, 7}}

	// result: [[4,2,3],[1,3,4],[5,1,6],[1,3,4]]
	// matrix := [][]int{
	// 	{20, -21, 14},
	// 	{-19, 4, 19},
	// 	{22, -47, 24},
	// 	{-19, 4, 19},
	// }

	// result: [[5,1,4],[1,2,3],[6,3,1]]
	// matrix := [][]int{
	// 	{7, 3, 6},
	// 	{1, 4, 5},
	// 	{9, 8, 2},
	// }

	// result: [[10,3,4,9,5,15,8],[12,4,2,10,1,13,14],[11,13,9,8,6,7,12],[2,10,15,4,9,3,15],[1,2,17,16,7,15,3],[5,14,18,11,10,8,4],[3,15,5,6,8,14,7]]
	matrix := [][]int{
		{-2,-35,-32,-5,-30,33,-12},
		{7,2,-43,4,-49,14,17},
		{4,23,-6,-15,-24,-17,6},
		{-47,20,39,-26,9,-44,39},
		{-50,-47,44,43,-22,33,-36},
		{-13,34,49,24,23,-2,-35},
		{-40,43,-22,-19,-4,23,-18},
	}

	// result: 
	// matrix := [][]int{}

	result := matrixRankTransform(matrix)
	fmt.Printf("result = %v\n", result)
}

