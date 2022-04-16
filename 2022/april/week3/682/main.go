package main
import (
	"fmt"
	"strconv"
)

const P = "+"
const C = "C"
const D = "D"

func calPoints(ops []string) int {
	scores := []int{}
	for _, str := range ops {
		if str == P {
			if len(scores) == 0 {
				continue
			} else if len(scores) == 1 {
				scores = append(scores, scores[0])
			} else {
				scores = append(scores, scores[len(scores) - 2] + scores[len(scores) - 1])
			}
		} else if str == C {
			if len(scores) >= 1 {
				scores = scores[:len(scores) - 1]
			}
		} else if str == D {
			if len(scores) >= 1 {
				scores = append(scores, scores[len(scores) - 1] * 2)
			}
		} else {
			num, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			scores = append(scores, num)
		}
	}

	result := int(0)
	for _, score := range scores {
		result += score
	}
	return result
}

func main() {
	// result: 30
	// ops := []string{"5","2","C","D","+"}

	// result: 27
	// ops := []string{"5","-2","4","C","D","9","+","+"}

	// result: 1
	// ops := []string{"1"}

	// result: 52
	ops := []string{"4","1","C","D","9","+","-1","10","+","-4"}

	// result: 
	// ops := []string{}

	result := calPoints(ops)
	fmt.Printf("result = %v\n", result)
}

