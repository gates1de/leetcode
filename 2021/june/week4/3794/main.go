package main
import (
	"fmt"
)

func removeDuplicates(s string) string {
	queue := []rune{rune(s[0])}
	for i, r := range s {
		if i == 0 {
			continue
		}
		if len(queue) == 0 {
			queue = append(queue, r)
			continue
		}
		if queue[len(queue) - 1] == r {
			queue = queue[:len(queue) - 1]
			continue
		}
		queue = append(queue, r)
	}

	if len(queue) >= 2 && queue[len(queue) - 2] == queue[len(queue) - 1] {
		queue = queue[:len(queue) - 2]
	}

	result := string(queue)
	return result
}

func main() {
	// result: ca
	// s := "abbaca"

	// result: ay
	// s := "azxxzy"

	// result: kuiviuka
	// s := "kuivvviuka"

	// result: 
	// s := "aaaaaaaa"

	// result: a
	s := "aaaaaaaaa"

	// result: 
	// s := ""

	result := removeDuplicates(s)
	fmt.Printf("result = %v\n", result)
}

