package main
import (
	"fmt"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	result := int(0)
	sum := make([]int, 1001)

	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			if abs(arr[j] - arr[k]) <= b {
				lj, rj := arr[j] - a, arr[j] + a
				lk, rk := arr[k] - c, arr[k] + c
				l := max(0, max(lj, lk))
				r := min(1000, min(rj, rk))

				if l <= r {
					if l == 0 {
						result += sum[r]
					} else {
						result += sum[r] - sum[l - 1]
					}
				}
			}
		}

		for k := arr[j]; k <= 1000; k++ {
			sum[k]++
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// arr := []int{3,0,1,1,9,7}
	// a := int(7)
	// b := int(2)
	// c := int(3)

	// result: 0
	arr := []int{1,1,2,2,3}
	a := int(0)
	b := int(0)
	c := int(1)

	// result: 
	// arr := []int{}
	// a := int(0)
	// b := int(0)
	// c := int(0)

	result := countGoodTriplets(arr, a, b ,c)
	fmt.Printf("result = %v\n", result)
}

