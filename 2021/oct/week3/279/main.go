package main
import (
	"fmt"
	"math"
)

const MAX = int(10000)

var res = make([]int, 16000);

func numSquares(n int) int {
    if res[n] > 0 {
        return res[n]
    }
    sqrt := sqrtInt(n)
    if n == sqrt * sqrt {
        return 1
    }
    for i := 1; i <= sqrt; i++ {
        res[i * i] = 1
    }
    for i := 2; i <= n; i++ {
        if res[i] > 0 {
            continue
        }
        sqrt = sqrtInt(i)
        min := 4
        for j := 1; j <= sqrt; j++ {
            if min > res[i - j * j] + 1 {
                min = res[i - j * j] + 1
            }
        }
        res[i] = min
    }
    return res[n]
}

func sqrtInt(a int) int {
    return int(math.Sqrt(float64(a)))
}

// Slow
func mySolution(n int) int {
	if n <= 3 {
		return n
	}

	result := MAX
	maxLength := int(1)
	for i := maxLength; i * i <= n; i++ {
		maxLength = i
	}
	// fmt.Printf("maxLength = %v\n", maxLength)
	helper(0, maxLength, n, &result)
	return result
}

func helper(count int, length int, n int, result *int) {
	// fmt.Printf("count = %v, length = %v, n = %v\n", count, length, n)
	if n == 0 {
		if *result > count {
			*result = count
		}
		return
	}

	if *result <= count {
		return
	}

	if n >= length * length {
		helper(count + 1, length, n - length * length, result)
	}

	downLength := length - 1
	for downLength > 0 && downLength * downLength > n {
		downLength--
	}
	if downLength <= 0 {
		return
	} else if downLength == 1 {
		helper(count + n, downLength, 0, result)
		return
	}
	helper(count, downLength, n, result)
	helper(count + 1, downLength, n - downLength * downLength, result)
}

func main() {
	// result: 3
	// n := int(12)

	// result: 2
	// n := int(13)

	// result: 2
	// n := int(3492)

	// result: 4
	n := int(9999)

	// result: 
	// n := int(0)

	result := numSquares(n)
	fmt.Printf("result = %v\n", result)
}

