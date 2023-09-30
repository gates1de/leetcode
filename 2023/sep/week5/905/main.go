package main
import (
	"fmt"
)

func sortArrayByParity(nums []int) []int {
    even := make([]int, 0, 2500)
    odd := make([]int, 0, 2500)

    for _, num := range nums {
        if num % 2 == 0 {
            even = append(even, num)
        } else {
            odd = append(odd, num)
        }
    }

    even = append(even, odd...)
    return even
}

func main() {
	// result: [2,4,3,1]
	// nums := []int{3,1,2,4}

	// result: [0]
	nums := []int{0}

	// result: 
	// nums := []int{}

	result := sortArrayByParity(nums)
	fmt.Printf("result = %v\n", result)
}

