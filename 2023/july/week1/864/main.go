package main
import (
	"fmt"
)

type Move struct {
    Data Cell
    Steps int
}

type Cell struct {
    I int
    J int
    Keys int
}

type Queue struct {
	Arr []Move
}

func (q *Queue) Push(v Move) {
	q.Arr = append(q.Arr, v)
}

func (q *Queue) Pop() Move {
	v := q.Arr[0]
	q.Arr = q.Arr[1:]
	return v
}

func (q *Queue) Size() int {
	return len(q.Arr)
}

func shortestPathAllKeys(grid []string) int {
    m := len(grid)
    n := len(grid[0])
    startI := int(-1)
    startJ := int(-1)
    keysCount := int(0)
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '@' {
                startI = i
                startJ = j
                continue
            }

            if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
                keysCount++
            }
        }
    }

    q := Queue{}
    seen := make(map[Cell]bool)
    start := Cell{I: startI, J: startJ, Keys: 0}
    seen[start] = true
    q.Push(Move{Data: start, Steps: 0})
    
    for q.Size() > 0 {
        curr := q.Pop()
        if curr.Data.Keys == ((2 << keysCount) - 2) {
            return curr.Steps
        }
        
        for _, dir := range dirs {
            nextI := curr.Data.I + dir[0]
            nextJ := curr.Data.J + dir[1]

            if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] != '#' {
                next := Cell{I: nextI, J: nextJ, Keys: curr.Data.Keys}
                if _, ok := seen[next]; ok {
                    continue
                }

                if grid[nextI][nextJ] >= 'A' && grid[nextI][nextJ] <= 'Z' {
                    bit := (2 << int(grid[nextI][nextJ] -'A'))
                    if curr.Data.Keys & bit == 0{
                        continue
                    }
                }

                if grid[nextI][nextJ] >= 'a' && grid[nextI][nextJ] <= 'z' {
                    bit := (2 << int(grid[nextI][nextJ] - 'a'))
                    next.Keys |= bit
                }

                seen[next]=true
                nextMove := Move{Data: next, Steps: curr.Steps + 1}
                q.Push(nextMove)
            }
        }
    }

    return -1
}

func main() {
	// result: 8
	// grid := []string{"@.a..","###.#","b.A.B"}

	// result: 6
	// grid := []string{"@..aA","..B#.","....b"}

	// result: -1
	grid := []string{"@Aa"}

	// result: 
	// grid := []string{}

	result := shortestPathAllKeys(grid)
	fmt.Printf("result = %v\n", result)
}

