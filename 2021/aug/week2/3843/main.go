package main
import (
	"fmt"
	"container/heap"
	"sort"
)

type Point struct {
    X int
    Y int
}

type UnionFind struct {
	Forest map[int]int
}

func NewUnion() UnionFind {
	return UnionFind{Forest: make(map[int]int)}
}

func (t *UnionFind) Add(x ...int) {
    for _, e := range x {
        t.Forest[e] = e
    }
}

func (t *UnionFind) Find(x int) int {
	if t.Forest[x] == x {
        return t.Forest[x]
	}
	// path compressed
    t.Forest[x] = t.Find(t.Forest[x])
    return t.Forest[x]
}

func (t *UnionFind) Union(x, y int) {
    t.Forest[t.Find(x)] = t.Find(y)
}

func (t *UnionFind) Group() map[int][]int {
	ret := make(map[int][]int)
	for k := range t.Forest {
		group := t.Find(k)
		ret[group] = append(ret[group], k)
	}
	return ret
}

func matrixRankTransform(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	rank := make([]int, m + n)

	// prepare matrix coordinate map
	coordMap := make(map[int][]Point)
	for y := range matrix {
		for x := range matrix[y] {
            coordMap[matrix[y][x]] = append(coordMap[matrix[y][x]], Point{X: x, Y: y})
        }
	}

	// sort map keys
	coordKeys := make([]int, 0, len(coordMap))
	for k := range coordMap {
		coordKeys = append(coordKeys, k)
	}
	sort.Ints(coordKeys)

	// iterate coordinate map keys
	for _, key := range coordKeys {
		dsu := NewUnion()

		for _, p := range coordMap[key] {
            dsu.Add(p.Y, p.X + n)
		}

		for _, p := range coordMap[key] {
			dsu.Union(p.Y, p.X + n)
		}

		for _, group := range dsu.Group() {
            max := rank[group[0]]
            for _, g := range group {
                if rank[g] > max {
                    max = rank[g]
                }
            }
			for _, g := range group {
				rank[g] = max + 1
			}
		}

		for _, p := range coordMap[key] {
			matrix[p.Y][p.X] = rank[p.Y]
        }
	}
	return matrix
}

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

// Wrong Answer
func ngSolution(matrix [][]int) [][]int {
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
	// matrix := [][]int{
	// 	{-2,-35,-32,-5,-30,33,-12},
	// 	{7,2,-43,4,-49,14,17},
	// 	{4,23,-6,-15,-24,-17,6},
	// 	{-47,20,39,-26,9,-44,39},
	// 	{-50,-47,44,43,-22,33,-36},
	// 	{-13,34,49,24,23,-2,-35},
	// 	{-40,43,-22,-19,-4,23,-18},
	// }

	// result: [[7,13,1,5,4,6,9,8],[8,11,2,10,1,12,14,9],[2,14,1,11,13,7,5,3],[3,19,16,12,14,7,10,13],[8,12,6,14,5,1,4,13],[2,16,15,17,4,18,3,14],[3,7,11,6,12,13,14,10],[16,19,18,3,15,2,11,17]]
    matrix := [][]int{
        {-23,20,-49,-30,-39,-28,-5,-14},
        {-19,4,-33,2,-47,28,43,-6},
        {-47,36,-49,6,17,-8,-21,-30},
        {-27,44,27,10,21,-8,3,14},
        {-19,12,-25,34,-27,-48,-37,14},
        {-47,40,23,46,-39,48,-41,18},
        {-27,-4,7,-10,9,36,43,2},
        {37,44,43,-38,29,-44,19,38},
    }

	// result: 
	// matrix := [][]int{}

	result := matrixRankTransform(matrix)
	fmt.Printf("result = %v\n", result)
}

