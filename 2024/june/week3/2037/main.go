package main
import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	result := int(0)
	for i, seat := range seats {
		result += abs(seat - students[i])
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 4
	// seats := []int{3,1,5}
	// students := []int{2,7,4}

	// result: 7
	// seats := []int{4,1,5,9}
	// students := []int{1,3,2,6}

	// result: 4
	seats := []int{2,2,6,6}
	students := []int{1,3,2,6}

	// result: 
	// seats := []int{}
	// students := []int{}

	result := minMovesToSeat(seats, students)
	fmt.Printf("result = %v\n", result)
}

