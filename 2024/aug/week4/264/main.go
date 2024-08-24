package main
import (
	"fmt"
)

func nthUglyNumber(n int) int {
	uglyNums := make([]int, n)
	uglyNums[0] = 1

	indexMultiTwo := int(0)
	indexMultiThree := int(0)
	indexMultiFive := int(0)

	nextMultiTwo := int(2)
	nextMultiThree := int(3)
	nextMultiFive := int(5)

	for i := 1; i < n; i++ {
		nextUglyNum := min(nextMultiTwo, min(nextMultiThree, nextMultiFive))
		uglyNums[i] = nextUglyNum

		if nextUglyNum == nextMultiTwo {
			indexMultiTwo++
			nextMultiTwo = uglyNums[indexMultiTwo] * 2
		}
		if nextUglyNum == nextMultiThree {
			indexMultiThree++
			nextMultiThree = uglyNums[indexMultiThree] * 3
		}
		if nextUglyNum == nextMultiFive {
			indexMultiFive++
			nextMultiFive = uglyNums[indexMultiFive] * 5
		}
	}

	return uglyNums[n - 1]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 12
	// n := int(10)

	// result: 1
	n := int(1)

	// result: 
	// n := int(0)

	result := nthUglyNumber(n)
	fmt.Printf("result = %v\n", result)
}

