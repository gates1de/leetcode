package main
import (
	"fmt"
	"strconv"
)

func openLock(deadends []string, target string) int {
    dead := make(map[int]struct{})

    for _, end := range deadends {
		num, _ := strconv.Atoi(end)
        dead[num] = struct{}{}
    }

    targetNum, _ := strconv.Atoi(target)
    offset := []func(int)int{}
    a := int(1)

    for i := 0; i < 4; i++ {
        b := a
        f1 := func(j int) int {
            if j / b % 10 == 9 {
                return j - b * 9
            } else {
                return j + b
            }
        }
        f2 := func(j int) int {
            if j / b % 10 == 0 {
                return j + b * 9
            } else {
                return j - b
            }
        }
        offset = append(offset, f1)
        offset = append(offset, f2)
        a *= 10
    }

    set := make(map[int]struct{})
    queue := []int{0}
    set[0] = struct{}{}
    result := int(0)

    for len(queue) > 0 {
        l := len(queue)

        for i := 0; i < l; i++ {
            if _, ok := dead[queue[i]]; ok {
                continue
            }
            if queue[i] == targetNum {
                return result
            }

            for _, f := range offset {
                newNum := f(queue[i])
                if _, ok := set[newNum]; ok {
                    continue
                }
                queue = append(queue, newNum)
                set[newNum] = struct{}{}
            }
        }

        queue = queue[l:]
        result++
    }
    
    return -1
}

func main() {
	// result: 6
	// deadends := []string{"0201","0101","0102","1212","2002"}
	// target := "0202"

	// result: 1
	// deadends := []string{"8888"}
	// target := "0009"

	// result: -1
	deadends := []string{"8887","8889","8878","8898","8788","8988","7888","9888"}
	target := "8888"

	// result: 
	// deadends := []string{}
	// target := ""

	result := openLock(deadends, target)
	fmt.Printf("result = %v\n", result)
}

