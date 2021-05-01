package main
import (
	"fmt"
)

func powerfulIntegers(x int, y int, bound int) []int {
	if bound <= 1 {
		return []int{}
	} else if bound == 2 {
		return []int{2}
	}

	sums := map[int]int{}
	for currentX := 1; currentX < bound; currentX *= x {
		for currentY := 1; currentX + currentY <= bound; currentY *= y {
			sums[currentX + currentY]++
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	result := []int{}
	for num, _ := range sums {
		result = append(result, num)
	}
	return result
}

// Time Limit Exceeded
func ngSolution(x int, y int, bound int) []int {
	if bound <= 1 {
		return []int{}
	} else if bound == 2 {
		return []int{2}
	}

	result := []int{}
	sums := map[int]int{}
	dfs(x, y, bound, 1, 1, sums)
	for num, _ := range sums {
		result = append(result, num)
	}
	return result
}

func dfs(x int, y int, bound int, currentX int, currentY int, sums map[int]int) {
	// fmt.Printf("currentX = %v, currentY = %v, sums = %v\n", currentX, currentY, sums)
	currentSum := currentX + currentY
	if currentSum > bound {
		return
	}
	sums[currentSum]++
	if x > 1 {
		dfs(x, y, bound, currentX * x, currentY, sums)
	}
	if y > 1 {
		dfs(x, y, bound, currentX, currentY * y, sums)
	}
}

func main() {
	// result: [2, 3, 4, 5, 7, 9, 10]
	// x := int(2)
	// y := int(3)
	// bound := int(10)

	// result: [2, 4, 6, 8, 10, 14]
	// x := int(3)
	// y := int(5)
	// bound := int(15)

	// result: [2, 3, 5, 9]
	// x := int(2)
	// y := int(1)
	// bound := int(10)

	// result: [33, 2, 3, 5, 65, 9, 17]
	// x := int(1)
	// y := int(2)
	// bound := int(100)

	// result: omit
	x := int(2)
	y := int(2)
	bound := int(400000)

	// result: []
	// x := int(0)
	// y := int(0)
	// bound := int(0)

	result := powerfulIntegers(x, y, bound)
	fmt.Printf("result = %v\n", result)
}

