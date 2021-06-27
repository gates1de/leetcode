package main
import (
	"fmt"
)

func candy(ratings []int) int {
	if len(ratings) {
		return 1
	}

	ratingsLen := len(ratings)
	candies := make([]int, ratingsLen)
	for i, _ := range candies {
		candies[i] = 1
	}

	// first child
	if ratings[0] > ratings[1] {
		candies[0] = 2
	}

	// last child
	if ratings[ratingsLen - 2] < ratings[ratingsLen - 1] {
		candies[ratingsLen - 1] = 2
	}

	for i := 1; i < ratingsLen - 1; i++ {
		j := ratingsLen - i - 1

		left := ratings[i - 1]
		rating := ratings[i]
		right := ratings[i + 1]
		if rating > left && rating > right {
			val := max(candies[i - 1], candies[i + 1]) + 1
			candies[i] = max(candies[i], val)
		} else if rating > left && rating <= right {
			candies[i] = max(candies[i], candies[i - 1] + 1)
		} else if rating <= left && rating > right {
			candies[i] = max(candies[i], candies[i + 1] + 1)
		}
		// fmt.Printf("i: candies = %v\n", candies)

		if j > 0 {
			left = ratings[j - 1]
			rating = ratings[j]
			right = ratings[j + 1]
			if rating > left && rating > right {
				val := max(candies[j - 1], candies[j + 1]) + 1
				candies[j] = max(candies[j], val)
			} else if rating > left && rating <= right {
				candies[j] = max(candies[j], candies[j - 1] + 1)
			} else if rating <= left && rating > right {
				candies[j] = max(candies[j], candies[j + 1] + 1)
			}
		}
		// fmt.Printf("j: candies = %v\n", candies)
	}

	// first child
	if ratings[0] > ratings[1] {
		candies[0] = max(2, candies[1] + 1)
	}

	// last child
	if ratings[ratingsLen - 2] < ratings[ratingsLen - 1] {
		candies[ratingsLen - 1] = max(2, candies[ratingsLen - 2] + 1)
	}

	return sum(candies)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 5
	// ratings := []int{1, 0, 2}

	// result: 4
	// ratings := []int{1, 2, 2}

	// result: 8
	// ratings := []int{1, 2, 2, 3, 2, 2}

	// result: 15
	// ratings := []int{13209, 19382, 1, 3812, 3812, 3812, 5, 13219, 20000}

	// result: 12
	ratings := []int{29, 51, 87, 87, 72, 12}

	// result: 
	// ratings := []int{}

	result := candy(ratings)
	fmt.Printf("result = %v\n", result)
}

