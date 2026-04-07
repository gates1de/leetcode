package main

import (
	"fmt"
)

type Robot struct {
	Moved  bool
	Index  int
	Pos    [][2]int
	Dir    []int
	DirMap map[int]string
}

func Constructor(width int, height int) Robot {
	robot := Robot{
		DirMap: map[int]string{
			0: "East",
			1: "North",
			2: "West",
			3: "South",
		},
	}

	for i := range width {
		robot.Pos = append(robot.Pos, [2]int{i, 0})
		robot.Dir = append(robot.Dir, 0)
	}
	for i := 1; i < height; i++ {
		robot.Pos = append(robot.Pos, [2]int{width - 1, i})
		robot.Dir = append(robot.Dir, 1)
	}
	for i := width - 2; i >= 0; i-- {
		robot.Pos = append(robot.Pos, [2]int{i, height - 1})
		robot.Dir = append(robot.Dir, 2)
	}
	for i := height - 2; i > 0; i-- {
		robot.Pos = append(robot.Pos, [2]int{0, i})
		robot.Dir = append(robot.Dir, 3)
	}
	robot.Dir[0] = 3

	return robot
}

func (this *Robot) Step(num int) {
	this.Moved = true
	this.Index = (this.Index + num) % len(this.Pos)
}

func (this *Robot) GetPos() []int {
	return []int{this.Pos[this.Index][0], this.Pos[this.Index][1]}
}

func (this *Robot) GetDir() string {
	if !this.Moved {
		return "East"
	}
	return this.DirMap[this.Dir[this.Index]]
}

func main() {
		// result: [null, null, null, [4, 0], "East", null, null, null, [1, 2], "West"]
		width := int(6)
		height := int(3)
		operations := [][]int{{0,2},{0,2},{1},{2},{0,2},{0,1},{0,4},{1},{2}}

		// result: []
		// width := int(0)
		// height := int(0)
		// operations := [][]int{{0,0,0},{1,0}}

		obj := Constructor(width, height)
		for _, ope := range operations {
			if ope[0] == 0 {
				obj.Step(ope[1])
			} else if ope[0] == 1 {
				fmt.Printf("obj.GetPos() = %v\n", obj.GetPos())
			} else if ope[0] == 2 {
				fmt.Printf("obj.GetDir() = %v\n", obj.GetDir())
			}
		}
}

