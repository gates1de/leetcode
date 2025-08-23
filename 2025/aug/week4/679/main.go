package main

import (
	"fmt"
	"math"
)

const eps = float64(0.001)

func judgePoint24(cards []int) bool {
	result := false
	arr := make([]float64, 0)
	for _, num := range cards {
		arr = append(arr, float64(num))
	}

	helper(arr, &result)
	return result
}

func helper(arr []float64, result *bool) {
	if *result {
		return
	}

	if len(arr) == 1 {
		if math.Abs(arr[0] - 24) < eps {
			*result = true
		}

		return
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			next := make([]float64, 0)
			p1 := arr[i]
			p2 := arr[j]

			next = append(next, p1 + p2, p1 - p2, p2 - p1, p1 * p2)

			if math.Abs(p2) > eps {
				next = append(next, p1 / p2)
			}
			if math.Abs(p1) > eps {
				next = append(next, p2 / p1)
			}

			arr = remove(arr, i)
			arr = remove(arr, j)

			for _, num := range next {
				arr = append(arr, num)
				newArr := make([]float64, len(arr))
				copy(newArr, arr)
				helper(newArr, result)
				arr = arr[:len(arr) - 1]
			}
			arr = add(arr, j, p2)
			arr = add(arr, i, p1)
		}
	}
}

func add(nums []float64, index int, num float64) []float64 {
	result := make([]float64, 0, len(nums) + 1)
	result = append(result, nums[:index]...)
	result = append(result, num)
	result = append(result, nums[index:]...)
	return result
}

func remove(nums []float64, i int) []float64 {
	return append(nums[:i], nums[i + 1:]...)
}

func main() {
	// result: true
	// cards := []int{4,1,8,7}

	// result: false
	cards := []int{1,2,1,2}

	// result: 
	// cards := []int{}

	result := judgePoint24(cards)
	fmt.Printf("result = %v\n", result)
}

