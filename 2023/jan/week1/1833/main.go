package main
import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	result := int(0)
	for _, cost := range costs {
		coins -= cost
		if coins < 0 {
			break
		}
		result++
	}
	return result
}

func main() {
	// result: 4
	// costs := []int{1,3,2,4,1}
	// coins := int(7)

	// result: 0
	// costs := []int{10,6,8,7,7,8}
	// coins := int(5)

	// result: 6
	costs := []int{1,6,3,1,2,5}
	coins := int(20)

	// result: 
	// costs := []int{}
	// coins := int(0)

	result := maxIceCream(costs, coins)
	fmt.Printf("result = %v\n", result)
}

