package main
import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
    stack := []int{}
    sum := 0
    num := 0
    sign := 1
    n := len(s)
    for i:=0; i<n;i++ {
        if s[i]>='0' && s[i]<='9' {
            num  = 10*num + int(s[i]-'0')
        } else if s[i] == '+' {
            sum = sum + sign *num
            sign = 1
            num = 0
        } else if s[i] == '-' {
            sum = sum + sign * num
            sign = -1
            num = 0
        } else if s[i] == '(' {
            stack = append(stack, sum)
            stack = append(stack, sign)
            sum = 0
            sign = 1
            num = 0
        } else if s[i] == ')' {
            sum = sum + sign * num
            sum *= stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            sum += stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            num = 0
        }
    }
    return sum + sign * num
}

// Slow & Use more memory
func mySolution(s string) int {
	if len(s) <= 2 {
		result, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return result
	}

	index := int(0)
	return helper(s, &index)
}

func helper(s string, index *int) int {
	queue := []string{}
	numQueue := []rune{}
	for ; *index < len(s); *index++ {
		r := rune(s[*index])
		if r == rune(' ') {
			continue
		}
		if r == rune('(') {
			*index++
			num := helper(s, index)
			numStr := strconv.Itoa(num)
			numQueue = append(numQueue, []rune(numStr)...)
			continue
		}
		if r == rune(')') {
			break
		}

		if r == rune('+') || r == rune('-') {
			if len(numQueue) > 0 {
				numStr := string(numQueue)
				queue = append(queue, numStr)
				numQueue = []rune{}
			}
			queue = append(queue, string(r))
		} else {
			numQueue = append(numQueue, r)
		}
	}

	if len(numQueue) > 0 {
		queue = append(queue, string(numQueue))
	}

	// fmt.Printf("queue = %v\n", queue)
	isNextPlus := true
	result := int(0)
	for _, str := range queue {
		if str == "+" {
			isNextPlus = true
		} else if str == "-" {
			isNextPlus = false
		} else {
			num, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			if isNextPlus {
				result += num
			} else {
				result -= num
			}
		}
	}
	return result
}

func main() {
	// result: 2
	// s := "1 + 1"

	// result: 3
	// s := "2-1 + 2"

	// result: 23
	// s := "(1+(4+5+2)-3)+(6+8)"

	// result: 38
	// s := "1 +(45+23-32+( 18-91 ) + 79 - (56 - (2 + 48))) + 1"

	// result: -1
	s := "-2+1"

	// result: 
	// s := ""

	result := calculate(s)
	fmt.Printf("result = %v\n", result)
}

