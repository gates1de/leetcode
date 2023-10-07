package main
import (
	"fmt"
)

func majorityElement(nums []int) []int {
	n := len(nums)
	candidate1 := int(0)
	candidate2 := int(1)
    count1 := int(0)
    count2 := int(0)
    
    for _, num := range nums {
        if num == candidate1 {
            count1++
        } else if num == candidate2 {
            count2++
        } else if count1 == 0 {
            candidate1 = num
            count1 = 1
        } else if count2 == 0 {
            candidate2 = num
            count2 = 1
        } else {
            count1--
            count2--
        }
    }

    count1 = 0
    count2 = 0
    for _, num := range nums {
        if num == candidate1 {
            count1++
        } else if num == candidate2 {
            count2++
        }
    }

	result := make([]int, 0, n)
    if count1 > n / 3 {
        result = append(result, candidate1)
    }
    if count2 > n / 3 {
        result = append(result, candidate2)
    }
    return result
}

func main() {
	// result: [3]
	// nums := []int{3,2,3}

	// result: [1]
	// nums := []int{1}

	// result: [1,2]
	nums := []int{1,2}

	// result: 
	// nums := []int{}

	result := majorityElement(nums)
	fmt.Printf("result = %v\n", result)
}

