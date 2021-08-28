package main
import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
    l := 0
    h := int(math.Sqrt(float64(c))) + 1

    for l <= h {
        sum := l * l + h * h
        if sum == c {
            return true
        }else if sum > c{
            h--
        }else {
            l++
        }
    }
    return false
}

// Slow & Use more memory
func mySolution1(c int) bool {
	m := map[int]bool{}
	for i := 0; i * i <= c; i++ {
		m[i * i] = true
	}
	for square, _ := range m {
		if m[c - square] {
			return true
		}
	}
	return false
}

// Too slow
func mySolution2(c int) bool {
	// fmt.Printf("c = %v\n", c)
	squares := [50000]int{}
	maxIndex := int(0)
	for i := 0; i * i <= c; i++ {
		squares[i] = i * i
		maxIndex = i
	}
	for i := maxIndex; i >= 0; i-- {
		remain := c - squares[i]
		// fmt.Printf("s1 = %v\n", squares[i])
		for j := 0; j <= i; j++ {
			// fmt.Printf("s2 = %v\n", squares[j])
			if remain == squares[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	// result: true
	// c := int(5)

	// result: false
	// c := int(3)

	// result: true
	// c := int(4)

	// result: true
	// c := int(2)

	// result: true
	// c := int(1)

	// result: true
	// c := int(140625 + 767376)

	// result: false
	c := int(2147483647)

	// result: 
	// c := int(0)

	result := judgeSquareSum(c)
	fmt.Printf("result = %v\n", result)
}

