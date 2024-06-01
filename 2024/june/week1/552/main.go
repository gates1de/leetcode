package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func checkRecord(n int) int {
	dpCurrent := make([][]int, 2)
	dpNext := make([][]int, 2)
	for i, _ := range dpCurrent {
		dpCurrent[i] = make([]int, 3)
		dpNext[i] = make([]int, 3)
	}

	dpCurrent[0][0] = 1

	for length := 0; length < n; length++ {
		for totalAbsences := 0; totalAbsences <= 1; totalAbsences++ {
			for consecutiveLates := 0; consecutiveLates <= 2; consecutiveLates++ {
				dpNext[totalAbsences][0] = (dpNext[totalAbsences][0] + dpCurrent[totalAbsences][consecutiveLates]) % modulo

				if totalAbsences < 1 {
					dpNext[totalAbsences + 1][0] = (dpNext[totalAbsences + 1][0] + dpCurrent[totalAbsences][consecutiveLates]) % modulo
				}

				if consecutiveLates < 2 {
					dpNext[totalAbsences][consecutiveLates + 1] = (dpNext[totalAbsences][consecutiveLates + 1] + dpCurrent[totalAbsences][consecutiveLates]) % modulo
				}
			}
		}

		copy(dpCurrent, dpNext)
		dpNext = make([][]int, 2)
		for i, _ := range dpNext {
			dpNext[i] = make([]int, 3)
		}
	}

	result := int(0)
	for totalAbsences := 0; totalAbsences <= 1; totalAbsences++ {
		for consecutiveLates := 0; consecutiveLates <= 2; consecutiveLates++ {
			result = (result + dpCurrent[totalAbsences][consecutiveLates]) % modulo
		}
	}

	return result
}

func main() {
	// result: 8
	// n := int(2)

	// result: 3
	// n := int(1)

	// result: 183236316
	n := int(10101)

	// result: 
	// n := int(0)

	result := checkRecord(n)
	fmt.Printf("result = %v\n", result)
}

