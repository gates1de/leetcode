package main
import (
	"fmt"
)

type MyCalendarThree struct {
	CountMap map[[2]int]int
	Max int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{CountMap: map[[2]int]int{}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	newCountMap := map[[2]int]int{}
	for key, val := range this.CountMap {
		newCountMap[key] = val
	}

	for schedule, count := range this.CountMap {
		s := schedule[0]
		e := schedule[1]
		newCount := count

		if s <= start && start < e {
			newCount = count + 1
			newCountMap[[2]int{start, min(e, end)}] = max(newCountMap[[2]int{start, min(e, end)}], newCount)
		} else if s < end && end <= e {
			newCount = count + 1
			newCountMap[[2]int{s, min(e, end)}] = max(newCountMap[[2]int{s, min(e, end)}], newCount)
		} else if start <= s && e < end {
			newCount = count + 1
			newCountMap[[2]int{s, e}] = max(newCountMap[[2]int{s, e}], newCount)
		}

		if this.Max < newCount {
			this.Max = newCount
		}
	}

	newCountMap[[2]int{start, end}] = max(1, newCountMap[[2]int{start, end}])
	if this.Max < 1 {
		this.Max = 1
	}
	this.CountMap = newCountMap

	return this.Max
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

func main() {
	// result: [null,1,1,2,3,3,3]
	// books := [][]int{{10,20},{50,60},{10,40},{5,15},{5,10},{25,55}}

	// result: [null,1,1,2,3,4,4]
	// books := [][]int{{10,20},{50,60},{10,40},{5,15},{5,11},{25,55}}

	// result:
	books := [][]int{{26,35},{26,32},{25,32},{18,26},{40,45},{19,26},{48,50},{1,6},{46,50},{11,18}}

	// result:
	// books := [][]int{}

	obj := Constructor()
	for _, book := range books {
		start := book[0]
		end := book[1]
		fmt.Printf("obj.Book(%v, %v) = %v\n", start, end, obj.Book(start, end))
	}
}
