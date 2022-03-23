package main
import (
	"fmt"
	"math"
)

func brokenCalc(startValue int, target int) int {
	if startValue == target {
		return 0
	}

	result := math.MaxInt32
	helper(target, 0, startValue, &result)
	return result
}

func helper(current int, count int, target int, result *int) {
	// fmt.Println(current, count, target)
	if current <= 0 || count >= *result {
		return
	}

	if current == target {
		if count < *result {
			*result = count
		}
		return
	} else if current < target {
		count += target - current
		if count < *result {
			*result = count
		}
		return
	}

	if current % 2 == 0 {
		helper(current / 2, count + 1, target, result)
	} else {
		helper(current + 1, count + 1, target, result)
	}
}

func main() {
	// result: 2
	// startValue := int(2)
	// target := int(3)

	// result: 2
	// startValue := int(5)
	// target := int(8)

	// result: 3
	// startValue := int(3)
	// target := int(10)

	// result: 39
	// startValue := int(1)
	// target := int(1000000000)

	// result: 99999999
	// startValue := int(1000000000)
	// target := int(1)

	// result: 4705961
	startValue := int(9411921)
	target := int(9411923)

	// result: 
	// startValue := int(0)
	// target := int(0)

	result := brokenCalc(startValue, target)
	fmt.Printf("result = %v\n", result)
}

