package main

import (
	"fmt"
)

func sumFourDivisors(nums []int) int {
	result := int(0)

	for _, val := range nums {
		num := val
		p := int(0)
		q := int(0)
		count := int(0)

		for i := 2; i * i <= num && count <= 2; i++ {
			if num % i == 0 {
				count++

				if count == 1 {
					p = i
				} else {
					q = i
				}

				for num % i == 0 {
					num /= i
				}
			}
		}

		if num > 1 {
			count++
			if count == 1 {
				p = num
			} else {
				q = num
			}
		}

		if count == 2 && p * q == val {
			result += 1 + p + q + val
		} else if count == 1 && p * p * p == val {
			result += 1 + p + p * p + val
		}
	}

	return result
}

func main() {
	// result: 32
	// nums := []int{21,4,7}

	// result: 64
	// nums := []int{21,21}

	// result: 0
	nums := []int{1,2,3,4,5}

	// result: 
	// nums := []int{}

	result := sumFourDivisors(nums)
	fmt.Printf("result = %v\n", result)
}

