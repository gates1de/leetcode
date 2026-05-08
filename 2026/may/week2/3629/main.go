package main

import (
	"fmt"
)

const MX = int(1000001)

var factors [MX][]int

func init() {
	for i := 2; i < MX; i++ {
		if len(factors[i]) == 0 {
			for j := i; j < MX; j += i {
				factors[j] = append(factors[j], i)
			}
		}
	}
}

func minJumps(nums []int) int {
	n := len(nums)
	edges := make(map[int][]int)
	for i, num := range nums {
		if len(factors[num]) == 1 {
			edges[num] = append(edges[num], i)
		}
	}

	result := int(0)
	seen := make([]bool, n)
	seen[n - 1] = true
	q := []int{n - 1}

	for {
		var q2 []int
		for _, i := range q {
			if i == 0 {
				return result
			}

			if i > 0 && !seen[i - 1] {
				seen[i - 1] = true
				q2 = append(q2, i - 1)
			}

			if i < n - 1 && !seen[i + 1] {
				seen[i + 1] = true
				q2 = append(q2, i + 1)
			}

			for _, p := range factors[nums[i]] {
				if list, ok := edges[p]; ok {
					for _, j := range list {
						if !seen[j] {
							seen[j] = true
							q2 = append(q2, j)
						}
					}

					delete(edges, p)
				}
			}
		}

		q = q2
		result++
	}
}

func main() {
	// result: 2
	// nums := []int{1,2,4,6}

	// result: 2
	// nums := []int{2,3,4,7,9}

	// result: 3
	nums := []int{4,6,5,8}

	// result: 
	// nums := []int{}

	result := minJumps(nums)
	fmt.Printf("result = %v\n", result)
}

