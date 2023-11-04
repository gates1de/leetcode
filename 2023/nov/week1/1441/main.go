package main
import (
	"fmt"
)

func buildArray(target []int, n int) []string {
	result := make([]string, 0, n)
	i := int(0)

	for _, num := range target {
		for i < num - 1 {
			result = append(result, "Push")
			result = append(result, "Pop")
			i++
		}

		result = append(result, "Push")
		i++
	}

	return result
}

func main() {
	// result: ["Push","Push","Pop","Push"]
	// target := []int{1,3}
	// n := int(3)

	// result: ["Push","Push","Push"]
	// target := []int{1,2,3}
	// n := int(3)

	// result: ["Push","Push"]
	target := []int{1,2}
	n := int(4)

	// result: 
	// target := []int{}
	// n := int(0)

	result := buildArray(target, n)
	fmt.Printf("result = %v\n", result)
}

