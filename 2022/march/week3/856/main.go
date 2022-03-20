package main
import (
	"fmt"
)

func scoreOfParentheses(s string) int {
    result, _ := counter(s, 0)
    return result
}

func counter(s string, currentIndex int) (int, int) {
    queue := []rune{}
    count := int(0)
    for i := currentIndex; i < len(s); i++ {
        r := rune(s[i])
        if r == rune('(') {
            if len(queue) == 0 {
                queue = append(queue, r)
                continue
            }

            innerCount, newIndex := counter(s, i)
            count += 2 * innerCount

            if newIndex == len(s) - 1 {
                return count, 0
            }

            i = newIndex
        } else if r == rune(')') {
            if len(queue) == 0 {
                return count, i
            }

            if queue[len(queue) - 1] == '(' {
                count++
            }
        }

        queue = []rune{}
    }

    return count, len(s)
}

func main() {
  // result: 1
  // s := "()"

  // result: 2
  // s := "(())"

  // result: 2
  // s := "()()"

  // result: 3
  s := "(())()"

  // result: 
  // s := ""

  result := scoreOfParentheses(s)
  fmt.Printf("result = %v\n", result)
}

