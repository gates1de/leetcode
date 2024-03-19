package main
import (
	"fmt"
)

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	maxCount := int(0)

	for _, task := range tasks {
		freq[task - 'A']++
		maxCount = max(maxCount, freq[task - 'A'])
	}

	time := (maxCount - 1) * (n + 1)
	for _, f := range freq {
		if f == maxCount {
			time++
		}
	}
    
	return max(len(tasks), time)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 8
	// tasks := []byte{'A','A','A','B','B','B'}
	// n := int(2)

	// result: 6
	// tasks := []byte{'A','C','A','B','D','B'}
	// n := int(1)

	// result: 10
	tasks := []byte{'A','A','A','B','B','B'}
	n := int(3)

	// result: 
	// tasks := []byte{}
	// n := int(0)

	result := leastInterval(tasks, n)
	fmt.Printf("result = %v\n", result)
}

