package main
import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
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
    numQueue := []byte{}

    for ; *index < len(s); *index++ {
        r := s[*index]
        if r == ' ' {
            continue
        }
        if r == '(' {
            *index++
            num := helper(s, index)
            numStr := strconv.Itoa(num)
            numQueue = append(numQueue, []byte(numStr)...)
            continue
        }
        if r == ')' {
            break
        }

        if r == '+' || r == '-' {
            if len(numQueue) > 0 {
                numStr := string(numQueue)
                queue = append(queue, numStr)
                numQueue = []byte{}
            }
            queue = append(queue, string(r))
        } else {
            numQueue = append(numQueue, r)
        }
    }

    if len(numQueue) > 0 {
        queue = append(queue, string(numQueue))
    }

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
	// s := " 2-1 + 2 "

	// result: 23
	s := "(1+(4+5+2)-3)+(6+8)"

	// result: 
	// s := ""

	result := calculate(s)
	fmt.Printf("result = %v\n", result)
}

