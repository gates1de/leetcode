package main
import (
	"fmt"
	"math"
)

func numsSameConsecDiff(n int, k int) []int {
	if n == 1 {
		return []int{0,1,2,3,4,5,6,7,8,9}
	}

	if k == 0 {
		result := []int{}
		num := int(0)
		for i := 0; i < n; i++ {
			num += 1 * int(math.Pow10(i))
		}
		for i := 1; i < 10; i++ {
			result = append(result, i * num)
		}
		return result
	}

	result := []int{1,2,3,4,5,6,7,8,9}
	for i := 1; i < n; i++ {
		next := []int{}
		for _, num := range result {
			if num % 10 - k >= 0 {
				next = append(next, num * 10 + num % 10 - k)
			}
			if num % 10 + k <= 9 {
				next = append(next, num * 10 + num % 10 + k)
			}
		}

		result = next
	}

	return result
}

// Wrong Answer
func ngSolution(n int, k int) []int {
	result := []int{}

	for i := 0; i < 10; i++ {
		if i >= 1 && i + k <= 9 {
			num := int(0)
			digit := int(1)
			pair := i + k
			for j := 1; j < int(10e+8); j *= 10 {
				if digit > n {
					break
				}

				num *= 10
				if digit % 2 == 0 {
					num += pair
				} else {
					num += i
				}

				digit++
			}

			result = append(result, num)
		}

		if k == 0 {
			continue
		}

		if i - k >= 0 {
			num := int(0)
			digit := int(1)
			pair := i - k
			for j := 1; j < int(10e+8); j *= 10 {
				if digit > n {
					break
				}

				num *= 10
				if digit % 2 == 0 {
					num += pair
				} else {
					num += i
				}

				digit++
			}

			result = append(result, num)
		}
	}

	return result
}

func main() {
	// result: [181,292,707,818,929]
	// n := int(3)
	// k := int(7)

	// result: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
	// n := int(2)
	// k := int(1)

	// result: [909090909]
	// n := int(9)
	// k := int(9)

	// result: [11,22,33,44,55,66,77,88,99]
	n := int(2)
	k := int(0)

	// result: 
	// n := int(0)
	// k := int(0)

	result := numsSameConsecDiff(n, k)
	fmt.Printf("result = %v\n", result)
}

