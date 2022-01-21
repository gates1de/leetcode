package main
import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	m := map[int]int{}
	n := len(gas)
	for i := 0; i < n; i++ {
		m[i] = -1
	}
	for i := 0; i < n; i++ {
		remainGas := helper(i + 1, i, gas[i], m, gas, cost)
		if remainGas >= 0 {
			return i
		}
	}

	return -1
}

func helper(index int, startIndex int, currentGas int, m map[int]int, gas []int, cost []int) int {
	// fmt.Printf("index = %v, startIndex = %v, currentGas = %v\n", index, startIndex, currentGas)
	if index - 1 < 0 {
		currentGas -= cost[len(cost) - 1]
	} else {
		currentGas -= cost[index - 1]
	}

	if index >= len(gas) {
		index = 0
	}

	if index == startIndex {
		return currentGas
	}
	if currentGas < 0 || currentGas <= m[index] {
		return -1
	}

	m[index] = currentGas
	currentGas += gas[index]

	return helper(index + 1, startIndex, currentGas, m, gas, cost)
}

func main() {
	// result: 3
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}

	// result: -1
	// gas := []int{2, 3, 4}
	// cost := []int{3, 4, 3}

	// result: 7
	gas := []int{1, 1, 1, 5, 5, 1, 1, 2, 12, 1}
	cost := []int{2, 3, 5, 5, 6, 1, 2, 2, 3, 1}

	// result: 
	// gas := []int{}
	// cost := []int{}

	result := canCompleteCircuit(gas, cost)
	fmt.Printf("result = %v\n", result)
}

