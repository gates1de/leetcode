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
	n := int(1)

	// result: 
	// n := int(0)

	result := generateMatrix(n)
	fmt.Printf("result = %v\n", result)
}

