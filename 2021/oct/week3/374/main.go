package main
import (
	"fmt"
)

// result: 6
// const N = int(10)
// const pick = int(6)

// result: 1
// const N = int(1)
// const pick = int(1)

// result: 1
// const N = int(2)
// const pick = int(1)

// result: 2
// const N = int(2)
// const pick = int(2)

// result: 34287916
const N = int(2147483647)
const pick = int(34287916)

// result: 
// n := int(0)
// pick := int(0)

func guessNumber(n int) int {
	return binarySearch(1, n)
}

func guess(num int) int {
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	}
	return 0
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

func main() {
	result := guessNumber(N)
	fmt.Printf("result = %v\n", result)
}

