package main
import (
	"fmt"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	visited := make(map[int]bool, 1000)
	visited[0] = true
	queue := make([]int, 0, 1000)
	for _, room := range rooms[0] {
		queue = append(queue, room)
	}

	for len(queue) > 0 {
		currentRoom := queue[0]
		queue = queue[1:]
		keys := rooms[currentRoom]
		visited[currentRoom] = true

		for _, key := range keys {
			if visited[key] {
				continue
			}
			queue = append(queue, key)
		}
	}

	return len(visited) == len(rooms)
}

func main() {
	// result: true
	// rooms := [][]int{{1},{2},{3},{}}

	// result: false
	rooms := [][]int{{1,3},{3,0,1},{2},{0}}

	// result: 
	// rooms := [][]int{}

	result := canVisitAllRooms(rooms)
	fmt.Printf("result = %v\n", result)
}

