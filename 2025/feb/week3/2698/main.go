package main
import (
	"fmt"
)

func punishmentNumber(n int) int {
	result := int(0)

	for i := 1; i <= n; i++ {
		squareNum := i * i

		if canPartition(squareNum, i) {
			result += squareNum
		}
	}

	return result
}

func canPartition(num int, target int) bool {
	if target < 0 || num < target {
		return false
	}

	if num == target {
		return true
	}

	return canPartition(num / 10, target - (num % 10)) ||
					canPartition(num / 100, target - (num % 100)) ||
					canPartition(num / 1000, target - (num % 1000))
}

func main() {
	// result: 182
	// n := int(10)

	// result: 1478
	n := int(37)

	// result: 
	// n := int(0)

	result := punishmentNumber(n)
	fmt.Printf("result = %v\n", result)
}

