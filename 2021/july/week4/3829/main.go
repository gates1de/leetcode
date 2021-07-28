package main
import (
	"fmt"
)

// REF: https://leetcode.com/problems/beautiful-array/discuss/793664/Divide-and-conquer-in-Golang-100-100
func beautifulArray(n int) []int {
	result := make([]int, n)
	if n <= 2 {
		for i := 1; i <= n; i++ {
			result[i - 1] = i
		}
		return result
	}

    mid := n / 2
    if n % 2 == 0 {
        tmp := beautifulArray(mid)
        for i := 0; i < mid; i++ {
            result[i] = tmp[i] * 2 - 1
            result[i + mid] = tmp[i] * 2
        }
    } else {
        tmp1 := beautifulArray(mid)
        for i := 0; i < mid; i++ {
            result[i + mid + 1] = tmp1[i] * 2
        }
        tmp2 := beautifulArray(mid + 1)
        for i := 0; i < mid + 1; i++ {
            result[i] = tmp2[i] * 2 - 1
        }
    }
    return result
}


func main() {
	// result: [2,1,4,3]
	// n := int(4)

	// result: [3,1,2,5,4]
	// n := int(5)

	// result: [1,5,3,2,6,4]
	n := int(6)

	result := beautifulArray(n)
	fmt.Printf("result = %v\n", result)
}

