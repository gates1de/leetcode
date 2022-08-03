package main
import (
	"fmt"
	"sort"
)

type MyCalendar struct {
	Bookings [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{Bookings: [][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	n := len(this.Bookings)
	result := true
	if n > 0 {
		insertIndex := sort.Search(n, func(i int) bool {
			if start < this.Bookings[i][0] {
				return true
			}
			return end < this.Bookings[i][1]
		})

		if insertIndex == n {
			pre := this.Bookings[n - 1]
			if (pre[0] <= start && start < pre[1]) || (pre[0] < end && end <= pre[1]) {
				result = false
			}
		} else if insertIndex == 0 {
			next := this.Bookings[0]
			if (next[0] <= start && start < next[1]) || (next[0] < end && end <= next[1]) || (start <= next[0] && next[1] <= end) {
				result = false
			}
		} else {
			pre := this.Bookings[insertIndex - 1]
			next := this.Bookings[insertIndex]

			if (pre[0] <= start && start < pre[1]) || (pre[0] < end && end <= pre[1]) {
				result = false
			}
			if (next[0] <= start && start < next[1]) || (next[0] < end && end <= next[1]) || (start <= next[0] && next[1] <= end) {
				result = false
			}
		}
	}

	if result {
		this.Bookings = append(this.Bookings, []int{start, end})
		sort.Slice(this.Bookings, func(a, b int) bool {
			if this.Bookings[a][0] < this.Bookings[b][0] {
				return true
			}
			return this.Bookings[a][1] < this.Bookings[b][1]
		})
	}

	return result
}

func main() {
	obj := Constructor()

	// result: [null, true, false, true]
	// bookings := [][]int{{10, 20}, {15, 25}, {20, 30}}

	// result: [null, true, false, false, false, true, true]
	// bookings := [][]int{{0, 100}, {10, 20}, {50, 100}, {99, 100}, {100, 101}, {101, 110}}

	// result: [null,true,false,true,true,false,true,false,true,false,false,false,false,false]
	// bookings := [][]int{{20, 29}, {13, 22}, {44, 50}, {1, 7}, {2, 10}, {14, 20}, {19, 25}, {36, 42}, {45, 50}, {47, 50}, {39, 45}, {44, 50}, {16, 25}}

	// result: [null,true,true,false,false,true,false,true,true,true,false]
	bookings := [][]int{{47,50},{33,41},{39,45},{33,42},{25,32},{26,35},{19,25},{3,8},{8,13},{18,27}}

	// result: 
	// bookings := [][]int{}

	for _, book := range bookings {
		start := book[0]
		end := book[1]
		fmt.Printf("obj.Book(%v, %v) = %v\n", start, end, obj.Book(start, end))
	}
}

