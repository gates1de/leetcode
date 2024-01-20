package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func sumSubarrayMins(arr []int) int {
    stack := [][2]int{[2]int{arr[0], 1}}
    result := arr[0]
    sum := arr[0]

    for i := 1; i < len(arr); i++ {
        v0 := arr[i]
        v1 := 1
        for len(stack) > 0 && v0 <= stack[len(stack) - 1][0] {
            v1 += stack[len(stack) - 1][1]
            sum -= stack[len(stack) - 1][0] * stack[len(stack) - 1][1]
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, [2]int{v0, v1})
        sum = (sum + v0 * v1) % modulo
        result = (result + sum) % modulo
    }

    return result
}

func main() {
	// result: 17
	// arr := []int{3,1,2,4}

	// result: 444
	arr := []int{11,81,94,43,3}

	// result: 
	// arr := []int{}

	result := sumSubarrayMins(arr)
	fmt.Printf("result = %v\n", result)
}

