package main

import (
	"fmt"
	"sort"
)

type Event struct {
	Y int
	L int
	Delta int
}

func separateSquares(squares [][]int) float64 {
	totalArea := int64(0)
	events := make([]Event, 0, len(squares) * 2)

	for _, sq := range squares {
		y := sq[1]
		l := sq[2]
		totalArea += int64(l) * int64(l)
		events = append(events, Event{Y: y, L: l, Delta: 1})
		events = append(events, Event{Y: y + l, L: l, Delta: -1})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Y < events[j].Y
	})

	coveredWidth := float64(0)
	currArea := float64(0)
	prevHeight := float64(0)

	for _, event := range events {
		y := event.Y
		l := event.L
		delta := event.Delta
		diff := float64(y) - prevHeight
		area := coveredWidth * diff

		if 2.0 * (currArea + area) >= float64(totalArea) {
				return prevHeight + (float64(totalArea) - 2.0 * currArea) / (2.0 * coveredWidth)
		}

		coveredWidth += float64(delta * l)
		currArea += area
		prevHeight = float64(y)
	}

	return 0.0
}

func main() {
	// result: 1.00000
	// squares := [][]int{{0,0,1}, {2,2,1}}

	// result: 1.16667
	squares := [][]int{{0,0,2}, {1,1,1}}

	// result: 
	// squares := [][]int{}

	result := separateSquares(squares)
	fmt.Printf("result = %v\n", result)
}

