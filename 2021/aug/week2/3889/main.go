package main
import (
	"fmt"
)

// REF: https://leetcode-cn.com/problems/remove-boxes/solution/yi-chu-he-zi-by-leetcode-solution/
func removeBoxes(boxes []int) int {
    dp := [100][100][100]int{}
    var calculatePoints func(boxes []int, l, r, k int) int
    calculatePoints = func(boxes []int, l, r, k int) int {
        if l > r {
            return 0
        }
        if dp[l][r][k] == 0 {
            dp[l][r][k] = calculatePoints(boxes, l, r - 1, 0) + (k + 1) * (k + 1)
            for i := l; i < r; i++ {
                if boxes[i] == boxes[r] {
                    dp[l][r][k] = max(dp[l][r][k], calculatePoints(boxes, l, i, k + 1) + calculatePoints(boxes, i + 1, r - 1, 0))
                }
            }
        }
        return dp[l][r][k]
    }
    return calculatePoints(boxes, 0, len(boxes) - 1, 0)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Time Limit Exceeded
func ngSolution(boxes []int) int {
	if len(boxes) == 1 {
		return 1
	}

	currentSum := int(0)
	boxesWithoutSingle := []int{}
	m := map[int]int{}
	for _, box := range boxes {
		m[box]++
	}
	for _, box := range boxes {
		count := m[box]
		if count == 1 {
			currentSum++
			continue
		}
		boxesWithoutSingle = append(boxesWithoutSingle, box)
	}

	result := currentSum
	pre := int(0)
	for i, box := range boxesWithoutSingle {
		if pre == box {
			continue
		}
		pre = box
		// fmt.Printf("-------------------- i = %v, boxes = %v -----------------\n", i, boxesWithoutSingle)
		copyBoxes := make([]int, len(boxesWithoutSingle))
		copy(copyBoxes, boxesWithoutSingle)
		helper(i, copyBoxes, currentSum, &result)
	}
	return result
}

func helper(index int, boxes []int, currentSum int, result *int) {
	if len(boxes) == 0 {
		return
	}

	targetBox := boxes[index]
	endIndex := index
	k := int(1)
	for i, box := range boxes[index + 1:] {
		if targetBox != box {
			break
		}
		k++
		endIndex = index + 1 + i
	}

	newBoxes := []int{}
	if len(boxes) > 1 {
		newBoxes = append(boxes[:index], boxes[endIndex + 1:]...)
	}
	currentSum += k * k
	// fmt.Printf("newBoxes = %v, currentSum = %v\n", newBoxes, currentSum)
	if *result < currentSum {
		*result = currentSum
	}
	pre := int(0)
	for i, box := range newBoxes {
		if pre == box {
			continue
		}
		pre = box
		copyNewBoxes := make([]int, len(newBoxes))
		copy(copyNewBoxes, newBoxes)
		helper(i, copyNewBoxes, currentSum, result)
	}
}

func main() {
	// result: 23
	// boxes := []int{1, 3, 2, 2, 2, 3, 4, 3, 1}

	// result: 9
	// boxes := []int{1, 1, 1}

	// result: 1
	boxes := []int{1}

	// result: 
	// boxes := []int{}

	result := removeBoxes(boxes)
	fmt.Printf("result = %v\n", result)
}

