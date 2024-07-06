package main
import (
	"fmt"
)

func passThePillow(n int, time int) int {
	fullRounds := time / (n - 1)
	extraTime := time % (n - 1)

	if fullRounds % 2 == 0 {
		return extraTime + 1
	}

	return n - extraTime
}

func main() {
	// result: 2
	// n := int(4)
	// time := int(5)

	// result: 3
	n := int(3)
	time := int(2)

	// result: 
	// n := int(0)
	// time := int(0)

	result := passThePillow(n, time)
	fmt.Printf("result = %v\n", result)
}

