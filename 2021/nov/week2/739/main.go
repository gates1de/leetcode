package main
import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
    type Tmp struct {
        Index int
        Val int
    }

    st := make([]Tmp, 0, len(temperatures))
    for i := len(temperatures) - 1; i >= 0; i-- {
        v := temperatures[i]

        for len(st) != 0 && st[len(st) - 1].Val <= v {
            st = st[:len(st) - 1]
        }

        if len(st) == 0 {
            temperatures[i] = 0
        } else {
            temperatures[i] = st[len(st) - 1].Index - i
        }

        st = append(st, Tmp{i, v})
    }

    return temperatures
}

// Slow & Use more memory
func mySolution(temperatures []int) []int {
	n := len(temperatures)
	queue := [][2]int{}
	result := make([]int, n)
	for i, num := range temperatures {
		for j := 0; j < len(queue); j++ {
			temp := queue[j]
			if num <= temp[1] {
				continue
			}
			result[temp[0]] = i - temp[0]
			queue = append(queue[:j], queue[j + 1:]...)
			j--
		}
		queue = append(queue, [2]int{i, num})
	}

	return result
}

func main() {
	// result: [1,1,4,2,1,1,0,0]
	// temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	// result: [1,1,1,0]
	// temperatures := []int{30, 40, 50, 60}

	// result: [1,1,0]
	temperatures := []int{30, 60, 90}

	// result: 
	// temperatures := []int{}

	result := dailyTemperatures(temperatures)
	fmt.Printf("result = %v\n", result)
}

