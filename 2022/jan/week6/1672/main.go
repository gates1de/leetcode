package main
import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {
	result := int(0)

	for _, account := range accounts {
		wealth := int(0)
		for _, num := range account {
			wealth += num
		}
		if result < wealth {
			result = wealth
		}
	}

	return result
}

func main() {
	// result: 6
	// accounts := [][]int{{1,2,3},{3,2,1}}

	// result: 10
	// accounts := [][]int{{1,5},{7,3},{3,5}}

	// result: 17
	accounts := [][]int{{2,8,7},{7,1,3},{1,9,5}}

	// result: 
	// accounts := [][]int{}

	result := maximumWealth(accounts)
	fmt.Printf("result = %v\n", result)
}

