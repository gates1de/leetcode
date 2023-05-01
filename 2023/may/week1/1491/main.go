package main
import (
	"fmt"
)

func average(salary []int) float64 {
	sum := float64(0)
	min := float64(1000001)
	max := float64(0)
	for _, val := range salary {
		floatVal := float64(val)
		if floatVal < min {
			min = floatVal
		}
		if floatVal > max {
			max = floatVal
		}
		sum += floatVal
	}

	return (sum - min - max) / float64(len(salary) - 2)
}

func main() {
	// result: 2500.00000
	// salary := []int{4000,3000,1000,2000}

	// result: 2000.00000
	salary := []int{1000,2000,3000}

	// result: 
	// salary := []int{}

	result := average(salary)
	fmt.Printf("result = %v\n", result)
}

