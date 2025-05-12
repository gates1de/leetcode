package main
import (
	"fmt"
	"sort"
)

func findEvenNumbers(digits []int) []int {
	n := len(digits)
	m := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || i == k {
					continue
				}

				num := digits[i] * 100 + digits[j] * 10 + digits[k]

				if num >= 100 && num%2 == 0 {
					m[num] = true
				}
			}
		}
	}

	result := make([]int, 0, len(m))
	for num := range m {
		result = append(result, num)
	}

	sort.Ints(result)
	return result
}

func main() {
	// result: [102,120,130,132,210,230,302,310,312,320]
	// digits := []int{2,1,3,0}

	// result: [222,228,282,288,822,828,882]
	// digits := []int{2,2,8,8,2}

	// result: []
	digits := []int{3,7,5}

	// result: []
	// digits := []int{}

	result := findEvenNumbers(digits)
	fmt.Printf("result = %v\n", result)
}

