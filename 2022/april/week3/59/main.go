package main
import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i, _ := range result {
		result[i] = make([]int, n)
	}

	val := int(1)
	current := [2]int{0, 0}
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex := int(0)
	for true {
		if val == n * n {
			result[current[0]][current[1]] = val
			break
		}

		result[current[0]][current[1]] = val
		val++
		nextY := current[0] + dirs[dirIndex][0]
		nextX := current[1] + dirs[dirIndex][1]

		if nextY < 0 || nextY == n || nextX < 0 || nextX == n {
			if dirIndex == 3 {
				dirIndex = 0
			} else {
				dirIndex++
			}
		} else {
			nextNum := result[nextY][nextX]
			if nextNum > 0 {
				if dirIndex == 3 {
					dirIndex = 0
				} else {
					dirIndex++
				}
			}
		}

		current[0] += dirs[dirIndex][0]
		current[1] += dirs[dirIndex][1]
	}

	return result
}

func main() {
	// result: [[1,2,3],[8,9,4],[7,6,5]]
	// n := int(3)

	// result: [[1]]
	// n := int(1)

	// result: [[1,2],[4,3]]
	// n := int(2)

	// result: [[1,2,3,4,5,6,7,8,9,10],[36,37,38,39,40,41,42,43,44,11],[35,64,65,66,67,68,69,70,45,12],[34,63,84,85,86,87,88,71,46,13],[33,62,83,96,97,98,89,72,47,14],[32,61,82,95,100,99,90,73,48,15],[31,60,81,94,93,92,91,74,49,16],[30,59,80,79,78,77,76,75,50,17],[29,58,57,56,55,54,53,52,51,18],[28,27,26,25,24,23,22,21,20,19]]
	n := int(10)

	// result: 
	// n := int(0)

	result := generateMatrix(n)
	fmt.Printf("result = %v\n", result)
}

