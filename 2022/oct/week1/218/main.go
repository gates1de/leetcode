package main
import (
	"fmt"
	"container/heap"
	"sort"
	"strconv"
)

type Skyline struct {
    Left int
	Right int
	Height int
    Index int
} 

type Heap []*Skyline

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool { return h[i].Height > h[j].Height }

func (h Heap) Swap(i, j int) { 
    h[i], h[j] = h[j], h[i]
    h[i].Index, h[j].Index = h[j].Index, h[i].Index
}
func (h *Heap) Push(x interface{}) {
    n := h.Len()
    skyline := x.(*Skyline)
    skyline.Index = n
    *h = append(*h, skyline)
}
func (h *Heap) Pop() interface{} { 
    n := h.Len()
    old := *h
    x := old[n - 1]
    *h = old[:n - 1]
    
    return x
}

func getSkyline(buildings [][]int) [][]int {
    m := map[int][]*Skyline{}

    for _, building := range buildings {
        skyline := &Skyline{
            Left: building[0],
            Right: building[1],
            Height: building[2],
        }

        m[skyline.Left] = append(m[skyline.Left], skyline)
        m[skyline.Right] = append(m[skyline.Right], skyline)
    }

    positions := []int{}

    for pos := range m {
        positions = append(positions, pos)
    }

    sort.Ints(positions)
    result := [][]int{}

    h := Heap{&Skyline{Left: -1, Right: -1, Height: 0}}

    for _, pos := range positions {
        curr := h[0].Height

        for _, skyline := range m[pos] {
            if skyline.Right > pos {
                heap.Push(&h, skyline)
            } else {
                heap.Remove(&h, skyline.Index)
            }
        }

        if h[0].Height != curr {
            result = append(result, []int{pos, h[0].Height})
        }
    }

    return result
}

// Wrong Answer
func ngSolution(buildings [][]int) [][]int {
	maxX := buildings[len(buildings) - 1][1]
	skyline := map[string]int{}

	for i, b1 := range buildings {
		start1 := b1[0]
		start1Str := strconv.Itoa(b1[0])
		end1 := b1[1]
		end1Str := strconv.Itoa(b1[1])
		height1 := b1[2]

		skyline[start1Str] = max(skyline[start1Str], height1)
		skyline[end1Str] = max(skyline[end1Str], 0)

		for j, b2 := range buildings {
			if i == j {
				continue
			}

			start2 := b2[0]
			end2 := b2[1]
			height2 := b2[2]

			if start2 <= start1 && start1 <= end2 {
				if height1 <= height2 {
					delete(skyline, start1Str)
				} else {
					skyline[start1Str] = max(skyline[start1Str], height1)
				}
			}

			if start2 <= end1 && end1 <= end2 {
				if height1 <= height2 {
					delete(skyline, end1Str)
				} else {
					skyline[end1Str] = max(skyline[end1Str], height2)
				}
			}

			if end1 < start1 {
				break
			}
		}
	}

	skyline[strconv.Itoa(maxX)] = 0

	result := [][]int{}
	for key, val := range skyline {
		x, _ := strconv.Atoi(key)
		result = append(result, []int{x, val})
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a][0] < result[b][0]
	})
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
	buildings := [][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}

	// result: [[0,3],[5,0]]
	// buildings := [][]int{{0,2,3},{2,5,3}}

	// result: 
	// buildings := [][]int{}

	result := getSkyline(buildings)
	fmt.Printf("result = %v\n", result)
}

