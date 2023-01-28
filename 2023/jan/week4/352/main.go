package main
import (
	"fmt"
	"sort"
)


type Range struct {
    Min int
	Max int
    Parent *Range
    D int
}

func newRange(val int) *Range {
    n := &Range{Max:val, Min:val}
    n.Parent = n
    n.D = 1
    return n
}

type Ranges []*Range

func (rs Ranges) Len() int {
    return len(rs)
}

func (rs Ranges) Swap(i, j int) {
    rs[i], rs[j] = rs[j], rs[i]
}

func (rs Ranges) Less(i, j int) bool {
    return rs[i].Min < rs[j].Min
}

type SummaryRanges struct {
    RangeMap map[int]*Range
    RootMap map[*Range]*Range
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		RangeMap: make(map[int]*Range),
		RootMap: make(map[*Range]*Range),
	}
}

func (this *SummaryRanges) AddNum(val int)  {
    if _, ok := this.RangeMap[val]; ok {
        return
    }

    n := newRange(val)
    this.RangeMap[val] = n
    this.RootMap[n] = n
    
    if r, ok := this.RangeMap[val - 1]; ok {
        this.Union(r, n)
    }

    r, ok := this.RangeMap[val+1]
    if !ok {
        return
    }
    this.Union(n, r)
}

func (this *SummaryRanges) Union(l, r *Range) {
    pl := find(l)
    pr := find(r)
    if pl.D > pr.D {
        pr.Parent = pl
        pl.Max = pr.Max
        this.RootMap[pr] = nil
    } else {
        if pl.D == pr.D {
            pr.D++
        }
        pl.Parent = pr
        pr.Min = pl.Min
        this.RootMap[pl] = nil
    }
}

func find (l *Range) *Range {
    if l.Parent == l {
        return l
    }
    p := find(l.Parent)
    l.Parent = p
    return p
}

func (this *SummaryRanges) GetIntervals() (results [][]int) {
    var ranges Ranges
    for _, root := range this.RootMap {
        if root != nil {
            ranges = append(ranges, root)
        }
    }
    if len(ranges) >= 2 {
        sort.Sort(ranges)
    }
    for _, r := range ranges {
        results = append(results, []int{r.Min, r.Max})
    }
    return
}

func main() {
	// result: [null, null, [[1, 1]], null, [[1, 1], [3, 3]], null, [[1, 1], [3, 3], [7, 7]], null, [[1, 3], [7, 7]], null, [[1, 3], [6, 7]]]
	operations := [][]int{{0, 1},{1, 0},{0, 3},{1, 0},{0, 7},{1, 0},{0, 2},{1, 0},{0, 6},{1, 0}}

	// result: 
	// operations := [][]int{}

	obj := Constructor()
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.AddNum(ope[1])
		} else {
			fmt.Printf("obj.GetIntervals() = %v\n", obj.GetIntervals())
		}
	}
}

