package main
import (
	"fmt"
	"math"
)

const modulo = int(1e9 + 7)

func colorTheGrid(m int, n int) int {
	valid := make(map[int][]int)
	maskEnd := int(math.Pow(3, float64(m)))

	for mask := 0; mask < maskEnd; mask++ {
		color := make([]int, m)
		mm := mask

		for i := 0; i < m; i++ {
			color[i] = mm % 3
			mm /= 3
		}

		isValid := true
		for i := 0; i < m-1; i++ {
			if color[i] == color[i + 1] {
				isValid = false
				break
			}
		}

		if isValid {
			valid[mask] = color
		}
	}

	adjacent := make(map[int][]int)
	for mask1 := range valid {
		for mask2 := range valid {
			isValid := true
			for i := 0; i < m; i++ {
				if valid[mask1][i] == valid[mask2][i] {
					isValid = false
					break
				}
			}

			if isValid {
				adjacent[mask1] = append(adjacent[mask1], mask2)
			}
		}
	}

	f := make(map[int]int)
	for mask := range valid {
		f[mask] = 1
	}

	for i := 1; i < n; i++ {
		tmp := make(map[int]int)
		for mask2 := range valid {
			for _, mask1 := range adjacent[mask2] {
				tmp[mask2] = (tmp[mask2] + f[mask1]) % modulo
			}
		}

		f = tmp
	}

	result := int(0)
	for _, num := range f {
		result = (result + num) % modulo
	}

	return result
}

func main() {
	// result: 3
	// m := int(1)
	// n := int(1)

	// result: 6
	// m := int(1)
	// n := int(2)

	// result: 580986
	m := int(5)
	n := int(5)

	// result: 
	// m := int(0)
	// n := int(0)

	result := colorTheGrid(m, n)
	fmt.Printf("result = %v\n", result)
}

