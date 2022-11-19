package main
import (
	"fmt"
)

func guessNumber(n int) int {
    return binarySearch(1, n)
}

func binarySearch(head int, tail int) int {
    for head <= tail {
        mid := head + (tail - head) / 2
        if guess(mid) == 1 {
            head = mid + 1
        } else if guess(mid) == -1 {
            tail = mid - 1
        } else {
            return mid
        }
    }
    return head
}

func guess(num int) int {
	// pick := int(6)
	// pick := int(1)
	pick := int(239876342)

	if num < pick {
		return 1
	} else if num > pick {
		return -1
	}

	return 0
}

func main() {
	// result: 6
	// n := int(10)

	// result: 1
	// n := int(2)

	// result: 1
	// n := int(1)

	// result: 239876342
	n := int(432791682)

	// result: 
	// n := int(0)

	result := guessNumber(n)
	fmt.Printf("result = %v\n", result)
}

