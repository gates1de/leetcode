package main
import (
	"fmt"
)

func minEatingSpeed(piles []int, h int) int {
	pileMax := int(0)
	total := int(0)
	for _, pile := range piles {
		total += pile
		if pile > pileMax {
			pileMax = pile
		}
	}

	minH := max(total / h, 1)
	k := -1
	for k != (pileMax + minH) >> 1 {
		k = (pileMax + minH) >> 1
		if isEatOver(piles, h, k) {
			pileMax = k
		} else {
			minH = k + 1
		}
	}
	return k
}

func isEatOver(piles []int, h int, k int) bool {
	for _, pile := range piles {
		cost := pile / k
		if pile % k != 0 {
			cost++
		}
		h -= cost
		if h < 0 {
			return false
		}
	}
	return true
}

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

// Time Limit Exceeded
func ngSolution(piles []int, h int) int {
	k := int(1)
	for true {
		currentH := int(0)
		for _, pile := range piles {
			currentH += pile / k
			if pile % k != 0 {
				currentH++
			}
			if currentH > h {
				break
			}
		}

		if currentH > h {
			k++
		} else {
			return k
		}
	}

	return piles[len(piles) - 1]
}

func main() {
	// result: 4
	// piles := []int{3, 6, 7, 11}
	// h := int(8)

	// result: 30
	// piles := []int{30, 11, 23, 4, 20}
	// h := int(5)

	// result: 23
	// piles := []int{30, 11, 23, 4, 20}
	// h := int(6)

	// result: 8
	// piles := []int{8,3,9,6,1,4,2,6,19,7,8,4,3,6,18}
	// h := int(20)

	// result: 3
	// piles := []int{8}
	// h := int(3)

	// result: 14
	// piles := []int{332484035,524908576,855865114,632922376,222257295,690155293,112677673,679580077,337406589,290818316,877337160,901728858,679284947,688210097,692137887,718203285,629455728,941802184}
	// h := int(823855818)

	// result: 51
	piles := []int{1, 1000, 50000, 10000}
	h := int(1200)

	// result: 
	// piles := []int{}
	// h := int(0)

	result := minEatingSpeed(piles, h)
	fmt.Printf("result = %v\n", result)
}

