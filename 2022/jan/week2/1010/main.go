package main
import (
	"fmt"
)

func numPairsDivisibleBy60(time []int) int {
    result := int(0)
    mp := make([]int, 60)
    for _, t := range time {
        mp[t % 60]++
    }
    result += ((mp[0] * (mp[0] - 1)) / 2) + ((mp[30] * (mp[30] - 1)) / 2)
    for i := 1; i < 30; i++ {
        result += mp[i] * mp[60 - i]
    }
    return result
}

// Slow & Use more memory
func mySolution(time []int) int {
	result := int(0)
	m := map[int]map[int]bool{}
	divsBy60 := []int{60, 120, 180, 240, 300, 360, 420, 480, 540, 600, 660, 720, 780, 840, 900, 960}
	for i, t := range time {
		for _, val := range divsBy60 {
			if m[val - t] == nil {
				m[val - t] = map[int]bool{}
			}
			m[val - t][i] = true
		}
	}

	for i, t := range time {
		for index, _ := range m[t] {
			if i < index {
				result++
			}
		}
	}
	return result
}

func main() {
	// result: 3
	// time := []int{30, 20, 150, 100, 40}

	// result: 3
	// time := []int{60, 60, 60}

	// result: 0
	// time := []int{60}

	// result: 7
	// time := []int{25, 35, 45, 55, 65, 75, 120, 85, 95, 15}

	// result: 96
	time := []int{4,92,96,48,42,91,28,64,100,90,47,34,89,59,87,11,33,42,53,42,36,33,21,70,93,1,33,60,59,94,72,87,34,91,19,50,44,27,15,65,49,4,21,10,96,67,31,76,79,61,69,11,12,80,69,21,87,50,82,78,51,37,31,39,72,48,73,58,91,72,7,78,9,42,6,79,38,13,86,95,68,52,35,54,84,52,85,93,24,90,74,58,58,32,67,29,39,48,76,100}

	// result: 
	// time := []int{}

	result := numPairsDivisibleBy60(time)
	fmt.Printf("result = %v\n", result)
}

