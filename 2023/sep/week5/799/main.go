package main
import (
	"fmt"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
    glasses := make([][]float64, 102)
    for i, _ := range glasses {
        glasses[i] = make([]float64, 102)
    }
    glasses[0][0] = float64(poured)
    for i := 0; i <= query_row; i++ {
        for j := 0; j <= i; j++ {
            q := float64(glasses[i][j] - 1) / float64(2)
            if q > 0 {
                glasses[i + 1][j] += q
                glasses[i + 1][j + 1] += q
            }
        }
    }
    return min(float64(1), glasses[query_row][query_glass])
}

func min(a, b float64) float64 {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 0.00000
	// poured := int(1)
	// query_row := int(1)
	// query_glass := int(1)

	// result: 0.50000
	// poured := int(2)
	// query_row := int(1)
	// query_glass := int(1)

	// result: 1.00000
	poured := int(100000009)
	query_row := int(33)
	query_glass := int(17)

	// result: 
	// poured := int(0)
	// query_row := int(0)
	// query_glass := int(0)

	result := champagneTower(poured, query_row, query_glass)
	fmt.Printf("result = %v\n", result)
}

