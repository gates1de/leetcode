package main
import (
	"fmt"
)

func minJumps(arr []int) int {
    m := map[int][]int{}
    for i, num := range arr {
        if m[num] == nil {
            m[num] = []int{}
        }
        m[num] = append(m[num], i)
    }

    numVisited := map[int]bool{}
    queue := []int{0}
    step := int(0)
    visited := make([]bool, len(arr))
    visited[0] = true
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            cursor := queue[i]
            if cursor == len(arr) - 1 {
                return step
            }

            num := arr[cursor]
            if !numVisited[num] {
                numVisited[num] = true
                for _, j := range m[num] {
                    if j == cursor {
                        continue
                    }
                    queue = append(queue, j)
                    visited[j] = true
                }
            }

            if cursor - 1 >= 0 && !visited[cursor - 1] {
                queue = append(queue, cursor - 1)
                visited[cursor - 1] = true
            }
            if cursor + 1 < len(arr) && !visited[cursor + 1] {
                queue = append(queue, cursor + 1)
                visited[cursor + 1] = true
			}
		}
		queue = queue[size:]
		step++
	}

	return -1
}

func main() {
	// result: 3
	// arr := []int{100,-23,-23,404,100,23,23,23,3,404}

	// result: 0
	// arr := []int{7}

	// result: 1
	arr := []int{7,6,9,6,9,6,9,7}

	// result: 
	// arr := []int{}

	result := minJumps(arr)
	fmt.Printf("result = %v\n", result)
}

