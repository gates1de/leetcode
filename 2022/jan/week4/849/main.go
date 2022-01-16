package main
import (
	"fmt"
)

func maxDistToClosest(seats []int) int {
	result := int(0)
	headZeroLen := int(0)
	tailZeroLen := int(0)
	isAppearedOneHead := false
	isAppearedOneTail := false

	i := int(0)
	j := len(seats) - 1 - i
	for i <= j {
		num := seats[i]
		tailNum := seats[j]

		if num == 0 {
			headZeroLen++
		} else {
			if !isAppearedOneHead {
				if headZeroLen > result {
					result = headZeroLen
				}
			} else {
				if (headZeroLen + 1) / 2 > result {
					result = (headZeroLen + 1) / 2
				}
			}
			isAppearedOneHead = true
			headZeroLen = 0
		}

		if i == j {
			break
		}

		if tailNum == 0 {
			tailZeroLen++
		} else {
			if !isAppearedOneTail {
				if tailZeroLen > result {
					result = tailZeroLen
				}
			} else {
				if (tailZeroLen + 1) / 2 > result {
					result = (tailZeroLen + 1) / 2
				}
			}
			isAppearedOneTail = true
			tailZeroLen = 0
		}

		i++
		j--
	}

	if !isAppearedOneHead || !isAppearedOneTail {
		if headZeroLen + tailZeroLen > result {
			return headZeroLen + tailZeroLen
		}
	}

	if (headZeroLen + tailZeroLen + 1) / 2 > result {
		result = (headZeroLen + tailZeroLen + 1) / 2
	}

	return result
}

func main() {
	// result: 2
	// seats := []int{1, 0, 0, 0, 1, 0, 1}

	// result: 3
	// seats := []int{1, 0, 0, 0}

	// result: 1
	// seats := []int{0, 1}

	// result: 4
	// seats := []int{0,1,0,1,1,1,0,1,1,0,0,1,0,0,0,1,0,0,0,1,0,1,0,0,1,0,1,0,0,0,1,0,0,0,0,0,0,0,1,1,1}

	// result: 6
	// seats := []int{0,0,0,0,0,0,1,0,1,1,1,0,1,1,0,0,1,0,0,0,1,0,0,0,1,0,1,0,0,1,0,1,0,0,0,1,0,0,0,0,0,0,0,1,1,1}

	// result: 7
	// seats := []int{0,0,0,0,0,1,0,1,0,1,0,0,1,0,1,0,0,0,1,0,0,0,0,0,0,0}

	// result: 8
	seats := []int{0,0,0,0,0,1,0,1,0,1,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0}

	// result: 
	// seats := []int{}

	result := maxDistToClosest(seats)
	fmt.Printf("result = %v\n", result)
}

