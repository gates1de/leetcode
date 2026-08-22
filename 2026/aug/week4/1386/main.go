package main

import (
	"fmt"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	reservedByRow := make(map[int]int)
	for _, seat := range reservedSeats {
		row, number := seat[0], seat[1]
		reservedByRow[row] |= 1 << number
	}

	groups := (n - len(reservedByRow)) * 2
	leftBlock := (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
	middleBlock := (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
	rightBlock := (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

	for _, reserved := range reservedByRow {
		if reserved&leftBlock == 0 {
			groups++
		}
		if reserved&rightBlock == 0 {
			groups++
		} else if reserved&leftBlock != 0 && reserved&middleBlock == 0 {
			groups++
		}
	}

	return groups
}

func main() {
	// result: 4
	// n := int(3)
	// reservedSeats := [][]int{{1,2},{1,3},{1,8},{2,6},{3,1},{3,10}}

	// result: 2
	// n := int(2)
	// reservedSeats := [][]int{{2,1},{1,8},{2,6}}

	// result: 4
	n := int(4)
	reservedSeats := [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}

	// result:
	// n := int(0)
	// reservedSeats := [][]int{}

	result := maxNumberOfFamilies(n, reservedSeats)
	fmt.Printf("result = %v\n", result)
}
