package main

import (
	"fmt"
)

func largestOverlap(img1 [][]int, img2 [][]int) int {
	type point struct {
		x int
		y int
	}

	points1 := make([]point, 0)
	points2 := make([]point, 0)
	for y := range img1 {
		for x := range img1[y] {
			if img1[y][x] == 1 {
				points1 = append(points1, point{x: x, y: y})
			}
			if img2[y][x] == 1 {
				points2 = append(points2, point{x: x, y: y})
			}
		}
	}

	width := 2*len(img1) - 1
	counts := make(map[int]int)
	result := int(0)
	for _, first := range points1 {
		for _, second := range points2 {
			dx := first.x - second.x + len(img1) - 1
			dy := first.y - second.y + len(img1) - 1
			key := dy * width + dx
			counts[key]++
			if counts[key] > result {
				result = counts[key]
			}
		}
	}

	return result
}

func main() {
	// result: 3
	// img1 := [][]int{{1,1,0},{0,1,0},{0,1,0}}
	// img2 := [][]int{{0,0,0},{0,1,1},{0,0,1}}

	// result: 1
	// img1 := [][]int{{1}}
	// img2 := [][]int{{1}}

	// result: 0
	img1 := [][]int{{0}}
	img2 := [][]int{{0}}

	// result:
	// img1 := [][]int{}
	// img2 := [][]int{}

	result := largestOverlap(img1, img2)
	fmt.Printf("result = %v\n", result)
}
