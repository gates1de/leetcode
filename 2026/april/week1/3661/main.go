package main

import (
	"fmt"
	"sort"
)

type RobotDist struct {
	Pos   int
	Dist  int
}

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	robotDist := make([]RobotDist, n)
	for i := 0; i < n; i++ {
		robotDist[i] = RobotDist{robots[i], distance[i]}
	}

	sort.Slice(robotDist, func(i, j int) bool {
		return robotDist[i].Pos < robotDist[j].Pos
	})
	sort.Ints(walls)

	m := len(walls)
	rightPtr, leftPtr, curPtr, robotPtr := int(0), int(0), int(0), int(0)
	prevRight, subLeft, subRight := int(0), int(0), int(0)

	for i := range n {
		robotPos := robotDist[i].Pos
		robotDistVal := robotDist[i].Dist
		for rightPtr < m && walls[rightPtr] <= robotPos {
			rightPtr++
		}
		pos1 := rightPtr

		for curPtr < m && walls[curPtr] < robotPos {
			curPtr++
		}
		pos2 := curPtr

		leftBound := robotPos - robotDistVal
		if i >= 1 {
			leftBound = max(robotPos - robotDistVal, robotDist[i - 1].Pos + 1)
		}

		for leftPtr < m && walls[leftPtr] < leftBound {
			leftPtr++
		}
		leftPos := leftPtr

		currentLeft := pos1 - leftPos
		rightBound := robotPos + robotDistVal
		if i < n - 1 {
			rightBound = min(robotPos+robotDistVal, robotDist[i + 1].Pos - 1)
		}

		for rightPtr < m && walls[rightPtr] <= rightBound {
			rightPtr++
		}
		rightPos := rightPtr
		currentRight := rightPos - pos2

		currentNum := 0
		if i > 0 {
			for robotPtr < m && walls[robotPtr] < robotDist[i - 1].Pos {
				robotPtr++
			}

			pos3 := robotPtr
			currentNum = pos1 - pos3
		}

		if i == 0 {
			subLeft = currentLeft
			subRight = currentRight
		} else {
			newSubLeft := max(subLeft + currentLeft, subRight - prevRight + min(currentLeft + prevRight, currentNum))
			newSubRight := max(subLeft + currentRight, subRight + currentRight)
			subLeft = newSubLeft
			subRight = newSubRight
		}

		prevRight = currentRight
	}

	return max(subLeft, subRight)
}

func main() {
	// result: 1
	// robots := []int{4}
	// distance := []int{3}
	// walls := []int{1,10}

	// result: 3
	// robots := []int{10,2}
	// distance := []int{5,1}
	// walls := []int{5,2,7}

	// result: 0
	robots := []int{1,2}
	distance := []int{100,1}
	walls := []int{10}

	// result: 
	// robots := []int{}
	// distance := []int{}
	// walls := []int{}

	result := maxWalls(robots, distance, walls)
	fmt.Printf("result = %v\n", result)
}

