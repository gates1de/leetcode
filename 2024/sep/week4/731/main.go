package main
import (
	"fmt"
)

type MyCalendarTwo struct {
	Bookings [][]int
	OverlapBookings [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{Bookings: make([][]int, 0), OverlapBookings: make([][]int, 0)}
}

func (this *MyCalendarTwo) DoesOverlap(start1 int, end1 int, start2 int, end2 int) bool {
	return max(start1, start2) < min(end1, end2)
}

func (this *MyCalendarTwo) GetOverlapped(start1 int, end1 int, start2 int, end2 int) []int {
	return []int{max(start1, start2), min(end1, end2)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, booking := range this.OverlapBookings {
		if this.DoesOverlap(booking[0], booking[1], start, end) {
			return false
		}
	}

	for _, booking := range this.Bookings {
		if this.DoesOverlap(booking[0], booking[1], start, end) {
			this.OverlapBookings = append(this.OverlapBookings, this.GetOverlapped(booking[0], booking[1], start, end))
		}
	}

	this.Bookings = append(this.Bookings, []int{start, end})
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func solution() int {
	return 0
}

func main() {
	// result: [null, true, true, true, false, true, true]
	books := [][]int{{10,20},{50,60},{10,40},{5,15},{5,10},{25,55}}

	// result: []
	// books := [][]int{}

	obj := Constructor()
	for _, book := range books {
		start := book[0]
		end := book[1]
		fmt.Printf("obj.Book(%v, %v) = %v\n", start, end, obj.Book(start, end))
	}
}
