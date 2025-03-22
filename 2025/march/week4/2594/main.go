package main
import (
	"fmt"
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	result := int64(1)
	high := int64(cars) * int64(cars) * int64(ranks[0])

	for result < high {
		mid := result + (high - result) / 2
		carsRepaired := int(0)

		for _, rank := range ranks {
			carsRepaired += int(math.Sqrt(1 * float64(mid / int64(rank))))
		}

		if carsRepaired < cars {
			result = mid + 1
		} else {
			high = mid
		}
	}

	return result
}

func main() {
	// result: 16
	// ranks := []int{4,2,3,1}
	// cars := int(10)

	// result: 16
	ranks := []int{5,1,8}
	cars := int(6)

	// result: 
	// ranks := []int{}
	// cars := int(0)

	result := repairCars(ranks, cars)
	fmt.Printf("result = %v\n", result)
}

